package analyser

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func ReturnAnalysis() map[string]map[int]float32 {

	data := readInCSV()

	sortingList := make(map[int][]DayContainer)
	resList := make(map[int]map[int][]int)
	for i := 0; i < 7; i++ {
		resList[i] = make(map[int][]int)
	}

	// round dates down to hour and sort them by weekday
	for _, con := range data {
		con.roundReadableDate()

		sortingList[int(con.weekd)] = append(sortingList[int(con.weekd)], con)
	}

	// take dates and sort by day and hour for usage
	for _, day := range sortingList {
		for _, val := range day {
			resList[int(val.weekd)][val.readableDate.Hour()] = append(resList[int(val.weekd)][val.readableDate.Hour()], val.usage)

		}
	}

	writableList := make(map[string]map[int]float32)
	for i := 0; i < 7; i++ {
		ind := strconv.Itoa(i)

		writableList[ind] = make(map[int]float32)
	}

	for i, weekday := range resList {
		for j, day := range weekday {
			sum := 0
			var average float32

			for _, hour := range day {
				sum += hour
			}

			average = float32(sum / len(day))

			ind := strconv.Itoa(i)
			writableList[ind][j] = average
		}
	}

	return writableList

}

func readInCSV() []DayContainer {
	var res []DayContainer

	csvLog, err := os.Open("./util_log.csv")
	if err != nil {
		log.Panicf("Failed to open csv file: %s", err)
	}

	defer csvLog.Close()

	csvLines, err := csv.NewReader(csvLog).ReadAll()

	for _, line := range csvLines {
		epo, _ := strconv.Atoi(line[0])
		usage, err := strconv.Atoi(line[1])
		if err != nil {
			fmt.Println(err)
		}

		readDate := time.Unix(int64(epo), 0)

		if usage == -1 {
			continue
		}

		con := DayContainer{
			epochDate:    epo,
			usage:        usage,
			readableDate: readDate,
		}

		con.weekd = con.readableDate.Weekday()

		res = append(res, con)
	}

	return res

}
