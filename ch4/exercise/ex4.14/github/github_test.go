package github

import (
	"testing"
)

func TestMilestoneEquals(t *testing.T) {
	var tests = []struct {
		m1   Milestone
		m2   *Milestone
		want bool
	}{
		{Milestone{}, &Milestone{}, true},
		{Milestone{}, &Milestone{ID: 42}, false},
		{Milestone{ID: 42}, &Milestone{ID: 42}, true},
		{Milestone{ID: 42}, &Milestone{ID: 43}, false},
	}

	for _, test := range tests {
		got := test.m1.Equals(test.m2)
		if got != test.want {
			t.Errorf("(%q).Equals(%q) = %t, want %t", test.m1, test.m2, got, test.want)
		}
	}
}

func TestUserEquals(t *testing.T) {
	var tests = []struct {
		u1   User
		u2   *User
		want bool
	}{
		{User{}, &User{}, true},
		{User{}, &User{ID: 42}, false},
		{User{ID: 42}, &User{ID: 42}, true},
		{User{ID: 42}, &User{ID: 43}, false},
	}

	for _, test := range tests {
		got := test.u1.Equals(test.u2)
		if got != test.want {
			t.Errorf("(%q).Equals(%q) = %t, want %t", test.u1, test.u2, got, test.want)
		}
	}
}
