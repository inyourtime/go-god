package logs

import (
	"gopher/src/coreplugins"
	"gopher/src/model"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	config := zap.NewProductionEncoderConfig()
	config.TimeKey = "timestamp"
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	// var err error
	// log, err = config.Build(zap.AddCallerSkip(1))
	// if err != nil {
	// 	panic(err)
	// }
	discordCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		zapcore.AddSync(NewDiscordWriter()),
		zap.ErrorLevel, // Send logs at the error level and above
	)
	// config.Core
	log = zap.New(discordCore, zap.AddCaller(), zap.AddCallerSkip(1))
	// zap.ReplaceGlobals(log)
	defer log.Sync()
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

func Error(message interface{}, fields ...zap.Field) {
	switch v := message.(type) {
	case error:
		log.Error(v.Error(), fields...)
	case string:
		log.Error(v, fields...)
	}
}

type DiscordWriter struct{}

func NewDiscordWriter() *DiscordWriter {
	return &DiscordWriter{}
}

func (w *DiscordWriter) Write(p []byte) (n int, err error) {
	message := string(p)
	ev := model.ServerEnvironment{
		Hostname: coreplugins.Ctx.Hostname(),
		Url: coreplugins.Ctx.OriginalURL(),
		Method: coreplugins.Ctx.Method(),
	}
	go coreplugins.WebhookSend(message, ev)
	return len(p), nil
}
