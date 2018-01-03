package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/julienschmidt/httprouter"
	"github.com/tangr206/gocommon"
	"github.com/urfave/negroni"

	"./SpiderServer"
)

func createHanler() http.Handler {
	router := httprouter.New()

	router.POST("/crawl/lazada", SpiderServer.CrawLazada)
	router.POST("/crawl/tokopedia", SpiderServer.CrawTokopedia)
	router.POST("/crawl/elevenia", SpiderServer.CrawElevenia)
	router.POST("/crawl/blibli", SpiderServer.CrawBlibli)
	router.POST("/crawl/bca", SpiderServer.CrawBCA)

	return router
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	gocommon.LoggerInit("./log/spider_server.log", 3600*24, 1024*1024*16, 10, 3)

	mux := createHanler()
	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(mux)

	port := fmt.Sprintf(":%v", SpiderServer.GetGlobalConfig().ListenPort)
	err := http.ListenAndServe(port, n)
	if nil != err {
		log.Fatal(err.Error())
	}
}
