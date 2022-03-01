package logger

import "testing"

func TestCLILogger(t *testing.T) {
	clog := NewCLILogger(6, 8)
	where := "asdf"
	clog.Trace("test").Msgf("Starting logging in %s", where)
	clog.Debug("12345678901234567890").Msgf("Starting logging in %s", where)
	clog.Info(where).Msgf("Starting logging in %s", where)
	clog.Warn(where).Msgf("Starting logging in %s", where)
	clog.Error(where).Msgf("Starting logging in %s", where)
	clog.Fatal(where).Msgf("Starting logging in %s", where)
	clog.Panic(where).Msgf("Starting logging in %s", where)

	clog.Highlight(where).Msgf("Starting logging in %s", where)
	clog.Success(where).Msgf("Starting logging in %s", where)
	clog.Announce(where).Msgf("Starting logging in %s", where)

	tbl := NewCLITable("S.No", "Name")
	tbl.Row("1", "Ukraine")
	tbl.Row("2", "Russia")
	tbl.Display()
}
