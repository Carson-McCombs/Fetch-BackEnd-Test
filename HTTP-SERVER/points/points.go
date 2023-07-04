package points

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

//Instead of having each of these functions returning a boolean variable that the receipt function
// would have to check before adding in its points to its calculation. I put the points rewarded from
// each criteria met within their respective functions. These could instead be stored within constant
// variables within the points package, but I didn't feel it was necessary.

var alphanumericRegex = regexp.MustCompile(`[a-zA-Z0-9]`) //regex pattern for alphanumeric characters and no symbols

func isMultipleOf(total float64, multiple float64) bool {
	decimalValues := math.Mod(total, multiple) //checks if the total modulus multiple is equal to 0 (in other words if the total divided by the multiple has a remainder of 0)
	return decimalValues == 0
}

func PointsForAlphanumericCharacters(name string) int {
	parsedName := alphanumericRegex.FindAllString(name, -1) //parses out all of the character that are not alphanumeric
	value := len(parsedName)                                //takes the length of the alphanumeric string
	fmt.Printf("Total Value From Number of Alphanumeric Characters in Retailer Name: %d\n", value)
	return value
}

func PointsForRoundDollarTotal(total float64) int {
	if isMultipleOf(total, 1) { //if the total is a multiple of 1.00
		fmt.Println("The Price Total is a Multiple of a Dollar (+50)")
		return 50
	}
	return 0
}

func PointsForMultipleOfAQuarter(total float64) int {
	if isMultipleOf(total, 0.25) { // if the total is a multiple of 0.25
		fmt.Println("The Price Total is a Multiple of 0.25 (+25)")
		return 25
	}
	return 0
}

func PointsPerTwoItems(itemsLength int) int {
	value := (itemsLength / 2) * 5 //value is equal to the number of items divided by two, rounded down, times 5
	fmt.Printf("Points from Number of Items is %d", value)
	return value
}

func PointsForItemCalculation(items []map[string]string) int {
	total := 0
	for _, item := range items { //for each item
		total += pointsForCalculateItemValue(item) //calculate the items value and add it to the total
	}
	fmt.Printf("All Items Give a Total Value of %d\n", total)
	return total
}

func pointsForCalculateItemValue(item map[string]string) int {
	descriptionLength := len(strings.Trim(item["ShortDescription"], " ")) //trims the description length
	if !isMultipleOf((float64(descriptionLength)), 3) {                   //if the trimmed description length is not a multiple of 3
		return 0
	}
	price, err := strconv.ParseFloat(item["Price"], 64) //parses the receipt price as a float
	if err != nil {                                     //if the price cannot be parsed
		return -99999
	}
	value := int(math.Ceil(price * 0.2)) //multiple the price by 0.2 and round up to the nearest integer
	fmt.Printf("Item %s is worth %d\n", item["ShortDescription"], value)
	return value
}

func PointsForPurchaseDateIsOdd(day int) int {
	if math.Mod((float64)(day), 2) == 1 { //if the date is ODD
		fmt.Println("Date is Odd (+6)")
		return 6
	}
	return 0
}

func PointsForPurchaseTimeIsBetweenTwoAndFourPM(hour int, minute int) int {
	if hour < 14 { //return 0 if time is before 2:00 pm
		return 0
	}
	if hour >= 16 { //return 0 if time is after 4:00 pm
		return 0
	}
	if hour == 14 && minute == 0 { //return 0 if time is at 2:00 pm exactly
		return 0
	}
	fmt.Println("Time is Between 2 and 4 pm (+10)")
	return 10
}
