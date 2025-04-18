// logger/logger.go
package logger

import (
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func InitLogger() {
	logFile, err := os.OpenFile("logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("cannot open log file: " + err.Error())
	}

	// Encoder cấu hình output format
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	// Ghi log ra file
	writeSyncer := zapcore.AddSync(logFile)

	core := zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel)

	Log = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

}

func Info(c *gin.Context, msg string, fields ...zap.Field) {
	Log.With(zap.String("request_id", getRequestID(c))).Info(msg, fields...)
}

func getRequestID(c *gin.Context) string {
	rid, ok := c.Get("request_id")
	if !ok {
		return "unknown"
	}
	return rid.(string)
}
