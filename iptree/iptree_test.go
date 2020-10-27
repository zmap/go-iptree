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

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

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

// randNumber returns a random integer between 1 and maxIP
func randNumber(maxIP int) int {
	return rand.Intn(maxIP-1) + 1
}

// benchmarkLookup is a crude iptree benchmark tool
// 1. Generates an iptree of IPs from 1.1.1.1 -> maxIP.maxIP.maxIP.maxIP
// 2. Generates a big list of random IPs to lookup.  Range of each octet is 1-maxIP
// 3. Resets the benchmark timer, so the setup step 1-2 don't count in the benchmark
// 4. Performs the lookups
// TODO - This is a pretty crude test.  Probably should use maps.
func benchmarkLookup(debugLevel int, maxIP int, lookups int, b *testing.B) {

	if debugLevel > 100 {
		fmt.Println("benchmarkLookup\tmaxIP:", maxIP, "\tlookups:", lookups)
	}

	var startMap = map[string]int{
		"a": 1,
		"b": 1,
		"c": 1,
		"d": 1,
	}

	var endMap = map[string]int{
		"a": maxIP,
		"b": maxIP,
		"c": maxIP,
		"d": maxIP,
	}

	var i int = 0
	ip := New()
	for a := startMap["a"]; a < endMap["a"]; a++ {
		for b := startMap["b"]; b < endMap["b"]; b++ {
			for c := startMap["c"]; c < endMap["c"]; c++ {
				for d := startMap["d"]; d < endMap["d"]; d++ {
					ipString := fmt.Sprintf("%d.%d.%d.%d/32", a, b, c, d)
					ipLabel := fmt.Sprintf("a=%d.b=%d.c=%d.d=%d", a, b, c, d)
					ip.AddByString(ipString, ipLabel)
					if debugLevel > 1000 {
						fmt.Println("i:", i, "\tipString:", ipString, "\tipLabel:", ipLabel)
					}
					i++
				}
			}
		}
	}
	if debugLevel > 10 {
		fmt.Println("ip trie loaded with:", i)
	}

	var lookupIPs = make([]string, lookups)
	var lookupLabels = make([]string, lookups)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < lookups; i++ {
		a := randNumber(maxIP)
		b := randNumber(maxIP)
		c := randNumber(maxIP)
		d := randNumber(maxIP)
		ipString := fmt.Sprintf("%d.%d.%d.%d", a, b, c, d)
		ipLabel := fmt.Sprintf("a=%d.b=%d.c=%d.d=%d", a, b, c, d)
		lookupIPs[i] = ipString
		lookupLabels[i] = ipLabel
	}

	b.ResetTimer()

	for i := 0; i < lookups; i++ {
		if val, _, _ := ip.GetByString(lookupIPs[i]); val != lookupLabels[i] {
			b.Error(i, "Test failed\texpected:", lookupLabels[i], "\tgot:", val)
		}
	}

}

// BenchmarkLookupA 1,000 lookups
func BenchmarkLookupA(b *testing.B) {
	benchmarkLookup(111, 30, 1e+3, b)
}

// BenchmarkLookupA 1,000,000 lookups
func BenchmarkLookupB(b *testing.B) {
	benchmarkLookup(111, 30, 1e+6, b)
}

// BenchmarkLookupA 10,000,000 lookups
func BenchmarkLookupC(b *testing.B) {
	benchmarkLookup(111, 30, 1e+7, b)
}
