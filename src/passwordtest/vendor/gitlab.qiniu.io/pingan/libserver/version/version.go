package version

import (
	"fmt"
	"time"
)

var (
	version     string
	gitCommitAt time.Time
)

func Set(rawVersion, rawCommitTime string) error {
	version = rawVersion
	t, err := parseGitCommitTime(rawCommitTime)
	if err != nil {
		return err
	}
	gitCommitAt = t
	return nil
}

func PrintVersion() {
	fmt.Println(version)
	fmt.Println("Built at", gitCommitAt.Local().Format(time.RubyDate))
}

func Version() string {
	return version
}

func GitCommitAt() time.Time {
	return gitCommitAt.UTC()
}

func parseGitCommitTime(commitTime string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339Nano, commitTime)
	if err != nil {
		if t, err = time.Parse("2006-01-02 15:04:05 -0700", commitTime); err != nil {
			return time.Now(), err
		}
	}
	return t, nil
}
