package config

import (
	"fmt"
	"time"
)

type Databases struct {
	Mysql *Mysql
	Redis *Redis
}

type Mysql struct {
	Addr     string `mapstructure:"addr" yaml:"addr"`
	Port     string `mapstructure:"addr" yaml:"port"`
	Db       string `mapstructure:"addr" yaml:"db"`
	UserName string `mapstructure:"username" yaml:"username"`
	PassWord string `mapstructure:"password" yaml:"password"`
	Charset  string `mapstructure:"charset" yaml:"charset"`

	ConnMaxIdleTime string `mapstructure:"connMaxIdleTime" yaml:"connMaxIdleTime"`
	ConnMaxLifeTime string `mapstructure:"connMaxLifeTime" yaml:"connMaxLifeTime"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns" yaml:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns" yaml:"maxOpenConns"`
}

func (m *Mysql) GetDsn() string {
	return fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=True&loc=Australia%%2FMelbourne",
		m.UserName, m.PassWord, m.Addr, m.Port, m.Db, m.Charset)
}

func (m *Mysql) GetConnMaxIdleTime() time.Duration {
	t, _ := time.ParseDuration(m.ConnMaxIdleTime)
	return t
}

func (m *Mysql) GetMaxOpenConns() int {
	return m.MaxIdleConns
}

func (m *Mysql) GetConnMaxLifetime() time.Duration {
	t, _ := time.ParseDuration(m.ConnMaxLifeTime)
	return t
}

func (m *Mysql) GetMaxIdleConns() int {
	return m.MaxIdleConns
}

type Redis struct {
	Addr     string `mapstructure:"addr" yaml:"addr"`
	Port     string `mapstructure:"port" yaml:"port"`
	Username string `mapstructure:"username" yaml:"username"`
	Password string `mapstructure:"password" yaml:"password"`
	Db       int    `mapstructure:"db" yaml:"db"`
	PoolSize int    `mapstructure:"poolSize" yaml:"poolSize"`
}

func (r *Redis) GetPoolSize() int {
	return r.PoolSize
}
