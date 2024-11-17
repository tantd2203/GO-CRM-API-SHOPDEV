package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`

	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"databases"`
}

func main() {
	// Khởi tạo Viper
	viper := viper.New()
	viper.AddConfigPath("./config/") // Đường dẫn tới thư mục chứa file config
	viper.SetConfigName("local")     // Tên file config (không bao gồm phần mở rộng)
	viper.SetConfigType("yaml")      // Loại file config

	// Đọc file config
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read configuration: %w\n", err))
	}

	// Lấy thông tin cổng server từ file config
	serverPort := viper.GetInt("server.port")
	fmt.Printf("Server Port: %d\n", serverPort)

	// Đọc toàn bộ cấu hình vào cấu trúc
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode configuration: %v\n", err)
	}

	// In ra thông tin để kiểm tra
	//	fmt.Printf("Loaded Config: %+v\n", config.Databases)

	for _, db := range config.Databases {
		fmt.Printf(" dabase user : %s,password:, %s", db.User, db.Password)

	}
}
