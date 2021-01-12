package mc

import (
	"errors"
	"hash/crc32"
	"sort"
	"strings"
)

var ErrNoServer = errors.New("no server available")

type Addrs []string

func (a Addrs) Len() int           { return len(a) }
func (a Addrs) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Addrs) Less(i, j int) bool { return strings.Compare(a[i], a[j]) < 0 }

type Selector struct {
	addrs Addrs
}

func (s *Selector) Set(addrs []string) {
	s.addrs = addrs
	sort.Sort(s.addrs)
}

func (s *Selector) Select(key string) (string, error) {
	n := len(s.addrs)
	if n == 0 {
		return "", ErrNoServer
	}
	if n == 1 {
		addr := s.addrs[0]
		return addr, nil
	}
	cs := crc32.ChecksumIEEE([]byte(key))
	addr := s.addrs[cs%uint32(n)]
	return addr, nil
}
