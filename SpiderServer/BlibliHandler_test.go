package SpiderServer

import (
	"fmt"
	"testing"
)

func TestLoadBlibli(t *testing.T) {
	bow, err := BlibliLogin("kristantok01@gmail.com", "b3l1bel1kv")
	CheckTestErr(err, t)
	fmt.Print(bow.Title())
	t.Log(bow.Title())

	ParseBlibliOrder(bow)
	ParseBlibliAccount(bow)

	if bow.Url().String() != LAZADA_ACCOUNT {
		t.Errorf("invalid res[%v]", bow.Url())
	}
}
