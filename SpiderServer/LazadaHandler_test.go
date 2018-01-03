package SpiderServer

import (
	"fmt"
	"testing"
)

func CheckTestErr(err error, t *testing.T) {
	if nil != err {
		t.Fatalf("err[%v]", err)
	}
}

func TestLoadLazada(t *testing.T) {
	bow, err := LazadaLogin(UNAME, PWD)
	CheckTestErr(err, t)
	fmt.Print(bow.Title())
	t.Log(bow.Title())

	ParseLazadaOrder(bow)
	ParseLazadaAccount(bow)

	if bow.Url().String() != LAZADA_LOGIN_SUCCESS {
		t.Errorf("invalid res[%v]", bow.Url())
	}

}
