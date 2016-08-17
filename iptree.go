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
	return nil
}

func (i *IPTree) AddByString(ipcidr string, v int) error {
	_, ipnet, err := net.ParseCIDR(ipcidr)
	if err != nil {
		return errors.New("invalid CIDR block")
	}
	return i.Add(ipnet, v)
}

func (i *IPTree) Get(ip net.IP) (int, error) {
	val := i.R.Find(ipToUint(ip), 32).Value
	return val.(int), nil
}

func (i *IPTree) GetByString(ipstr string) (int, error) {
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return 0, errors.New("invalid IP address")
	}
	return i.Get(ip)
}
