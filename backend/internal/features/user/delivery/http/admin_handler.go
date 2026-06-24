package httpdelivery

import (
	"github.com/b1g-nguyx/strangerchat-backend/internal/common/response"
	"github.com/b1g-nguyx/strangerchat-backend/internal/features/user/usecase"
	"github.com/gofiber/fiber/v2"
)

type adminRoutes struct {
	userUseCase *user.UseCase
}

// NewAdminRoutes sets up Admin-related User APIs
func NewAdminRoutes(handler fiber.Router, uc *user.UseCase) {
	r := &adminRoutes{userUseCase: uc}

	h := handler.Group("/admin/users")
	{
		h.Get("/", r.getUsers)
	}
}

// getUsers godoc
// @Summary      Get list of users
// @Description  Get users with filtering (Admin only)
// @Tags         Admin User
// @Accept       json
// @Produce      json
// @Success      200      {object}  response.APIResponse
// @Failure      500      {object}  response.APIResponse
// @Router       /admin/users [get]
func (r *adminRoutes) getUsers(c *fiber.Ctx) error {
	// Parse queries to filters map
	filters := make(map[string]any)
	c.Context().QueryArgs().VisitAll(func(key, value []byte) {
		filters[string(key)] = string(value)
	})

	users, err := r.userUseCase.GetUsers(c.Context(), filters)
	if err != nil {
		return response.Error(c, err)
	}

	return response.Success(c, fiber.StatusOK, "Success", users)
}
