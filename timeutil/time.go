package timeutil

import "time"

const (
	Layout = "2006-01-02 15:04:05"
)

func Unix() int64 {
	return time.Now().Unix()
}

func UnixNano() int64 {
	return time.Now().UnixNano()
}

func UnixMilli() int64 {
	return UnixNano() / int64(time.Millisecond)
}

func FromUnix(u int64) time.Time {
	return time.Unix(u, 0)
}

func Parse(s string, layout ...string) (time.Time, error) {
	lay := Layout
	if len(layout) > 0 {
		lay = layout[0]
	}
	return time.ParseInLocation(s, lay, time.Local)
}

func Today(t ...time.Time) time.Time {
	var y int
	var m time.Month
	var d int
	if len(t) > 0 {
		y, m, d = t[0].Date()
	} else {
		y, m, d = time.Now().Date()
	}
	return time.Date(y, m, d, 0, 0, 0, 0, time.Local)
}
