package main

import (
	_ "image/png"
	"log"
	"os"
)

const (
	ps5DiscEdition     = "https://www.amazon.de/dp/B08H93ZRK9/ref=cm_sw_em_r_mt_dp_E5DA355RKGHMCTGWFDB7"
	ps5BundleRatchet   = "https://www.amazon.de/dp/B095Z1QGWJ/ref=cm_sw_r_em_api_i_1R5XGJXCXVG29E8E4EMY"
	ps5BundleSpiderman = "https://www.amazon.de/dp/B091DZ8WZQ/ref=cm_sw_r_em_api_i_2CHR32K4930APTDW9BS3"
	ps5DigitalEdition  = "https://www.amazon.de/dp/B08H98GVK8/ref=cm_sw_r_em_api_i_HG47A8Q2XVDJ59WCTQ8C"
	myBasketURL        = "https://www.amazon.de/-/en/gp/cart/view.html"

	// Changes these const values with yours:
	lastname           = "Wayne"                // Change this with your lastname
	yourAmazonEmail    = "brucewayne@gmail.com" // Change this with your Amazon E-mail
	yourAmazonPassword = "doYouBl33d?"          // Change this with your Amazon Password
)

func main() {
	urls := [4]string{ps5DiscEdition, ps5BundleSpiderman, ps5BundleRatchet, ps5DigitalEdition}

	for {
		for index, product := range urls {
			log.Println("Trying to buy product No ", index)
			log.Println("----------------------------")
			err := amazonBuy(product)
			if err != nil {
				log.Println(err)
				log.Println("Couldn't buy this product :sadface:")
			} else {
				log.Println("No error. Product must have been purchased!")
				log.Println("Exiting now ... !")
				os.Exit(0)
			}
		}
		log.Println("Nothing bought. Retrying again with all the products from scratch...")
	}
}
