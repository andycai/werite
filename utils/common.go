package utils

import (
	"math"

	"github.com/andycai/werite/constant"
)

func CalcPagination(count int64) (int, bool) {
	var (
		hasPagination   bool
		totalPagination int
	)
	if count > 0 && (count/int64(constant.NUM_PER_PAGE) > 0) {
		pageDivision := float64(count) / float64(constant.NUM_PER_PAGE)
		totalPagination = int(math.Ceil(pageDivision))
		hasPagination = true
	}

	return totalPagination, hasPagination
}
