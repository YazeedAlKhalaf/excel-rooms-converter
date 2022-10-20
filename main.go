package main

import (
	"fmt"
	"rooms-excel-converter/internal"

	"github.com/elliotchance/orderedmap/v2"
	"github.com/xuri/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("ROOMS.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, err := f.GetRows("Sheet4")
	if err != nil {
		fmt.Println(err)
		return
	}

	rowsNormalized := make([][]string, 0)
	for _, row := range rows {
		rowToNormalize := row

		if len(rowToNormalize) == 0 || rowToNormalize[0] == "Times" {
			continue
		}

		if rowToNormalize[0] == "Days" {
			rowToNormalize[0] = "time_start"
			rowToNormalize[1] = "time_end"
		}
		rowsNormalized = append(rowsNormalized, rowToNormalize)
	}

	indexWithRooms := orderedmap.NewOrderedMap[int, string]()
	for index, row := range rowsNormalized {
		if row[1] == "Room:" {
			roomName := row[2]
			indexWithRooms.Set(index, roomName)
		}
	}

	roomsRowsCount := make(map[string]int)
	for el := indexWithRooms.Front(); el != nil; el = el.Next() {
		if el == nil || el.Next() == nil {
			continue
		}

		roomsRowsCount[(*el).Value] = ((*el).Next().Key - (*el).Key) - 1
	}

	rooms := make(map[string]*internal.Room)
	lastRoomName := ""
	for index, row := range rowsNormalized {
		if row[1] == "Room:" {
			lastRoomName = row[2]

			if rooms[lastRoomName] == nil {
				rooms[lastRoomName] = &internal.Room{
					Name: lastRoomName,
				}
			} else {
				rooms[lastRoomName].SetName(lastRoomName)
			}

			continue
		}

		if row[0] == "time_start" {
			continue
		}

		if isNotDayName(row[3]) {
			timeStart, err := internal.ParseTimeOfDay(row[0])
			if err != nil {
				fmt.Println("index:", index, "error parsing time start:", err)
				continue
			}

			timeEnd, err := internal.ParseTimeOfDay(row[1])
			if err != nil {
				fmt.Println("index:", index, "error parsing time end:", err)
				continue
			}

			// make sure each row is 8 cells long.
			if len(row) < 8 {
				for i := len(row); i < 8; i++ {
					row = append(row, "")
				}
			}

			rooms[lastRoomName].AppendTimeSlotToSunday(internal.TimeSlot{
				TimeStart:  *timeStart,
				TimeEnd:    *timeEnd,
				CourseName: row[3],
			})
			rooms[lastRoomName].AppendTimeSlotToMonday(internal.TimeSlot{
				TimeStart:  *timeStart,
				TimeEnd:    *timeEnd,
				CourseName: row[4],
			})
			rooms[lastRoomName].AppendTimeSlotToTuesday(internal.TimeSlot{
				TimeStart:  *timeStart,
				TimeEnd:    *timeEnd,
				CourseName: row[5],
			})
			rooms[lastRoomName].AppendTimeSlotToWednesday(internal.TimeSlot{
				TimeStart:  *timeStart,
				TimeEnd:    *timeEnd,
				CourseName: row[6],
			})
			rooms[lastRoomName].AppendTimeSlotToThursday(internal.TimeSlot{
				TimeStart:  *timeStart,
				TimeEnd:    *timeEnd,
				CourseName: row[7],
			})
		}
	}
}

func isNotDayName(s string) bool {
	if s == "Sunday" || s == "Monday" || s == "Tuesday" || s == "Wednesday" || s == "Thursday" {
		return false
	}

	return true
}
