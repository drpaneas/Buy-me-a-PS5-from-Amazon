package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
	"os"
	"time"
)

// addToBasket() description
// You might wonder why I didn't click the "Add to Basket" button instead... a good question
// this is on purpose, because the "Add to Basket" might (sometimes) open a opop-up and ask you to buy additional
// stuff, such as extended guarantee, etc ... until to make up your mind and decide, nothing is getting added to the
// basket -- so you have to follow the pop-up dialog (if this appears). To avoid this confusing situation,
// I decided to click the "Buy now" which immediately adds the product in your basket without asking questions.
func addToBasket(ctx *context.Context, res *[]byte) error {
	log.Println("Click the 'Buy now' button in order to add the item to the basket")
	err := chromedp.Run(*ctx,
		chromedp.Click(`#buy-now-button`, chromedp.ByQuery),
		chromedp.Sleep(time.Second*5),
		chromedp.CaptureScreenshot(res),
	)

	_ = os.WriteFile("buynow.png", *res, 0644) // skip err check for saving image file to disk

	return err

}
