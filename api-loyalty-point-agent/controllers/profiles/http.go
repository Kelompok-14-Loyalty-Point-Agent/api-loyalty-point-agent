package profiles

import (
	"api-loyalty-point-agent/businesses/profiles"
	"api-loyalty-point-agent/controllers"
	"api-loyalty-point-agent/controllers/profiles/request"
	"api-loyalty-point-agent/controllers/profiles/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ProfileController represents the controller for user profiles.
type ProfileController struct {
	profileUsecase profiles.Usecase
}

// NewProfileController creates a new instance of ProfileController.
func NewProfileController(profileUC profiles.Usecase) *ProfileController {
	return &ProfileController{
		profileUsecase: profileUC,
	}
}

// GetAll retrieves all user profiles.
// @Summary Retrieves all user profiles
// @Description Retrieves all user profiles
// @Tags profiles
// @Accept json
// @Produce json
// @Security BearerToken
// @Success 200 {object} controllers.Response[[]response.Profile] "success"
// @Failure 400 {object} controllers.Response[string] "failed"
// @Router /profiles/profilesAll [get]
func (cc *ProfileController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	profilesData, err := cc.profileUsecase.GetAll(ctx)
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed tp fetch data", "")
	}

	profiles := make([]response.Profile, 0, len(profilesData))
	for _, profile := range profilesData {
		profiles = append(profiles, response.FromDomain(profile))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all profiles", profiles)
}

// GetByID retrieves a user profile by ID.
// @Summary Retrieves a user profile by ID
// @Description Retrieves a user profile with the given ID
// @Tags profiles
// @Accept json
// @Produce json
// @Security BearerToken
// @Param id path string true "Profile ID"
// @Success 200 {object} controllers.Response[response.Profile] "success"
// @Failure 400 {object} controllers.Response[string] "failed"
// @Router /profiles/{id} [get]
func (cc *ProfileController) GetByID(c echo.Context) error {
	var profileID string = c.Param("id")
	ctx := c.Request().Context()

	profile, err := cc.profileUsecase.GetByID(ctx, profileID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "profile not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "profile found", response.FromDomain(profile))
}

// Update updates the profile information and changes the user's password.
// @Summary Update the profile information and change the user's password
// @Description Update the profile information and change the user's password
// @Tags profiles
// @Accept json
// @Produce json
// @Security BearerToken
// @Param id path string true "Profile ID"
// @Param profile body request.UpdateProfileRequest true "Profile data"
// @Param password body request.ChangePasswordRequest true "New password data"
// @Success 200 {object} controllers.Response[response.Profile] "success"
// @Failure 400 {object} controllers.Response[string] "failed"
// @Router /profiles/{id} [put]
func (cc *ProfileController) Update(c echo.Context) error {
	var profileID string = c.Param("id")
	profileInput := request.UpdateProfileRequest{}
	passwordInput := request.ChangePasswordRequest{}
	ctx := c.Request().Context()

	if err := c.Bind(&profileInput); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid profile request", "")
	}

	if err := c.Bind(&passwordInput); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid password request", "")
	}

	if err := profileInput.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "profile validation failed", "")
	}

	if err := passwordInput.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "password validation failed", "")
	}

	profile, err := cc.profileUsecase.GetByID(ctx, profileID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", err.Error(), "")
	}

	// Update profile data
	profile.Name = profileInput.Name
	profile.Address = profileInput.Address

	// Update password
	profile.Password = passwordInput.Password

	updatedProfile, err := cc.profileUsecase.Update(ctx, profileInput.ToDomain(), profileID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "update profile failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "profile updated and password changed", response.FromDomain(updatedProfile))
}

// Delete deletes a user profile.
// @Summary Deletes a user profile
// @Description Deletes a user profile with the given ID
// @Tags profiles
// @Accept json
// @Produce json
// @Security BearerToken
// @Param id path string true "Profile ID"
// @Success 200 {object} controllers.Response[string] "success"
// @Failure 400 {object} controllers.Response[string] "failed"
// @Router /profiles/{id} [delete]
func (cc *ProfileController) Delete(c echo.Context) error {
	var profileID string = c.Param("id")
	ctx := c.Request().Context()

	err := cc.profileUsecase.Delete(ctx, profileID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "delete profile failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "profile deleted", "")
}
