package main

import "testing"

func TestDoStuff(t *testing.T) {
	doStuff()
	if 1 != 1 {
		t.Fail()
	}
}
