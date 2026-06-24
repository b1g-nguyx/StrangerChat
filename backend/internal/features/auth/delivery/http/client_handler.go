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
	}
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
		User:         ToUserDTO(newUser),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

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
		User:         ToUserDTO(user),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	// 5. Return standard Response
	return response.Success(c, fiber.StatusOK, "Login successfully", authData)
}
