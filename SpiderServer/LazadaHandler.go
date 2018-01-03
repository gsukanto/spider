package SpiderServer

import (
	"encoding/json"
	"errors"
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

func CrawLazada(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	param := dt.CrawParam{}
	if err := json.NewDecoder(r.Body).Decode(&param); nil != err {
		gocommon.LogDetailf("login err[%v]", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	bow, err := LazadaLogin(param.GetUsername(), param.GetPassword())
	if nil != err {
		gocommon.LogDetailf("login err[%v]", err)
		respondWithError(w, http.StatusBadRequest, "username or password not matched")
		return
	}

	crawRes := dt.CrawResult{Platform: proto.String(r.RequestURI)}
	crawRes.OrderList = ParseLazadaOrder(bow)
	crawRes.Account = ParseLazadaAccount(bow) // will open new page in ParseAccount
	gocommon.LogDetailf("get orderList[%v]", crawRes)
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
func genLazadaOrderMeta(order *dt.Order, idx int, text string) {
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

func ParseLazadaOrder(bow *browser.Browser) (orderList []*dt.Order) {
	orderList = make([]*dt.Order, 0)
	bow.Dom().Find(".recent-orders__table-row").Each(func(idx int, row *goquery.Selection) {
		orderMeta := &dt.Order{}
		row.Find(".recent-orders__table-cell").Each(func(idy int, cell *goquery.Selection) {
			text := strings.TrimSpace(cell.Text())
			genLazadaOrderMeta(orderMeta, idy, text)
		})
		orderList = append(orderList, orderMeta)
	})
	return orderList
}

func LazadaLogin(username, pwd string) (bow *browser.Browser, err error) {
	bow = surf.NewBrowser()
	if err = bow.Open(LAZADA_LOGIN); nil != err {
		return
	}

	fm, fmerr := bow.Form("form#form-account-login")
	if nil != fmerr {
		err = fmerr
		gocommon.LogWarningf("err %v", err)
		return
	}

	fm.Input("LoginForm[email]", username)
	fm.Input("LoginForm[password]", pwd)
	if fm.Submit() != nil {
		gocommon.LogWarningf("err %v", err)
		return
	}

	if bow.Url().String() != LAZADA_LOGIN_SUCCESS {
		err = errors.New("login error")
		return
	}
	return
}

func ParseLazadaAccount(bow *browser.Browser) (account *dt.Account) {
	account = &dt.Account{}

	bow.Find("div.account-dashboard__box").Each(func(idx int, row *goquery.Selection) {
		address := row.Find("p").Text()
		account.Address = append(account.Address, address)
	})

	bow.Open(LAZADA_ACCOUNT)
	bow.Find("form.css-my-account-form").Each(func(idx int, row *goquery.Selection) {
		row.Find("div.css-my-account-form__right").Each(func(idy int, div *goquery.Selection) {
			genLazadaUser(div, account, idy, div.Text())
		})
	})

	gocommon.LogDetailf("get Account[%v]", account)
	return account
}

func genLazadaUser(div *goquery.Selection, account *dt.Account, idx int, text string) {
	switch idx {
	case 0:
		account.Email = &text
	case 1:
		account.PhoneNumber = &text
	case 2:
		val, _ := div.Find("input#EditForm_first_name").Attr("value")
		account.Name = &val
	case 3:
		div.Find("input").Each(func(idx int, input *goquery.Selection) {
			_, exist := input.Attr("checked")
			if true == exist {
				tmp, _ := input.Attr("value")
				account.Gender = &tmp
			}
		})
	case 4:
		dateStr := ""
		div.Find("option").Each(func(ida int, opt *goquery.Selection) {
			if _, exist := opt.Attr("selected"); true == exist {
				tmp, _ := opt.Attr("value")
				dateStr += tmp
			}
		})
		account.Birth = &dateStr

	}
	return
}
