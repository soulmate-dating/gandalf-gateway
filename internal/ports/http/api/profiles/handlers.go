package profiles

import (
	"errors"
	"github.com/soulmate-dating/gandalf-gateway/internal/app/clients/profile"
	"github.com/soulmate-dating/gandalf-gateway/internal/ports/http/response"
	"net/http"

	"github.com/TobbyMax/validator"
	"github.com/labstack/echo/v4"
)

var ErrParameterNotFound = errors.New("necessary parameters not provided")

func createProfile(client profile.ProfileServiceClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		var reqBody Profile
		userID := c.Param("user_id")
		err := c.Bind(&reqBody)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.Error(err))
		}

		p, err := client.CreateProfile(
			c.Request().Context(),
			&profile.CreateProfileRequest{
				Id: userID,
				PersonalInfo: &profile.PersonalInfo{
					FirstName:        reqBody.FirstName,
					LastName:         reqBody.LastName,
					BirthDate:        reqBody.BirthDate,
					Sex:              reqBody.Sex,
					PreferredPartner: reqBody.PreferredPartner,
					Intention:        reqBody.Intention,
					Height:           reqBody.Height,
					HasChildren:      reqBody.HasChildren,
					FamilyPlans:      reqBody.FamilyPlans,
					Location:         reqBody.Location,
					DrinksAlcohol:    reqBody.DrinksAlcohol,
					Smokes:           reqBody.Smokes,
				},
			},
		)

		if err != nil {
			switch {
			case errors.As(err, &validator.ValidationErrors{}):
				return c.JSON(http.StatusBadRequest, response.Error(err))
			default:
				return c.JSON(http.StatusInternalServerError, response.Error(err))
			}
		}
		return c.JSON(http.StatusCreated, response.Success(NewProfile(p)))
	}
}

func getProfile(client profile.ProfileServiceClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Param("user_id")
		p, err := client.GetProfile(
			c.Request().Context(),
			&profile.GetProfileRequest{Id: userID},
		)

		if err != nil {
			switch {
			case errors.As(err, &validator.ValidationErrors{}):
				return c.JSON(http.StatusBadRequest, response.Error(err))
			default:
				return c.JSON(http.StatusInternalServerError, response.Error(err))
			}
		}
		return c.JSON(http.StatusOK, response.Success(NewProfile(p)))
	}
}

func updateProfile(client profile.ProfileServiceClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		var reqBody Profile
		userID := c.Param("user_id")
		err := c.Bind(&reqBody)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.Error(err))
		}

		p, err := client.CreateProfile(
			c.Request().Context(),
			&profile.CreateProfileRequest{
				Id: userID,
				PersonalInfo: &profile.PersonalInfo{
					FirstName:        reqBody.FirstName,
					LastName:         reqBody.LastName,
					BirthDate:        reqBody.BirthDate,
					Sex:              reqBody.Sex,
					PreferredPartner: reqBody.PreferredPartner,
					Intention:        reqBody.Intention,
					Height:           reqBody.Height,
					HasChildren:      reqBody.HasChildren,
					FamilyPlans:      reqBody.FamilyPlans,
					Location:         reqBody.Location,
					DrinksAlcohol:    reqBody.DrinksAlcohol,
					Smokes:           reqBody.Smokes,
				},
			},
		)

		if err != nil {
			switch {
			case errors.As(err, &validator.ValidationErrors{}):
				return c.JSON(http.StatusBadRequest, response.Error(err))
			default:
				return c.JSON(http.StatusInternalServerError, response.Error(err))
			}
		}
		return c.JSON(http.StatusCreated, response.Success(NewProfile(p)))
	}
}
