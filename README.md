stringset [![Build Status][1]][2]
=========

A simple set implementation for strings in Go.

## Godoc

### PACKAGE

    package stringset
    
    import "github.com/yanatan16/gostringset"

Package stringset implements a very simple Set implementation for
strings with a map[string]bool.

### TYPES

    type Set map[string]bool

Set contains a set of strings.

    func New(elements ...string) Set

Create a new Set.

    func (s Set) Add(elements ...string)

Add elements to a set.

    func (s Set) Any(values ...string) bool

Test whether any of the values are in a set.

    func (s Set) Difference(t Set)

Difference will do an in-place difference of t from s.

    func (s Set) Empty() bool

Test whether a set is empty.

    func (s Set) In(values ...string) bool

Test whether all values are in a set.

    func (s Set) Intersection(t Set)

Intersection will do an in-place intersection of t and s.

    func (s Set) Len() int

Get the length of the set.

    func (s Set) Remove(elements ...string)

Remove elements from a set.

    func (s Set) Union(t Set)

Union will do an in-place union of s and t.

## License

MIT License found in the LICENSE file. 


  [1]: https://travis-ci.org/yanatan16/gostringset.png?branch=master
  [2]: http://travis-ci.org/yanatan16/gostringset