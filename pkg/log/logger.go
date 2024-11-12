package log

import (
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	L    *logrus.Logger
	once sync.Once
)

// InitLogger 在函数运行时初始化 logger
func InitLogger() *logrus.Logger {
	once.Do(func() {
		L = logrus.New()
		L.SetFormatter(&logrus.JSONFormatter{})
	})
	return L
}
