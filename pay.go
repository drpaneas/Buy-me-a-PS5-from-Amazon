package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
	"os"
	"time"
)

func pay(ctx *context.Context, res *[]byte) error {
	log.Println("Place your Order and Pay")

	err := chromedp.Run(*ctx,
		chromedp.Click(`#submitOrderButtonId > span > input`, chromedp.ByQuery),
		chromedp.Sleep(time.Second*5),
		chromedp.CaptureScreenshot(res),
	)

	_ = os.WriteFile("order.png", *res, 0644) // skip err check saving image file to disk
	return err
}
