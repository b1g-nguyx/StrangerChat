package httpdelivery

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	auth_usecase "github.com/b1g-nguyx/strangerchat-backend/internal/features/auth/usecase"
	"github.com/b1g-nguyx/strangerchat-backend/internal/common/apperror"
	"github.com/b1g-nguyx/strangerchat-backend/internal/common/response"
)

// Initialize the Validator instance (Can be moved to a shared config file later)
var validate = validator.New()

type clientRoutes struct {
	authUseCase *auth_usecase.AuthUseCase
}

// NewClientRoutes sets up Authentication APIs for regular users
func NewClientRoutes(handler fiber.Router, uc *auth_usecase.AuthUseCase) {
	r := &clientRoutes{authUseCase: uc}

	h := handler.Group("/auth")
	{
		h.Post("/register", r.register)
		h.Post("/login", r.login)
		h.Post("/refresh", r.refresh)
		h.Post("/logout", r.logout)
	}
}

func setRefreshTokenCookie(c *fiber.Ctx, token string) {
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    token,
		Path:     "/",
		MaxAge:   7 * 24 * 60 * 60, // 7 days in seconds
		Secure:   false,            // Set to true in production with HTTPS
		HTTPOnly: true,
		SameSite: "Lax",
	})
}

func clearRefreshTokenCookie(c *fiber.Ctx) {
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		Secure:   false,
		HTTPOnly: true,
		SameSite: "Lax",
	})
}

// register godoc
// @Summary      Register a new user
// @Description  Create a new user account with username, email, and password
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request  body      RegisterRequest  true  "Register Payload"
// @Success      201      {object}  response.APIResponse{data=AuthData}
// @Failure      400      {object}  response.APIResponse
// @Router       /auth/register [post]
func (r *clientRoutes) register(c *fiber.Ctx) error {
	var req RegisterRequest

	// 1. Parse JSON body into the Request struct
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, apperror.ErrInvalidJSON)
	}

	// 2. Validate data based on struct tags
	if err := validate.Struct(req); err != nil {
		return response.Error(c, apperror.ErrValidation)
	}

	// 3. Call Usecase
	newUser, accessToken, refreshToken, err := r.authUseCase.Register(c.Context(), req.Username, req.Email, req.Password)
	if err != nil {
		// Usecase should ideally return AppError too, but if it doesn't,
		// our response.Error will automatically fallback to 500.
		// If you want to force it to a specific error:
		// return response.Error(c, apperror.ErrEmailExists)
		return response.Error(c, err)
	}

	// 4. Map Entity to DTO and package it with tokens
	authData := AuthData{
		User:        ToUserDTO(newUser),
		AccessToken: accessToken,
	}

	// Set refresh token in HttpOnly Cookie
	setRefreshTokenCookie(c, refreshToken)

	// 5. Return standard Response
	return response.Success(c, fiber.StatusCreated, "Registered successfully", authData)
}

// login godoc
// @Summary      Login method
// @Description  Login in system with username password
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request  body      LoginRequest  true  "Login Payload"
// @Success      200      {object}  response.APIResponse{data=AuthData}
// @Failure      400      {object}  response.APIResponse
// @Failure      401      {object}  response.APIResponse
// @Router       /auth/login [post]
func (r *clientRoutes) login(c *fiber.Ctx) error {
	var req LoginRequest

	// 1. Parse JSON body into the Request struct
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, apperror.ErrInvalidJSON)
	}

	// 2. Validate data based on struct tags
	if err := validate.Struct(req); err != nil {
		return response.Error(c, apperror.ErrValidation)
	}

	// 3. Call Usecase
	user, accessToken, refreshToken, err := r.authUseCase.Login(c.Context(), req.Email, req.Password)
	if err != nil {
		return response.Error(c, apperror.ErrInvalidEmailPwd)
	}

	// 4. Map Entity to DTO and package it with tokens
	authData := AuthData{
		User:        ToUserDTO(user),
		AccessToken: accessToken,
	}

	// Set refresh token in HttpOnly Cookie
	setRefreshTokenCookie(c, refreshToken)

	// 5. Return standard Response
	return response.Success(c, fiber.StatusOK, "Login successfully", authData)
}

// refresh godoc
// @Summary      Refresh token
// @Description  Get a new access token using a valid refresh token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        Cookie   header    string  true  "refresh_token cookie"
// @Success      200      {object}  response.APIResponse{data=AuthData}
// @Failure      401      {object}  response.APIResponse
// @Router       /auth/refresh [post]
func (r *clientRoutes) refresh(c *fiber.Ctx) error {
	// 1. Get refresh token from cookie
	refreshToken := c.Cookies("refresh_token")
	if refreshToken == "" {
		return response.Error(c, apperror.ErrUnauthorized)
	}

	// 2. Call Usecase
	user, newAccessToken, newRefreshToken, err := r.authUseCase.Refresh(c.Context(), refreshToken)
	if err != nil {
		clearRefreshTokenCookie(c)
		return response.Error(c, apperror.ErrUnauthorized)
	}

	// 3. Map Entity to DTO and package it
	authData := AuthData{
		User:        ToUserDTO(user),
		AccessToken: newAccessToken,
	}

	// 4. Set new refresh token in HttpOnly Cookie
	setRefreshTokenCookie(c, newRefreshToken)

	// 5. Return standard Response
	return response.Success(c, fiber.StatusOK, "Token refreshed successfully", authData)
}

// logout godoc
// @Summary      Logout
// @Description  Logout user and clear refresh token cookie
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Success      200      {object}  response.APIResponse
// @Router       /auth/logout [post]
func (r *clientRoutes) logout(c *fiber.Ctx) error {
	clearRefreshTokenCookie(c)
	return response.Success(c, fiber.StatusOK, "Logout successfully", nil)
}
