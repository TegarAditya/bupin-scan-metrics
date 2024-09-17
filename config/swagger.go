package config

import "github.com/gofiber/contrib/swagger"

func SwaggerConfig() swagger.Config {
	return swagger.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.json",
		Path:     "/",
		Title:    "Swagger API Docs",
	}
}
