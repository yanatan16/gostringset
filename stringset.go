package gostringset

type Set map[string]bool

func New() Set {
  s := make(map[string]bool)
  return s
}

func (s Set) Add(v string) {
  s[v] = true
}

func (s Set) Remove(v string) {
  delete(s, v)
}

func (s Set) In(v string) bool {
  _, ok := s[v]
  return ok
}

func (s Set) Union(t Set) {
  for v, _ := range t {
    s[v] = true
  }
}

func (s Set) Difference(t Set) {
  for v, _ := range t {
    delete(s, v)
  }
}

func (s Set) Intersection(t Set) {
  for v, _ := range s {
    if !t.In(v) {
      delete(s, v)
    }
  }
}

func (s Set) UnionL(t []string) {
  for _, v := range t {
    s[v] = true
  }
}

func (s Set) DifferenceL(t []string) {
  for _, v := range t {
    delete(s, v)
  }
}

func (s Set) IntersectionL(t []string) {
  for _, v := range s {
    if !t.In(v) {
      delete(s, v)
    }
  }
}