module gostudy

go 1.16

replace (
	github.com/gin-gonic/gin v1.7.3 => /Users/xiaxin/YayProjects/opensource/gin
	github.com/xiaxin/logger => /Users/xiaxin/YayProjects/logger
	go.uber.org/zap v1.21.0 => /Users/xiaxin/YayProjects/opensource/zap
)

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/go-playground/validator/v10 v10.8.0 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/mattn/go-isatty v0.0.13 // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/opentracing/opentracing-go v1.2.0
	github.com/uber/jaeger-client-go v2.30.0+incompatible
	github.com/ugorji/go v1.2.6 // indirect
	github.com/xiaxin/logger v0.0.0-00010101000000-000000000000
	go.uber.org/zap v1.21.0
)
