package middleware

import (
	"fmt"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
	relog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	scr,err:=os.OpenFile(utils.LogFilePath,os.O_RDWR|os.O_CREATE,0755)
	if err!=nil{
		fmt.Println("err",err)
	}
	logger := logrus.New()
	logger.Out = scr
	logger.SetLevel(logrus.DebugLevel)

	logWriter, _ := relog.New(
		utils.LogFilePath+"%Y%m%d.log",
		relog.WithMaxAge(7*24*time.Hour),
		relog.WithRotationTime(24*time.Hour),
		//relog.WithLinkName(linkName),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger.AddHook(Hook)

	return func(c *gin.Context) {
		startTime:=time.Now()
		c.Next()
		stopTime :=time.Since(startTime)
		spendTime:=fmt.Sprintf("%d ms",int(math.Ceil(float64(stopTime.Nanoseconds())/1000000.0)))//ms
		hostName,err:=os.Hostname()
		if err !=nil{
			hostName="unknown"
		}
		statusCode := c.Writer.Status()
		clientIp :=c.ClientIP()
		userAgent :=c.Request.UserAgent()
		dataSize := c.Writer.Size()
		if dataSize<0{
			dataSize = 0
		}
		method :=c.Request.Method
		path := c.Request.RequestURI

		entry:=logger.WithFields(logrus.Fields{
			"Hostname":hostName,
			"Status":statusCode,
			"SpendTime":spendTime,
			"IP":clientIp,
			"Method":method,
			"Path":path,
			"DataSize":dataSize,
			"Agent":userAgent,
		})
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}

}
