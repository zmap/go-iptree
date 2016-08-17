package iptree

import "testing"

func TestCreate(t *testing.T) {
	i := New()
	if i == nil {
		t.Error("new doesn't work")
	}
}

func TestExactValues(t *testing.T) {
	ip := New()
	ip.AddByString("1.2.3.4/32", 1)
	ip.AddByString("1.2.3.5/32", 2)
	if val, _, _ := ip.GetByString("1.2.3.4"); val != 1 {
		t.Error("Does not set exact value correctly.")
	}
	if val, _, _ := ip.GetByString("1.2.3.5"); val != 2 {
		t.Error("Does not set exact value correctly.")
	}
}

func TestCovering(t *testing.T) {
	ip := New()
	ip.AddByString("0.0.0.0/0", 1)
	if val, found, _ := ip.GetByString("1.2.3.4"); !found {
		t.Error("Values within covering value not found.")
	} else if val != 1 {
		t.Error("Value within covering set not correct")
	}
}
