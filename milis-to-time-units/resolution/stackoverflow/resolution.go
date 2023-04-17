package stackoverflow

import (
	"math"

	"github.com/vituchon/labora-golang-course/milis-to-time-units/resolution"
)

func SplitSecondsInDaysHoursMinutesAndSeconds(durationInSeconds int) resolution.Result {
	durationInMilis := durationInSeconds * 1000
	return splitMsInDaysHoursMinutesAndSeconds(durationInMilis)
}

// https://go.dev/play/p/fA_rb-znoeu
func splitMsInDaysHoursMinutesAndSeconds(durationInMilis int) resolution.Result {
	const (
		second int = 1000
		minute int = second * 60
		hour   int = minute * 60
		day    int = hour * 24
	)
	days := math.Floor(float64(durationInMilis / (day)))
	hours := math.Floor(float64((durationInMilis % (day)) / (hour)))
	minutes := math.Floor(float64((durationInMilis % (hour)) / (minute)))
	seconds := math.Floor(float64((durationInMilis % (minute)) / (second)))

	return resolution.Result{
		Days:    int(days),
		Hours:   int(hours),
		Minutes: int(minutes),
		Seconds: int(seconds),
	}
}

// https://stackoverflow.com/questions/76023850/javascript-function-returns-time-in-negative

/*
function splitInDaysHoursMinutesAndSeconds(durationInMilis) {

    let second = 1000;
    let minute = second * 60;
    let hour = minute * 60;
    let day = hour * 24;

    let d = Math.floor(durationInMilis / (day));
    let h = Math.floor((durationInMilis % (day)) / (hour));
    let m = Math.floor((durationInMilis % (hour)) / (minute));
    let s = Math.floor((durationInMilis % (minute)) / (second));

    return {
        d: d,
        h: h,
        m: m,
        s: s
    }
}*/
