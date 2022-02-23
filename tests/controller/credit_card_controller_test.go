package controller

import "testing"

func TestPay(t *testing.T) {
	got := 10
	want := 5

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
