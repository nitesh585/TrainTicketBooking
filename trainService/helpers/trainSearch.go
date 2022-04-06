package helper

import (
	models "trainService/models"
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

func FilterDetailsOnWeekdayAvailability(
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
func findElementIndex(ele string, array []string) int {
	for idx, st := range array {
		if ele == st {
			return idx
		}
	}

	return -1
}

func isEqualArray(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func getPriceByClasses(i, j int, avlClasses []string) []int {
	basePriceTable := make(map[string]int)
	basePriceTable["2S"] = 60
	basePriceTable["SL"] = 100
	basePriceTable["1A"] = 300
	basePriceTable["2A"] = 250
	basePriceTable["3A"] = 200

	priceRateTable := make(map[string]int)
	priceRateTable["2S"] = 30
	priceRateTable["SL"] = 50
	priceRateTable["1A"] = 100
	priceRateTable["2A"] = 90
	priceRateTable["3A"] = 80

	var priceArray []int
	for _, ele := range avlClasses {
		priceArray = append(
			priceArray,
			basePriceTable[ele]+(j-i)*priceRateTable[ele],
		)
	}
	return priceArray
}

func CalculatePriceMany(
	source, destination string,
	trainDetails []models.Train,
) []models.Train {
	for idx := range trainDetails {
		i := findElementIndex(source, trainDetails[idx].Stations)
		j := findElementIndex(destination, trainDetails[idx].Stations)

		trainDetails[idx].Price = getPriceByClasses(i, j, trainDetails[idx].AvlClasses)
	}
	return trainDetails
}

func CalculatePriceOne(
	source, destination string,
	trainDetails models.Train,
) models.Train {
	i := findElementIndex(source, trainDetails.Stations)
	j := findElementIndex(destination, trainDetails.Stations)

	trainDetails.Price = getPriceByClasses(i, j, trainDetails.AvlClasses)
	return trainDetails
}
