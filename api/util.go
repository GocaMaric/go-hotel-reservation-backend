package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go-hotel-reservation-backend/types"
)

func getAuthUser(c *fiber.Ctx) (*types.User, error) {
	user, ok := c.Context().UserValue("user").(*types.User)
	if !ok {
		return nil, fmt.Errorf("unauthorized")
	}
	return user, nil
}
