package SpiderServer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/gogo/protobuf/proto"
	"github.com/headzoo/surf"
	"github.com/headzoo/surf/browser"
	"github.com/julienschmidt/httprouter"
	"github.com/tangr206/gocommon"
	dt "spider_data.pb"
)

func CrawElevenia(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	param := dt.CrawParam{}
	if err := json.NewDecoder(r.Body).Decode(&param); nil != err {
		gocommon.LogDetailf("login err[%v]", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	username := param.GetUsername()
	password := param.GetPassword()

	bow, err := EleveniaLogin(username, password)
	if nil != err {
		gocommon.LogDetailf("login err[%v]", err)
		respondWithError(w, http.StatusBadRequest, "username or password not matched")
		return
	}

	crawRes := dt.CrawResult{Platform: proto.String(r.RequestURI)}
	crawRes.OrderList = ParseEleveniaOrder(bow)
	crawRes.Account, _ = ParseEleAccount(bow)
	gocommon.LogDetailf("get res[%v]", crawRes)
	respondWithJSON(w, http.StatusAccepted, crawRes)
}

func EleveniaLogin(username, pwd string) (bow *browser.Browser, err error) {

	bow = surf.NewBrowser()
	if err = bow.Open(ELE_LOGIN); nil != err {
		return
	}
	fmt.Println(bow.Title())
	fmt.Println(bow.Url().String())

	login_url := ELE_LOGIN_REAL
	form := url.Values{}
	form.Add("authMethod", "login")
	form.Add("email", username)
	form.Add("ordScrtNO", pwd)
	form.Add("isNonMember", "authlogin")
	form.Add("loginName", username)
	form.Add("passWord", pwd)
	err = bow.PostForm(login_url, form)
	if nil != err {
		return
	}
	fmt.Println(bow.Title())
	fmt.Println(bow.Url().String())

	return
}

func genEleAccount(td *goquery.Selection, account *dt.Account, idx int, text string) {

	switch idx {
	case 0:
		account.Email = &text
	case 1:
	case 2:
		account.Name = &text
	case 3:
		phone := ""
		td.Find("input").Each(func(idx int, input *goquery.Selection) {
			tmp, _ := input.Attr("value")
			phone += tmp
		})
		account.PhoneNumber = &phone
	case 4:
		td.Find("input").Each(func(idx int, input *goquery.Selection) {
			_, exist := input.Attr("checked")
			if true == exist {
				tmp, _ := input.Attr("value")
				account.Gender = &tmp
			}
		})
	case 5:
		dateStr := ""
		td.Find("option").Each(func(ida int, opt *goquery.Selection) {
			if _, exist := opt.Attr("selected"); true == exist {
				tmp, _ := opt.Attr("value")
				dateStr += tmp
			}
		})
		account.Birth = &dateStr

	case 6:
		td.Find("input.loingInput").Each(func(_ int, input *goquery.Selection) {
			tmp, _ := input.Attr("value")
			account.Address = append(account.Address, tmp)
		})

	}
}

func ParseEleAccount(bow *browser.Browser) (account *dt.Account, err error) {

	if err = bow.Open(ELE_ACCOUNT); nil != err {
		gocommon.LogWarningf("open %v err [%v]", ELE_ACCOUNT, err)
		return
	}

	account = &dt.Account{}
	bow.Find("table.tableDefault").Each(func(idx int, row *goquery.Selection) {
		row.Find("td").Each(func(idy int, td *goquery.Selection) {
			genEleAccount(td, account, idy, td.Text())
		})
	})
	gocommon.LogDetailf("get account %v", account)
	return
}

// TODO
func ParseEleveniaOrder(bow *browser.Browser) (orderList []*dt.Order) {

	form := url.Values{}
	form.Add("type", "orderList")
	form.Add("pageNumber", "1")
	form.Add("rows", "10")
	form.Add("shDateFrom", "20160101")
	form.Add("shDateTo", "20161211")
	form.Add("ver", "02")
	form.Add("isFromMain", "false")
	form.Add("CHOICEMENU", "A01")
	form.Add("isSSL", "Y")
	form.Add("PCID", "14814325364622003075420")
	err := bow.PostForm(ELE_ORDER, form)
	if nil != err {
		gocommon.LogWarningf("err %v", err)
		return
	}

	gocommon.LogDetailf("html2 [%v]\n", bow.Title())

	orderList = make([]*dt.Order, 0)
	bow.Find("table.shoppingList").Find("tbody").Find("tr").Each(func(idx int, row *goquery.Selection) {
		order := &dt.Order{}
		genEleveniaOrderMeta2(row, order)
		if order.GetOrdersn() == "" && len(orderList) >= 1 {
			order.Ordersn = proto.String(orderList[len(orderList)-1].GetOrdersn())
		}
		orderList = append(orderList, order)
	})
	return orderList
}

func genEleveniaOrderMeta2(row *goquery.Selection, order *dt.Order) {
	orderNum := row.Find("td.orderNum")
	ordersn := orderNum.Find("span").Text()
	order.Ordersn = &ordersn
	dateStr := orderNum.Find("strong").Text()
	order.Date = &dateStr

	text1 := row.Find("td.prod2").Text()
	order.Desc = &text1

	text2 := row.Find("td.price").Text()
	order.Amount = &text2

	text3 := row.Find("td.terms").Find("div.delivery").Text()
	order.Delivery = &text3

	text4 := row.Find("td.payment").Text()
	order.Status = &text4
}
