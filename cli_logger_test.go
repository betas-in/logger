package logger

import "testing"

func TestCLILogger(t *testing.T) {
	NewCLILogger("TRACE")
	where := "test"
	Announcef("Starting logging in %s", where)

	tbl := NewCLITable("S.No", "Name")
	tbl.Row("1", "Ukraine")
	tbl.Row("2", "Russia")
	tbl.Display()
}
