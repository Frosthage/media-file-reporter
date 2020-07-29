package formats

import (
	"fmt"
	"math"
	"time"
)

func GetDuration(d time.Duration) string {

	seconds := int(math.Round(d.Seconds()))

	s := seconds % 60
	seconds -= s

	m := seconds % 3600
	m /= 60

	seconds -= seconds % 3600

	h := seconds / 3600

	return fmt.Sprintf("%02d:%02d:%02d ", h, m, s)
}
