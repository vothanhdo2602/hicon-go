package natsio

import (
	"fmt"
	"strings"
)

func GenerateReqrepSubject(channel string) string {
	return fmt.Sprintf("%s.reqrep.%s", stream, channel)
}

func GenerateJetstreamSubject(channel, action string) string {
	return fmt.Sprintf("%s.jetstream.%s.%s", stream, channel, action)
}

func GenerateQueueNameFromSubject(subject string) string {
	return strings.ReplaceAll(subject, ".", "_")
}

func GetUpdateConfigSubject() string {
	return GenerateReqrepSubject(UpdateConfigChannel)
}

func GetFindByPrimaryKeysSubject() string {
	return GenerateReqrepSubject(FindByPrimaryKeysChannel)
}
