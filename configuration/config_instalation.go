package configuration

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type ConfigStruct struct {
	MousePosX, MousePosY int64
	Duration             int64
	Delay                int64
}

func MakeConfig() (config ConfigStruct) {
	eofErr := false
	confMap := make(map[string]string)
	relativePath, _ := filepath.Abs("./configuration/settings")
	file, err := os.Open(relativePath)
	if err != nil {
		fmt.Println(err)
		return makeDefaultConfig()
	}
	readerFile := bufio.NewReader(file)

	for !eofErr {
		line, _, err := readerFile.ReadLine()
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
				return makeDefaultConfig()
			}
			if err == io.EOF {
				eofErr = true
				break
			}
		}
		stringArrBuf := strings.Split(string(line), "-")
		confMap[strings.Trim(stringArrBuf[0], " ")] = strings.Trim(stringArrBuf[1], " ")
	}
	config = makeParametrizedConfig(confMap)
	file.Close()
	return
}

func makeDefaultConfig() ConfigStruct {
	return ConfigStruct{MousePosX: 900, MousePosY: 900, Duration: 10, Delay: 3}
}

func makeParametrizedConfig(confMap map[string]string) ConfigStruct {
	var bufConfX, bufConfY, bufConfDuration, bufConfDelay int64
	var buf string
	var ok bool
	buf, ok = confMap["mousePosX"]
	if !ok {
		return makeDefaultConfig()
	} else {
		bufConfX, _ = strconv.ParseInt(buf, 10, 32)
	}
	buf, ok = confMap["mousePosY"]
	if !ok {
		return makeDefaultConfig()
	} else {
		bufConfY, _ = strconv.ParseInt(buf, 10, 32)
	}
	buf, ok = confMap["duration"]
	if !ok {
		return makeDefaultConfig()
	} else {
		bufConfDuration, _ = strconv.ParseInt(buf, 10, 32)
	}
	buf, ok = confMap["delay"]
	if !ok {
		return makeDefaultConfig()
	} else {
		bufConfDelay, _ = strconv.ParseInt(buf, 10, 32)
	}
	return ConfigStruct{MousePosX: bufConfX, MousePosY: bufConfY, Duration: bufConfDuration, Delay: bufConfDelay}
}
