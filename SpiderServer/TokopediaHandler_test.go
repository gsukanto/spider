package SpiderServer

import (
	"fmt"
	"testing"
)

func TestLoadTokopedia(t *testing.T) {
	bow, err := TokopediaLogin("kristantok01@gmail.com", "T0k0p3di4kv")
	CheckTestErr(err, t)
	fmt.Print(bow.Title())
	t.Log(bow.Title())

	userid := ParseTokoUserid(bow)
	ParseTokopediaAccount(userid, bow)

	if bow.Url().String() != TOKO_SUCCESS {
		t.Errorf("invalid res[%v]", bow.Url())
	}

}
