package cache

import (
	"bookstore/common/utils"
	"fmt"
)

const (
	BatchBookListCacheKey = "BatchBookListCacheKey:"
)

func GetBookListCacheKeys(ids []int) string {

	strIds := utils.SliceIntToString(ids)

	return fmt.Sprintf("%s:%s", BatchBookListCacheKey, strIds)
}
