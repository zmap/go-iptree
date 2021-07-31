/*
 * ZGrab Copyright 2016 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

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
	if val, _, _ := ip.GetByString("1.2.3.4"); val.(int) != 1 {
		t.Error("Does not set exact value correctly.")
	}
	if val, _, _ := ip.GetByString("1.2.3.5"); val.(int) != 2 {
		t.Error("Does not set exact value correctly.")
	}
}

func TestDelete(t *testing.T) {
	ip := New()
	ip.AddByString("1.2.3.4/24", 1)
	ip.AddByString("1.2.3.5/32", 2)
	if val, _, _ := ip.GetByString("1.2.3.4"); val.(int) != 1 {
		t.Error("Does not set exact value correctly.")
	}
	if val, _, _ := ip.GetByString("1.2.3.5"); val.(int) != 2 {
		t.Error("Does not set exact value correctly.")
	}
	if err := ip.DeleteByString("1.2.3.4/24"); err != nil {
		t.Error(err)
	}
	if _, found, _ := ip.GetByString("1.2.3.4"); found {
		t.Error("Found deleted value.")
	}
	if val, _, _ := ip.GetByString("1.2.3.5"); val.(int) != 2 {
		t.Error("Does not set exact value correctly.")
	}
}

func TestCovering(t *testing.T) {
	ip := New()
	ip.AddByString("0.0.0.0/0", 1)
	if val, found, _ := ip.GetByString("1.2.3.4"); !found {
		t.Error("Values within covering value not found.")
	} else if val.(int) != 1 {
		t.Error("Value within covering set not correct")
	}
}

func TestMultiple(t *testing.T) {
	ip := New()
	ip.AddByString("0.0.0.0/0", 0)
	ip.AddByString("141.212.120.0/24", 3)
	if val, found, _ := ip.GetByString("1.2.3.4"); !found {
		t.Error("Values within covering value not found.")
	} else if val.(int) != 0 {
		t.Error("Value within covering set not correct")
	}
	if val, _, _ := ip.GetByString("141.212.120.15"); val.(int) != 3 {
		t.Error("Value within subset not correct")
	}
}

func TestFailingSubnet(t *testing.T) {
	ip := New()
	ip.AddByString("115.254.0.0/17", 3)
	ip.AddByString("115.254.0.0/22", 1)
	if val, _, _ := ip.GetByString("115.254.115.198"); val.(int) != 3 {
		t.Error("Value within subset not correct")
	}
	if val, _, _ := ip.GetByString("115.254.0.198"); val.(int) != 1 {
		t.Error("Value within subset not correct")
	}
}
