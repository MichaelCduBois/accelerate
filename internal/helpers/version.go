package helpers

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func GetGoVersion() (string, error) {
	cmd := exec.Command("go", "version")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Unable to get go version: ", err)
		return "", err
	}
	outputString := string(output)
	outStrParts := strings.Split(outputString, " ")
	if len(outStrParts) < 3 {
		return "", errors.New("unable to parse go version")
	}
	rawGoVersion := outStrParts[2]
	return string(rawGoVersion[2:]), nil
}
