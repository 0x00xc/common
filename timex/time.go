package timex

import "time"

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
	lay := "2006-01-02 15:04:05"
	if len(layout) > 0 {
		lay = layout[0]
	}
	return time.Parse(s, lay)
}
