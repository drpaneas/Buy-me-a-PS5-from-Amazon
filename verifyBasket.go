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

func verifyBasket(ctx *context.Context, res *[]byte, myBasketURL string, expectedAddedItems int) error {
	errMoreThanOne := errors.New(" -> [ERROR] There are more than one items in the basket (see basket2.png)")
	errItemScrape := errors.New(" -> [ERROR] Could not scrape the '#sc-subtotal-label-buybox' element")
	HTMLElement := `#sc-subtotal-label-buybox`

	// All the cases for the basket
	var expectedAddedItemsString string
	if expectedAddedItems == 0 {
		HTMLElement = `#sc-active-cart > div > div > div > h1`
		expectedAddedItemsString = "Your Amazon Basket is empty."
	}
	if expectedAddedItems == 1 {
		expectedAddedItemsString = "1 item"
	}
	if expectedAddedItems > 1 {
		expectedAddedItemsString = fmt.Sprintf("%v items", expectedAddedItems)
	}
	errItemEmpty := fmt.Errorf(" -> [ERROR] the '%v' element is empty", HTMLElement)

	log.Println("Visiting my Amazon basket page ...")

	err := chromedp.Run(*ctx,
		chromedp.Navigate(myBasketURL),
		chromedp.WaitReady("body"),
		chromedp.Sleep(time.Second*2),
		chromedp.CaptureScreenshot(res),
	)

	_ = os.WriteFile("basket2.png", *res, 0644) // skip err check for saving image file to disk

	if err != nil {
		return err
	}

	log.Println(" --> [OK] Basket page has been loaded (see basket.png)")

	// Check the number of the existing basket items is matching the expected number of basket items

	log.Printf("Check if there are %v item(s) into the basket\n", expectedAddedItems)
	var item string
	if err = chromedp.Run(*ctx, chromedp.Text(HTMLElement, &item, chromedp.NodeVisible, chromedp.ByQuery)); err != nil {
		return errItemScrape
	}

	if len(item) > 0 {
		if strings.Contains(item, expectedAddedItemsString) {
			log.Println(" -> [OK] Basket has ", expectedAddedItems, " number of item(s)")
		} else {
			return errMoreThanOne
		}
	} else {
		return errItemEmpty
	}

	return err
}
