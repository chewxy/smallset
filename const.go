package smallset

const licence = `/*
The MIT License (MIT)

Copyright (c) 2016 Xuanyi Chew (chewxy [AT] gmail.com)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
`
const basic = `
//{{.Name}}Set is a set of {{.Pointer}}{{.Name}}
type {{.Name}}Set []{{.Pointer}}{{.Name}}

// New{{.Name}}Set creates a new set of {{.Pointer}}{{.Name}}, given an input of any {{.Pointer}}{{.Name}}
func New{{.Name}}Set(a ...{{.Pointer}}{{.Name}}) {{.Name}}Set {
	var set {{.Name}}Set

	for _, v := range a {
		set = set.Add(v)
	}

	return set
}

// ToSlice returns the elements of the current set as a slice
func (set {{.Name}}Set) ToSlice() []{{.Pointer}}{{.Name}} {
	return []{{.Pointer}}{{.Name}}(set)
}



// ContainsALl determines if all the wanted items are already in set
func (set {{.Name}}Set) ContainsAll(ws ...{{.Pointer}}{{.Name}}) bool {
	for _, w := range ws {
		if !set.Contains(w) {
			return false
		}
	}
	return true
}

// Add adds an item into the set, and then returns a new set. If the item already exists, it returns the current set
func (set {{.Name}}Set) Add(item {{.Pointer}}{{.Name}}) {{.Name}}Set {
	if set.Contains(item) {
		return set
	}
	set = append(set, item)
	return set
}

// IsSubSetOf checks if the current set is a subset of the other set.
func (set {{.Name}}Set) IsSubsetOf(other {{.Name}}Set) bool {
	if len(set) > len(other) {
		return false
	}

	for _, v := range set {
		if !other.Contains(v) {
			return false
		}
	}

	return true
}

// IsSupersetOf checks if the current set is a superset of the other set
func (set {{.Name}}Set) IsSupersetOf(other {{.Name}}Set) bool {
	return other.IsSubsetOf(set)
}

// Intersect performs an intersection between two sets - only items that exist in both are returned
func (set {{.Name}}Set) Intersect(other {{.Name}}Set) {{.Name}}Set {
	retVal := make({{.Name}}Set, 0)
	for _, o := range other {
		if set.Contains(o) {
			retVal = append(retVal, o)
		}
	}
	return retVal
}

//Union joins both sets together, keeping only unique items
func (set {{.Name}}Set) Union(other {{.Name}}Set) {{.Name}}Set {
	retVal := make({{.Name}}Set, len(set))
	copy(retVal, set)
	for _, o := range other {
		if !retVal.Contains(o) {
			retVal = append(retVal, o)
		}
	}
	return retVal
}

// Difference returns a new set with items in the current set but not in the other set. 
// Equivalent to  (set - other)
func (set {{.Name}}Set) Difference(other {{.Name}}Set) {{.Name}}Set {
	retVal := make({{.Name}}Set, 0)
	for _, v := range set {
		if !other.Contains(v) {
			retVal = append(retVal, v)
		}
	}
	return retVal
}

// SymmetricDifference is the set of items that is not in each either set.
func (set {{.Name}}Set) SymmetricDifference(other {{.Name}}Set) {{.Name}}Set {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Equals compares two sets and checks if it is the same
func (set {{.Name}}Set) Equals(other {{.Name}}Set) bool {
	if len(set) != len(other) {
		return false
	}

	for _, v := range set {
		if !other.Contains(v) {
			return false
		}
	}

	return true
}

// String for stuff
func (set {{.Name}}Set) String() string {
	var buf bytes.Buffer
	buf.WriteString("{{.Name}}Set[")
	for i, v := range set {
		if i == len(set) - 1{
			fmt.Fprintf(&buf, "%v", v)
		}else{
			fmt.Fprintf(&buf, "%v, ", v)
		}
	}
	buf.WriteString("]")
	return buf.String()
}
`

const defaultEqTempl = `// Contains determines if an item is in the set already
func (set {{.Name}}Set) Contains(w {{.Pointer}}{{.Name}}) bool {
	for _, v := range set {
		if v == w {
			return true
		}
	}
	return false
}
`

const customEqTempl = `// Contains determines if an item is in the set already. This method uses a custom equality method
func (set {{.Name}}Set) Contains(w {{.Pointer}}{{.Name}}) bool {
	for _, v := range set {
		if v.{{.EqFn}}(w) {
			return true
		}
	}
	return false
}
`
