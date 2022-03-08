package helper

import (
	models "rail/models/train"
)

var weekDayToKey = make(map[string]string)

func IsTrainAvailableOnGivenWeekDay(weekday string, trainRecord models.Train) bool {
	switch weekday {
	case "Tuesday":
		return trainRecord.RunningTue == "Y"
	case "Wednesday":
		return trainRecord.RunningWed == "Y"
	case "Thursday":
		return trainRecord.RunningThr == "Y"
	case "Friday":
		return trainRecord.RunningFri == "Y"
	case "Saturday":
		return trainRecord.RunningSat == "Y"
	case "Sunday":
		return trainRecord.RunningSun == "Y"

	default:
		return trainRecord.RunningMon == "Y"
	}
}

func FilterDetailsOnWeekdayAwailability(
	weekday string,
	trainDetails []models.Train,
) []models.Train {
	var filteredTrainDetails []models.Train
	for _, rec := range trainDetails {
		if IsTrainAvailableOnGivenWeekDay(weekday, rec) {
			filteredTrainDetails = append(filteredTrainDetails, rec)
		}
	}

	return filteredTrainDetails
}
