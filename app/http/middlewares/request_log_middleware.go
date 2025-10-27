package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func RequestLogMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		start := time.Now()
		err := ctx.Next()
		end := time.Now()

		entry := logrus.WithFields(logrus.Fields{
			"ip":      ctx.IP(),
			"latency": end.Sub(start).String(),
			"method":  ctx.Method(),
			"path":    ctx.Path(),
			"query":   ctx.Context().QueryArgs().String(),
			"status":  ctx.Response().StatusCode(),
		})

		if err != nil {
			logrus.WithField("error", err.Error())
			if ctx.Response().StatusCode() == 500 {
				entry.Error(err.Error())
			} else {
				entry.Warn(err.Error())
			}

			return err
		}

		entry.Info("request finished")

		return nil
	}
}
