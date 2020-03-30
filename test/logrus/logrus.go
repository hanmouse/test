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

// Print Level 정의
const (
	_ = iota
	PrintLevelCritical
	PrintLevelMajor
	PrintLevelMinor
	PrintLevelData
	PrintLevelComment
)

type configurations struct {
	logFormat  string
	printLevel int
	// classification: ERROR, EVENT, INFO
	classifications []string
}

var config = configurations{
	logFormat:       "[%time%|%blockName%|%srcFileName%|%printLevel%|%codeLine%|%classification%|%logCode%]%msg%\n",
	printLevel:      PrintLevelCritical,
	classifications: []string{"ERROR", "INFO", "EVENT"},
}

func main() {

	smsfLogger := logrus.New()

	smsfLogger.SetFormatter(&easy.Formatter{
		TimestampFormat: "20060102|15:04:05.999",
		LogFormat:       config.logFormat,
	})

	smsfLoggerWithFields := smsfLogger.WithFields(logrus.Fields{
		"blockName": "usmsf",
		"logCode":   0,
	})

	WriteLog(smsfLoggerWithFields, "EVENT", PrintLevelCritical, "%s %d occurred", "Strange Error", 999)
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

// WriteLog 로그를 쓴다.
func WriteLog(entry *logrus.Entry, classification string, printLevel int, format string, args ...interface{}) {

	if !isClassificationEnabled(classification) || (printLevel > config.printLevel) {
		return
	}

	srcInfo := makeSourceInfo()

	logger := entry.WithFields(logrus.Fields{
		"classification": classification,
		"printLevel":     printLevel,
		"srcFileName":    srcInfo["file"],
		"codeLine":       srcInfo["line"],
	})

	logger.Errorf(format, args...)
}

func isClassificationEnabled(classification string) bool {
	for _, v := range config.classifications {
		if classification == v {
			return true
		}
	}
	return false
}
