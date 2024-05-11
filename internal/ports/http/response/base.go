package response

import (
	"github.com/labstack/echo/v4"
	"github.com/soulmate-dating/gandalf-gateway/pkg/grpc2http"
	"google.golang.org/grpc/status"
	"net/http"
)

// Response is a generic response structure.
type Response struct {
	Data  any `json:"data"`
	Error any `json:"error"`
}

func Error(message string) Response {
	return Response{Data: nil, Error: message}
}

func Success(data any) Response {
	return Response{Data: data, Error: nil}
}

func MapError(c echo.Context, err error) error {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			return c.JSON(grpc2http.TransformCode(e.Code()), Error(e.Message()))
		} else {
			return c.JSON(http.StatusInternalServerError, Error(err.Error()))
		}
	}
	return nil
}
