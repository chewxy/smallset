# smallset
A fast slice-backed set generator for small sets in Go.  Use with [`gen`](https://github.com/clipperhouse/gen)

#Why?#

There exists already fairly decent set implementations (the [default `gen` implementaion](https://github.com/clipperhouse/set) is one). So why something so simple? It turns out after some benchmark and performance analysis, I discovered that I had many small sets, each with fewer than 100 elements in them. A `map` based set would take quite a bit longer than a slice based set.

Here are some quick and dirty benchmarks:

```
PASS
BenchmarkSmallSet1-8    	200000000	         8.31 ns/op
BenchmarkSmallSet5-8    	30000000	        40.9 ns/op
BenchmarkSmallSet10-8   	10000000	       114 ns/op
BenchmarkSmallSet50-8   	 1000000	      1649 ns/op
BenchmarkSmallSet100-8  	  300000	      4839 ns/op
BenchmarkSmallSet500-8  	   20000	     79464 ns/op
BenchmarkSmallSet1000-8 	    5000	    293561 ns/op
BenchmarkSmallSet5000-8 	     200	   7172039 ns/op
BenchmarkSmallSet10000-8	      50	  28184651 ns/op
BenchmarkSet1-8         	30000000	        35.2 ns/op
BenchmarkSet5-8         	10000000	       197 ns/op
BenchmarkSet10-8        	 3000000	       480 ns/op
BenchmarkSet50-8        	  500000	      2579 ns/op
BenchmarkSet100-8       	  300000	      5202 ns/op
BenchmarkSet500-8       	   50000	     27204 ns/op
BenchmarkSet1000-8      	   30000	     55601 ns/op
BenchmarkSet5000-8      	    5000	    292904 ns/op
BenchmarkSet10000-8     	    2000	    590174 ns/op
```

Performed on a i7-2600K CPU @ 3.40GHz with 24GiB of RAM, and `GOMAXPROCS` = 8

#Notes#

This is an important caveat: this set structure is **NOT** performant for sets that have more than 150+ elements. It works best with small sets.

##API Differences ##
The API for smallsets are not the same as the API for the [default `gen` implementaion](https://github.com/clipperhouse/set) of sets. The reason is since this set is slice backed, I thought it'd be wise to reuse the slice semantics. Consider the following:

```Go
	ss := NewThingSet([]int{1,2,3,4}...) // slice based set, this implementation
	ms := NewThingMapSet([]{1,2,3,4}...) // map based set, which is the common one

	// this is the difference in adding an item to the set
	ss = ss.Add(5) // note that Add() returns a set, the same way append() does it
	ms.Add(5) 
````

* `Remove()` is not implemented. For the current work I do, I didn't need to remove things from sets. I may come back and add this in the future, or send a pull request
* `Equals()` is used instead of `Equal()`. Grammar is important.

##Hacky TagValues##
The `TagValue` is hacked to allow custom equality methods. See To Use below.

#To Use#

The gen annotation syntax is rather standard: 

```// +gen [*] tag:"Value, Value[T,T]" anothertag```

Here's how to use `smallset` with gen:

```go

// +gen smallset
type MyStruct struct{}
```

If you have a very complicated struct that requires custom equality methods, this can be achieved (albeit rather hackily) with this:

```go
// +gen smallset:"eq_Equals"
type MyStruct struct{}
```

This will generate this code in the `Contains()` method:

```go
func (set MyStructSet) Contains(want MyStruct) bool {
	for _, v := range set {
		if v.Equals(want){
			return true
		}
	}
	return false
}
```

The `eq_` is a reserved TagValue that allows you to pass in your Equality method. Anything after the first underscore is considered to be the method name. Here's one more example to make it clearer:

If you do this (note the `EQ`):

```go
// +gen smallset:"eq_EQ"
type MyStruct struct{}
```

This code will be generated (note that `.EQ(want)` is now the method):

```go
func (set MyStructSet) Contains(want MyStruct) bool {
	for _, v := range set {
		if v.EQ(want){
			return true
		}
	}
	return false
}
```

If no values are passed in, the default equality comparator (`==`) is used, generating code that looks like this:
```go
func (set MyStructSet) Contains(want MyStruct) bool {
	for _, v := range set {
		if v == want {
			return true
		}
	}
	return false
}
```


Now obviously this is quite janky and hacky. But it works for now. Feel free to send a pull request if you want to fix this