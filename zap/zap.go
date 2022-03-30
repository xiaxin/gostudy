package main

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {

	// logger, _ := zap.NewProduction()
	// logger, _ := zap.NewDevelopment()

	logger, _ := New()
	sugar := logger.Named("AppName").Sugar().With(zap.String("one", "two"))

	defer sugar.Sync()

	for i := 0; i < 100; i++ {
		if i%5 == 0 {
			sugar.Debugf("i [%d]", i)
		} else {
			sugar.Infof("i [%d]", i)
		}
		time.Sleep(1 * time.Second)
	}

}

func New() (*zap.Logger, error) {
	fields := make(map[string]interface{})
	fields["a"] = "b"

	return zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: false,
		// Sampling: &zap.SamplingConfig{
		// 	Initial: 5,

		// 	Thereafter: 10,
		// 	Hook: func(e zapcore.Entry, sd zapcore.SamplingDecision) {
		// 		println(e.Message, sd)
		// 	},
		// },

		// Caller
		DisableCaller: true,

		// Stacktrace
		DisableStacktrace: false,

		// Encoder
		Encoding: "console",
		EncoderConfig: zapcore.EncoderConfig{
			FunctionKey:    zapcore.OmitKey,
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeDuration: zapcore.SecondsDurationEncoder,

			// Time
			TimeKey: "T",
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format("2006-01-02 15:04:05"))
			},
			// Level
			LevelKey: "L",
			// LevelEncoder
			//   - 小写 LowercaseLevelEncoder or LowercaseColorLevelEncoder
			//   - 大写 CapitalLevelEncoder   or CapitalColorLevelEncoder
			EncodeLevel: zapcore.CapitalLevelEncoder,

			// Name is the name of the
			NameKey: "N",

			// Caller
			CallerKey: "C",
			// CallerEncoder FullCallerEncoder or ShortCallerEncoder
			EncodeCaller: zapcore.ShortCallerEncoder,

			// Message
			MessageKey: "M",

			// Stack
			StacktraceKey: "S",
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		InitialFields:    fields,
	}.Build()
}
