# How the code works

If you are curious, here's how it works ...

It basically uses the `chromedp` library to click through Amazon website.
The logic can be found at the function `amazonBuy()` and it consists of several small steps (each step has its own `$functionName.go` file) which they come is a screenshot (so you can debug what's actually happening)

```go
// 1. Login and accept cookies
if err := login(&ctx, &res); err != nil {
    log.Println(" -> [ERROR]: Login failed. Retrying to login again")
    return err
}
```

It first visits `Amazon.de` (that's the one I've tested) and it accepts the cookies.
I've found out after testing that login is behaving better after accepting the cookies, for example it doesn't ask me for CAPTCHA (fingerscrossed).

```go
// 2. Verify that the previous login process was successful
if err := verifyLogin(&ctx, lastname); err != nil {
    return err
}
```

Then it verifies you have actually logged into your account by checking your lastname.
For example, if you login now to your Amazon.de account you will see it says: "Hello, $lastname". This is the check it performs to make sure you signed it.

```go
// 3. Make sure there is 0 items in the basket
if err := verifyBasket(&ctx, &res, myBasketURL, 0); err != nil {
    return err
}
```

Then it goes to your basket and makes sure is empty!
This is because the script is not going to cherry-pick individual items from your basket.
If you have any items already there, handing leftovers, the script will exit and it won't even try to check for PS5.
So, make sure your basket is empty before starting this.

```go
// 4. Check if product is available for purchase (in Stock)
if err := isAvailable(url, &ctx, &res, lastname); err != nil {
    return err
}
```

Then it visits the PS5 links in sequence (one by one -- and not in parallel).
It checks (once again) if you are still remaining logged in, using again your lastname, but this time it uses another field, that is `Lieferung an $YOURLASTNAME` which corrersponds to your default shipping address (I suppose your shipping address is registered with your own lastname).
It also tries to see if the product is *in stock*, by checking if the buttons "Add to Basket" and "Buy Now" do exist. If not, then you cannot buy this product, thus it's not available.


```go
// 5. Optional: Scrape metadata. If this fails, no worries, just keep going ...
    scrapeMetadata(&ctx)
```

Then it scrapes some interesting metadata for your eyes only (it doesn't do any cross-validation to make sure this is PS5 product or anything like that -- so make sure the URL links are correct).

```go
// 6. Add the product to your basket by clicking "Buy Now" button
if err := addToBasket(&ctx, &res); err != nil {
 return err
}
```

Then it clicks "Buy Now" instead of "Add to Basket".
This is done on purpose because sometimes when you click ton "Add to Basket" there is a probability when a pop-up dialog appears out of nowhere and asks you to buy more extra stuff (e.g. extetend guarantee, etc).
Unfortunatelly I don't know how to handle this pop-up situation with `chromedp`, so I chose to go with "Buy Now" button, which immediately adds the product in your Basket -- no questions asked, no pop-ups.
Having said that, it would make sense to setup the "Buy Now" functionality with your account, if you haven't already done so.

```go
// 7. Verify if one product is actually added to the basket (assume it's the one we want)
if err := verifyBasket(&ctx, &res, myBasketURL, 1); err != nil {
    return err
}
```

Then it visits the basket and checks there's 1 item (the one it added just earlier).
Beware: it doesn't verifies if this is actually the PS5, but it assumes it is because it clicked the button at the link (make sure the links are valid, otherwise it will buy you wrong stuff).

```go
// 8. Go to the checkout page
if err := goToCheckout(&ctx, &res); err != nil {
    return err
}
```

It goes to the "Checkout" page. Simple as that.

```go
// 9. Buy the product
if err := pay(&ctx, &res); err != nil {
    return err
}
```

And finally it buys the product with your default payment method and default shipping address.
Make sure your payment method doesn't require any manual steps and it just works non-interactively.
