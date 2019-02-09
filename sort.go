package p_sort

import (
	"errors"
	"sort"
)

type SortClass struct {
}

var Sort = SortClass{}

const (
	SortClass_ASC  int = 0
	SortClass_DESC int = 1
)

func (this *SortClass) SortIntSlice(ints []int, order int) []int {
	if order == SortClass_ASC {
		sort.Ints(ints)
	} else if order == SortClass_DESC {
		sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	} else {
		panic(errors.New(`order params illigal`))
	}
	return ints
}

func (this *SortClass) SortFloat64Slice(floats []float64, order int) []float64 {
	if order == SortClass_ASC {
		sort.Float64s(floats)
	} else if order == SortClass_DESC {
		sort.Sort(sort.Reverse(sort.Float64Slice(floats)))
	} else {
		panic(errors.New(`order params illigal`))
	}
	return floats
}

func (this *SortClass) SortStringSlice(strings []string, order int) []string {
	if order == SortClass_ASC {
		sort.Strings(strings)
	} else if order == SortClass_DESC {
		sort.Sort(sort.Reverse(sort.StringSlice(strings)))
	} else {
		panic(errors.New(`order params illigal`))
	}
	return strings
}
