package receipt

import (
	points "HTTP-SERVER/points"
	"encoding/json"
	"testing"
)

type alphanumericTest struct {
	name           string
	expectedOutput int
}

var alphanumericTests = []alphanumericTest{
	{"Target", 6},
	{"M&M Corner Market", 14},
	{"Walmart   ", 7},
	{"Chil-fil-A", 8},
	{"2-Good-4-Dinner~", 12},
}

func TestPointsForAlphanumericCharacter(t *testing.T) {
	for _, test := range alphanumericTests {
		output := points.PointsForAlphanumericCharacters(test.name)
		if output != test.expectedOutput {
			t.Errorf("Output of %d but expected output of %d for the number of alphanumeric characters - '%s'.", output, test.expectedOutput, test.name)
		}
	}
}

type isRoundDollarTest struct {
	value          float64
	expectedOutput int
}

var isRoundDolarTests = []isRoundDollarTest{
	{1, 50},
	{4, 50},
	{0, 50},
	{1.25, 0},
	{3.40, 0},
	{1921.1492, 0},
	{1492, 50},
}

func TestPointsForIsRoundDollarTotal(t *testing.T) {
	for _, test := range isRoundDolarTests {
		output := points.PointsForRoundDollarTotal(test.value)
		if output != test.expectedOutput {
			t.Errorf("Output of %d but expected output of %d for the points rewarded from having a round dollar total - '%f'.", output, test.expectedOutput, test.value)
		}
	}
}

type isMultipleOfAQuarterTest struct {
	value          float64
	expectedOutput int
}

var isMultipleOfAQuarterTests = []isMultipleOfAQuarterTest{
	{1, 25},
	{4, 25},
	{0, 25},
	{1.25, 25},
	{3.40, 0},
	{7.75, 25},
	{5.50, 25},
	{1921.1492, 0},
	{1492, 25},
}

func TestPointsForIsMultipleOfAQuarter(t *testing.T) {
	for _, test := range isMultipleOfAQuarterTests {
		output := points.PointsForMultipleOfAQuarter(test.value)
		if output != test.expectedOutput {
			t.Errorf("Output of %d but expected output of %d for the points rewarded from having a total be a multiple of 0.25 - '%f'.", output, test.expectedOutput, test.value)
		}
	}
}

type pointsPerTwoItemsTest struct {
	numberOfItems  int
	expectedOutput int
}

var pointsPerTwoItemsTests = []pointsPerTwoItemsTest{
	{1, 0},
	{5, 10},
	{3, 5},
	{61, 150},
	{2, 5},
	{4, 10},
	{10, 25},
}

func TestPointsPerTwoItems(t *testing.T) {
	for _, test := range pointsPerTwoItemsTests {
		output := points.PointsPerTwoItems(test.numberOfItems)
		if output != test.expectedOutput {
			t.Errorf("Output of %d but expected output of %d for the points rewarded per two items - number of items '%d'.", output, test.expectedOutput, test.numberOfItems)
		}
	}
}

type receiptItem struct {
	ShortDescription string
	Price            string
}

type receiptItemPointsCalculationTest struct {
	items          []receiptItem
	expectedOutput int
}

var receiptItemPointsCalculationTests = []receiptItemPointsCalculationTest{
	{
		[]receiptItem{
			{"Mountain Dew 12PK", "6.49"},
			{"Emils Cheese Pizza", "12.25"},
			{"Knorr Creamy Chicken", "1.26"},
			{"Doritos Nacho Cheese", "3.35"},
			{"   Klarbrunn 12-PK 12 FL OZ  ", "12.00"},
		},
		6,
	},
	{
		[]receiptItem{
			{"Gatorade", "2.25"},
			{"Gatorade", "2.25"},
			{"Gatorade", "2.25"},
			{"Gatorade", "2.25"},
		},
		0,
	},
	{
		[]receiptItem{
			{"Milos Tea", "5.00"},
		},
		1,
	},
	{
		[]receiptItem{
			{"Starburst", "3.00"},
		},
		1,
	},
	{
		[]receiptItem{
			{"KitKats", "2.00"},
		},
		0,
	},
}

func TestPointsForItemCalculation(t *testing.T) {
	for _, test := range receiptItemPointsCalculationTests {
		//same conversion used in receipt to calculate points with incurring circular dependencies
		receiptItemsData, _ := json.Marshal(test.items)
		receiptItemsInterface := []map[string]string{}
		json.Unmarshal(receiptItemsData, &receiptItemsInterface)

		output := points.PointsForItemCalculation(receiptItemsInterface)
		if output != test.expectedOutput {
			t.Errorf("Output of %d but expected output of %d for the points rewarded per item if the trimmed length of the item description is a multiple of three, then calculate a value for the item - '%v'.", output, test.expectedOutput, receiptItemsInterface)
		}
	}
}

type purchaseDateIsOddTest struct {
	day            int
	expectedOutput int
}

var purchaseDateIsOddTests = []purchaseDateIsOddTest{
	{1, 6},
	{3, 6},
	{31, 6},
	{15, 6},
	{2, 0},
	{4, 0},
	{20, 0},
	{30, 0},
}

func TestPointsIfPurchaseDateIsOdd(t *testing.T) {
	for _, test := range purchaseDateIsOddTests {
		output := points.PointsForPurchaseDateIsOdd(test.day)
		if output != test.expectedOutput {
			t.Errorf("Output of %d but expected output of %d for the points rewarded if the date is odd - day '%d'.", output, test.expectedOutput, test.day)
		}
	}
}

type timeBetweenTwoAndFourPMTest struct {
	hour           int
	minute         int
	expectedOutput int
}

var timeBetweenTwoAndFourPMTests = []timeBetweenTwoAndFourPMTest{
	{2, 1, 0},
	{14, 0, 0},
	{14, 21, 10},
	{15, 15, 10},
	{16, 0, 0},
	{16, 12, 0},
	{22, 59, 0},
	{8, 16, 0},
}

func TestPointsIfTimeBetweenTwoAndFourPM(t *testing.T) {
	for _, test := range timeBetweenTwoAndFourPMTests {
		output := points.PointsForPurchaseTimeIsBetweenTwoAndFourPM(test.hour, test.minute)
		if output != test.expectedOutput {
			t.Errorf("Output of %d but expected output of %d for the points rewarded if the time is between 2 and 4 PM (or 14 and 16 ) - hour '%d', minute '%d' .", output, test.expectedOutput, test.hour, test.minute)
		}
	}
}
