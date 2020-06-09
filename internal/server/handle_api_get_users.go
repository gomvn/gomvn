package server

import (
	"time"

	"github.com/gofiber/fiber"

	"github.com/gomvn/gomvn/internal/entity"
)

func (s *Server) handleApiGetUsers(c *fiber.Ctx) {
	limit := getQueryUint64(c, "limit", 50)
	offset := getQueryUint64(c, "offset", 0)

	users, count, err := s.us.GetAll(limit, offset)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		return
	}

	c.JSON(&apiGetUsersResponse{
		Total: count,
		Items: mapToApiGetUsersItem(users),
	})
}

func mapToApiGetUsersItem(users []entity.User) []apiGetUsersItem {
	items := make([]apiGetUsersItem, len(users))
	for i, user := range users {
		items[i] = apiGetUsersItem{
			Id:        user.ID,
			Name:      user.Name,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Paths:     mapToApiGetUsersPathItem(user.Paths),
		}
	}
	return items
}

func mapToApiGetUsersPathItem(paths []entity.Path) []apiGetUsersPathItem {
	items := make([]apiGetUsersPathItem, len(paths))
	for i, path := range paths {
		items[i] = apiGetUsersPathItem{
			Path:      path.Path,
			Deploy:    path.Deploy,
			CreatedAt: path.CreatedAt,
			UpdatedAt: path.UpdatedAt,
		}
	}
	return items
}

type apiGetUsersResponse struct {
	Total uint64            `json:"total"`
	Items []apiGetUsersItem `json:"items"`
}

type apiGetUsersItem struct {
	Id        uint                  `json:"id"`
	Name      string                `json:"name"`
	CreatedAt time.Time             `json:"createdAt"`
	UpdatedAt time.Time             `json:"updatedAt"`
	Paths     []apiGetUsersPathItem `json:"allowed"`
}

type apiGetUsersPathItem struct {
	Path      string    `json:"name"`
	Deploy    bool      `json:"deploy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
