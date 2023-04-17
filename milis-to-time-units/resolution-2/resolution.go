package resolution_home_made

import (
	"math"
)

type Result struct {
	days    int
	hours   int
	minutes int
	seconds int
}

func SplitSecondsInDaysHoursMinutesAndSeconds(durationInSeconds int) Result {
	const (
		secondsPerMinute int = 60
		secondsPerHour   int = secondsPerMinute * 60
		secondsPerDay    int = secondsPerHour * 24
	)

	remainingSeconds := durationInSeconds
	days := math.Floor(float64(remainingSeconds / (secondsPerDay)))

	remainingSeconds = remainingSeconds % (secondsPerDay)
	hours := math.Floor(float64((remainingSeconds / secondsPerHour)))

	remainingSeconds = remainingSeconds % (secondsPerHour)
	minutes := math.Floor(float64((remainingSeconds / secondsPerMinute)))

	seconds := remainingSeconds % (secondsPerMinute)

	return Result{
		days:    int(days),
		hours:   int(hours),
		minutes: int(minutes),
		seconds: int(seconds),
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
