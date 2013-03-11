package stringset

import (
	"testing"
)

type mytester struct {
	*testing.T
}

func (t mytester) in(s Set, v string) {
	if !s.In(v) {
		t.Error(v + " is In the set!")
	}
}

func (t mytester) notin(s Set, v string) {
	if s.In(v) {
		t.Error(v + " isn't In the set!")
	}
}

func TestStringSet(T *testing.T) {
	t := mytester{T}

	s := New("abc", "123", "howdy")

	t.notin(s, "nothere")
	t.in(s, "abc")

	s.Add("def", "456")

	t.in(s, "def")
	t.in(s, "123")

	s.Remove("def")

	t.notin(s, "def")

	if !s.In("abc", "123", "howdy") {
		t.Error("All In didn't work")
	}

	if s.Any("abcas", "123adf", "howdydadf") {
		t.Error("Any In (fail) didn't work")
	}

	if !s.Any("abc", "123adf", "howdydadf") {
		t.Error("Any In (fail) didn't work")
	}

	if s.Empty() {
		t.Error("Set reports empty when its not!")
	}

	if !New().Empty() {
		t.Error("Empty new Set doesn't report empty.")
	}

	if s.Len() != 4 {
		t.Error("Length of set isn't right.")
	}

	s.Add("abc")
	if s.Len() != 4 {
		t.Error("Length changed after noop add.")
	}

	s.Union(New("abc", "def"))

	t.in(s, "def")

	s.Difference(New("adfa", "123"))
	t.notin(s, "123")
	t.notin(s, "adfa")

	r := New("abc", "nothere")
	s.Intersection(r)

	t.notin(s, "123")
	t.in(r, "nothere")

}
