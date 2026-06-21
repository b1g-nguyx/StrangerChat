package auth

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/b1g-nguyx/strangerchat-backend/internal/usecase/user"
)

// Initialize the Validator instance (Can be moved to a shared config file later)
var validate = validator.New()

type authRoutes struct {
	userUseCase *user.UseCase
}

// NewAuthRoutes sets up Authentication-related APIs
func NewAuthRoutes(handler fiber.Router, uc *user.UseCase) {
	r := &authRoutes{userUseCase: uc}

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
// @Success      201      {object}  AuthResponse
// @Failure      400      {object}  ErrorResponse
// @Router       /auth/register [post]
func (r *authRoutes) register(c *fiber.Ctx) error {
	var req RegisterRequest

	// 1. Parse JSON body into the Request DTO
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Invalid JSON format",
		})
	}

	// 2. Validate data based on struct tags (Prevents SQLi, invalid emails, XSS, etc.)
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Validation failed: " + err.Error(),
		})
	}

	// 3. Call Usecase (Data is 100% clean at this point)
	newUser, accessToken, refreshToken, err := r.userUseCase.Register(c.Context(), req.Username, req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	// 4. Return standard Response DTO
	return c.Status(fiber.StatusCreated).JSON(AuthResponse{
		Message:      "Registered successfully",
		Data:         &newUser,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

// login godoc
// @Summary      Login method
// @Description  Login in system with username password
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request  body      LoginRequest  true  "Login Payload"
// @Success      200      {object}  AuthResponse
// @Failure      400      {object}  ErrorResponse
// @Router       /auth/login [post]
func (r *authRoutes) login(c *fiber.Ctx) error {
	var req LoginRequest
	// 1. Parse JSON body into the Request DTO
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Invalid JSON format",
		})
	}

	// 2. Validate data based on struct tags (Prevents SQLi, invalid emails, XSS, etc.)
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Validation failed: " + err.Error(),
		})
	}

	// 3. Call Usecase (Data is 100% clean at this point)
	user, accessToken, refreshToken, err := r.userUseCase.Login(c.Context(), req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	// 4. Return standard Response DTO
	return c.Status(fiber.StatusCreated).JSON(AuthResponse{
		Message:      "Login successfully",
		Data:         &user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
