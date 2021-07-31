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
	"net"

	"github.com/asergeyev/nradix"
)

type IPTree struct {
	R *nradix.Tree
}

func ipToUint(ip net.IP) uint32 {
	return uint32(ip[0])<<24 | uint32(ip[1])<<16 | uint32(ip[2])<<8 | uint32(ip[3])
}

func New() *IPTree {
	t := new(IPTree)
	t.R = nradix.NewTree(0)
	return t
}

func (i *IPTree) Add(cidr *net.IPNet, v interface{}) error {
	return i.R.AddCIDR(cidr.String(), v)
}

func (i *IPTree) AddByString(ipcidr string, v interface{}) error {
	return i.R.AddCIDR(ipcidr, v)
}

func (i *IPTree) Get(ip net.IP) (interface{}, bool, error) {
	v, err := i.R.FindCIDR(ip.String())
	if v != nil {
		return v, true, err
	} else {
		return v, false, err
	}
}

func (i *IPTree) GetByString(ipstr string) (interface{}, bool, error) {
	v, err := i.R.FindCIDR(ipstr)
	if v != nil {
		return v, true, err
	} else {
		return v, false, err
	}
}

func (i *IPTree) DeleteByString(ipstr string) error {
	return i.R.DeleteCIDR(ipstr)
}
