package middleware

import (
	"reflect"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/sirupsen/logrus"
)

// LoggerToLogrusGlobal logs message using global instance of StandardLogger
func LoggerToLogrusGlobal(next actor.ActorFunc) actor.ActorFunc {
	return LoggerToLogrusInstance(logrus.StandardLogger())(next)
}

// LoggerToLogrusInstance logs message using logger instance passed as parameter
func LoggerToLogrusInstance(logger logrus.FieldLogger) func(next actor.ActorFunc) actor.ActorFunc {
	return func(next actor.ActorFunc) actor.ActorFunc {
		fn := func(c actor.Context) {
			message := c.Message()
			logger.WithFields(logrus.Fields{
				"actor":   c.Self,
				"message": reflect.TypeOf(message),
				"content": message,
			}).Debug("received message")
			next(c)
		}

		return fn
	}
}
