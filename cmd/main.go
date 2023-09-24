package main

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	logFile, err := os.Create("../log/example.log")
	if err != nil {
		log.Fatalf("os create file error:%s", err.Error())
	}
	config := zap.NewProductionEncoderConfig() //產生一個production環境的 encoder的配置,用於配置zap logger的格式和行為
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	config.EncodeCaller = zapcore.FullCallerEncoder
	encoder := zapcore.NewConsoleEncoder(config)     //將上面encoder config放入console
	multiWriteSyncer := zapcore.NewMultiWriteSyncer( //設定同時寫入log的地方
		zapcore.AddSync(os.Stdout), //設置寫入資料的地方 os.Stdout(通常指控制台介面或terminal)
		zapcore.AddSync(logFile),   //  //設置寫入log的地方,這邊寫入example.log
	)

	level := zapcore.DebugLevel                               //debug level 適合測試環境
	core := zapcore.NewCore(encoder, multiWriteSyncer, level) //創建一個設定好的核心
	option := zap.AddCaller()                                 //設置logger的option ,新增調用者的路徑,行數,函數名
	logger := zap.New(core, option)                           //產一個設定好內容格式的logger來記住log

	logger.Info("設置訊息紀錄類的log")    //info代表訊息類
	logger.Error("設置ERROR訊息的log") //error代表錯誤訊息

}
