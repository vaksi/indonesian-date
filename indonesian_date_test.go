package indonesiandate

import (
	"fmt"
	"testing"
	"time"
)

func TestIndonesianDate_Format(t *testing.T) {
	tID := New(time.Now())
	fmt.Println("FormatIDDateFull: " + tID.Format(FormatIDDateFull))
	fmt.Println("FormatIDDateWithDayFull: " + tID.Format(FormatIDDateWithDayFull))
	fmt.Println("FormatIDDateWithTimeFull: " + tID.Format(FormatIDDateWithTimeFull))
	fmt.Println("FormatIDMonthYear: " + tID.Format(FormatIDMonthYear))
}
