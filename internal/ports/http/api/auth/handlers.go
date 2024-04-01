package auth

import (
	"errors"
	"github.com/soulmate-dating/gandalf-gateway/internal/app/clients/auth"
	"github.com/soulmate-dating/gandalf-gateway/internal/ports/http/response"
	"net/http"

	"github.com/TobbyMax/validator"
	"github.com/labstack/echo/v4"

	"github.com/soulmate-dating/gandalf-gateway/internal/app"
)

var ErrParameterNotFound = errors.New("necessary parameters not provided")

func signup(client auth.AuthServiceClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		var reqBody loginRequest
		err := c.Bind(&reqBody)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.Error(err))
		}

		user, err := client.SignUp(
			c.Request().Context(),
			&auth.SignUpRequest{Email: reqBody.Email, Password: reqBody.Password},
		)

		if err != nil {
			switch {
			case errors.As(err, &validator.ValidationErrors{}):
				return c.JSON(http.StatusBadRequest, response.Error(err))
			default:
				return c.JSON(http.StatusInternalServerError, response.Error(err))
			}
		}
		return c.JSON(http.StatusCreated, response.Success(NewUser(user)))
	}
}

func login(client auth.AuthServiceClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		var reqBody loginRequest
		err := c.Bind(&reqBody)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.Error(err))
		}

		user, err := client.Login(
			c.Request().Context(),
			&auth.LoginRequest{Email: reqBody.Email, Password: reqBody.Password},
		)

		if err != nil {
			switch {
			case errors.As(err, &validator.ValidationErrors{}):
				return c.JSON(http.StatusBadRequest, response.Error(err))
			default:
				return c.JSON(http.StatusInternalServerError, response.Error(err))
			}
		}
		c.Response().Header().Set("Access-Token", user.AccessToken)
		c.Response().Header().Set("Refresh-Token", user.RefreshToken)
		return c.JSON(http.StatusOK, response.Success(NewUser(user)))
	}
}

func refresh(client auth.AuthServiceClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		refreshToken := c.Request().Header.Get("Refresh-Token")
		if refreshToken == "" {
			return c.JSON(http.StatusBadRequest, response.Error(ErrParameterNotFound))
		}

		user, err := client.Refresh(c.Request().Context(), &auth.RefreshRequest{
			RefreshToken: refreshToken,
		})
		if err != nil {
			switch {
			case errors.Is(err, app.ErrForbidden):
				return c.JSON(http.StatusForbidden, response.Error(err))
			default:
				return c.JSON(http.StatusInternalServerError, response.Error(err))
			}
		}
		c.Response().Header().Set("Access-Token", user.AccessToken)
		c.Response().Header().Set("Refresh-Token", user.RefreshToken)
		return c.JSON(http.StatusOK, response.Success(NewUser(user)))
	}
}

func logout(client auth.AuthServiceClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		accessToken := c.Request().Header.Get("Access-Token")
		if accessToken == "" {
			return c.JSON(http.StatusBadRequest, response.Error(ErrParameterNotFound))
		}

		_, err := client.Logout(c.Request().Context(), &auth.LogoutRequest{
			AccessToken: accessToken,
		})
		if err != nil {
			switch {
			case errors.Is(err, app.ErrForbidden):
				return c.JSON(http.StatusForbidden, response.Error(err))
			default:
				return c.JSON(http.StatusInternalServerError, response.Error(err))
			}
		}
		return c.JSON(http.StatusOK, response.Success(nil))
	}
}
