package main

import (
	"path/filepath"
	"runtime"

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

func main() {

	smsfLogger := logrus.New()

	smsfLogger.SetFormatter(&easy.Formatter{
		TimestampFormat: "20060102|15:04:05.999",
		LogFormat:       "[%time%|%blockName%|%srcFileName%|%printLevel%|%codeLine%|%classification%|%logCode%]%msg%\n",
	})

	const blockName = "SMSF"
	const printLevel = 2
	const classification = "EVENT"
	const logCode = 0

	srcFileName, codeLine := getSrcFileInfo()

	smsfLoggerWithFields := smsfLogger.WithFields(logrus.Fields{
		"blockName":   blockName,
		"srcFileName": srcFileName,
		"codeLine":    codeLine,
		"logCode":     logCode,
	})

	writeLog(smsfLoggerWithFields, classification, printLevel, "%s %d occurred", "Strange Error", 999)
}

// TODO line은 원하는 정보와 다르군...
func getSrcFileInfo() (string, int) {

	_, fileName, line, ok := runtime.Caller(0)
	if !ok {
		fileName = "unknown"
		line = 0
	}

	return filepath.Base(fileName), line
}

func writeLog(entry *logrus.Entry, classification string, printLevel int, format string, args ...interface{}) {

	logger2 := entry.WithFields(logrus.Fields{
		"classification": classification,
		"printLevel":     printLevel,
	})

	logger2.Errorf(format, args...)
}
