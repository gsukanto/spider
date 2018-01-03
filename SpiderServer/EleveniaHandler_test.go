package SpiderServer

import (
	"fmt"
	"testing"
)

func TestLoadElevenia(t *testing.T) {
	bow, err := EleveniaLogin(UNAME, PWD)
	CheckTestErr(err, t)
	fmt.Print(bow.Title())
	t.Log(bow.Title())

	ParseEleveniaOrder(bow)
	ParseEleAccount(bow)

	if bow.Url().String() != TOKO_SUCCESS {
		t.Errorf("invalid res[%v]", bow.Url())
	}
}
