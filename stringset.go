// Package stringset implements a very simple Set implementation for strings with a map[string]bool.
package stringset

// Set contains a set of strings.
type Set map[string]bool

// Create a new Set.
func New() Set {
  s := make(map[string]bool)
  return s
}

// Add a value to a set.
func (s Set) Add(v string) {
  s[v] = true
}

// Remove a value from a set.
func (s Set) Remove(v string) {
  delete(s, v)
}

// Test whether a value is in a set.
func (s Set) In(v string) bool {
  _, ok := s[v]
  return ok
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

// UnionL will do an in-place union of s and t, where t is a slice of strings.
func (s Set) UnionL(t []string) {
  for _, v := range t {
    s[v] = true
  }
}

// DifferenceL will do an in-place difference of t from s, where t is a slice of strings.
func (s Set) DifferenceL(t []string) {
  for _, v := range t {
    delete(s, v)
  }
}

// IntersectionL will do an in-place intersection of t and s, where t is a slice of strings.
func (s Set) IntersectionL(t []string) {
  for _, v := range s {
    if !t.In(v) {
      delete(s, v)
    }
  }
}