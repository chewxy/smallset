package main

import (
	"math/rand"
	"testing"
)

func newThings(cnt int) []Thing {
	retVal := make([]Thing, cnt)
	for i := 0; i < cnt; i++ {
		retVal[i] = Thing(rand.Int())
	}
	return retVal
}

func newThangs(cnt int) []Thang {
	retVal := make([]Thang, cnt)
	for i := 0; i < cnt; i++ {
		retVal[i] = Thang(rand.Int())
	}
	return retVal
}

func benchThingContains(i int, b *testing.B) {
	l := newThings(i)
	s := NewThingSet(l...)

	for n := 0; n < b.N; n++ {
		for _, v := range l {
			s.Contains(v)
		}
	}
}

func benchThangContains(i int, b *testing.B) {
	l := newThangs(i)
	s := NewThangSet(l...)

	for n := 0; n < b.N; n++ {
		for _, v := range l {
			s.Contains(v)
		}
	}
}

func BenchmarkThingSet_Contains1(b *testing.B)      { benchThingContains(1, b) }
func BenchmarkThingSet_Contains10(b *testing.B)     { benchThingContains(10, b) }
func BenchmarkThingSet_Contains50(b *testing.B)     { benchThingContains(50, b) }
func BenchmarkThingSet_Contains100(b *testing.B)    { benchThingContains(100, b) }
func BenchmarkThingSet_Contains1000(b *testing.B)   { benchThingContains(1000, b) }
func BenchmarkThingSet_Contains10000(b *testing.B)  { benchThingContains(10000, b) }
func BenchmarkThingSet_Contains100000(b *testing.B) { benchThingContains(100000, b) }

func BenchmarkThangSet_Contains1(b *testing.B)      { benchThangContains(1, b) }
func BenchmarkThangSet_Contains10(b *testing.B)     { benchThangContains(10, b) }
func BenchmarkThangSet_Contains50(b *testing.B)     { benchThangContains(50, b) }
func BenchmarkThangSet_Contains100(b *testing.B)    { benchThangContains(100, b) }
func BenchmarkThangSet_Contains1000(b *testing.B)   { benchThangContains(1000, b) }
func BenchmarkThangSet_Contains10000(b *testing.B)  { benchThangContains(10000, b) }
func BenchmarkThangSet_Contains100000(b *testing.B) { benchThangContains(100000, b) }

func benchThingIntersect(i, j int, b *testing.B) {
	l1 := newThings(i)
	l2 := newThings(j)

	s1 := NewThingSet(l1...)
	s2 := NewThingSet(l2...)

	for n := 0; n < b.N; n++ {
		s1.Intersect(s2)
	}
}

func benchThangIntersect(i, j int, b *testing.B) {
	l1 := newThangs(i)
	l2 := newThangs(j)

	s1 := NewThangSet(l1...)
	s2 := NewThangSet(l2...)

	for n := 0; n < b.N; n++ {
		s1.Intersect(s2)
	}
}

func BenchmarkThingSet_Intersect1(b *testing.B)      { benchThingIntersect(1, 1, b) }
func BenchmarkThingSet_Intersect10(b *testing.B)     { benchThingIntersect(10, 5, b) }
func BenchmarkThingSet_Intersect50(b *testing.B)     { benchThingIntersect(50, 25, b) }
func BenchmarkThingSet_Intersect100(b *testing.B)    { benchThingIntersect(100, 50, b) }
func BenchmarkThingSet_Intersect1000(b *testing.B)   { benchThingIntersect(1000, 500, b) }
func BenchmarkThingSet_Intersect10000(b *testing.B)  { benchThingIntersect(10000, 5000, b) }
func BenchmarkThingSet_Intersect100000(b *testing.B) { benchThingIntersect(100000, 50000, b) }

func BenchmarkThangSet_Intersect1(b *testing.B)      { benchThangIntersect(1, 1, b) }
func BenchmarkThangSet_Intersect10(b *testing.B)     { benchThangIntersect(10, 5, b) }
func BenchmarkThangSet_Intersect50(b *testing.B)     { benchThangIntersect(50, 25, b) }
func BenchmarkThangSet_Intersect100(b *testing.B)    { benchThangIntersect(100, 50, b) }
func BenchmarkThangSet_Intersect1000(b *testing.B)   { benchThangIntersect(1000, 500, b) }
func BenchmarkThangSet_Intersect10000(b *testing.B)  { benchThangIntersect(10000, 5000, b) }
func BenchmarkThangSet_Intersect100000(b *testing.B) { benchThangIntersect(100000, 50000, b) }
