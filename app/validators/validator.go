package validators

import (
	"golang-fiber-base-project/app/exceptions"
	"log"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// type ValidatorInterface interface {
// 	registerCustomRules()
// }

type Validator struct {
	v  *validator.Validate
	DB *gorm.DB
}

func NewValidator(DB *gorm.DB) *Validator {
	validator := &Validator{
		v:  validator.New(),
		DB: DB,
	}
	validator.registerCustomRules()

	return validator
}

func (val *Validator) Struct(s any) error {
	return val.v.Struct(s)
}

func ValidateBody[T any](v *Validator, ctx *fiber.Ctx) (*T, error) {
	var body T

	var err error
	if err = ctx.BodyParser(&body); err != nil {
		return nil, exceptions.NewBadRequestException()
	}

	if err = v.Struct(body); err != nil {
		return nil, exceptions.NewValidationException(v.FormatValidationErrors(err, body))
	}

	return &body, nil
}

// FormatValidationErrors converts validator errors into a readable map
func (val *Validator) FormatValidationErrors(err error, obj any) map[string][]string {
	out := make(map[string][]string)
	if err == nil {
		return out
	}

	ves, ok := err.(validator.ValidationErrors)
	if !ok {
		out["error"] = []string{err.Error()}
		return out
	}

	t := reflectType(obj)

	for _, fe := range ves {
		jsonName := jsonFieldName(t, fe.StructField())
		msg := val.translateError(fe)
		out[jsonName] = append(out[jsonName], msg)
	}

	return out
}

func (val *Validator) registerCustomRules() {
	var err error

	// register custom rules
	if err = val.v.RegisterValidation("unique", uniqueValidator(val.DB)); err != nil {
		log.Panicf("unique validator error: %v", err)
	}
	if err = val.v.RegisterValidation("exists", existsValidator(val.DB)); err != nil {
		log.Panicf("unique validator error: %v", err)
	}
}

func (val *Validator) translateError(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "is required"
	case "email":
		return "must be a valid email"
	case "min":
		return "minimum " + fe.Param() + " characters"
	case "unique":
		return "must be unique"
	case "exists":
		return "must exist"
	default:
		return fe.Error()
	}
}

// Reflection helpers
func reflectType(obj any) reflect.Type {
	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

func jsonFieldName(t reflect.Type, fieldName string) string {
	jsonName := strings.ToLower(fieldName)
	if f, found := t.FieldByName(fieldName); found {
		tag := f.Tag.Get("json")
		if tag != "" {
			jsonName = strings.Split(tag, ",")[0]
		}
	}
	return jsonName
}
