package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/HotPotatoC/unsplash-clone/actions"
	"github.com/HotPotatoC/unsplash-clone/pkg/database"
	"github.com/gofiber/fiber/v2"
)

// FiberEngine struct
type FiberEngine struct {
	ctx  context.Context
	App  *fiber.App
	db   *database.MongoHandler
	addr string
}

// NewFiberApp creates a new fiber app
func NewFiberApp(ctx context.Context, db *database.MongoHandler, addr string, config ...fiber.Config) *FiberEngine {
	return &FiberEngine{
		ctx:  ctx,
		App:  fiber.New(config...),
		db:   db,
		addr: addr,
	}
}

// Start runs the fiber server
func (f FiberEngine) Start() {
	if !fiber.IsChild() {
		fmt.Printf("[Master %d] Process started", os.Getppid())
		defer fmt.Printf("[Master %d] Exiting program...\n", os.Getppid())
	}

	go func() {
		if err := f.App.Listen(fmt.Sprintf("%s", f.addr)); err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
	_ = f.App.Shutdown()

	f.db.Disconnect(f.ctx)
}

// SetupHandlers registers the routes
func (f FiberEngine) SetupHandlers() {
	api := f.App.Group("/api")
	api.Get("/images", f.buildListAllImagesAction())
	api.Post("/images", f.buildCreateNewImage())
	api.Delete("/images/:imageID", f.buildDeletePhotoAction())
}

func (f FiberEngine) buildListAllImagesAction() fiber.Handler {
	return func(c *fiber.Ctx) error {
		action := actions.NewListAllImagesAction(f.ctx, f.db)

		return action.Execute(c)
	}
}

func (f FiberEngine) buildCreateNewImage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		action := actions.NewCreateNewImageAction(f.ctx, f.db)

		return action.Execute(c)
	}
}

func (f FiberEngine) buildDeletePhotoAction() fiber.Handler {
	return func(c *fiber.Ctx) error {
		action := actions.NewDeleteImageAction(f.ctx, f.db)

		return action.Execute(c)
	}
}
