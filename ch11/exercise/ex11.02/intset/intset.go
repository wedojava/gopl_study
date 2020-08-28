package intset

import (
	"bytes"
	"fmt"
)

// An BitIntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type BitIntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *BitIntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *BitIntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) { // ?
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit // ?
}

func (s *BitIntSet) AddAll(nums ...int) {
	for _, n := range nums {
		s.Add(n)
	}
}

// UnionWith sets s to the union of s and t
func (s *BitIntSet) UnionWith(t IntSet) {
	if bis, ok := t.(*BitIntSet); ok {
		for i, tword := range bis.words {
			if i < len(s.words) { // ?
				s.words[i] |= tword // ?
			} else {
				s.words = append(s.words, tword)
			}
		}
	} else {
		for _, i := range t.Ints() {
			s.Add(i)
		}
	}
}

func popcount(x uint64) int {
	count := 0
	for x != 0 {
		count++
		x &= x - 1 // ?
	}
	return count
}

func (s *BitIntSet) Len() int {
	count := 0
	for _, word := range s.words {
		count += popcount(word)
	}
	return count
}

func (s *BitIntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	s.words[word] &^= 1 << bit // AND NOT
	// s.words[word]^(s.words[word] & 1<<bit) // AND first then XOR
	// The &^ operator is bit clear (AND NOT):
	// z = x &^ y, each bit of z is 0 if the corresponding bit of y is 1; otherwise it equals the corresponding bit of x.
	// x=00100010
	// y=00000110
	// z=x&y=00000010
	// x^z=00100010^00000010=00100000
}

func (s *BitIntSet) Clear() {
	for i := range s.words {
		s.words[i] = 0
	}
}

func (s *BitIntSet) Copy() IntSet {
	new := &BitIntSet{}
	new.words = make([]uint64, len(s.words))
	copy(new.words, s.words)
	return new
}

// String returns the set as a string of the form "{1 2 3}"
func (s *BitIntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *BitIntSet) Ints() []int {
	var ints []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			// 左移j位和word与运算，是1说明这个位置上word的值是1
			if word&(1<<uint(j)) != 0 {
				// 64*i+j: 64*1+3 = 67: 67/64 = 1 余 3
				ints = append(ints, 64*i+j)
			}
		}
	}
	return ints
}
