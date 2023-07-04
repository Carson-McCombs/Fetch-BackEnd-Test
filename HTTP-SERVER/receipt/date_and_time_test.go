package receipt

import "testing"

type createDateTest struct {
	dateString     string
	expectedOutput Date
}

var createDateTests = []createDateTest{
	{"2022-01-01", Date{2022, 01, 01}},
	{"2022-03-20", Date{2022, 03, 20}},
	{"2021-05-12", Date{2021, 05, 12}},
	{"2020-11-28", Date{2020, 11, 28}},
}

func TestCreateDate(t *testing.T) {
	for _, test := range createDateTests {
		output := *createDate(test.dateString)
		if test.expectedOutput.Day != output.Day || test.expectedOutput.Month != output.Month || test.expectedOutput.Year != output.Year {
			t.Errorf("Output of %v but expected output of %v given the string of %s", output, test.expectedOutput, test.dateString)
		}
	}
}

type createTimeTest struct {
	timeString     string
	expectedOutput Time
}

var createTimeTests = []createTimeTest{
	{"4:00", Time{4, 0}},
	{"5:01", Time{5, 1}},
	{"12:10", Time{12, 10}},
	{"16:40", Time{16, 40}},
}

func TestCreateTime(t *testing.T) {
	for _, test := range createTimeTests {
		output := *createTime(test.timeString)
		if test.expectedOutput.Hour != output.Hour || test.expectedOutput.Minute != output.Minute {
			t.Errorf("Output of %v but expected output of %v given the string of %s", output, test.expectedOutput, test.timeString)
		}
	}
}
