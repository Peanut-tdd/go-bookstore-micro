package svc

import (
	"bookstore/model/source"
	"bookstore/rpc/add/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	BookModel   source.BookModel
	RedisClient *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.BookDataBase.DataSource)
	redisConn, _ := redis.NewRedis(c.Redis)
	return &ServiceContext{
		Config:      c,
		BookModel:   source.NewBookModel(sqlConn),
		RedisClient: redisConn,
	}
}
