package validators

func IsValidPaymentMethod(method string) bool {
	validMethods := []string{"paypal", "stripe"} // integrate more payments
	for _, m := range validMethods {
		if method == m {
			return true
		}
	}
	return false
}
