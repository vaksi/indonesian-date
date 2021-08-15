package indonesiandate

import (
	"fmt"
	"time"
)

type IndonesianDate struct {
	t time.Time
}

func New(t time.Time) *IndonesianDate {
	return &IndonesianDate{
		t: t,
	}
}

const (
	FormatIDDateFull         = "2 Januari 2006"
	FormatIDDateSmall        = "2 Jan 2006"
	FormatIDDateWithDayFull  = "Rabu, 2 Januari 2006"
	FormatIDMonthYear        = "Januari 2006"
	FormatIDDateWithTimeFull = "2 Januari 2006 00:00:00"
)

func (idTime IndonesianDate) Format(layout string) string {
	switch layout {
	case FormatIDDateFull:
		return fmt.Sprintf("%d %s %d",
			idTime.t.Day(), months[idTime.t.Month()-1], idTime.t.Year(),
		)
	case FormatIDDateWithDayFull:
		return fmt.Sprintf("%s, %d %s %d",
			days[idTime.t.Weekday()], idTime.t.Day(), months[idTime.t.Month()-1], idTime.t.Year(),
		)
	case FormatIDDateSmall:
		return fmt.Sprintf("%d %s %d",
			idTime.t.Day(), months[idTime.t.Month()-1][:3], idTime.t.Year(),
		)
	case FormatIDDateWithTimeFull:
		return fmt.Sprintf("%d %s %d  %02d:%02d:%02d",
			idTime.t.Day(), months[idTime.t.Month()-1], idTime.t.Year(), idTime.t.Hour(), idTime.t.Minute(), idTime.t.Second())
	case FormatIDMonthYear:
		return fmt.Sprintf("%s %d",
			months[idTime.t.Month()-1], idTime.t.Year(),
		)
	default:
		return ""
	}
}

var days = [...]string{
	"Minggu", "Senin", "Selasa", "Rabu", "Kamis", "Jum'at", "Sabtu",
}

var months = [...]string{
	"Januari", "Febuari", "Maret", "April", "Mei", "Juni",
	"Juli", "Agustus", "September", "Oktober", "November", "Desember",
}
