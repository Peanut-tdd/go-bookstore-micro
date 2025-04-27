package book

import (
	"bookstore/common/cache"
	"bookstore/model/source"
	"bookstore/rpc/add/internal/svc"
	"context"
	"encoding/json"
)

func GetBatchBookListCache(ctx context.Context, svcCtx svc.ServiceContext, bookIds []int) (resp []source.Book, err error) {

	cacheKey := cache.GetBookListCacheKeys(bookIds)

	jsonData, err := svcCtx.RedisClient.GetCtx(ctx, cacheKey)
	if err != nil {

	}

	if jsonData != "" {
		err = json.Unmarshal([]byte(jsonData), &resp)
	}

	return

}

func SetBatchBookListCache(ctx context.Context, svcCtx svc.ServiceContext, bookIds []int, data []source.Book) (err error) {

	cacheKeys := cache.GetBookListCacheKeys(bookIds)

	jsonData, err := json.Marshal(data)
	if err != nil {

		return
	}

	err = svcCtx.RedisClient.SetCtx(ctx, cacheKeys, string(jsonData))

	if err != nil {

	}

	return
}
