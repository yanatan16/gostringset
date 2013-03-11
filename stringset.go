// Package stringset implements a very simple Set implementation for strings with a map[string]bool.
package stringset

// Set contains a set of strings.
type Set map[string]bool

// Create a new Set.
func New(elements...string) Set {
  s := make(Set)
  s.Add(elements...)
  return s
}

// Add elements to a set.
func (s Set) Add(elements ...string) {
  for _, v := range elements {
    s[v] = true
  }
}

// Remove elements from a set.
func (s Set) Remove(elements ...string) {
  for _, v := range elements {
    delete(s, v)
  }
}

// Test whether all values are in a set.
func (s Set) In(values ...string) bool {
  for _, v := range values {
    if _, ok := s[v]; !ok {
      return false
    }
  }
  return true
}

// Test whether any of the values are in a set.
func (s Set) Any(values ...string) bool {
  for _, v := range values {
    if _, ok := s[v]; ok {
      return true
    }
  }
  return false
}

// Test whether a set is empty.
func (s Set) Empty() bool {
  return len(s) == 0
}

// Get the length of the set.
func (s Set) Len() int {
  return len(s)
}

// Union will do an in-place union of s and t.
func (s Set) Union(t Set) {
  for v, _ := range t {
    s[v] = true
  }
}

// Difference will do an in-place difference of t from s.
func (s Set) Difference(t Set) {
  for v, _ := range t {
    delete(s, v)
  }
}

// Intersection will do an in-place intersection of t and s.
func (s Set) Intersection(t Set) {
  for v, _ := range s {
    if !t.In(v) {
      delete(s, v)
    }
  }
}