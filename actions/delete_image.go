package actions

import (
	"context"

	"github.com/HotPotatoC/unsplash-clone/entity"
	"github.com/HotPotatoC/unsplash-clone/pkg/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DeleteImageAction dependencies
type DeleteImageAction struct {
	ctx   context.Context
	mongo *database.MongoHandler
}

// NewDeleteImageAction constructs a new delete image action
func NewDeleteImageAction(ctx context.Context, mongo *database.MongoHandler) DeleteImageAction {
	return DeleteImageAction{
		ctx:   ctx,
		mongo: mongo,
	}
}

// Execute creates the handler
func (a DeleteImageAction) Execute(c *fiber.Ctx) error {
	imageID := c.Params("imageID")

	id, err := primitive.ObjectIDFromHex(imageID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "There was a problem on our side",
		})
	}

	count, err := a.mongo.Delete(a.ctx, entity.ImageCollectionName, bson.M{
		"_id": id,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "There was a problem on our side",
		})
	}

	if count < 1 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Did not found image",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully deleted a image",
	})
}
