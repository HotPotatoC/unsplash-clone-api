package actions

import (
	"context"
	"fmt"

	"github.com/HotPotatoC/unsplash-clone/entity"
	"github.com/HotPotatoC/unsplash-clone/pkg/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// ListAllPhotosAction dependencies
type ListAllPhotosAction struct {
	ctx   context.Context
	mongo *database.MongoHandler
}

// NewListAllPhotosAction constructs a new list all photos action
func NewListAllPhotosAction(ctx context.Context, mongo *database.MongoHandler) ListAllPhotosAction {
	return ListAllPhotosAction{
		ctx:   ctx,
		mongo: mongo,
	}
}

// Execute creates the handler
func (a ListAllPhotosAction) Execute(c *fiber.Ctx) error {
	var photos []entity.Photo

	if err := a.mongo.FindAll(a.ctx, entity.PhotoCollectionName, bson.M{}, &photos); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "There was a problem on our side",
		})
	}

	fmt.Println(photos)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"total_items": len(photos),
		"items":       photos,
	})
}
