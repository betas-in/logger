package logger

import "testing"

func TestLogger(t *testing.T) {
	l := NewLogger(-1, true)
	where := "test"
	l.Debug(where).Msgf("Test logging in %s", where)
}
