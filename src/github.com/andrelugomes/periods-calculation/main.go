package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"time"

	"github.com/andrelugomes/time-periods/periods"
)

func main() {
	file, err := os.Open("issues.csv")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	rows, err := read(file)
	if err != nil {
		panic(err)
	}

	fmt.Println("################################### PURE DATA ################################################")
	assigners := mapAssignersToPeriod(rows)
	calculateDuration(assigners)

	fmt.Println("#################################### NORMALIZED DATA ########################################")
	periods := normalizeConflicts(assigners)
	result := calculateDuration(periods)

	fmt.Println("#################################### RESULT ################################################")
	fmt.Println(result)
}

func read(rs io.ReadSeeker) (rows [][]string, err error) {
	// Skip first row (line)
	row1, err := bufio.NewReader(rs).ReadSlice('\n')
	if err != nil {
		return nil, err
	}
	_, err = rs.Seek(int64(len(row1)), io.SeekStart)
	if err != nil {
		return nil, err
	}

	// Read remaining rows
	r := csv.NewReader(rs)
	rows, err = r.ReadAll()
	return
}

func mapAssignersToPeriod(rows [][]string) (assigners map[string][]periods.Period) {
	assigners = make(map[string][]periods.Period)

	for _, row := range rows {
		assigners[row[10]] = append(assigners[row[10]], periods.Period{Start: toTime(row[3]), End: toTime(row[4])})
	}
	return
}

func calculateDuration(assigners map[string][]periods.Period) map[string]float64 {
	result := make(map[string]float64)
	for assignee, periodsArr := range assigners {
		realHours := 0.0
		for _, period := range periodsArr {
			start := period.Start
			end := period.End
			hours := end.Sub(start).Hours()
			hoursWithoutWeekend := hours - float64(periods.Weekends(period)*24)
			realHours = realHours + realHour(hoursWithoutWeekend)
			fmt.Println(assignee, start, end, hours, hoursWithoutWeekend, realHours)
		}
		result[assignee] = math.RoundToEven(realHours)
	}
	return result
}

func realHour(h float64) float64 {
	if h > 8 { return (h / 24) * 8} else {return h}
}

func normalizeConflicts(assigners map[string][]periods.Period) (aux map[string][]periods.Period) {
	aux = make(map[string][]periods.Period)

	for assignee, p := range assigners {
		aux[assignee] = periods.Normalize(p)
	}
	return
}

func toTime(date string) time.Time {
	layout := "02/01/2006 15:04:05"
	time, err := time.Parse(layout, date)

	if err != nil {
		panic(err)
	}
	return time
}
