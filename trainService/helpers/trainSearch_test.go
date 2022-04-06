package helper

import (
	"testing"
	models "trainService/models"
)

// table driven test
func TestIsTrainAvailableOnGivenWeekDay(t *testing.T) {
	trainRecord := models.Train{
		RunningMon: "Y",
		RunningTue: "N",
		RunningWed: "N",
		RunningThr: "N",
		RunningFri: "N",
		RunningSat: "Y",
		RunningSun: "Y",
	}
	cases := []struct {
		desc        string
		weekday     string
		trainRecord models.Train
		expected    bool
	}{
		{
			"TestWeekAvailable",
			"Monday",
			trainRecord,
			true,
		},
		{
			"TestWeekNotAvailable",
			"Friday",
			trainRecord,
			false,
		},
	}

	for _, tc := range cases {
		actual := IsTrainAvailableOnGivenWeekDay(tc.weekday, tc.trainRecord)
		if actual != tc.expected {
			t.Errorf("actual = %t not equals to epected = %t", actual, tc.expected)
		}
	}
}

func TestFilterDetailsOnWeekdayAwailability(t *testing.T) {
	trainRecords := []models.Train{{
		RunningMon: "Y",
		RunningTue: "N",
		RunningWed: "N",
		RunningThr: "N",
		RunningFri: "N",
		RunningSat: "Y",
		RunningSun: "Y",
	},
		{
			RunningMon: "N",
			RunningTue: "N",
			RunningWed: "N",
			RunningThr: "N",
			RunningFri: "N",
			RunningSat: "Y",
			RunningSun: "Y",
		},
		{
			RunningMon: "Y",
			RunningTue: "N",
			RunningWed: "N",
			RunningThr: "N",
			RunningFri: "N",
			RunningSat: "Y",
			RunningSun: "Y",
		},
	}

	cases := []struct {
		desc        string
		weekday     string
		trainRecord []models.Train
		expectedLen int
	}{
		{
			"TestFewTrainAvailability",
			"Monday",
			trainRecords,
			2,
		},
		{
			"TestNoTrainAvailability",
			"Wednesday",
			trainRecords,
			0,
		},
		{
			"TestFullTrainAvailability",
			"Sunday",
			trainRecords,
			3,
		},
	}

	for _, tc := range cases {
		actualLength := len(FilterDetailsOnWeekdayAvailability(tc.weekday, tc.trainRecord))
		if actualLength != tc.expectedLen {
			t.Errorf(
				"actualLength -%v is not equals to expectedLength - %v",
				actualLength,
				tc.expectedLen,
			)
		}

	}

}

func TestFindElementIndex(t *testing.T) {
	arr := []string{"world", "hello"}

	cases := []struct {
		ele      string
		array    []string
		expected int
	}{
		{"hello", arr, 1},
		{"me", arr, -1},
	}

	for _, tc := range cases {
		actual := findElementIndex(tc.ele, tc.array)
		if actual != tc.expected {
			t.Errorf("actual - %d not equals to expected - %d", actual, tc.expected)
		}

	}
}

func TestGetPriceByClasses(t *testing.T) {
	avlClasses := []string{"2S", "SL", "1A", "2A"}
	expectedAns1 := []int{120, 200, 500, 430}
	cases := []struct {
		i          int
		j          int
		avlClasses []string
		expected   []int
	}{
		{0, 2, avlClasses, expectedAns1},
	}

	for _, tc := range cases {
		actual := getPriceByClasses(tc.i, tc.j, tc.avlClasses)
		if !isEqualArray(actual, tc.expected) {
			t.Errorf("actual - %v not equals to expected - %v", actual, tc.expected)
		}
	}
}
