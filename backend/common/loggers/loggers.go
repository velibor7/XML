package loggers

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var authenticationLogger = logrus.New()
var gatewayLogger = logrus.New()
var profileLogger = logrus.New()
var postLogger = logrus.New()
var commentLogger = logrus.New()
var connectionLogger = logrus.New()
var jobLogger = logrus.New()

func NewConnectionLogger() *logrus.Logger {
	connectionLogger.SetLevel(logrus.InfoLevel)
	connectionLogger.SetReportCaller(true)
	connectionLogger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000Z",
	})
	multiWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   "../../logs/connection_service/connection.log",
		MaxSize:    1,
		MaxBackups: 0,
		MaxAge:     28,
		Compress:   true,
	})
	connectionLogger.SetOutput(multiWriter)
	return connectionLogger
}

func NewCommentLogger() *logrus.Logger {
	commentLogger.SetLevel(logrus.InfoLevel)
	commentLogger.SetReportCaller(true)
	commentLogger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000Z",
	})
	multiWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   "../../logs/comment_service/comment.log",
		MaxSize:    1,
		MaxBackups: 0,
		MaxAge:     28,
		Compress:   true,
	})
	commentLogger.SetOutput(multiWriter)
	return commentLogger
}

func NewPostLogger() *logrus.Logger {
	postLogger.SetLevel(logrus.InfoLevel)
	postLogger.SetReportCaller(true)
	postLogger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2021-02-02T16:04:12.000Z",
	})
	multiWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   "../../logs/post_service/post.log",
		MaxSize:    1,
		MaxBackups: 0,
		MaxAge:     28,
		Compress:   true,
	})
	postLogger.SetOutput(multiWriter)
	return postLogger
}

func NewProfileLogger() *logrus.Logger {
	profileLogger.SetLevel(logrus.InfoLevel)
	profileLogger.SetReportCaller(true)
	profileLogger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000Z",
	})
	multiWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   "../../logs/profile_service/profile.log",
		MaxSize:    1,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	})
	profileLogger.SetOutput(multiWriter)
	return profileLogger
}

func NewAuthenticationLogger() *logrus.Logger {
	authenticationLogger.SetLevel(logrus.InfoLevel)
	authenticationLogger.SetReportCaller(true)
	authenticationLogger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000Z",
	})
	multiWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   "../../logs/authentication_service/authentication.log",
		MaxSize:    1,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	})
	authenticationLogger.SetOutput(multiWriter)
	return authenticationLogger
}

func NewGatewayLogger() *logrus.Logger {
	gatewayLogger.SetLevel(logrus.InfoLevel)
	gatewayLogger.SetReportCaller(true)
	gatewayLogger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000Z",
	})
	multiWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   "../../logs/api_gateway/api_gateway.log",
		MaxSize:    1,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	})
	gatewayLogger.SetOutput(multiWriter)
	return gatewayLogger
}

func NewJobLogger() *logrus.Logger {
	jobLogger.SetLevel(logrus.InfoLevel)
	jobLogger.SetReportCaller(true)
	jobLogger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000Z",
	})
	multiWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   "../../logs/job_service/job.log",
		MaxSize:    1,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	})
	jobLogger.SetOutput(multiWriter)
	return jobLogger
}
