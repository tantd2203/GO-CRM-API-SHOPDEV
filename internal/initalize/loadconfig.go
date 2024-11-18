package initalize

import (
	"GO-CRM-API-SHOPDEV/global"
	"fmt"
	v "github.com/spf13/viper"
)

func LoadConfig() {
	// Khởi tạo Viper
	viper := v.New()
	viper.AddConfigPath("./config/") // Đường dẫn tới thư mục chứa file config
	viper.SetConfigName("local")     // Tên file config (không bao gồm phần mở rộng)
	viper.SetConfigType("yaml")      // Loại file config

	// Đọc file config
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read configuration: %w\n", err))
	}
	// Đọc toàn bộ cấu hình vào cấu trúc
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode configuration: %v\n", err)
	}

}
