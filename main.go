package main

import (
	"fmt"
	"time"

	"github.com/rajajamal/golang-framework/configs"
	"github.com/rajajamal/golang-framework/models"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
	"github.com/joho/godotenv"
	"golang.org/x/sync/singleflight"
)

var sFlight singleflight.Group
var key string = "course_singleflight"
var sharing = false

func init() {
	godotenv.Load()
	configs.Load()
	configs.Db.AutoMigrate(
		&models.User{},
	)
}

func main() {
	// listener, err := net.Listen("tcp", fmt.Sprintf(":%d", configs.Env.DB_PORT))
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	if configs.Env.Debug {
		app.Use(logger.New())
	}

	app.Use(recover.New())
	app.Use(compress.New()) // Content-Length minimal 200
	app.Use(helmet.New())

	app.Get("/metrics", monitor.New(monitor.Config{}))
	app.Get("/single-flight", func(c *fiber.Ctx) error {
		_, err, _ := sFlight.Do(key, func() (interface{}, error) {
			if !sharing {
				fmt.Println("Harusnya cuma sekali diprint")
				time.Sleep(17 * time.Second)
			}

			return true, nil
		})

		sharing = true
		if err != nil {
			return c.JSON(map[string]interface{}{
				"message": err.Error(),
			})
		}

		return c.JSON(map[string]interface{}{
			"message": sharing,
		})
	})

	// routes.RegisterUserRoutes(app)

	app.Listen(fmt.Sprintf(":%d", configs.Env.AppPort))
}
