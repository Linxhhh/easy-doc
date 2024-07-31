package core

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogRequest struct {
	LogPath  string
	AppName  string
	NoDate   bool    // 不按照日期划分日志
	NoErr    bool    // 不单独存放 Err 日志信息
	NoGlobal bool    // 不替换全局 logger
}

type LogFormatter struct{}

// Format 实现 Formatter(entry *logrus.Entry) ([]byte, error) 接口
func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 根据不同 level 去展示颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	// 自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		// 自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		// 自定义输出格式
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s\n", timestamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

// DateHook 按照时间将日志信息写入 gvd 文件
type DateHook struct {
	file     *os.File
	fileDate string
	LogPath  string
	AppName  string
}

func (DateHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook DateHook) Fire(entry *logrus.Entry) error {
	timer := entry.Time.Format("2006-01-02")
	line, _ := entry.String()
	if hook.fileDate == timer {
		hook.file.Write([]byte(line))
		return nil
	}
	// 时间不等
	hook.file.Close()
	os.MkdirAll(path.Join(hook.LogPath, timer), os.ModePerm)
	filename := path.Join(hook.LogPath, timer, hook.AppName + ".log")

	hook.file, _ = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	hook.fileDate = timer
	hook.file.Write([]byte(line))
	return nil
}

// ErrorHook 将 error 级别信息写入 err 文件
type ErrorHook struct {
	file     *os.File
	fileDate string
	LogPath  string
	AppName  string
}

func (ErrorHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel}
}

func (hook ErrorHook) Fire(entry *logrus.Entry) error {
	timer := entry.Time.Format("2006-01-02")
	line, _ := entry.String()
	if hook.fileDate == timer {
		hook.file.Write([]byte(line))
		return nil
	}
	// 时间不等
	hook.file.Close()
	os.MkdirAll(path.Join(hook.LogPath, timer), os.ModePerm)
	filename := path.Join(hook.LogPath, timer, "err.log")

	hook.file, _ = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	hook.fileDate = timer
	hook.file.Write([]byte(line))
	return nil
}

func InitLogger(requestList ...LogRequest) *logrus.Logger {
	var request LogRequest
	if len(requestList) > 0 {
		request = requestList[0]
	}
	if request.LogPath == "" {
		request.LogPath = "logs"
	}
	if request.AppName == "" {
		request.AppName = "gvd"
	}

	mLog := logrus.New()               // 新建一个实例
	mLog.SetOutput(os.Stdout)          // 设置输出类型
	mLog.SetReportCaller(true)         // 开启返回函数名和行号
	mLog.SetFormatter(&LogFormatter{}) // 设置自定义 Formatter
	mLog.SetLevel(logrus.DebugLevel)   // 设置最低的 Level

	if !request.NoDate {
		mLog.AddHook(&DateHook{        // 设置全级别的 Hook
			LogPath: request.LogPath,
			AppName: request.AppName,
		})      
	}
	if !request.NoErr {
		mLog.AddHook(&ErrorHook{       // 设置 error 级别的 Hook
			LogPath: request.LogPath,
			AppName: request.AppName,
		})
	}
	if !request.NoGlobal {
		InitDefaultLogger()            // 设置全局 logger
	}
	return mLog
}

func InitDefaultLogger() {
	// 全局 logger
	logrus.SetOutput(os.Stdout)
	logrus.SetReportCaller(true)         // 开启返回函数名和行号
	logrus.SetFormatter(&LogFormatter{}) // 设置自定义 Formatter
	logrus.SetLevel(logrus.DebugLevel)   // 设置最低的 Level
}
