package SpiderServer

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gogo/protobuf/proto"
	"github.com/headzoo/surf"
	"github.com/headzoo/surf/browser"
	"github.com/julienschmidt/httprouter"
	"github.com/tangr206/gocommon"
	dt "spider_data.pb"
)

func CrawBCA(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	param := dt.CrawParam{}
	if err := json.NewDecoder(r.Body).Decode(&param); nil != err {
		gocommon.LogDetailf("login err[%v]", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	username := param.GetUsername()
	password := param.GetPassword()

	bow, err := BCALogin(username, password)
	if nil != err {
		gocommon.LogDetailf("login err[%v]", err)
		respondWithError(w, http.StatusBadRequest, "username or password not matched")
		return
	}

	crawRes := dt.CrawResult{Platform: proto.String(r.RequestURI)}
	crawRes.Account = ParseBCAAccount(bow)
	crawRes.TransHis = ParseBCATransHis(bow)
	gocommon.LogDetailf("get res[%v]", crawRes)
	respondWithJSON(w, http.StatusAccepted, crawRes)
}

func BCALogin(username, pwd string) (bow *browser.Browser, err error) {
	bow = surf.NewBrowser()
	if err = bow.Open(BCA_LOGIN); nil != err {
		return
	}

	fm, fmerr := bow.Form("form")
	if nil != fmerr {
		err = fmerr
		gocommon.LogWarningf("err %v", err)
		return
	}

	fm.Input("value(user_id)", username)
	fm.Input("value(pswd)", pwd)
	if fm.Submit() != nil {
		gocommon.LogWarningf("err %v", err)
		return
	}
	return
}

func ParseBCAAccount(bow *browser.Browser) (account *dt.Account) {
	content := "application/x-www-form-urlencoded"
	body := bytes.NewReader([]byte(""))
	bow.Post(BCA_BALANCE, content, body)

	extinfo := ""
	account = &dt.Account{}
	bow.Find("table").Each(func(idx int, row *goquery.Selection) {
		row.Find("tr").Each(func(idy int, tr *goquery.Selection) {
			extinfo += "|" + tr.Text()
		})
	})

	extinfoFormated := ParseExtInfo(extinfo)
	resByte, _ := json.Marshal(extinfoFormated)
	account.ExtralInfo = append(account.ExtralInfo, string(resByte))

	if len(extinfoFormated) == 3 {
		dataList := extinfoFormated[2]
		if len(dataList) == 4 {
			account.AccountNumber = proto.String(dataList[0])
			account.AccountType = proto.String(dataList[1])
			account.Currency = proto.String(dataList[2])
			account.Balance = proto.String(dataList[3])
		}
	}

	gocommon.LogDetailf("get account[%v]", account)
	return
}

func ParseExtInfo(str string) [][]string {
	seq1 := strings.Split(str, "|")

	resList := make([][]string, 0)
	for _, item := range seq1 {
		tempList := make([]string, 0)
		seq2 := strings.Split(item, "\n")
		for _, meta := range seq2 {
			metaStr := (strings.TrimSpace(meta))
			if len(metaStr) != 0 {
				tempList = append(tempList, metaStr)
			}
		}
		if len(tempList) != 0 {
			resList = append(resList, tempList)
		}
	}
	return resList
}

/*
value(D1):0
value(r1):2
value(x):1
value(fDt):0109
value(tDt):3009
value(submit1):View Account Statement

value(D1):0
value(r1):1
value(startDt):06
value(startMt):10
value(startYr):2017
value(endDt):24
value(endMt):10
value(endYr):2017
value(fDt):
value(tDt):
value(submit1):View Account Statement

https://ibank.klikbca.com/accountstmt.do?value(actions)=acctstmtview
*/

func ParseBCATransHis(bow *browser.Browser) (transHis *dt.BCA) {

	form := url.Values{}
	form.Add("value(D1)", "0")
	form.Add("value(r1)", "1")
	form.Add("value(startDt)", "06")
	form.Add("value(startMt)", "10")
	form.Add("value(startYr)", "2017") // TODO
	form.Add("value(endDt)", "24")     // TODO
	form.Add("value(endMt)", "10")     // TODO
	form.Add("value(endYr)", "2017")   // TODO
	form.Add("value(fDt)", "")
	form.Add("value(tDt)", "")
	form.Add("value(submit1)", "View Account Statement")

	if err := bow.PostForm(BCA_TRANS, form); nil != err {
		gocommon.LogDetailf("post err[%v]", err)
		return
	}

	transHisList := ParseBcaTransBody(bow.Body())
	resB, _ := json.Marshal(transHisList)
	transHis = &dt.BCA{}
	transHis.Extinfo = proto.String(string(resB))
	if len(transHisList) > 4 {
		transHis.TransHistory = transHisList[4]
	}
	gocommon.LogDetailf("get transHis[%v]", transHis)
	return
}

func ParseBcaTransBody(body string) (resList [][]string) {
	resList = make([][]string, 0)
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader([]byte(body)))
	if nil != err {
		gocommon.LogWarningf("new err%v", err)
		return
	}

	doc.Find("table").Each(func(idx int, row *goquery.Selection) {
		resList = append(resList, make([]string, 0))
		tabIdx := len(resList) - 1
		row.Find("tr").Each(func(idy int, tr *goquery.Selection) {
			curStr := tr.Text()
			curSeq := strings.Split(curStr, "\n")
			tmpList := make([]string, 0)
			for _, item := range curSeq {
				item := strings.TrimSpace(item)
				if len(item) != 0 {
					tmpList = append(tmpList, item)
				}
			}
			if len(tmpList) != 0 {
				rowStr := strings.Join(tmpList, ",")
				resList[tabIdx] = append(resList[tabIdx], rowStr)
			}
		})
	})

	return
}
