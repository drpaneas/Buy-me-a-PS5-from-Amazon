# Does it work?

I still haven't got my PS5 -- hopefully this script works someday and buys me one :D
I tries to buy one out of a couple of PS5 standalone and bundles (see the links in the `main.go` and also feel free to replace them if they are outdated).

for example:

```go
ps5DiscEdition     = "https://www.amazon.de/dp/B08H93ZRK9/ref=cm_sw_em_r_mt_dp_E5DA355RKGHMCTGWFDB7"
ps5BundleRatchet   = "https://www.amazon.de/dp/B095Z1QGWJ/ref=cm_sw_r_em_api_i_1R5XGJXCXVG29E8E4EMY"
ps5BundleSpiderman = "https://www.amazon.de/dp/B091DZ8WZQ/ref=cm_sw_r_em_api_i_2CHR32K4930APTDW9BS3"
ps5DigitalEdition  = "https://www.amazon.de/dp/B08H98GVK8/ref=cm_sw_r_em_api_i_HG47A8Q2XVDJ59WCTQ8C"
myBasketURL        = "https://www.amazon.de/-/en/gp/cart/view.html"
```

Feel free to add more or delete.
Make sure you reflect this changes to the code as well:

```go
urls := [4]string{ps5DiscEdition, ps5BundleSpiderman, ps5BundleRatchet, ps5DigitalEdition}
```