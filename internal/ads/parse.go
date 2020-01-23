package ads

import (
	"bufio"
	"os"
	"peeka/internal/component/encode"
	"strings"
)

type UserInfo struct {
	Name     string
	BatchNum string
	Username string
	Password string
	Org      string
}

func ReadUserFromCsvToStruct(filename string) ([]UserInfo, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	var line []byte
	var record []string
	var userinfos []UserInfo
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	for scanner.Scan() {
		line, _ = encode.DecodeGBK(scanner.Bytes())
		record = strings.Split(string(line), ",")
		userinfos = append(userinfos, UserInfo{
			Name:     record[0],
			BatchNum: record[1],
			Username: record[3],
			Password: record[4],
			Org:      record[5],
		})
	}
	return userinfos, nil
}
