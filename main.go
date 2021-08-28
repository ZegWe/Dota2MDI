// Command screenshot is a chromedp example demonstrating how to take a
// screenshot of a specific element and of the entire browser viewport.
package main

import (
	"context"
	"log"
	"net/http"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
)

var chromeCtx context.Context
var url string

func main() {
	// create context
	allocator, cancel := chromedp.NewRemoteAllocator(
			context.Background(),
			"ws://chrome:9222",
		)
	defer cancel()
	chromeCtx, _ = chromedp.NewContext(
		allocator,
		// chromedp.WithDebugf(log.Printf),
	)
	// chromedp.Run(ctx, chromedp.Tasks{})
	url = `https://www.dotabuff.com/matches/`
	log.Printf("initialized")
	// capture screenshot of an element
	http.HandleFunc("/match", handle)
	http.ListenAndServe(":8080", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
		log.Printf("start")
		var buf []byte
		id := r.URL.Query().Get("id")
		log.Printf(id)
		if _, err := chromedp.RunResponse(chromeCtx, elementScreenshot(url+id, `.match-show`, &buf)); err != nil {
			log.Println(err)
			w.Header().Set("Content-Type", "text")
			w.Write([]byte("wrong"))
			return
		}
		log.Printf("get")
		
		w.Header().Set("Content-Type", "image/png")
		w.Write(buf)
		log.Printf("ok")
	}
// elementScreenshot takes a screenshot of a specific element.
func elementScreenshot(urlstr, sel string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Emulate(device.IPadPro),
		chromedp.Navigate(urlstr),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible),
	}
}