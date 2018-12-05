package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"time"
)

type Record struct {
	Start time.Time
	End   time.Time
}
type RecordSet struct {
	ID    int
	Sleep []Record
	Total float64
}

func NewRecordSet(id int) *RecordSet {
	data := RecordSet{id, make([]Record, 0), 0}
	return &data
}

func main() {
	data, _ := ioutil.ReadFile("./data.txt")
	lines := strings.Split(string(data), "\n")
	fmt.Printf("Input parsed, %d lines\n", len(lines))

	//due to date format, plain string sort can be used
	sort.Strings(lines)

	// prepare the world
	guards := make(map[int]*RecordSet)

	var rec *RecordSet
	var istart, iend time.Time
	var err error
	for _, line := range lines {
		if strings.Contains(line, "Guard") {
			var iguard int
			var ok bool
			fmt.Sscanf(line[19:], "Guard #%d begins shift", &iguard)
			rec, ok = guards[iguard]
			if !ok {
				rec = NewRecordSet(iguard)
				guards[iguard] = rec
			}
		} else if strings.Contains(line, "wakes up") {
			iend, err = time.Parse("2006-01-02 15:04", line[1:17])
			if err != nil {
				fmt.Printf("Problem with date parsing %s\n", line)
			}
			rec.Sleep = append(rec.Sleep, Record{istart, iend})
		} else if strings.Contains(line, "falls asleep") {
			istart, err = time.Parse("2006-01-02 15:04", line[1:17])
			if err != nil {
				fmt.Printf("Problem with date parsing %s\n", line)
			}
		} else {
			fmt.Println("Problem with state parsing")
		}
	}

	maxTime := 0.0
	var sleppyGuard *RecordSet
	for _, guard := range guards {
		for _, timeRec := range guard.Sleep {
			guard.Total += timeRec.End.Sub(timeRec.Start).Minutes()
		}

		if guard.Total >= maxTime {
			sleppyGuard = guard
			maxTime = guard.Total
		}
	}

	maxCount := 0
	minute := 0
	for i := 0; i < 59; i++ {
		count := 0
		for _, timeRec := range sleppyGuard.Sleep {
			if timeRec.Start.Minute() <= i && timeRec.End.Minute() > i {
				count++
			}
		}
		if count > maxCount {
			minute = i
			maxCount = count
		}
	}

	fmt.Printf("Sleepy guard %d, %f minutes, favorite %d, count %d, key %d \n", sleppyGuard.ID, sleppyGuard.Total, minute, maxCount, sleppyGuard.ID*minute)

	maxCount = 0
	minute = 0
	sleppyGuard = nil
	for i := 0; i < 59; i++ {
		for _, guard := range guards {
			count := 0
			for _, timeRec := range guard.Sleep {
				if timeRec.Start.Minute() <= i && timeRec.End.Minute() > i {
					count++
				}
			}

			if count > maxCount {
				minute = i
				maxCount = count
				sleppyGuard = guard
			}
		}
	}

	fmt.Printf("Sleepy guard %d, %f minutes, favorite %d, count %d, key %d \n", sleppyGuard.ID, sleppyGuard.Total, minute, maxCount, sleppyGuard.ID*minute)

}
