package user_test

import (
	"apps/user"
	"testing"
)

func TestToBson(t *testing.T) {
	user := *user.New("Xuong", "Tran Trieu")
	t.Log(user.ToBson())
}
