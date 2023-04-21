package main

import "fmt"

func CalculateDiscount(currentPurchase float64, purchaseHistory []float64) (float64, error) {
	var resp, soma, media float64
	currentPurchase = 1000
	purchaseHistory = []float64{}
	for i := 0; i < len(purchaseHistory); i++ {
		soma += purchaseHistory[i]
	}
	media = soma / float64(len(purchaseHistory))
	if media > 1000 {
		resp = currentPurchase * 0.80
	} else if purchaseHistory[0] == 0 {
		resp = currentPurchase * 0.90
	} else if currentPurchase > 1000 {
		resp = currentPurchase * 0.90
	} else if currentPurchase <= 1000 && currentPurchase > 500 {
		resp = currentPurchase * 0.95
	} else if currentPurchase <= 500 {
		resp = currentPurchase * 0.98
	}
	fmt.Printf("%f", resp)
	return resp, nil
}
