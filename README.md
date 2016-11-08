# smallset
A fast slice-backed set generator for small sets in Go.  Use with [`gen`](https://github.com/clipperhouse/gen)

#Why?#

There exists already fairly decent set implementations (the [default `gen` implementaion](https://github.com/clipperhouse/set) is one). So why something so simple? It turns out after some benchmark and performance analysis, I discovered that I had many small sets, each with fewer than 50 elements in them. A `map` based set would take quite a bit longer than a slice based set.

Here are some quick and dirty benchmarks:

```
BenchmarkThingSet_Contains1-8         	300000000	         4.90 ns/op	       0 B/op	       0 allocs/op
BenchmarkThingSet_Contains10-8        	30000000	        59.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkThingSet_Contains50-8        	 2000000	       980 ns/op	       0 B/op	       0 allocs/op
BenchmarkThingSet_Contains100-8       	  500000	      3483 ns/op	       0 B/op	       0 allocs/op
BenchmarkThingSet_Contains1000-8      	    5000	    280934 ns/op	       4 B/op	       0 allocs/op
BenchmarkThingSet_Contains10000-8     	      50	  27922636 ns/op	    9354 B/op	       0 allocs/op
BenchmarkThingSet_Contains100000-8    	       1	5465854463 ns/op	 5459936 B/op	      48 allocs/op
BenchmarkThangSet_Contains1-8         	200000000	         7.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkThangSet_Contains10-8        	10000000	       148 ns/op	       0 B/op	       0 allocs/op
BenchmarkThangSet_Contains50-8        	 2000000	       819 ns/op	       0 B/op	       0 allocs/op
BenchmarkThangSet_Contains100-8       	 1000000	      1634 ns/op	       0 B/op	       0 allocs/op
BenchmarkThangSet_Contains1000-8      	  100000	     22320 ns/op	       0 B/op	       0 allocs/op
BenchmarkThangSet_Contains10000-8     	    5000	    281176 ns/op	      92 B/op	       0 allocs/op
BenchmarkThangSet_Contains100000-8    	     500	   3463190 ns/op	    8122 B/op	      12 allocs/op

BenchmarkThingSet_Intersect1-8        	100000000	        16.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkThingSet_Intersect10-8       	20000000	        67.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkThingSet_Intersect50-8       	 2000000	       982 ns/op	       0 B/op	       0 allocs/op
BenchmarkThingSet_Intersect100-8      	  500000	      3298 ns/op	       0 B/op	       0 allocs/op
BenchmarkThingSet_Intersect1000-8     	    5000	    279414 ns/op	       7 B/op	       0 allocs/op
BenchmarkThingSet_Intersect10000-8    	      50	  28132173 ns/op	   13301 B/op	       0 allocs/op
BenchmarkThingSet_Intersect100000-8   	       1	6217674098 ns/op	 8160768 B/op	      59 allocs/op
BenchmarkThangSet_Intersect1-8        	10000000	       213 ns/op	      48 B/op	       1 allocs/op
BenchmarkThangSet_Intersect10-8       	 5000000	       363 ns/op	      48 B/op	       1 allocs/op
BenchmarkThangSet_Intersect50-8       	 1000000	      1118 ns/op	      48 B/op	       1 allocs/op
BenchmarkThangSet_Intersect100-8      	 1000000	      2190 ns/op	      48 B/op	       1 allocs/op
BenchmarkThangSet_Intersect1000-8     	  100000	     19856 ns/op	      48 B/op	       1 allocs/op
BenchmarkThangSet_Intersect10000-8    	   10000	    203680 ns/op	     117 B/op	       1 allocs/op
BenchmarkThangSet_Intersect100000-8   	     500	   2343606 ns/op	   12252 B/op	      19 allocs/op
PASS
ok  	github.com/chewxy/smallset/test	62.084s

```

Performed on a i7-2600K CPU @ 3.40GHz with 24GiB of RAM, and `GOMAXPROCS` = 8

#Notes#

This is an important caveat: this set structure is **NOT** performant for sets that have more than 150+ elements. It works best with small sets (<= 100 items).

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