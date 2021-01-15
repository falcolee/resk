package base

import (
	"fmt"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/mattn/go-colorable"
	"github.com/prometheus/common/log"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/tietang/props/v3/kvs"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"io"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func init() {
	formatter := &prefixed.TextFormatter{}
	formatter.ForceColors = true
	formatter.DisableColors = false
	formatter.ForceFormatting = true
	formatter.SetColorScheme(&prefixed.ColorScheme{
		InfoLevelStyle:  "green",
		WarnLevelStyle:  "yellow",
		ErrorLevelStyle: "red",
		FatalLevelStyle: "41",
		PanicLevelStyle: "41",
		DebugLevelStyle: "blue",
		PrefixStyle:     "cyan",
		TimestampStyle:  "37",
	})

	formatter.FullTimestamp = true
	formatter.TimestampFormat = "2006-01-02.15.04:05.000000"
	logrus.SetFormatter(formatter)
	logrus.SetOutput(colorable.NewColorableStderr())
	logrus.SetReportCaller(true)
}

var log_writer io.Writer

func InitLog(conf kvs.ConfigSource) {
	level, err := logrus.ParseLevel(conf.GetDefault("log.level", "info"))
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)
	//配置日志输出目录

	logDir := conf.GetDefault("log.dir", "./logs")
	logPath := logDir
	logFilePath, _ := filepath.Abs(logPath)
	logrus.Infof("log dir: %s", logFilePath)
	logFileName := conf.GetDefault("log.file.name", "red-envelop")
	maxAge := conf.GetDurationDefault("log.max.age", time.Hour*24)
	rotationTime := conf.GetDurationDefault("log.rotation.time", time.Hour*1)
	_ = os.MkdirAll(logPath, os.ModePerm)
	baseLogPath := path.Join(logPath, logFileName)
	//设置滚动日志输出writer
	writer, err := rotatelogs.New(
		strings.TrimSuffix(baseLogPath, ".log")+".%Y%m%d%H.log",
		rotatelogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		logrus.Errorf("config local file system logger error. %+v", err)
	}

	//设置日志文件输出的日志格式
	formatter := &logrus.TextFormatter{}
	formatter.CallerPrettyfier = func(frame *runtime.Frame) (function string, file string) {
		function = frame.Function
		dir, filename := path.Split(frame.File)
		f := path.Base(dir)
		return function, fmt.Sprintf("%s/%s:%d", f, filename, frame.Line)
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, formatter)

	log.AddHook(lfHook)

	log_writer = writer
}
