package blacklist

import (
	"bufio"
	"os"
	"strings"

	"github.com/zmap/iptree/iptree"
)

type Blacklist struct {
	T *iptree.IPTree
}

func New() *Blacklist {
	t := new(Blacklist)
	t.T = iptree.New()
	t.T.AddByString("0.0.0.0/0", 0)
	return t
}

func (i *IPTree) AddEntry(cidr string) error {
	return i.T.AddByString(cidr, 1)
}

func (i *IPTree) ParseFromFile(path string, v int) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		i.AddEntry(words[0])
	}
	if err := scanner.Err(); err != nil {
		return err
	}
}

func (i *IPTree) IsBlacklisted(ip string) (bool, error) {
	r, _, err := i.GetByString(ip)
	if err != nil {
		return false, err
	}
	if r == 0 {
		return false
	} else {
		return true
	}
}
