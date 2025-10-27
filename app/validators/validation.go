package validators

// import (
// 	"fmt"
// 	"log"
// 	"reflect"
// 	"strings"

// 	"github.com/go-playground/validator/v10"
// 	"github.com/gofiber/fiber/v2"
// 	"gorm.io/gorm"
// )

// var validate *validator.Validate

// func InitValidator(db *gorm.DB) {
// 	validate = validator.New()

// 	// Use json tag name instead of field name in error messages
// 	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
// 		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]

// 		if name == "-" || name == "" {
// 			return field.Name
// 		}

// 		return name
// 	})

// 	var err error

// 	// register custom rules
// 	if err = validate.RegisterValidation("unique", uniqueValidator(db)); err != nil {
// 		log.Panicf("unique validator error: %v", err)
// 	}
// 	if err = validate.RegisterValidation("exists", existsValidator(db)); err != nil {
// 		log.Panicf("unique validator error: %v", err)
// 	}
// }

// func ValidateStruct(s any) map[string]string {
// 	err := validate.Struct(s)
// 	if err == nil {
// 		return nil
// 	}

// 	errors := make(map[string]string)
// 	for _, er := range err.(validator.ValidationErrors) {
// 		errors[er.Field()] = getErrorMessage(er)
// 		// field := er.Field()
// 		// errors[field] = append(errors[field], getErrorMessage(e))
// 	}

// 	return errors
// }

// func getErrorMessage(err validator.FieldError) string {
// 	switch err.Tag() {
// 	case "required":
// 		return fmt.Sprintf("%s is required", err.Field())
// 	case "email":
// 		return fmt.Sprintf("%s must be a valid email address", err.Field())
// 	case "min":
// 		return fmt.Sprintf("%s must be at least %s characters long", err.Field(), err.Param())
// 	case "boolean":
// 		return fmt.Sprintf("%s must be a boolean value", err.Field())
// 	case "unique":
// 		return fmt.Sprintf("%s must be unique", err.Field())
// 	case "exists":
// 		return fmt.Sprintf("%s does not exist", err.Field())
// 	default:
// 		return fmt.Sprintf("%s is not valid", err.Field())
// 	}

// }

// func ValidateBody[T any](c *fiber.Ctx) (*T, any) {
// 	var body T

// 	if err := c.BodyParser(&body); err != nil {
// 		return nil, err
// 		// return nil, c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
// 		// 	"error": "Invalid request body",
// 		// })
// 	}

// 	if errs := ValidateStruct(body); errs != nil {
// 		return nil, errs
// 		// return nil, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 		// 	"error": errs,
// 		// })
// 	}

// 	return &body, nil
// }
