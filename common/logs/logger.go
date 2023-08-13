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

func Init() {
	var err error
	file, err = os.OpenFile(consts.LogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("无法打开日志文件：", err)
	}
	// 关闭文件会无法写入日志
	// defer file.Close()
	once.Do(func() {
		Logger = log.New(os.Stdout, "INFO", log.Ldate|log.Ltime|log.Lshortfile)
		LocalLogger = log.New(file, "INFO", log.Ldate|log.Ltime|log.Lshortfile)
	})
}

// Info 详情
func Infos(args ...interface{}) {
	Logger.SetPrefix("[INFO]")
	Logger.Println(args...)
	LocalLogger.SetPrefix("[INFO]")
	LocalLogger.Println(args...)
}

// Danger 错误 为什么不命名为 error？避免和 error 类型重名
func Error(args ...interface{}) {
	Logger.SetPrefix("[ERROR]")
	Logger.Println(args...)
	LocalLogger.SetPrefix("[ERROR]")
	LocalLogger.Println(args...)
	LocalLogger.Println("堆栈信息：")
	LocalLogger.Println(string(debug.Stack()))
}
