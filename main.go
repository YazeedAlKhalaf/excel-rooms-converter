package main

import (
	"encoding/json"
	"fmt"
	"os"
	"rooms-excel-converter/internal"
	"sort"

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

	rows, err := f.GetRows("tweeq-schedule")
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

			if row[3] != "" {
				rooms[lastRoomName].AppendTimeSlotToSunday(internal.TimeSlot{
					TimeStart:  *timeStart,
					TimeEnd:    *timeEnd,
					CourseName: row[3],
				})
			}

			if row[4] != "" {
				rooms[lastRoomName].AppendTimeSlotToMonday(internal.TimeSlot{
					TimeStart:  *timeStart,
					TimeEnd:    *timeEnd,
					CourseName: row[4],
				})
			}

			if row[5] != "" {
				rooms[lastRoomName].AppendTimeSlotToTuesday(internal.TimeSlot{
					TimeStart:  *timeStart,
					TimeEnd:    *timeEnd,
					CourseName: row[5],
				})
			}

			if row[6] != "" {
				rooms[lastRoomName].AppendTimeSlotToWednesday(internal.TimeSlot{
					TimeStart:  *timeStart,
					TimeEnd:    *timeEnd,
					CourseName: row[6],
				})
			}

			if row[7] != "" {
				rooms[lastRoomName].AppendTimeSlotToThursday(internal.TimeSlot{
					TimeStart:  *timeStart,
					TimeEnd:    *timeEnd,
					CourseName: row[7],
				})
			}
		}
	}

	for _, r := range rooms {
		sort.Slice(r.Sunday, func(i, j int) bool {
			return r.Sunday[i].TimeEnd.IsBefore(&r.Sunday[j].TimeStart)
		})
		sort.Slice(r.Monday, func(i, j int) bool {
			return r.Monday[i].TimeEnd.IsBefore(&r.Monday[j].TimeStart)
		})
		sort.Slice(r.Tuesday, func(i, j int) bool {
			return r.Tuesday[i].TimeEnd.IsBefore(&r.Tuesday[j].TimeStart)
		})
		sort.Slice(r.Wednesday, func(i, j int) bool {
			return r.Wednesday[i].TimeEnd.IsBefore(&r.Wednesday[j].TimeStart)
		})
		sort.Slice(r.Thursday, func(i, j int) bool {
			return r.Thursday[i].TimeEnd.IsBefore(&r.Thursday[j].TimeStart)
		})
	}

	for _, r := range rooms {
		// originalSundayLen := len(r.Sunday)
		// for i, ts := range r.Sunday {
		// 	currentSundayLen := len(r.Sunday)
		// 	if currentSundayLen == i+1+(currentSundayLen-originalSundayLen) {
		// 		break
		// 	}

		// 	ts = r.Sunday[i+(currentSundayLen-originalSundayLen)]
		// 	nts := r.Sunday[i+1+(currentSundayLen-originalSundayLen)]
		// 	if !ts.TimeEnd.IsEqual(&nts.TimeStart) {
		// 		freeTS := internal.TimeSlot{
		// 			TimeStart:  ts.TimeEnd,
		// 			TimeEnd:    nts.TimeStart,
		// 			CourseName: "",
		// 		}
		// 		r.Sunday = internal.InsertToSlice(r.Sunday, freeTS, i+1+(currentSundayLen-originalSundayLen))
		// 	}
		// }
		r.Sunday = addBreaksToTimeSlots(r.Sunday)
		r.Monday = addBreaksToTimeSlots(r.Monday)
		r.Tuesday = addBreaksToTimeSlots(r.Tuesday)
		r.Wednesday = addBreaksToTimeSlots(r.Wednesday)
		r.Thursday = addBreaksToTimeSlots(r.Thursday)
	}

	roomsJsonByte, err := json.Marshal(rooms)
	if err != nil {
		fmt.Println(err)
		return
	}

	os.WriteFile("rooms.json", roomsJsonByte, 0644)
}

func addBreaksToTimeSlots(tss []internal.TimeSlot) []internal.TimeSlot {
	originalTSSLen := len(tss)
	for i, ts := range tss {
		currentTSSLen := len(tss)
		if currentTSSLen == i+1+(currentTSSLen-originalTSSLen) {
			break
		}

		ts = tss[i+(currentTSSLen-originalTSSLen)]
		nts := tss[i+1+(currentTSSLen-originalTSSLen)]
		if !ts.TimeEnd.IsEqual(&nts.TimeStart) {
			freeTS := internal.TimeSlot{
				TimeStart:  ts.TimeEnd,
				TimeEnd:    nts.TimeStart,
				CourseName: "",
			}
			tss = internal.InsertToSlice(tss, freeTS, i+1+(currentTSSLen-originalTSSLen))
		}
	}

	return tss
}
