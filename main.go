package main

import (
	"encoding/json"
	"fmt"
	"os"
	"rooms-excel-converter/internal"

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

	rooms := make(map[string]*internal.Room)
	lastRoomName := ""
	for index, row := range rows {
		if len(row) == 0 || row[0] == "Times" || row[0] == "Days" {
			continue
		}

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

		if !internal.IsDayName(row[3]) {
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

	roomsJsonByte, err := json.Marshal(rooms)
	if err != nil {
		fmt.Println(err)
		return
	}

	os.WriteFile("rooms.json", roomsJsonByte, 0644)
}
