// Command screenshot is a chromedp example demonstrating how to take a
// screenshot of a specific element and of the entire browser viewport.
package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/tebeka/selenium"
)

var url string = `https://www.dotabuff.com/matches/`

func main() {
	wd, _ := selenium.NewRemote(selenium.Capabilities{}, "http://phantomjs:9222")
	id, _ := wd.CurrentWindowHandle()
	wd.ResizeWindow(id, 1200, 1080)

	log.Printf("initialized")
	http.HandleFunc("/match", handle(wd))
	http.ListenAndServe(":8080", nil)
}

func handle(wd selenium.WebDriver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				buf := make([]byte, 1<<16)
				runtime.Stack(buf, true)
				ret := fmt.Sprintf("error: %v\n%v", err, string(buf))
				w.Header().Set("Content-Type", "text")
				w.Write([]byte(ret))
				log.Print(ret)
				return
			}
		}()
		log.Printf("start")
		id := r.URL.Query().Get("id")
		log.Printf(id)
		err := wd.Get("https://www.dotabuff.com/matches/6153272077")
		if err != nil {
			panic(err)
		}
		log.Printf("navigate")
		buf, err := wd.Screenshot()
		if err != nil {
			panic(err)
		}
		log.Printf("get")

		w.Header().Set("Content-Type", "image/png")
		w.Write(buf)
		log.Printf("ok")
	}
}