package SpiderServer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/tangr206/gocommon"
)

var (
	g_cfg = &Config{}
	P     = fmt.Printf
)

func GetGlobalConfig() *Config {
	return g_cfg
}

func init() {
	if err := g_cfg.FromFile("conf/spider_server.xml"); nil != err {
		log.Fatal(err)
	}
	log.Println("load config done")

	initConstUrl()
}

var (
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
)

func initConstUrl() {

	LAZADA_LOGIN = g_cfg.LAZADA_LOGIN
	LAZADA_LOGIN_SUCCESS = g_cfg.LAZADA_LOGIN_SUCCESS
	LAZADA_ACCOUNT = g_cfg.LAZADA_ACCOUNT

	TOKO_LOGIN = g_cfg.TOKO_LOGIN
	TOKO_SUCCESS = g_cfg.TOKO_SUCCESS
	TOKO_ACCOUNT = g_cfg.TOKO_ACCOUNT
	TOKO_ADDRESS = g_cfg.TOKO_ADDRESS

	ELE_LOGIN = g_cfg.ELE_LOGIN
	ELE_LOGIN_REAL = g_cfg.ELE_LOGIN_REAL
	ELE_ACCOUNT = g_cfg.ELE_ACCOUNT
	ELE_ORDER = g_cfg.ELE_ORDER

	BUKA_LOGIN = g_cfg.BUKA_LOGIN
	BUKA_SUCCESS = g_cfg.BUKA_SUCCESS

	BLIBLI_LOGIN = g_cfg.BLIBLI_LOGIN
	BLIBLI_USER_LOGIN = g_cfg.BLIBLI_USER_LOGIN
	BLIBLI_ORDER = g_cfg.BLIBLI_ORDER
	BLIBLI_PROFILE = g_cfg.BLIBLI_PROFILE
	BLIBLI_ADDRESS = g_cfg.BLIBLI_ADDRESS

	BCA_LOGIN = g_cfg.BCA_LOGIN
	BCA_BALANCE = g_cfg.BCA_BALANCE
	BCA_TRANS = g_cfg.BCA_TRANS
}

var (
	BLIBLI_ORDER_LIST = []string{"https://www.blibli.com/member/order/cancel-order?page=1",
		"https://www.blibli.com/member/order/past-order?page=1",
		"https://www.blibli.com/member/order/open-order?page=1",
	}
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	gocommon.LogWarningf("returnning code[%v], message[%v]", code, message)
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Write(response)
}
