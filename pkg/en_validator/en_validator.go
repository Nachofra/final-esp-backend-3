package en_validator

import (
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
	"strings"
)

// Validator is a struc for validations
type Validator struct {
	Validate   *validator.Validate
	Translator ut.Translator
}

// Get sets the validator used in the whole application (this is a bit hardcoded, but I don't have time)
// A lot of this could be parametrized in the future, like language, validator params in New() func and settings
// default translations accordingly
func Get() *Validator {

	// Initializing translation to english, this
	english := en.New()
	uni := ut.New(english, english)

	// In the future, this could be known or extracted from http 'Accept-Language' header
	trans, found := uni.GetTranslator("en")
	if found != true {
		panic("translation for validator not found")
	}

	v := validator.New(validator.WithRequiredStructEnabled())

	err := entranslations.RegisterDefaultTranslations(v, trans)
	if err != nil {
		panic(err)
	}

	return &Validator{
		Validate:   v,
		Translator: trans,
	}
}

func (v *Validator) Translate(errs validator.ValidationErrors) string {
	output := errs.Translate(v.Translator)

	var msg string
	for key, value := range output {
		msg += fmt.Sprintf("%s: %s - ", key, value)
	}
	msg = strings.Trim(msg, "- ")

	return msg
}
