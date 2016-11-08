// Generated by: main
// TypeWriter: set
// Directive: +gen on Thang

package main

// Set is a modification of https://github.com/deckarep/golang-set
// The MIT License (MIT)
// Copyright (c) 2013 Ralph Caraveo (deckarep@gmail.com)

// ThangSet is the primary type that represents a set
type ThangSet map[Thang]struct{}

// NewThangSet creates and returns a reference to an empty set.
func NewThangSet(a ...Thang) ThangSet {
	s := make(ThangSet)
	for _, i := range a {
		s.Add(i)
	}
	return s
}

// ToSlice returns the elements of the current set as a slice
func (set ThangSet) ToSlice() []Thang {
	var s []Thang
	for v := range set {
		s = append(s, v)
	}
	return s
}

// Add adds an item to the current set if it doesn't already exist in the set.
func (set ThangSet) Add(i Thang) bool {
	_, found := set[i]
	set[i] = struct{}{}
	return !found //False if it existed already
}

// Contains determines if a given item is already in the set.
func (set ThangSet) Contains(i Thang) bool {
	_, found := set[i]
	return found
}

// ContainsAll determines if the given items are all in the set
func (set ThangSet) ContainsAll(i ...Thang) bool {
	for _, v := range i {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

// IsSubset determines if every item in the other set is in this set.
func (set ThangSet) IsSubset(other ThangSet) bool {
	for elem := range set {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

// IsSuperset determines if every item of this set is in the other set.
func (set ThangSet) IsSuperset(other ThangSet) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set ThangSet) Union(other ThangSet) ThangSet {
	unionedSet := NewThangSet()

	for elem := range set {
		unionedSet.Add(elem)
	}
	for elem := range other {
		unionedSet.Add(elem)
	}
	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set ThangSet) Intersect(other ThangSet) ThangSet {
	intersection := NewThangSet()
	// loop over smaller set
	if set.Cardinality() < other.Cardinality() {
		for elem := range set {
			if other.Contains(elem) {
				intersection.Add(elem)
			}
		}
	} else {
		for elem := range other {
			if set.Contains(elem) {
				intersection.Add(elem)
			}
		}
	}
	return intersection
}

// Difference returns a new set with items in the current set but not in the other set
func (set ThangSet) Difference(other ThangSet) ThangSet {
	differencedSet := NewThangSet()
	for elem := range set {
		if !other.Contains(elem) {
			differencedSet.Add(elem)
		}
	}
	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set ThangSet) SymmetricDifference(other ThangSet) ThangSet {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Clear clears the entire set to be the empty set.
func (set *ThangSet) Clear() {
	*set = make(ThangSet)
}

// Remove allows the removal of a single item in the set.
func (set ThangSet) Remove(i Thang) {
	delete(set, i)
}

// Cardinality returns how many items are currently in the set.
func (set ThangSet) Cardinality() int {
	return len(set)
}

// Iter returns a channel of type Thang that you can range over.
func (set ThangSet) Iter() <-chan Thang {
	ch := make(chan Thang)
	go func() {
		for elem := range set {
			ch <- elem
		}
		close(ch)
	}()

	return ch
}

// Equal determines if two sets are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (set ThangSet) Equal(other ThangSet) bool {
	if set.Cardinality() != other.Cardinality() {
		return false
	}
	for elem := range set {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

// Clone returns a clone of the set.
// Does NOT clone the underlying elements.
func (set ThangSet) Clone() ThangSet {
	clonedSet := NewThangSet()
	for elem := range set {
		clonedSet.Add(elem)
	}
	return clonedSet
}