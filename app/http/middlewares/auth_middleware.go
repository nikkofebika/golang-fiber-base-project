package middlewares

import (
	"fmt"
	"golang-fiber-base-project/app/exceptions"
	"golang-fiber-base-project/app/helpers"
	"golang-fiber-base-project/config"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(cfg *config.AppConfig) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		bearerToken := ctx.Get("Authorization")
		if bearerToken == "" {
			return exceptions.NewUnauthorizedException()
		}
		fmt.Println("bearerToken", bearerToken)

		tokens := strings.Split(bearerToken, " ")
		if len(tokens) != 2 || tokens[0] != "Bearer" {
			return exceptions.NewUnauthorizedException()
		}
		fmt.Println("tokens", tokens)

		jwtToken, err := helpers.ValidateToken(tokens[1], cfg.JWTSecret)
		if err != nil {
			return exceptions.NewUnauthorizedException()
		}
		fmt.Println("jwtToken asu", jwtToken)

		id, err := helpers.ExtractUserID(jwtToken)
		fmt.Println("OPO IKI", err)
		if err != nil {
			return exceptions.NewUnauthorizedException()
		}
		fmt.Println("id", id)

		ctx.Locals("user_id", id)
		return ctx.Next()
	}
}
