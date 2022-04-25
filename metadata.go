package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
)

func scrapeMetadata(ctx *context.Context) {
	log.Println("Printing metadata, if possible")

	// ---------------------------------------------------------------- //
	var productTitle string
	if err := chromedp.Run(*ctx, chromedp.Text(`#productTitle`, &productTitle, chromedp.NodeVisible, chromedp.ByQuery)); err != nil {
		log.Println(" -> [DEBUG] Could not scrape the '#productTitle' element.", err)
	}

	if len(productTitle) > 0 {
		log.Println(" -> [INFO] Product is", productTitle)
	} else {
		log.Println(" -> [DEBUG] Empty name for ProductTitle")
	}

	// ---------------------------------------------------------------- //

	var productPrice string
	priceTag := `#corePrice_feature_div > div > span > span.a-offscreen`
	if err := chromedp.Run(*ctx, chromedp.Text(priceTag, &productPrice, chromedp.NodeVisible, chromedp.ByQuery)); err != nil {
		log.Println(" -> [DEBUG]  Could not scrape the '#corePrice_feature_div > div > span > span.a-offscreen' element.", err)
	}

	if len(productTitle) > 0 {
		log.Println(" -> [INFO] Product price is", productPrice, " Euro")
	} else {
		log.Println(" -> [DEBUG] Empty name for productPrice")
	}

	// ---------------------------------------------------------------- //

	var productDelivery string
	if err := chromedp.Run(*ctx, chromedp.Text(`#contextualIngressPtLabel_deliveryShortLine`, &productDelivery, chromedp.NodeVisible, chromedp.ByQuery)); err != nil {
		log.Println(" -> [DEBUG] Could not scrape the '#contextualIngressPtLabel_deliveryShortLine", err)
	}

	if len(productTitle) > 0 {
		log.Println(" -> [INFO] Shipping: ", productDelivery)
	} else {
		log.Println(" -> [DEBUG] Empty name for productDelivery")
	}
}
