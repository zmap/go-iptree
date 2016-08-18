package blacklist

import (
	"bufio"
	"os"
	"strings"

	"github.com/zmap/go-iptree/iptree"
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

func (b *Blacklist) AddEntry(cidr string) error {
	return b.T.AddByString(cidr, 1)
}

func (b *Blacklist) ParseFromFile(path string, v int) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		b.AddEntry(words[0])
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func (b *Blacklist) IsBlacklisted(ip string) (bool, error) {
	r, _, err := b.T.GetByString(ip)
	if err != nil {
		return false, err
	}
	if r == 0 {
		return false, nil
	} else {
		return true, nil
	}
}
