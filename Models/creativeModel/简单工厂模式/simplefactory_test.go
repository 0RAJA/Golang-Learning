package simple_factory

import "testing"

func TestT1(t *testing.T) {
	api := NewAPI(1)
	s := api.Say("Tom")
	if s != "T1:Tom" {
		t.Fatal("T1 test fail")
	}
}

func TestT2(t *testing.T) {
	api := NewAPI(2)
	s := api.Say("Tom")
	if s != "T2:Tom" {
		t.Fatal("T2 test fail")
	}
}
