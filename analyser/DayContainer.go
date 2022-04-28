package analyser

import "time"

type DayContainer struct {
	epochDate    int
	usage        int
	readableDate time.Time
	weekd        time.Weekday
}

func (con *DayContainer) roundReadableDate() {
	con.readableDate = time.Date(con.readableDate.Year(), con.readableDate.Month(), con.readableDate.Day(),
		con.readableDate.Hour(), 0, 0, 0, con.readableDate.Location())

}
