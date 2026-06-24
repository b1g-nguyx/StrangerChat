package httpdelivery

import (
	"github.com/gofiber/fiber/v2"

	"github.com/b1g-nguyx/strangerchat-backend/internal/common/apperror"
	"github.com/b1g-nguyx/strangerchat-backend/internal/common/response"
	auth_usecase "github.com/b1g-nguyx/strangerchat-backend/internal/features/auth/usecase"
)

type adminRoutes struct {
	authUseCase *auth_usecase.AuthUseCase
}

// NewAdminRoutes sets up Authentication APIs for Admin users
func NewAdminRoutes(handler fiber.Router, uc *auth_usecase.AuthUseCase) {
	r := &adminRoutes{authUseCase: uc}

	h := handler.Group("/admin/auth")
	{
		h.Post("/login", r.adminLogin)
	}
}

// adminLogin godoc
// @Summary      Admin Login method
// @Description  Login to admin dashboard (requires admin role)
// @Tags         Admin Auth
// @Accept       json
// @Produce      json
// @Param        request  body      AdminLoginRequest  true  "Admin Login Payload"
// @Success      200      {object}  response.APIResponse{data=AdminAuthData}
// @Failure      401      {object}  response.APIResponse
// @Router       /admin/auth/login [post]
func (r *adminRoutes) adminLogin(c *fiber.Ctx) error {
	var req AdminLoginRequest

	// 1. Parse JSON body
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, apperror.ErrInvalidJSON)
	}

	// 2. Validate
	if err := validate.Struct(req); err != nil {
		return response.Error(c, apperror.ErrValidation)
	}

	// 3. TODO: Implement r.authUseCase.AdminLogin() that checks role == "ADMIN"
	// user, token, _, err := r.authUseCase.AdminLogin(c.Context(), req.Email, req.Password, req.OTP)
	
	// For illustration purposes, returning a mock success
	mockData := AdminAuthData{
		AccessToken: "mock-admin-token",
		Role:        "SUPER_ADMIN",
	}

	return response.Success(c, fiber.StatusOK, "Admin login successful", mockData)
}
