package log

import (
	"errors"
	"runtime"

	"go.uber.org/zap"
)

type Line struct {
	FileName string
	Line     int
	FuncName string
}

func GetCurrentLine() (*Line, error) {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		line := &Line{
			FileName: file,
			Line:     line,
			FuncName: runtime.FuncForPC(pc).Name(),
		}
		return line, nil
	}
	return nil, errors.New("failed to get current line")
}

func (l *Line) String() string {
	return l.FileName + ":" + string(rune(l.Line)) + " " + l.FuncName
}

var logger *zap.SugaredLogger

func AppLogger() *zap.SugaredLogger {
	if logger != nil {
		return logger
	}
	l, _ := zap.NewProduction()
	defer l.Sync()

	undo := zap.ReplaceGlobals(l)
	defer undo()

	logger = l.Sugar()
	return logger
}
