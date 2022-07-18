package tempus

import "time"

func Now() time.Time {
	return time.Now().UTC().Truncate(time.Millisecond)
}
