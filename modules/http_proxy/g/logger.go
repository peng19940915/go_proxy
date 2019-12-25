package g

import (
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"path"
	"time"
)

func ConfigLocalFilesystemLogger(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) {
	baseLogPath := path.Join(logPath, logFileName)
	writer, err := rotatelogs.New(
		baseLogPath+"_%Y%m%d.log",
		rotatelogs.WithLinkName(baseLogPath),      // windows环境下无法生成软连接
		rotatelogs.WithMaxAge(maxAge),             // 文件保存最大时间
		rotatelogs.WithRotationTime(rotationTime), //日志切割时间间隔
	)
	if err != nil {
		logrus.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.TextFormatter{})
	logrus.AddHook(lfHook)
}

func InitLog() {
	logPath := Config().Log.LogPath
	level := Config().Log.LogLevel
	logFilename := Config().Log.FileName
	maxAge := Config().Log.MaxAge * time.Hour * 24
	rotationTime := Config().Log.RotationTime * time.Hour * 24
	// 配置日志
	ConfigLocalFilesystemLogger(logPath, logFilename, maxAge, rotationTime)
	// 配置级别
	switch level {
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	default:
		logrus.Fatal("log conf only allow [info, debug, warn], please check your confguire")
	}
}
