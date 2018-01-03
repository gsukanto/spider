package SpiderServer

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Config struct {
	XMLName    xml.Name `xml:"config"`
	ListenPort string

	LogLevel          int32
	LogPath           string
	LogRotateInterval int64
	LogSizeMax        int64
	LogCountMax       int64
	LogRequest        bool
	LogResponse       bool
	UseNewLog         bool
	NeedSqlLog        bool
	LogIDMethod       int32

	RunningLevel string
	StopCmd      []int32

	ReadTimeout  int64
	WriteTimeout int64

	LAZADA_LOGIN         string
	LAZADA_LOGIN_SUCCESS string
	LAZADA_ACCOUNT       string
	TOKO_LOGIN           string
	TOKO_SUCCESS         string
	TOKO_ACCOUNT         string
	TOKO_ADDRESS         string
	ELE_LOGIN            string
	ELE_LOGIN_REAL       string
	ELE_ACCOUNT          string
	ELE_ORDER            string
	BUKA_LOGIN           string
	BUKA_SUCCESS         string
	BLIBLI_LOGIN         string
	BLIBLI_USER_LOGIN    string
	BLIBLI_ORDER         string
	BLIBLI_PROFILE       string
	BLIBLI_ADDRESS       string
	BCA_LOGIN            string
	BCA_BALANCE          string
	BCA_TRANS            string
}

func (config Config) String() string {
	return fmt.Sprintf(`TCP listen at "%s"\n`, config.ListenPort)
}

func (config *Config) FromFile(filename string) (err error) {
	configFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	err = xml.Unmarshal(configFile, config)

	return
}
