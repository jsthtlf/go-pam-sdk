package recorder

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jsthtlf/go-pam-sdk/pkg/logger"
	"github.com/jsthtlf/go-pam-sdk/pkg/model"
	"github.com/jsthtlf/go-pam-sdk/pkg/storage"
	"github.com/jsthtlf/go-pam-sdk/pkg/utils"
)

type sessionProvider interface {
	CreateSessionCommand(commands []*model.Command) (err error)
	UploadReplay(sid, gZipFile string) error

	SessionFinished(sid string, time time.Time) error
	FinishReply(sid string) error
}

type oldReplay struct {
	SessionId string
	IsGzip    bool
	Version   utils.ReplayVersion
}

func UploadOldReplays(replayFolderPath string, replayConfig model.ReplayConfig, p sessionProvider) {
	replayStorage := storage.NewReplayStorage(p, replayConfig)
	allRemainFiles := make(map[string]oldReplay)
	_ = filepath.Walk(replayFolderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		if replayInfo, ok := parseReplayFilename(info.Name()); ok {
			if err2 := p.SessionFinished(replayInfo.SessionId, info.ModTime()); err2 != nil {
				logger.Error("Remain session finish failed: ", err2)
				return nil
			}
			allRemainFiles[path] = replayInfo
		}
		return nil
	})

	for absPath, remainFile := range allRemainFiles {
		absGzPath := absPath
		if !remainFile.IsGzip {
			switch remainFile.Version {
			case utils.Version2:
				if err := validateRemainReplayFile(absPath); err != nil {
					continue
				}
				absGzPath = absPath + utils.SuffixReplayGz
			case utils.Version3:
				absGzPath = absPath + utils.SuffixGz
			default:
				absGzPath = absPath + utils.SuffixGz
			}

			if err := utils.CompressToGzipFile(absPath, absGzPath); err != nil {
				logger.Error("Remain replay compress failed: ", err)
				continue
			}
			_ = os.Remove(absPath)
		}
		Target, _ := filepath.Rel(replayFolderPath, absGzPath)
		logger.Infof("Upload remain replay (%s) to storage: %s", absGzPath, replayStorage.TypeName())
		if err2 := replayStorage.Upload(absGzPath, Target); err2 != nil {
			logger.Errorf("Upload remain replay file (%s) failed: %+v", absGzPath, err2)
			continue
		}
		if err := p.FinishReply(remainFile.SessionId); err != nil {
			logger.Errorf("Finish remain replay session (%s) failed: %+v", remainFile.SessionId, err)
			continue
		}
		_ = os.Remove(absGzPath)
		logger.Infof("Upload remain replay (%s) is success", absGzPath)
	}
	logger.Info("Upload remain replay is done")
}

func validateRemainReplayFile(path string) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	tmp := make([]byte, 1)
	_, err = f.Seek(-1, 2)
	if err != nil {
		return err
	}
	_, err = f.Read(tmp)
	if err != nil {
		return err
	}
	switch string(tmp) {
	case "}":
		return nil
	case ",":
		_, err = f.Write([]byte(`"0":""}`))
	default:
		_, err = f.Write([]byte(`}`))
	}
	return err
}

func parseReplayFilename(filename string) (replay oldReplay, ok bool) {
	if len(filename) == 36 {
		replay.SessionId = filename
		replay.Version = utils.Version2
		ok = true
		return
	}
	if replay.SessionId, replay.Version, ok = isReplayFile(filename); ok {
		replay.IsGzip = isGzipFile(filename)
	}
	return
}

func isGzipFile(filename string) bool {
	return strings.HasSuffix(filename, utils.SuffixGz)
}

func isReplayFile(filename string) (id string, version utils.ReplayVersion, ok bool) {
	suffixesMap := map[string]utils.ReplayVersion{
		utils.SuffixCast:     utils.Version3,
		utils.SuffixCastGz:   utils.Version3,
		utils.SuffixReplayGz: utils.Version2}
	for suffix := range suffixesMap {
		if strings.HasSuffix(filename, suffix) {
			sidName := strings.Split(filename, ".")[0]
			if len(sidName) == 36 {
				id = sidName
				version = suffixesMap[suffix]
				ok = true
				return
			}
		}
	}
	return
}
