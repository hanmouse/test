package main

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

/*
삼성향 로그 포맷:

[<yymmdd>|<hh:mm:ss.msec>|<Block>|<Thread>|<PrintLevel>|<code_line>|<Classification>|<LogCode>] log_string
- Block: 로그 발생 블록명
- Thread: 소스 파일 이름
- PrintLevel: 1 ~ 5 (1: Critical, 2: Major(Default), 3: Minor, 4: Data, 5: Comment)
- Classification: ERROR, EVENT, INFO
- LogCode: 0

[200203|20:53:25.619|DBCS.K|FMCDBG_FMCNanoMsgHandler.cpp|1|271|INFO|0]waitForMessageForSub()- recv _msg_IomDiskUsageReq_IomAgent
*/

// Print Level 정의
const (
	_ = iota
	PrintLevelCritical
	PrintLevelMajor
	PrintLevelMinor
	PrintLevelData
	PrintLevelComment
)

type logConfig struct {
	// printLevel map의 key는 classification 이름 (e.g. "SELFDIAG", "ERROR")
	printLevel map[string]int
}

type configurations struct {
	alarm logConfig
	diag  logConfig
	app   logConfig
}

var config = configurations{
	alarm: logConfig{
		printLevel: map[string]int{
			"FAULT": PrintLevelMajor,
		},
	},
	diag: logConfig{
		printLevel: map[string]int{
			"SELFDIAG": PrintLevelMinor,
			"RESET":    PrintLevelMinor,
			"INIT":     PrintLevelMinor,
			"CONFIG":   PrintLevelMinor,
		},
	},
	app: logConfig{
		printLevel: map[string]int{
			"ERROR": PrintLevelMajor,
			"EVENT": PrintLevelMinor,
			"INFO":  PrintLevelData,
		},
	},
}

// SamsungLogger : 삼성향 logger 정의
type SamsungLogger struct {
	logger *logrus.Logger
	entry  *logrus.Entry
	config *logConfig
}

// SamsungLoggers : 삼성향 logger들의 집합
type SamsungLoggers struct {
	alarm SamsungLogger
	diag  SamsungLogger
	app   SamsungLogger
}

func (s *SamsungLogger) init(config *logConfig) {

	s.config = config

	s.logger = logrus.New()

	s.logger.SetFormatter(&easy.Formatter{
		TimestampFormat: "20060102|15:04:05.999",
		LogFormat:       "[%time%|%blockName%|%srcFileName%|%printLevel%|%codeLine%|%classification%|%logCode%]%msg%\n",
	})

	s.entry = s.logger.WithFields(logrus.Fields{
		"blockName": filepath.Base(os.Args[0]),
		"logCode":   0,
	})
}

// Init : logger들을 초기화한다.
func (loggers *SamsungLoggers) Init() {

	loggers.alarm.init(&config.alarm)
	loggers.diag.init(&config.diag)
	loggers.app.init(&config.app)
}

// AlarmLogger : alarm logger를 반환한다.
func (loggers *SamsungLoggers) AlarmLogger() *SamsungLogger {
	return &loggers.alarm
}

// DiagLogger : diag logger를 반환한다.
func (loggers *SamsungLoggers) DiagLogger() *SamsungLogger {
	return &loggers.diag
}

// AppLogger : app logger를 반환한다.
func (loggers *SamsungLoggers) AppLogger() *SamsungLogger {
	return &loggers.app
}

// Log : 로그를 기록한다.
func (s *SamsungLogger) Log(classification string, printLevel int, format string, args ...interface{}) {

	_, present := s.config.printLevel[classification]
	if !present || (printLevel > s.config.printLevel[classification]) {
		return
	}

	srcInfo := makeSourceInfo()

	logger := s.entry.WithFields(logrus.Fields{
		"classification": classification,
		"printLevel":     printLevel,
		"srcFileName":    srcInfo["file"],
		"codeLine":       srcInfo["line"],
	})

	logger.Errorf(format, args...)
}

func main() {

	var loggers SamsungLoggers

	loggers.Init()

	alarmLogger := loggers.AlarmLogger()
	diagLogger := loggers.DiagLogger()
	appLogger := loggers.AppLogger()

	alarmLogger.Log("FAULT", PrintLevelCritical, "Alarm!!!: code %#v", 2)
	diagLogger.Log("RESET", PrintLevelMajor, "System Reset")
	appLogger.Log("EVENT", PrintLevelMinor, "%s %d occurred", "Strange Error", 999)
	appLogger.Log("UNKNOWN", PrintLevelMinor, "%s %d occurred", "Strange Error", 999)
}

func makeSourceInfo() logrus.Fields {

	for i := 2; i < 50; i++ {
		_, file, line, ok := runtime.Caller(i)
		if ok {
			if strings.HasSuffix(file, "_noposline.go") {
				continue
			}
			return logrus.Fields{"file": filepath.Base(file), "line": line}
		}
		break
	}

	return logrus.Fields{"file": "unknown"}
}
