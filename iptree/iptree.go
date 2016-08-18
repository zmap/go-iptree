/*
 * IPTree Copyright 2016 Regents of the University of Michigan
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
	"errors"
	"net"

	"github.com/miekg/bitradix"
)

type IPTree struct {
	R *bitradix.Radix32
}

func ipToUint(ip net.IP) uint32 {
	return uint32(ip[0])<<24 | uint32(ip[1])<<16 | uint32(ip[2])<<8 | uint32(ip[3])
}

func New() *IPTree {
	t := new(IPTree)
	t.R = bitradix.New32()
	return t
}

func (i *IPTree) Add(cidr *net.IPNet, v int) error {
	size, _ := cidr.Mask.Size()
	i.R.Insert(ipToUint(cidr.IP.To4()), size, v)
	return nil
}

func (i *IPTree) AddByString(ipcidr string, v int) error {
	_, ipnet, err := net.ParseCIDR(ipcidr)
	if err != nil {
		return errors.New("invalid CIDR block")
	}
	return i.Add(ipnet, v)
}

func (i *IPTree) Get(ip net.IP) (int, bool, error) {
	val := i.R.Find(ipToUint(ip.To4()), 32)
	if val == nil {
		return 0, false, nil
	}
	return val.Value.(int), true, nil
}

func (i *IPTree) GetByString(ipstr string) (int, bool, error) {
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return 0, false, errors.New("invalid IP address")
	}
	return i.Get(ip)
}
