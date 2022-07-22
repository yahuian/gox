package validatex

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	enLocales "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

// https://github.com/go-playground/validator/tree/v9/_examples/gin-upgrading-overriding

type defaultValidator struct {
	validate   *validator.Validate
	translator ut.Translator
}

var instance *defaultValidator

var _ binding.StructValidator = &defaultValidator{}

func (v *defaultValidator) ValidateStruct(obj interface{}) error {
	if kindOfData(obj) == reflect.Struct {
		if err := v.validate.Struct(obj); err != nil {
			return errors.New(pretty(v.translator, err))
		}
	}
	return nil
}

func (v *defaultValidator) Engine() interface{} {
	return v.validate
}

func kindOfData(data interface{}) reflect.Kind {
	value := reflect.ValueOf(data)
	valueType := value.Kind()
	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}

type option struct {
	Gin bool // whether in gin framework
}

type WithOption func(opt *option)

func WithGin() WithOption {
	return func(opt *option) {
		opt.Gin = true
	}
}

func Init(opts ...WithOption) error {
	opt := new(option)
	for _, f := range opts {
		f(opt)
	}

	v := new(defaultValidator)
	v.validate = validator.New()

	// register trans
	english := enLocales.New()
	uni := ut.New(english, english)
	v.translator, _ = uni.GetTranslator("en")
	if err := enTranslations.RegisterDefaultTranslations(v.validate, v.translator); err != nil {
		return fmt.Errorf("register translation err: %w", err)
	}

	// replace gin default validate
	if opt.Gin {
		v.validate.SetTagName("validate")
		binding.Validator = v
	}

	instance = v

	return nil
}

func Struct(s any) error {
	err := instance.validate.Struct(s)
	if err != nil {
		return errors.New(pretty(instance.translator, err))
	}
	return nil
}

func Val(s any, name, tag string) error {
	err := instance.validate.Var(s, tag)
	if err != nil {
		return errors.New(name + pretty(instance.translator, err))
	}
	return nil
}

func pretty(t ut.Translator, err error) string {
	errorsMap, ok := err.(validator.ValidationErrors) //nolint:errorlint
	if !ok {
		return err.Error()
	}
	result := make([]string, 0, len(errorsMap))
	for _, e := range errorsMap {
		result = append(result, e.Translate(t))
	}
	return strings.Join(result, ", ")
}
