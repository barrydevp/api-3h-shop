package parser

import (
	"errors"
	"strings"
)

const ErrorMsg = "invalid mysql url"

func ParseMysqlUrl(url string) (string, error) {
	protocolAndRest := strings.Split(url, "://")
	if len(protocolAndRest) != 2 {
		return "", errors.New(ErrorMsg)
	}

	_ = protocolAndRest[0]
	hostAndRest := strings.Split(protocolAndRest[1], "/")

	if len(hostAndRest) != 2 {
		return "", errors.New(ErrorMsg)
	}

	host := strings.Split(hostAndRest[0], "@")
	if len(host) != 2 {
		return "", errors.New(ErrorMsg)
	}

	userCredentials := host[0]

	address := host[1]

	dbNameAndAttributes := strings.Split(hostAndRest[1], "?")
	if len(dbNameAndAttributes) < 1 {
		return "", errors.New(ErrorMsg)
	}
	dbName := dbNameAndAttributes[0]

	returnUrl := userCredentials + "@tcp(" + address + ")/" + dbName

	return returnUrl, nil
}
