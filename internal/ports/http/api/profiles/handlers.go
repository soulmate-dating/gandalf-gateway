package profiles

import (
	"bytes"
	"errors"
	"github.com/soulmate-dating/gandalf-gateway/internal/app/clients/profiles"
	"github.com/soulmate-dating/gandalf-gateway/internal/ports/http/response"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/TobbyMax/validator"
	"github.com/labstack/echo/v4"
)

var ErrParameterNotFound = errors.New("necessary parameters not provided")

func createProfile(client profiles.ProfileServiceClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		var reqBody Profile
		userID := c.Param("user_id")
		err := c.Bind(&reqBody)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.Error(err))
		}

		p, err := client.CreateProfile(
			c.Request().Context(),
			&profiles.CreateProfileRequest{
				Id: userID,
				PersonalInfo: &profiles.PersonalInfo{
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

func getProfile(client profiles.ProfileServiceClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Param("user_id")
		p, err := client.GetProfile(
			c.Request().Context(),
			&profiles.GetProfileRequest{Id: userID},
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

func updateProfile(client profiles.ProfileServiceClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		var reqBody Profile
		userID := c.Param("user_id")
		err := c.Bind(&reqBody)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.Error(err))
		}

		p, err := client.CreateProfile(
			c.Request().Context(),
			&profiles.CreateProfileRequest{
				Id: userID,
				PersonalInfo: &profiles.PersonalInfo{
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
		return c.JSON(http.StatusOK, response.Success(NewProfile(p)))
	}
}

func getFullProfile(client profiles.ProfileServiceClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Param("user_id")
		p, err := client.GetFullProfile(
			c.Request().Context(),
			&profiles.GetProfileRequest{Id: userID},
		)

		if err != nil {
			switch {
			case errors.As(err, &validator.ValidationErrors{}):
				return c.JSON(http.StatusBadRequest, response.Error(err))
			default:
				return c.JSON(http.StatusInternalServerError, response.Error(err))
			}
		}
		return c.JSON(http.StatusOK, response.Success(NewFullProfile(p)))
	}
}

func getRecommendation(client profiles.ProfileServiceClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Param("user_id")
		p, err := client.GetRandomProfilePreferredByUser(
			c.Request().Context(),
			&profiles.GetRandomProfilePreferredByUserRequest{UserId: userID},
		)

		if err != nil {
			switch {
			case errors.As(err, &validator.ValidationErrors{}):
				return c.JSON(http.StatusBadRequest, response.Error(err))
			default:
				return c.JSON(http.StatusInternalServerError, response.Error(err))
			}
		}
		return c.JSON(http.StatusOK, response.Success(NewFullProfile(p)))
	}
}

func getPrompts(client profiles.ProfileServiceClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Param("user_id")
		prompts, err := client.GetPrompts(
			c.Request().Context(),
			&profiles.GetPromptsRequest{UserId: userID},
		)

		if err != nil {
			switch {
			case errors.As(err, &validator.ValidationErrors{}):
				return c.JSON(http.StatusBadRequest, response.Error(err))
			default:
				return c.JSON(http.StatusInternalServerError, response.Error(err))
			}
		}
		return c.JSON(http.StatusOK, response.Success(Prompts(prompts)))
	}
}

func createPrompt(client profiles.ProfileServiceClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		var reqBody []Prompt
		userID := c.Param("user_id")
		err := c.Bind(&reqBody)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.Error(err))
		}

		prompts, err := client.AddPrompts(
			c.Request().Context(),
			&profiles.AddPromptsRequest{
				UserId:  userID,
				Prompts: mapPrompts(reqBody),
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
		return c.JSON(http.StatusCreated, response.Success(Prompts(prompts)))
	}
}

func updatePrompt(client profiles.ProfileServiceClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		var reqBody Prompt
		userID := c.Param("user_id")
		promptID := c.Param("prompt_id")
		err := c.Bind(&reqBody)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.Error(err))
		}

		prompt, err := client.UpdatePrompt(
			c.Request().Context(),
			&profiles.UpdatePromptRequest{
				UserId: userID,
				Prompt: &profiles.Prompt{
					Id:       promptID,
					Question: reqBody.Question,
					Content:  reqBody.Content,
					Position: reqBody.Position,
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
		return c.JSON(http.StatusOK, response.Success(NewPrompt(prompt.Prompt)))
	}
}

func updateFilePrompt(client profiles.ProfileServiceClient) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var (
			buf bytes.Buffer
		)
		userID := c.Param("user_id")
		promptID := c.Param("prompt_id")

		file, err := c.FormFile("file")
		if err != nil {
			return err
		}
		data, err := file.Open()
		if err != nil {
			return err
		}
		defer func(data multipart.File) {
			if err != nil {
				err = data.Close()
			}
		}(data)
		_, err = io.Copy(&buf, data)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.Error(err))
		}

		prompt, err := client.UpdateFilePrompt(
			c.Request().Context(),
			&profiles.UpdateFilePromptRequest{
				UserId:   userID,
				Id:       promptID,
				Question: c.FormValue("question"),
				Content:  buf.Bytes(),
				Type:     c.FormValue("type"),
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
		return c.JSON(http.StatusOK, response.Success(NewPrompt(prompt.Prompt)))
	}
}

func createFilePrompt(client profiles.ProfileServiceClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			buf bytes.Buffer
		)
		userID := c.Param("user_id")
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}
		data, err := file.Open()
		if err != nil {
			return err
		}
		defer func(data multipart.File) {
			if err != nil {
				err = data.Close()
			}
		}(data)
		_, err = io.Copy(&buf, data)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.Error(err))
		}

		prompt, err := client.AddFilePrompt(
			c.Request().Context(),
			&profiles.AddFilePromptRequest{
				UserId:   userID,
				Question: c.FormValue("question"),
				Content:  buf.Bytes(),
				Type:     c.FormValue("type"),
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
		return c.JSON(http.StatusOK, response.Success(NewPrompt(prompt.Prompt)))
	}
}
