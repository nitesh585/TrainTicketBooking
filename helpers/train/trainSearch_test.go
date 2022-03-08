package helper

import (
	models "rail/models/train"
	"testing"
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
		actualLength := len(FilterDetailsOnWeekdayAwailability(tc.weekday, tc.trainRecord))
		if actualLength != tc.expectedLen {
			t.Errorf(
				"actualLength -%v is not equals to expectedLength - %v",
				actualLength,
				tc.expectedLen,
			)
		}

	}

}
