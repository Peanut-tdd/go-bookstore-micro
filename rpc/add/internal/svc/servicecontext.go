package svc

import (
	"bookstore/model/book"
	"bookstore/rpc/add/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	BookModel   book.BookModel
	RedisClient *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.BookDataBase.DataSource)
	redisConn, _ := redis.NewRedis(c.RedisConfig)
	return &ServiceContext{
		Config:      c,
		BookModel:   book.NewBookModel(sqlConn),
		RedisClient: redisConn,
	}
}
