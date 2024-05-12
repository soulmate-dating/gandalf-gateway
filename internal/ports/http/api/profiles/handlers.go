package profiles

import (
	"bytes"
	"errors"
	"github.com/soulmate-dating/gandalf-gateway/internal/app/clients/profiles"
	"github.com/soulmate-dating/gandalf-gateway/internal/ports/http/middleware"
	"github.com/soulmate-dating/gandalf-gateway/internal/ports/http/response"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/labstack/echo/v4"
)

var ErrParameterNotFound = errors.New("necessary parameters not provided")

// @Summary Create user profile
// @Description This can only be done by the logged in user.
// @Tags profiles
// @ID createProfileByUserId
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <access_token>" "Authorization header"
// @Param user_id path string true "User ID"
// @Param body body Profile true "Create user profile"
// @Success 201 {object} response.Response{data=Profile,error=nil} "Profile created"
// @Failure 500 {object} response.Response{data=nil,error=string} "Internal server error"
// @Router /users/{user_id}/profile [post]
func createProfile(client profiles.ProfileServiceClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		var reqBody Profile
		userID := c.Param("user_id")
		authID := c.Get(middleware.AuthIDKey).(string)
		if authID != userID {
			return c.JSON(http.StatusForbidden, response.Error("Access denied"))
		}
		err := c.Bind(&reqBody)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.Error(err.Error()))
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
			return response.MapError(c, err)
		}
		return c.JSON(http.StatusCreated, response.Success(NewProfile(p)))
	}
}

// @Summary Get user profile
// @Description 'This can only be done by the logged in user.'
// @Tags profiles
// @ID getUserProfileById
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <access_token>" "Authorization header"
// @Param user_id path string true "User id"
// @Success 200 {object} response.Response{data=Profile,error=nil} "Profile found"
// @Failure 500 {object} response.Response{data=nil,error=string} "Internal server error"
// @Router /users/{user_id}/profile [get]
func getProfile(client profiles.ProfileServiceClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Param("user_id")
		p, err := client.GetProfile(
			c.Request().Context(),
			&profiles.GetProfileRequest{Id: userID},
		)

		if err != nil {
			return response.MapError(c, err)
		}
		return c.JSON(http.StatusOK, response.Success(NewProfile(p)))
	}
}

// @Summary Update user profile
// @Description Should be replaced with a PATCH request.
// @Tags profiles
// @ID updateProfileByUserId
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <access_token>" "Authorization header"
// @Param user_id path string true "User ID"
// @Param body body Profile true "Update user profile"
// @Success 200 {object} response.Response{data=Profile,error=nil} "Profile Updated"
// @Failure 500 {object} response.Response{data=nil,error=string} "Internal server error"
// @Router /users/{user_id}/profile [put]
func updateProfile(client profiles.ProfileServiceClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		var reqBody Profile
		userID := c.Param("user_id")
		authID := c.Get(middleware.AuthIDKey).(string)
		if authID != userID {
			return c.JSON(http.StatusForbidden, response.Error("Access denied"))
		}
		err := c.Bind(&reqBody)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.Error(err.Error()))
		}

		p, err := client.UpdateProfile(
			c.Request().Context(),
			&profiles.UpdateProfileRequest{
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
			return response.MapError(c, err)
		}
		return c.JSON(http.StatusOK, response.Success(NewProfile(p)))
	}
}

// @Summary Get full user profile
// @Description 'This can only be done by the logged in user.'
// @Tags profiles
// @ID getFullProfileByUserId
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <access_token>" "Authorization header"
// @Param user_id path string true "User id"
// @Success 200 {object} response.Response{data=FullProfile,error=nil} "Full profile found"'
// @Failure 500 {object} response.Response{data=nil,error=string} "Internal server error"
// @Router /users/{user_id}/profile/full [get]
func getFullProfile(client profiles.ProfileServiceClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Param("user_id")
		p, err := client.GetFullProfile(
			c.Request().Context(),
			&profiles.GetProfileRequest{Id: userID},
		)

		if err != nil {
			return response.MapError(c, err)
		}
		return c.JSON(http.StatusOK, response.Success(NewFullProfile(p)))
	}
}

// @Summary Get random profile recommendation
// @Description 'This can only be done by the logged in user.'
// @Tags profiles
// @ID getRecommendationByUserId
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <access_token>" "Authorization header"
// @Param user_id path string true "User id"
// @Success 200 {object} response.Response{data=FullProfile,error=nil} "Recommendation found"
// @Failure 500 {object} response.Response{data=nil,error=string} "Internal server error"
// @Router /users/{user_id}/profile/recommendation [get]
func getRecommendation(client profiles.ProfileServiceClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Param("user_id")
		authID := c.Get(middleware.AuthIDKey).(string)
		if authID != userID {
			return c.JSON(http.StatusForbidden, response.Error("Access denied"))
		}
		p, err := client.GetRandomProfilePreferredByUser(
			c.Request().Context(),
			&profiles.GetRandomProfilePreferredByUserRequest{UserId: userID},
		)

		if err != nil {
			return response.MapError(c, err)
		}
		return c.JSON(http.StatusOK, response.Success(NewFullProfile(p)))
	}
}

// @Summary Get user prompts
// @Description 'This can only be done by the logged in user.'
// @Tags prompts
// @ID getPromptsByUserId
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <access_token>" "Authorization header"
// @Param user_id path string true "User id"
// @Success 200 {object} response.Response{data=[]Prompt,error=nil} "Prompts found"
// @Failure 500 {object} response.Response{data=nil,error=string} "Internal server error"
// @Router /users/{user_id}/prompts [get]
func getPrompts(client profiles.ProfileServiceClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Param("user_id")
		prompts, err := client.GetPrompts(
			c.Request().Context(),
			&profiles.GetPromptsRequest{UserId: userID},
		)

		if err != nil {
			return response.MapError(c, err)
		}
		return c.JSON(http.StatusOK, response.Success(Prompts(prompts)))
	}
}

// @Summary Create text prompts
// @Description 'This can only be done by the logged in user.'
// @Tags prompts
// @ID createPromptsByUserId
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <access_token>" "Authorization header"
// @Param user_id path string true "User id"
// @Param body body Prompt true "Create user text prompt"
// @Success 201 {object} response.Response{data=Prompt,error=nil} "Prompts created"
// @Failure 500 {object} response.Response{data=nil,error=string} "Internal server error"
// @Router /users/{user_id}/prompts/text [post]
func createPrompt(client profiles.ProfileServiceClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		var reqBody Prompt
		userID := c.Param("user_id")
		authID := c.Get(middleware.AuthIDKey).(string)
		if authID != userID {
			return c.JSON(http.StatusForbidden, response.Error("Access denied"))
		}
		err := c.Bind(&reqBody)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.Error(err.Error()))
		}

		prompts, err := client.AddPrompts(
			c.Request().Context(),
			&profiles.AddPromptsRequest{
				UserId:  userID,
				Prompts: mapPrompts([]Prompt{reqBody}),
			},
		)
		if err != nil {
			return response.MapError(c, err)
		}
		return c.JSON(http.StatusCreated, response.Success(NewPrompt(prompts.Prompts[0])))
	}
}

// @Summary Update text prompt
// @Description 'This can only be done by the logged in user.'
// @Tags prompts
// @ID updatePromptByUserId
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <access_token>" "Authorization header"
// @Param user_id path string true "User id"
// @Param prompt_id path string true "Prompt id"
// @Param body body Prompt true "Update user prompt"
// @Success 200 {object} response.Response{data=Prompt,error=nil} "Prompt updated"
// @Failure 500 {object} response.Response{data=nil,error=string} "Internal server error"
// @Router /users/{user_id}/prompts/text/{prompt_id} [put]
func updatePrompt(client profiles.ProfileServiceClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		var reqBody Prompt
		userID := c.Param("user_id")
		authID := c.Get(middleware.AuthIDKey).(string)
		if authID != userID {
			return c.JSON(http.StatusForbidden, response.Error("Access denied"))
		}
		promptID := c.Param("prompt_id")
		err := c.Bind(&reqBody)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.Error(err.Error()))
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
					Type:     "text",
				},
			},
		)

		if err != nil {
			return response.MapError(c, err)
		}
		return c.JSON(http.StatusOK, response.Success(NewPrompt(prompt.Prompt)))
	}
}

// @Summary Update file prompt
// @Description 'This can only be done by the logged-in user.'
// @Tags prompts
// @ID updateFilePromptByUserId
// @Accept multipart/form-data
// @Produce json
// @Param Authorization header string true "Bearer <access_token>" "Authorization header"
// @Param user_id path string true "User id"
// @Param prompt_id path string true "Prompt id"
// @Param question formData string true "Prompt question"
// @Param type formData string true "Prompt type"
// @Param file formData file true "Prompt file"
// @Success 201 {object} response.Response{data=Prompt,error=nil} "File prompt created"
// @Failure 500 {object} response.Response{data=nil,error=string} "Internal server error"
// @Router /users/{user_id}/prompts/file/{prompt_id} [put]
func updateFilePrompt(client profiles.ProfileServiceClient) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var (
			buf bytes.Buffer
		)
		userID := c.Param("user_id")
		authID := c.Get(middleware.AuthIDKey).(string)
		if authID != userID {
			return c.JSON(http.StatusForbidden, response.Error("Access denied"))
		}
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
			return c.JSON(http.StatusBadRequest, response.Error(err.Error()))
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
			return response.MapError(c, err)
		}
		return c.JSON(http.StatusOK, response.Success(NewPrompt(prompt.Prompt)))
	}
}

// @Summary Create file prompt
// @Description 'This can only be done by the logged in user.'
// @Tags prompts
// @ID createFilePromptByUserId
// @Accept multipart/form-data
// @Produce json
// @Param Authorization header string true "Bearer <access_token>" "Authorization header"
// @Param user_id path string true "User id"
// @Param question formData string true "Prompt question"
// @Param type formData string true "Prompt type"
// @Param file formData file true "Prompt file"
// @Success 201 {object} response.Response{data=Prompt,error=nil} "File prompt created"
// @Failure 500 {object} response.Response{data=nil,error=string} "Internal server error"
// @Router /users/{user_id}/prompts/file [post]
func createFilePrompt(client profiles.ProfileServiceClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			buf bytes.Buffer
		)
		userID := c.Param("user_id")
		authID := c.Get(middleware.AuthIDKey).(string)
		if authID != userID {
			return c.JSON(http.StatusForbidden, response.Error("Access denied"))
		}
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
			return c.JSON(http.StatusBadRequest, response.Error(err.Error()))
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
			return response.MapError(c, err)
		}
		return c.JSON(http.StatusOK, response.Success(NewPrompt(prompt.Prompt)))
	}
}
