package actions

import (
	"context"
	"errors"

	"github.com/HotPotatoC/unsplash-clone/entity"
	"github.com/HotPotatoC/unsplash-clone/pkg/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// DeleteImageAction dependencies
type DeleteImageAction struct {
	ctx   context.Context
	mongo *database.MongoHandler
}

type deleteImageInput struct {
	Password string `json:"password,omitempty" validate:"omitempty"`
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
	var input deleteImageInput
	var image entity.Image

	imageID := c.Params("imageID")

	_ = c.BodyParser(&input)

	id, err := primitive.ObjectIDFromHex(imageID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "There was a problem on our side",
		})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "There was a problem on our side",
		})
	}

	err = a.mongo.FindOne(a.ctx, entity.ImageCollectionName, bson.M{
		"_id": id,
	}, &image)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Did not found image",
		})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "There was a problem on our side",
		})
	}

	if image.Password != "" && input.Password == "" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "A password is required to delete this image",
			"input":   input,
		})
	}

	if image.Password != "" && image.Password != input.Password {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Invalid password",
		})
	}

	_, err = a.mongo.Delete(a.ctx, entity.ImageCollectionName, bson.M{
		"_id": id,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "There was a problem on our side",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully deleted a image",
	})
}
