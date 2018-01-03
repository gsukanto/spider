package SpiderServer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gogo/protobuf/proto"
	"github.com/headzoo/surf"
	"github.com/headzoo/surf/browser"
	"github.com/julienschmidt/httprouter"
	"github.com/tangr206/gocommon"
	dt "spider_data.pb"
)

func CrawTokopedia(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	param := dt.CrawParam{}
	if err := json.NewDecoder(r.Body).Decode(&param); nil != err {
		gocommon.LogDetailf("login err[%v]", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	username := param.GetUsername()
	password := param.GetPassword()

	bow, err := TokopediaLogin(username, password)
	if nil != err {
		gocommon.LogDetailf("login err[%v]", err)
		respondWithError(w, http.StatusBadRequest, "username or password not matched")
		return
	}

	userid := ParseTokoUserid(bow)
	crawRes := dt.CrawResult{Platform: proto.String(r.RequestURI)}
	crawRes.Account = ParseTokopediaAccount(userid, bow)
	gocommon.LogDetailf("get res[%v]", crawRes)
	respondWithJSON(w, http.StatusAccepted, crawRes)
}

/*
gg ------0, MoonVerified Account %!v(MISSING) %!v(MISSING)
gg ------1, test@advance.aiUbah Email %!v(MISSING) %!v(MISSING)
gg ------2, 085921421828Ubah Nomor %!v(MISSING) %!v(MISSING)*/

func genTokopediaOrderMeta(order *dt.Order, idx int, text string) {
	switch idx {
	case 0:
		order.Ordersn = &text
	case 1:
		order.Date = &text
	case 2:
		order.Amount = &text
	case 3:
		order.Status = &text
	}
}

func ParseTokopediaAccount(userid string, bow *browser.Browser) (account *dt.Account) {
	bow.Open(fmt.Sprintf(TOKO_ACCOUNT, userid))
	gocommon.LogDetailf("html2 [%v]\n", bow.Title())
	account = &dt.Account{}
	bow.Find("div.span9").Each(func(idx int, row *goquery.Selection) {
		genTokoUser(account, idx, row.Text())
	})

	// Address
	bow.Open(fmt.Sprintf(TOKO_ADDRESS, userid))
	gocommon.LogDetailf("html2 [%v]\n", bow.Title())
	bow.Find("ul.address-content").Each(func(idx int, row *goquery.Selection) {
		account.Address = append(account.Address, row.Text())
	})

	gocommon.LogDetailf("get Account[%v]", account)
	return account
}

func genTokoUser(account *dt.Account, idx int, text string) {
	switch idx {
	case 0:
		account.Name = &text
	case 1:
		account.Email = &text
	case 2:
		account.PhoneNumber = &text
	}
	return
}

func TokopediaLogin(username, pwd string) (bow *browser.Browser, err error) {
	bow = surf.NewBrowser()
	if err = bow.Open(TOKO_LOGIN); nil != err {
		return
	}
	fm, fmerr := bow.Form("form#header-frm-login")
	if nil != fmerr {
		err = fmerr
		gocommon.LogWarningf("err %v", err)
		return
	}

	fm.Input("email", username)
	fm.Input("password", pwd)
	if err = fm.Submit(); nil != err {
		gocommon.LogWarningf("err %v", err)
		return
	}

	gocommon.LogDetailf("login success[%v]", bow.Url().String())
	return
}

func ParseTokoUserid(bow *browser.Browser) string {
	userid := ""
	bow.Find("div#side-profile").Each(func(idx int, row *goquery.Selection) {
		str, _ := row.Find("a").Attr("href")
		gocommon.LogDetailf("get gguid %v, %v %v\n", idx, row.Nodes[0].Attr, str)
		seq := strings.Split(str, "/")
		if len(seq) > 1 {
			userid = seq[len(seq)-1]
		}
	})

	gocommon.LogDetailf("get userid:%v", userid)
	return userid
}
