package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func amazonBuy(url string) error {
	var res []byte

	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithBrowserOption(), chromedp.WithLogf(log.Printf))

	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)

	defer cancel()

	// 1. Login and accept cookies
	if err := login(&ctx, &res); err != nil {
		log.Println(" -> [ERROR]: Login failed. Retrying to login again")
		return err
	}

	log.Println(" -> [OK] Login process completed without errors (see submit.png screenshot)")

	// 2. Verify that the previous login process was successful
	if err := verifyLogin(&ctx, lastname); err != nil {
		return err
	}

	log.Println(" -> [OK] Login has been verified")

	// 3. Make sure there is 0 items in the basket
	if err := verifyBasket(&ctx, &res, myBasketURL, 0); err != nil {
		return err
	}

	// 4. Check if product is available for purchase (in Stock)
	if err := isAvailable(url, &ctx, &res, lastname); err != nil {
		return err
	}

	log.Println(" -> [OK] Product is available (both 'Add to Basket' and 'Buy now' buttons exist")

	// 5. Optional: Scrape metadata. If this fails, no worries, just keep going ...
	scrapeMetadata(&ctx)

	// 6. Add the product to your basket by clicking "Buy Now" button
	if err := addToBasket(&ctx, &res); err != nil {
		return err
	}

	log.Println(" -> [OK] Item has --- probably --- been added to the basket (see buynow.png)")

	// 7. Verify if one product is actually added to the basket (assume it's the one we want)
	if err := verifyBasket(&ctx, &res, myBasketURL, 1); err != nil {
		return err
	}

	log.Println(" -> [FYI] Assume this basket's added item is the one we are interested into (see basket2.png)")

	// 8. Go to the checkout page
	if err := goToCheckout(&ctx, &res); err != nil {
		return err
	}

	log.Println("-> [OK] Clicked 'Proceed to Checkout'. (see checkout.png)")

	// 9. Buy the product
	if err := pay(&ctx, &res); err != nil {
		return err
	}

	log.Println("-> [OK] Placed my order and paid out'. (see order.png)")
	return nil
}
