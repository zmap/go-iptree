package blacklist

import "testing"

func TestCreate(t *testing.T) {
	i := New()
	if i == nil {
		t.Error("new doesn't work")
	}
}

func TestBlacklist(t *testing.T) {
	bl := New()
	bl.ParseFromFile("example-blacklist.conf")
	if val, _ := bl.IsBlacklisted("1.2.3.4"); val != false {
		t.Error("Does not set exact value correctly.")
	}
	if val, _ := bl.IsBlacklisted("148.73.0.0"); val != true {
		t.Error("Does not set exact value correctly.")
	}

}
