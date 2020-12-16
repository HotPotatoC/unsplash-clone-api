package actions

import (
	"context"

	"github.com/HotPotatoC/unsplash-clone/entity"
	"github.com/HotPotatoC/unsplash-clone/pkg/database"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/labstack/gommon/log"
)

// CreateNewImageAction dependencies
type CreateNewImageAction struct {
	ctx   context.Context
	mongo *database.MongoHandler
}

type createNewImageInput struct {
	Label    string `json:"label" validate:"required,max=128"`
	URL      string `json:"url" validate:"required,url"`
	Password string `json:"password" validate:"omitempty,min=6"`
}

type createNewImageErrorOutput struct {
	FailedField string
	Tag         string
	Value       string
}

// NewCreateNewImageAction constructs a new create image action
func NewCreateNewImageAction(ctx context.Context, mongo *database.MongoHandler) CreateNewImageAction {
	return CreateNewImageAction{
		ctx:   ctx,
		mongo: mongo,
	}
}

// Execute creates the handler
func (a CreateNewImageAction) Execute(c *fiber.Ctx) error {
	var input createNewImageInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if errors := a.validateInput(&input); errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}

	if err := a.mongo.Store(a.ctx, entity.ImageCollectionName, input); err != nil {
		log.Info(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "There was a problem on our side",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Successfully created a new photo",
	})
}

func (a CreateNewImageAction) validateInput(input *createNewImageInput) []*createNewImageErrorOutput {
	var errors []*createNewImageErrorOutput
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element createNewImageErrorOutput
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
