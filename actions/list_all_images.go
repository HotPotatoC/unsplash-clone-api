package actions

import (
	"context"

	"github.com/HotPotatoC/unsplash-clone/entity"
	"github.com/HotPotatoC/unsplash-clone/pkg/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// ListAllImagesAction dependencies
type ListAllImagesAction struct {
	ctx   context.Context
	mongo *database.MongoHandler
}

type listAllImagesOutput struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	URL   string `json:"url"`
}

// NewListAllImagesAction constructs a new list all photos action
func NewListAllImagesAction(ctx context.Context, mongo *database.MongoHandler) ListAllImagesAction {
	return ListAllImagesAction{
		ctx:   ctx,
		mongo: mongo,
	}
}

// Execute creates the handler
func (a ListAllImagesAction) Execute(c *fiber.Ctx) error {
	var images []entity.Image

	if err := a.mongo.FindAll(a.ctx, entity.ImageCollectionName, bson.M{}, &images); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "There was a problem on our side",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"total_items": len(images),
		"items":       a.output(images),
	})
}

// output handles the display of the data
func (a ListAllImagesAction) output(images []entity.Image) []listAllImagesOutput {
	output := make([]listAllImagesOutput, 0)

	for _, image := range images {
		output = append(output, listAllImagesOutput{
			ID:    image.ID.Hex(),
			Label: image.Label,
			URL:   image.URL,
		})
	}

	return output
}
