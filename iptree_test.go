package iptree

import "testing"
import "iptree"

func TestCreate(t *testing.T) {
	_ := iptree.New()
}

func TestOne(t *testing.T) {
	ip := iptree.New()
	ip.AddByString("0.0.0.0/0", 0)
	ip.AddByString("1.2.3.4/32", 1)
	if ip.GetByString("1.2.3.4") != 1 {
		t.Error("Does not set values correctly.")
	}
	if ip.GetByString("1.2.3.3") != 0 {
		t.Error("Does not set values correctly.")
	}
	if ip.GetByString("4.3.2.1") != 0 {
		t.Error("Does not set values correctly.")
	}

}
