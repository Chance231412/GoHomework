package config

type Logger struct {
	LogLevel   string `mapstucture:"logLevel" yaml:"logLevel"`
	SavePath   string `mapstucture:"savePath" yaml:"savePath"`
	MaxSize    int    `mapstucture:"maxSize" yaml:"maxSize"`
	MaxBackups int    `mapstucture:"maxBackups" yaml:"maxBackups"`
	MaxAge     int    `mapstucture:"maxAge" yaml:"maxAge"`
	IsCompress bool   `mapstucture:"isCompress" yaml:"isCompress"`
}
