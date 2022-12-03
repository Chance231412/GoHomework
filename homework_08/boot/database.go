package boot

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"homework_08/app/global"
	g "homework_08/app/global"
	"time"
)

func MysqlSetup() {
	config := global.Config.Databases.Mysql

	dsn := config.GetDsn()
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		g.Logger.Fatal("Initialize database failed.", zap.Error(err))
	}

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, _ := db.DB()

	//SetConnMaxIdleTime 设置最大空闲时间
	//释放空闲时间超过最大空闲时间的数据库连接,以避免因为没有释放数据库连接而引起的数据库连接遗漏
	sqlDB.SetConnMaxIdleTime(config.GetConnMaxIdleTime())

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(config.GetMaxOpenConns())

	//SetMaxIdleConns 用于设置连接池中空闲连接的最大数量
	sqlDB.SetMaxIdleConns(config.GetMaxIdleConns())

	//判断数据库连接是否成功
	err = sqlDB.Ping()
	if err != nil {
		g.Logger.Fatal("connect to mysql db failed.", zap.Error(err))
	}

	g.MysqlDB = db

	g.Logger.Info("initialize mysql successfully!")
}

func RedisSetup() {
	config := g.Config.Databases.Redis

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Addr, config.Port),
		Username: config.Username,
		Password: config.Password,
		DB:       config.Db,
		PoolSize: config.PoolSize,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		g.Logger.Fatal("connect to redis instance failed.", zap.Error(err))
	}

	g.Rdb = rdb

	g.Logger.Info("initialize redis client successfully!")
}
