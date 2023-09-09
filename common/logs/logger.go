package logs

import (
	"log"
	"os"
	"runtime/debug"
	"sync"

	"wechatGpt/common/consts"
)

var Logger *log.Logger
var LocalLogger *log.Logger
var once sync.Once
var file *os.File

func Init(needLocalStore bool) error {
	if needLocalStore == true {
		var err error
		file, err = os.OpenFile(consts.LogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal("无法打开日志文件：", err)
			return err
		}
		LocalLogger = log.New(file, "INFO", log.Ldate|log.Ltime|log.Lshortfile)
	}

	// 关闭文件会无法写入日志
	// defer file.Close()
	Logger = log.New(os.Stdout, "INFO", log.Ldate|log.Ltime|log.Lshortfile)

	return nil
}

// Info 详情
func Info(format string, v ...any) {
	Logger.SetPrefix("[INFO]")
	Logger.Printf(format, v...)
}

// Error 错误
func Error(format string, v ...any) {
	Logger.SetPrefix("[ERROR]")
	Logger.Printf(format, v...)
	if LocalLogger != nil {
		LocalLogger.SetPrefix("[ERROR]")
		LocalLogger.Printf(format, v...)
		LocalLogger.Printf("堆栈信息：%v", string(debug.Stack()))
	}
}
