package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	decisionMessage := " is clearly the better choice."

	if (option1 < option2) {
		decisionMessage = option1 + decisionMessage
	} else {
		decisionMessage = option2 + decisionMessage
	}
	
	return decisionMessage
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	price := originalPrice
	
	if (age >= 10) {
		price = originalPrice * 0.5
	} else if (age < 3) {
		price = originalPrice * 0.8
	} else {
		price = originalPrice * 0.7
	}

	return price
}
