package middlewares

import (
	"Go2/tracing"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func SpanStarter() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.UserContext()
		ctx, span := tracing.StartSpan(ctx, c.Method()+" "+c.Path())
		c.SetUserContext(ctx)
		err := c.Next()
		finishSpan(span, c.Response().StatusCode(), err)

		return err
	}
}

func finishSpan(span trace.Span, statusCode int, err error) {
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
	} else if statusCode >= 400 {
		span.SetStatus(codes.Error, fmt.Sprintf("HTTP %d", statusCode))
	}
	span.End()
}
