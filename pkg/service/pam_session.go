package service

import (
	"fmt"

	"github.com/jsthtlf/go-pam-sdk/pkg/common"
	"github.com/jsthtlf/go-pam-sdk/pkg/model"
)

func (s *PAMService) Upload(sessionID, gZipFile string) error {
	version := model.ParseReplayVersion(gZipFile, model.Version3)
	return s.UploadReplay(sessionID, gZipFile, version)
}

func (s *PAMService) UploadReplay(sid, gZipFile string, version model.ReplayVersion) error {
	var res map[string]interface{}
	Url := fmt.Sprintf(SessionReplayURL, sid)
	fields := make(map[string]string)
	fields["version"] = string(version)
	return s.authClient.PostFileWithFields(Url, gZipFile, fields, &res)
}

func (s *PAMService) FinishReply(sid string) error {
	data := map[string]bool{"has_replay": true}
	return s.sessionPatch(sid, data)
}

func (s *PAMService) CreateSession(sess model.Session) error {
	_, err := s.authClient.Post(SessionListURL, sess, nil)
	return err
}

func (s *PAMService) SessionSuccess(sid string) error {
	data := map[string]bool{
		"is_success": true,
	}
	return s.sessionPatch(sid, data)
}

func (s *PAMService) SessionFailed(sid string, err error) error {
	data := map[string]bool{
		"is_success": false,
	}
	return s.sessionPatch(sid, data)
}
func (s *PAMService) SessionDisconnect(sid string) error {
	return s.SessionFinished(sid, common.NewNowUTCTime())
}

func (s *PAMService) SessionFinished(sid string, time common.UTCTime) error {
	data := map[string]interface{}{
		"is_finished": true,
		"date_end":    time,
	}
	return s.sessionPatch(sid, data)
}

func (s *PAMService) sessionPatch(sid string, data interface{}) error {
	Url := fmt.Sprintf(SessionDetailURL, sid)
	_, err := s.authClient.Patch(Url, data, nil)
	return err
}

func (s *PAMService) GetSessionById(sid string) (data model.Session, err error) {
	reqURL := fmt.Sprintf(SessionDetailURL, sid)
	_, err = s.authClient.Get(reqURL, &data)
	return
}

func (s *PAMService) CreateSessionTicketRelation(sid, ticketId string) (err error) {
	data := map[string]string{
		"session": sid,
		"ticket":  ticketId,
	}
	_, err = s.authClient.Post(TicketSessionURL, data, nil)
	return
}
