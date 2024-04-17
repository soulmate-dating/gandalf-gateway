package auth

import (
	"errors"
	"github.com/soulmate-dating/gandalf-gateway/internal/app/clients/auth"
	"github.com/soulmate-dating/gandalf-gateway/internal/ports/http/response"
	"net/http"
	"strings"

	"github.com/TobbyMax/validator"
	"github.com/labstack/echo/v4"

	"github.com/soulmate-dating/gandalf-gateway/internal/app"
)

const BearerPrefix = "Bearer "

var ErrParameterNotFound = errors.New("necessary parameters not provided")

// @Summary Sign up a new user
// @Description Registers a new user with the provided credentials.
// @Tags auth
// @ID signUpUser
// @Accept json
// @Produce json
// @Param body body Credentials true "User credentials"
// @Success 201 {object} response.Response{data=User,error=nil}"User created"
// @Failure 500 {object} response.Response{data=nil,error=string} "Internal server error"
// @Router /auth/signup [post]
func signup(client auth.AuthServiceClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		var reqBody Credentials
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

// @Summary Log in a user
// @Description Authenticates a user and returns an access token.
// @Tags auth
// @ID loginUser
// @Accept json
// @Produce json
// @Param body body Credentials true "Log in a user"
// @Success 200 {object} response.Response{data=User,error=nil} "User logged in"
// @Failure 500 {object} response.Response{data=nil,error=string} "Internal server error"
// @Router /auth/login [post]
func login(client auth.AuthServiceClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		var reqBody Credentials
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
		return c.JSON(http.StatusOK, response.Success(NewUser(user)))
	}
}

// @Summary Refresh access token
// @Description Refreshes the access token.
// @Tags auth
// @ID refreshToken
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <refresh_token>" "Authorization header"
// @Success 200 {object} response.Response{data=User,error=nil} "User refreshed"
// @Failure 500 {object} response.Response{data=nil,error=string} "Internal server error"
// @Router /auth/refresh [get]
func refresh(client auth.AuthServiceClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if strings.HasPrefix(authHeader, BearerPrefix) == false {
			return c.JSON(http.StatusForbidden, response.Error(errors.New("wrong authorization header format")))
		}
		refreshToken := authHeader[len(BearerPrefix):]
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

// @Summary Log out a user
// @Description Logs out a user by invalidating the access token.
// @Tags auth
// @ID logoutUser
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <access_token>" "Authorization header"
// @Success 200 {object} response.Response{data=nil,error=nil} "User logged out"
// @Failure 500 {object} response.Response{data=nil,error=string} "Internal server error"
// @Router /auth/logout [post]
func logout(client auth.AuthServiceClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if strings.HasPrefix(authHeader, BearerPrefix) == false {
			return c.JSON(http.StatusForbidden, response.Error(errors.New("wrong authorization header format")))
		}
		accessToken := authHeader[len(BearerPrefix):]
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
