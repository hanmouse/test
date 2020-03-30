package main

import (
	"camel.uangel.com/ua5g/ulib.git/implements/logruslogger"
	"camel.uangel.com/ua5g/ulib.git/loggerfactory"
	"camel.uangel.com/ua5g/ulib.git/uconf"
	"camel.uangel.com/ua5g/ulib.git/ulog"
	"github.com/csgura/di"
	"github.com/savsgio/go-logger"
)

func main() {
	logger := ulog.GetLogger("com.uangel")
	ulog.SetFormatter(&log.TextFormatter{
		TimestampFormat:           "060102 15:04:05.999",
		FullTimestamp:             true,
		EnvironmentOverrideColors: true,
	})

	logger.Warn("hello world")
}

type LogrusModule struct {
}

func (this *LogrusModule) Configure(binder *di.Binder) {
	binder.BindProvider((*ulog.LoggerFactory)(nil), func(injector di.Injector) interface{} {
		cfg := injector.GetInstance((*uconf.Config)(nil)).(uconf.Config)

		configFile := cfg.GetString("logrus.config-file")

		if configFile != "" {
			ulog.Info("Use LoggerFactory")
			logfactory := loggerfactory.NewLogFactory(configFile, logruslogger.CreateLogrusLogger)
			ulog.SetLoggerFactory(logfactory)
			return (ulog.LoggerFactory)(logfactory)
		}

		logger.Info("Not Use LoggerFactory")

		logfactory := func(loggerName string) ulog.Logger {
			ulog.Info("Create new Logger : %s", loggerName)
			var logger = log.New()
			if cfg.GetString("logrus.format") == "json" {
				logger.Info("Use JSON Formatter")
				logger.SetFormatter(&log.JSONFormatter{})
			} else {
				logger.SetFormatter(&log.TextFormatter{
					TimestampFormat:           "060102 15:04:05.999",
					EnvironmentOverrideColors: true,
					FullTimestamp:             true,
				})
			}

			return logruslogger.NewLogrusLogger(logger, cfg.GetBoolean("logrus.print-caller"))
		}
		ulog.SetLoggerFactory(logfactory)
		return (ulog.LoggerFactory)(logfactory)

	}).AsEagerSingleton()
}
