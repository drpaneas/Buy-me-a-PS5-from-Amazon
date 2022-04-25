package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"os"
	"strings"
	"time"
)

func isAvailable(url string, ctx *context.Context, res *[]byte, lastname string) error {
	errCannotScrapeBody := errors.New("-> [ERROR] could not scrape the 'body' element")
	errNotAvailable := errors.New("-> [ERROR] Product is unavailable")
	errLoginFailed := errors.New("-> [ERROR] Login has stopped working although " +
		"it worked fine previously (see loggedin2.png)")
	errBodyIsEmpty := errors.New(" -> [ERROR] For some reason the <body> HTML element is empty")
	log.Println("Checking availability for ", url, " ...")

	var body string
	err := chromedp.Run(*ctx,
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
		chromedp.Sleep(time.Second*2),
		chromedp.CaptureScreenshot(res),
	)
	_ = os.WriteFile("loggedin2.png", *res, 0644) // skip error check for saving image file to disk

	if err != nil {
		return err
	}

	log.Println(" -> [OK] Product page has been loaded")

	log.Println("Testing if you still remain logged in that new page")

	if err = chromedp.Run(*ctx, chromedp.Text(`body`, &body, chromedp.NodeVisible, chromedp.ByQuery)); err != nil {
		return errCannotScrapeBody
	}

	if len(body) > 0 {
		if strings.Contains(body, fmt.Sprintf("Deliver to %s", lastname)) {
			log.Println(" -> [OK] Login still works fine (see: loggedin2.png)")
			if !(strings.Contains(body, "Add to Basket") && strings.Contains(body, "Buy now")) {
				return errNotAvailable
			}
		} else {
			log.Println("Debugging 2nd failed login:")
			log.Println(body)

			return errLoginFailed
		}
	} else {
		return errBodyIsEmpty
	}

	return err
}
