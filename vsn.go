package vsn

import (
	"sort"
	"strconv"
	"strings"
)

type Version string

func (v Version) String() string {
	return string(v)
}

func (v Version) split() []string {
	return strings.Split(v.String(), ".")
}

func (v Version) index(i int) string {
	s := v.split()
	if (len(s) - 1) < i {
		return ""
	}
	return s[i]
}

func (v Version) indexI(i int) int64 {
	s := v.index(i)
	if s == "" {
		return 0
	}

	i6, _ := strconv.ParseInt(s, 0, 64)
	return i6
}

func (v Version) Len() int {
	return len(v.split())
}

func (v Version) Major() int64 {
	return v.indexI(0)
}

func (v Version) Minor() int64 {
	return v.indexI(1)
}

func (v Version) Release() int64 {
	return v.indexI(2)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	if y > x {
		return y
	}
	return x
}

func (v Version) Cmp(v2 Version) int {
	s := max(v.Len(), v2.Len())
	for i := 0; i < s; i++ {
		l1 := v.indexI(i)
		l2 := v2.indexI(i)
		if l1 == l2 {
			continue
		}

		if l2 < l1 {
			return -1
		}

		if l2 > l1 {
			return 1
		}
	}

	return 0
}

type Sorter []Version

func (s Sorter) Len() int {
	return len(s)
}

func (s Sorter) Less(i, j int) bool {
	y := s[i].Cmp(s[j]) == -1
	return y
}

func (s Sorter) Swap(i, j int) {
	vi := s[i]
	vj := s[j]

	s[i] = vj
	s[j] = vi
}

func NewestVersion(input []Version) Version {
	sv := SortVersions(input)
	if len(sv) == 0 {
		return ""
	}

	return sv[0]
}

func SortVersions(input []Version) []Version {
	srt := Sorter(input)
	sort.Sort(srt)
	return []Version(srt)
}
