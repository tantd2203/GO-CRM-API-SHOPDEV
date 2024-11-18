package initalize

import (
	"GO-CRM-API-SHOPDEV/global"
	"fmt"
	"go.uber.org/zap"
)

func Run() {

	// Flow LoadConfig -> Init  Logger -> load database
	LoadConfig()
	fmt.Println("Load configuration mysql : ", global.Config.Mysql.Dbname)
	InitLogger()

	global.Logger.Info("Config Log ok !", zap.String("ok", "success"))
	InitMysql()

	r := InitRouter()

	r.Run()

}
