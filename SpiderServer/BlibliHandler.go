package SpiderServer

import (
	"bytes"
	"encoding/json"
	"html"
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

func CrawBlibli(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	username := ps.ByName("username")
	password := ps.ByName("password")

	bow, err := BlibliLogin(username, password)
	if nil != err {
		gocommon.LogDetailf("login err[%v]", err)
		respondWithError(w, http.StatusBadRequest, "username or password not matched")
		return
	}

	crawRes := dt.CrawResult{Platform: proto.String(r.RequestURI)}
	crawRes.OrderList = ParseBlibliOrder(bow)
	crawRes.Account = ParseBlibliAccount(bow)
	gocommon.LogDetailf("get crawRes[%v]", crawRes)
	respondWithJSON(w, http.StatusAccepted, crawRes)
}

/*
		gg ------0, 0 Pesanan
	gg ------0, 1 Di pesan pada
	gg ------0, 2 Total
	gg ------0, 3 Status
	gg ------1, 0 3534571442
	gg ------1, 1 17/01/2017
	gg ------1, 2 RP 44.000
	gg ------1, 3 Telah Diterima

	                                        Terkirim pada 25/01/2017
	gg ------1, 4 Lacak pesanan sayaPengembalian
	gg ------2, 0 349798891
	gg ------2, 1 17/11/2016
	gg ------2, 2 RP 119.950
*/
func genBlibliOrderMeta(order *dt.Order, idx int, text string) {
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

func ParseBlibliOrder(bow *browser.Browser) (orderList []*dt.Order) {

	for _, order_url := range BLIBLI_ORDER_LIST {
		order := &dt.Order{}
		if err := bow.Open(order_url); nil != err {
			P("open err[%v]", err)
			return
		}
		orderExt := html.UnescapeString(bow.Body())
		order.ExtralInfo = append(order.ExtralInfo, orderExt)
		orderList = append(orderList, order)
	}
	gocommon.LogDetailf("get orders[%v]", orderList)
	return
}

func ParseBlibliAccount(bow *browser.Browser) (account *dt.Account) {
	bow.Open(BLIBLI_PROFILE)
	gocommon.LogDetailf("Now [%v]\n", bow.Url())
	account = &dt.Account{}
	bow.Find("section.content-section").Each(func(idx int, row *goquery.Selection) {
		row.Find("script").Each(func(idy int, cell *goquery.Selection) {
			text := strings.TrimSpace(cell.Text())
			if true == strings.Contains(text, "inputAddress") {
				account.ExtralInfo = append(account.ExtralInfo, text)
			}
		})
	})
	gocommon.LogDetailf("account[%v]", account)
	return
}

func BlibliLogin(username, pwd string) (bow *browser.Browser, err error) {
	bow = surf.NewBrowser()
	if err = bow.Open(BLIBLI_LOGIN); nil != err {
		return
	}

	bodyMap := map[string]string{
		"username": username,
		"password": pwd,
	}
	b, err := json.Marshal(bodyMap)
	if nil != err {
		gocommon.LogWarningf("err %v", err)
		return
	}
	body := bytes.NewReader(b)
	bow.AddRequestHeader("Accept", "application/json, text/plain, */*")
	content := "application/json;charset=UTF-8"
	err = bow.Post(BLIBLI_USER_LOGIN, content, body)
	if nil != err {
		return
	}
	gocommon.LogDetailf("login success")
	return
}
