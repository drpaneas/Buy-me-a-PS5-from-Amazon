package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"strings"
)

var ErrLastnameNotFound = errors.New(" -> [ERROR] Couldn't find your lastname. Login to Amazon.de has failed.")

func verifyLogin(ctx *context.Context, lastname string) error {
	log.Println("Testing if login is actually correct (using your lastname as valid data ...")
	errBodyEmpty := errors.New(" -> [ERROR] The <body> HTML element is empty. Login to Amazon.de has failed.")

	var body string // HTML <body> element

	err := chromedp.Run(*ctx, chromedp.Text(`body`, &body, chromedp.NodeVisible, chromedp.ByQuery))

	// Check if body element is populated
	if len(body) > 0 {
		if !strings.Contains(body, fmt.Sprintf("Hallo, %s", lastname)) {
			err = ErrLastnameNotFound
		}
	} else {
		err = errBodyEmpty
	}

	return err
}
