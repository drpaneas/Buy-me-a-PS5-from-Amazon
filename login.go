package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
	"os"
	"time"
)

func login(ctx *context.Context, res *[]byte) error {
	log.Println("Visit Amazon.de and accept cookies...")
	// Accepting the cookies seems to behave better (e.g. avoid asking for CAPTCHA)
	err := chromedp.Run(*ctx,
		chromedp.Navigate("https://amazon.de"),
		chromedp.WaitReady("body"),
		chromedp.Click("#sp-cc-accept", chromedp.ByQuery),
		chromedp.Sleep(time.Second*1),
		chromedp.CaptureScreenshot(res),
	)

	_ = os.WriteFile("homepage.png", *res, 0644) // skip err check for saving the image file

	if err != nil {
		log.Println(" -> [ERROR]: The body element of the home page is not loading (see homepage.png)")
		return err
	}

	log.Println(" -> [OK]: The body element of the home page has loaded (see homepage.png)")
	log.Println("Click Sign-in")

	err = chromedp.Run(*ctx,
		chromedp.Click(`a[data-nav-role="signin"]`, chromedp.ByQuery),
		chromedp.Sleep(time.Second*2),
		chromedp.CaptureScreenshot(res),
	)

	_ = os.WriteFile("signing.png", *res, 0644) // skip err check for saving the image file

	if err != nil {
		// If it failed the first time, try reloading
		if err2 := chromedp.Run(*ctx, chromedp.Reload(), chromedp.Click(`a[data-nav-role="signin"]`, chromedp.ByQuery),
			chromedp.Sleep(time.Second*2), chromedp.CaptureScreenshot(res)); err2 != nil {
			log.Println(" -> [ERROR]: Cannot click to signin button. Probably issue with captcha (see signing.png)")
			return err
		}

	}

	log.Println(" -> [OK]: Sign-in button has been pressed (see signing.png)")
	log.Println("Type your e-mail and click 'Continue' button")

	err = chromedp.Run(*ctx,
		chromedp.SetValue(`ap_email`, yourAmazonEmail, chromedp.ByID),
		chromedp.Click(`continue`, chromedp.ByID),
		chromedp.Sleep(time.Second*1),
		chromedp.CaptureScreenshot(res),
	)

	_ = os.WriteFile("continue.png", *res, 0644) // skip err check for saving the image file

	if err != nil {
		log.Println(" -> [ERROR]: Cannot click to Continue button (see continue.png)")
		return err
	}

	log.Println(" -> [OK]: Continue button has been pressed (see continue.png)")
	log.Println("Type your password and click 'Submit' button")

	err = chromedp.Run(*ctx,
		chromedp.SetValue(`ap_password`, yourAmazonPassword, chromedp.ByID),
		chromedp.Click(`signInSubmit`, chromedp.ByID),
		chromedp.Sleep(time.Second*2),
		chromedp.CaptureScreenshot(res),
	)

	// Take a screenshot
	_ = os.WriteFile("submit.png", *res, 0644) // skip err check for saving the image file

	if err != nil {
		log.Println(" -> [ERROR]: Cannot click to Submit button (see submit.png)")
	}

	return err
}
