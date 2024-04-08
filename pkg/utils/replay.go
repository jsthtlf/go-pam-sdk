package utils

import "strings"

type ReplayVersion string

const (
	UnKnown  ReplayVersion = ""
	Version2 ReplayVersion = "2"
	Version3 ReplayVersion = "3"
)

const (
	SuffixReplayGz = ".replay.gz"
	SuffixCastGz   = ".cast.gz"
	SuffixCast     = ".cast"
	SuffixGz       = ".gz"
)

var SuffixMap = map[ReplayVersion]string{
	Version2: SuffixReplayGz,
	Version3: SuffixCastGz,
}

func ParseReplayVersion(gzFile string, defaultValue ReplayVersion) ReplayVersion {
	for version, suffix := range SuffixMap {
		if strings.HasSuffix(gzFile, suffix) {
			return version
		}
	}
	return defaultValue
}
