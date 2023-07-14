package logger

import "go.uber.org/zap"

var logger *zap.SugaredLogger

type Fields map[string]interface{}

func NewZapLogger(){
	log, err := zap.NewProduction()
	if err != nil{
		panic(err)
	}
	
	sugar := log.Sugar()
	defer log.Sync()

	logger = sugar
}

func Infof(format string, args ...interface{}){
	logger.Infof(format,args...)
}

func Warnf(format string, args ...interface{}){
	logger.Warnf(format,args...)
}

func Errorf(format string, args ...interface{}){
	logger.Errorf(format,args...)
}

func Fatal(args ...interface{}){
	logger.Fatal(args)
}

func WithFields(fields Fields) *zap.SugaredLogger{
	var f = make([]interface{},0)
	for index, field := range fields{
		f = append(f, index)
		f = append(f, field)
	}
	log := logger.With(f...)
	return log
}