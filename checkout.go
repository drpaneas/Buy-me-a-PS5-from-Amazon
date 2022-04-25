package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
	"os"
	"time"
)

func goToCheckout(ctx *context.Context, res *[]byte) error {
	log.Println("Proceed to Checkout")

	err := chromedp.Run(*ctx,
		chromedp.Click(`#sc-buy-box-ptc-button > span > input`, chromedp.ByQuery),
		chromedp.Sleep(time.Second*2),
		chromedp.CaptureScreenshot(res),
	)

	_ = os.WriteFile("checkout.png", *res, 0644) // skip err check saving image file to disk

	return err
}
