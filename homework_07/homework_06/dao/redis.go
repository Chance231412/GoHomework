package dao

import (
	"fmt"
	"github.com/go-redis/redis"
	"homework06/model"
	"log"
)

var Rdb *redis.Client

//ConnRedis 链接redis数据库
func ConnRedis()error {
	Rdb = redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	_, err := Rdb.Ping().Result()
	if err != nil {
		log.Print(err)
		return err
	}
	fmt.Println("redis 链接成功")
	return nil
}

//SelectUserByName 在缓存中查找用户,并获取用户信息
func SelectUserByName(name string,user *model.User)bool {
	data, err := Rdb.HGetAll(name).Result()
	if err!=nil || len(data)==0 {
		//查询mysql
		fmt.Println(err)
		ok := MysqlSelectUserByName(name,user)
		if ok {
			//如果mysql中存在则添加到缓存中
			// 初始化hash数据的多个字段值
			batchData := make(map[string]interface{})
			batchData["name"] = name
			batchData["password"] = user.Password
			batchData["phone"] = user.Phone
			//添加用户信息到缓存
			err = Rdb.HMSet(name, batchData).Err()
			if err!=nil {
				fmt.Println(err)
				return false
			}
			return true
		}
		return false
	}
	user.Name=data["name"]
	user.Phone=data["phone"]
	user.Password=data["password"]
	return true
}

func UpdateData(key,filed,value string)error {
	err := Rdb.HSet(key,filed,value).Err()
	return err
}