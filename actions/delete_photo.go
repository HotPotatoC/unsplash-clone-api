package actions

import (
	"context"

	"github.com/HotPotatoC/unsplash-clone/entity"
	"github.com/HotPotatoC/unsplash-clone/pkg/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DeletePhotoAction dependencies
type DeletePhotoAction struct {
	ctx   context.Context
	mongo *database.MongoHandler
}

// NewDeletePhotoAction constructs a new list all photos action
func NewDeletePhotoAction(ctx context.Context, mongo *database.MongoHandler) DeletePhotoAction {
	return DeletePhotoAction{
		ctx:   ctx,
		mongo: mongo,
	}
}

// Execute creates the handler
func (a DeletePhotoAction) Execute(c *fiber.Ctx) error {
	photoID := c.Params("photoID")

	id, err := primitive.ObjectIDFromHex(photoID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "There was a problem on our side",
		})
	}

	count, err := a.mongo.Delete(a.ctx, entity.PhotoCollectionName, bson.M{
		"_id": id,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "There was a problem on our side",
		})
	}

	if count < 1 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Did not found photo",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully deleted a photo",
	})
}
