package main

import (
	"github.com/shagtv/newsgrabber/news"
	"github.com/shagtv/newsgrabber/router"
)

func main() {
	r := router.New()
	a := news.New()

	go a.Serve()
	r.Run()
}
