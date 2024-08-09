package common

import (
	"time"

	"github.com/google/uuid"
)

func ExtractTimestampFromUUIDString(uuidStr string) time.Time {
	uuid := uuid.MustParse(uuidStr)

	t := uuid.Time()
	sec, nsec := t.UnixTime()
	timeStamp := time.Unix(sec, nsec)
	return timeStamp
}

func UuidStringToTimeString(uuidStr string) string {
	uuid := uuid.MustParse(uuidStr)

	t := uuid.Time()
	sec, nsec := t.UnixTime()
	timeStamp := time.Unix(sec, nsec)
	return TimeToString(timeStamp)
}

func TimeToString(time time.Time) string {
	return time.UTC().Format("2006-01-02T15:04:05Z07:00")
}
