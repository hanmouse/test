package main

import (
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

func main() {

	smsfLogger := logrus.New()

	smsfLogger.SetFormatter(&easy.Formatter{
		TimestampFormat: "20060102|15:04:05.999",
		LogFormat:       "[%time%|%blockName%|%srcFileName%|%printLevel%|%codeLine%|%classification%|%logCode%]%msg%\n",
	})

	smsfLoggerWithFields := smsfLogger.WithFields(logrus.Fields{
		"blockName": "usmsf",
		"logCode":   0,
	})

	/*
		classification := "EVENT"
		printLevel := 2
	*/

	writeLog(smsfLoggerWithFields, "EVENT", 2, "%s %d occurred", "Strange Error", 999)
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

func writeLog(entry *logrus.Entry, classification string, printLevel int, format string, args ...interface{}) {

	srcInfo := makeSourceInfo()

	logger2 := entry.WithFields(logrus.Fields{
		"classification": classification,
		"printLevel":     printLevel,
		"srcFileName":    srcInfo["file"],
		"codeLine":       srcInfo["line"],
	})

	logger2.Errorf(format, args...)
}
