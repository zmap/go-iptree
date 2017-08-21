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

func (b *Blacklist) ParseFromFile(path string) error {
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
	if r.(int) == 0 {
		return false, nil
	} else {
		return true, nil
	}
}
