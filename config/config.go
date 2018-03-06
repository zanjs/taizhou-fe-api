package config

import (
	"fmt"

	"github.com/jinzhu/configor"
)

// Config is
var Config = struct {
	APP struct {
		Name string
		Port string `default:"8080"`
	}
	Redis struct {
		Host     string
		Port     string `default:"6379"`
		Password string `default:"root"`
		Expire   string `default:"120"`
	}
	DB struct {
		Driver   string
		Host     string
		Port     string `default:"3306"`
		Name     string
		User     string `default:"root"`
		Password string `required:"true" env:"DBPassword"`
	}
	JWT struct {
		Secret      string
		AdminSecret string
	}
	Upload struct {
		Path string `default:"uploads/m"`
		Ext  string `default:".jpg|.jpeg|.png|.gif"`
	}
	Media struct {
		Host string
		Path string
	}
}{}

func init() {
	configor.Load(&Config, "app.yml")
	fmt.Printf("config port : %#v", Config)
}
