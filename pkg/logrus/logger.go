package logrus

import (
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"time"
)

func MakeLogEntry(c echo.Context) *log.Entry {
	if c == nil {
		return log.WithFields(log.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	return log.WithFields(log.Fields{
		"at":         time.Now().Format("2006-01-02 15:04:05"),
		"method":     c.Request().Method,
		"status":     c.Response().Status,
		"uri":        c.Request().URL.String(),
		"ip":         c.Request().RemoteAddr,
		"user-agent": c.Request().UserAgent(),
	})
}
