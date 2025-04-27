package utils

import (
	"sort"
	"strconv"
)

func SliceIntToString(in []int) string {

	sort.Ints(in)

	str := ""
	for _, v := range in {
		str += strconv.Itoa(v)
	}

	return str
}
