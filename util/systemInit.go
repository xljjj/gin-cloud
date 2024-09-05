package util

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config app:", viper.Get("app"))
	JWTKey = []byte(viper.GetString("app.JWTKey")) // 设置jwt密钥
	fmt.Println("config mysql:", viper.Get("mysql"))
	fmt.Println("config redis:", viper.Get("redis"))
}
