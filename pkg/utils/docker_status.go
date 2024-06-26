package utils

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

type Mem struct {
	LimitUsage uint64
	Usage      uint64

	Stats MemStat
}

func (m Mem) Percent() float64 {
	if m.LimitUsage != 0 {
		return m.MemUsageNoCache() / float64(m.LimitUsage) * 100
	}
	return -1
}

func (m Mem) MemUsageNoCache() float64 {
	// cgroup v1
	if v, isCgroup1 := m.Stats["total_inactive_file"]; isCgroup1 && v < m.Usage {
		return float64(m.Usage - v)
	}
	// cgroup v2
	if v := m.Stats["inactive_file"]; v < m.Usage {
		return float64(m.Usage - v)
	}
	return float64(m.Usage)
}

type MemStat map[string]uint64

/*
	/sys/fs/cgroup/memory/memory.limit_in_bytes

	/sys/fs/cgroup/memory/memory.usage_in_bytes

	/sys/fs/cgroup/memory/memory.stat
*/

func CGroupMem() (Mem, error) {
	stat, err := parseMemStat()
	if err != nil {
		return Mem{}, err
	}
	limitUsage, err := parseMemLimit()
	if err != nil {
		return Mem{}, err
	}
	usage, err := parseMemUsage()
	if err != nil {
		return Mem{}, err
	}
	return Mem{
		LimitUsage: limitUsage,
		Usage:      usage,
		Stats:      stat,
	}, nil
}

var (
	ErrLines = errors.New("not correct line format")
)

func parseMemStat() (MemStat, error) {
	lines, err := ReadFileLines("/sys/fs/cgroup/memory/memory.stat")
	if err != nil {
		return nil, err
	}
	return ParseMemStat(lines)
}

func parseMemLimit() (uint64, error) {
	lines, err := ReadFileLines("/sys/fs/cgroup/memory/memory.limit_in_bytes")
	if err != nil {
		return 0, err
	}
	return ParseMemLimit(lines)
}

func parseMemUsage() (uint64, error) {
	lines, err := ReadFileLines("/sys/fs/cgroup/memory/memory.usage_in_bytes")
	if err != nil {
		return 0, err
	}
	return ParseMemUsage(lines)
}

func ParseMemStat(lines []string) (MemStat, error) {
	var mem = make(MemStat)
	for i := range lines {
		line := lines[i]
		fields := strings.Split(line, " ")
		if len(fields) != 2 {
			continue
		}
		value, err2 := strconv.ParseUint(fields[1], 10, 64)
		if err2 != nil {
			return nil, err2
		}
		name := fields[0]
		mem[name] = value
	}
	return mem, nil
}

func ParseMemLimit(lines []string) (uint64, error) {
	if len(lines) != 1 {
		return 0, ErrLines
	}
	return strconv.ParseUint(lines[0], 10, 64)
}

func ParseMemUsage(lines []string) (uint64, error) {
	if len(lines) != 1 {
		return 0, ErrLines
	}
	return strconv.ParseUint(lines[0], 10, 64)
}

func ReadFileLines(path string) ([]string, error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	reader := bufio.NewReader(fd)
	lines := make([]string, 0, 10)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return lines, nil
			}
			return nil, err
		}
		lines = append(lines, strings.TrimSpace(line))
	}
}
