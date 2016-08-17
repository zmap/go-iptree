package iptree

import "testing"

func TestCreate(t *testing.T) {
	i := New()
	if i == nil {
		t.Error("new doesn't work")
	}
}

func TestOne(t *testing.T) {
	ip := New()
	ip.AddByString("0.0.0.0/0", 0)
	ip.AddByString("1.2.3.4/32", 1)
	if val, _ := ip.GetByString("1.2.3.4"); val != 1 {
		t.Error("Does not set exact value correctly.")
	}
	if val, _ := ip.GetByString("1.2.3.3"); val != 0 {
		t.Error("Add impacts other nodes.")
	}
}
