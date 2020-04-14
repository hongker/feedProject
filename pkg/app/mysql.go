package app

import (
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var db *gorm.DB
var redisClient *redis.Client

func init() {
	ConnectDB()
	ConnectRedis()


}

func ConnectDB()  {
	// dsn的格式为 用户名:密码/tcp(主机地址)/数据库名称?charset=字符格式
	dsn := "root:123456@tcp(10.0.75.2:3307)/feed?charset=utf8&parseTime=True&loc=Local"
	conn, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db = conn
}

func ConnectRedis()  {
	redisClient = redis.NewClient(&redis.Options{
		Addr:        "10.0.75.2:6379",
		Password:    "123456",
		PoolSize:    100,
		MaxRetries:  3,
		IdleTimeout: time.Second * 3,
	})

	if err := redisClient.Ping().Err(); err != nil {
		panic(err)
	}
}

func DB() *gorm.DB {
	db.LogMode(true)
	return db
}

func Redis() *redis.Client {
	return redisClient
}
