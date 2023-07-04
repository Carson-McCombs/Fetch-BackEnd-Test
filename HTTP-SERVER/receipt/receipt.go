package receipt

import (
	"HTTP-SERVER/points"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

// Stores all of the JSON Receipt data
type Receipt struct {
	//ID           uuid.UUID used for debugging purposes
	Retailer     string
	PurchaseDate string
	PurchaseTime string
	Total        string
	Items        []ReceiptItem
}

// Stores the JSON Receipt Item data
type ReceiptItem struct {
	ShortDescription string
	Price            string
}

// Stores the points corresponding to the receipt with the corresponding ID (only accessible on receipt initial post).
// The only data stored within server memory
type ReceiptPoints struct {
	ID     uuid.UUID
	Points int
}

// Stores data containing the parsed purchase time
type Time struct {
	Hour   int
	Minute int
}

// Parses the PurchaseTime string into a Time object
func createTime(s string) *Time {
	time := new(Time)
	ary := strings.Split(s, ":")
	hour, err := strconv.ParseInt(ary[0], 10, 32)
	if err != nil {
		fmt.Println("Error: hour could not be parsed -> ", err)
	}
	minute, err := strconv.ParseInt(ary[1], 10, 32)
	if err != nil {
		fmt.Println("Error: minute could not be parsed -> ", err)
	}
	time.Hour = int(hour)
	time.Minute = int(minute)
	return time
}

// Stores data containing the parsed purchase date
type Date struct {
	Year  int
	Month int
	Day   int
}

// Parses the PurchaseDate string into a Data object
func createDate(s string) *Date {
	date := new(Date)
	ary := strings.Split(s, "-")
	year, err := strconv.ParseInt(ary[0], 10, 32)
	if err != nil {
		fmt.Println("Error: year could not be parsed -> ", err)
	}
	month, err := strconv.ParseInt(ary[1], 10, 32)
	if err != nil {
		fmt.Println("Error: month could not be parsed -> ", err)
	}
	day, err := strconv.ParseInt(ary[2], 10, 32)
	if err != nil {
		fmt.Println("Error: day could not be parsed -> ", err)
	}
	date.Year = int(year)
	date.Month = int(month)
	date.Day = int(day)
	return date
}

// Calculates the point total for a given receipt and sets the points variable within the receipt object on completion.
// To calculate the point total, each method within the points package is called (each corresponding to a given criteria in the prompt)
func (receipt *Receipt) CalculatePoints() int {
	totalPoints := 0

	//parses the receipt total as a float
	receiptTotal, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil { //if there is an error, print it out and return the total as -1
		fmt.Printf("Error: Parsing Total As Float64 %s", err)
		return -1
	}
	//parse the date and time objects
	date := createDate(receipt.PurchaseDate)
	time := createTime(receipt.PurchaseTime)

	//converts list of Receipt Items to an array of Map[string]string to prevent circular dependency whilst not moving the objects out of the Receipt package
	receiptItemsData, _ := json.Marshal(receipt.Items)
	receiptItemsInterface := []map[string]string{}
	json.Unmarshal(receiptItemsData, &receiptItemsInterface)

	totalPoints += points.PointsForAlphanumericCharacters(receipt.Retailer)
	totalPoints += points.PointsForRoundDollarTotal(receiptTotal)
	totalPoints += points.PointsForMultipleOfAQuarter(receiptTotal)
	totalPoints += points.PointsPerTwoItems(len(receipt.Items))

	totalPoints += points.PointsForItemCalculation(receiptItemsInterface)
	totalPoints += points.PointsForPurchaseDateIsOdd(date.Day)
	totalPoints += points.PointsForPurchaseTimeIsBetweenTwoAndFourPM(time.Hour, time.Minute)
	fmt.Println("Total Receipt Points: ", totalPoints)
	return totalPoints
}
