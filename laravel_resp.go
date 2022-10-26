package laravel_resp

import (
	"errors"
	"strings"

	"github.com/abdullahPrasetio/validation_formatter_laravel_go/lang"
	"github.com/abdullahPrasetio/validation_formatter_laravel_go/utils"
	"github.com/go-playground/validator/v10"
)

func getErrorMsg(fe validator.FieldError, language string) string {
	lang.NewLang(language)
	messages := lang.LangMessage
	message := messages[fe.Tag()]
	if len(message) <= 0 {
		return "Unknown error"
	}

	return message
}

func messageReplacer(message string, fe validator.FieldError) string {
	newMessage := strings.Replace(message, ":attribute", fe.Field(), 1)
	newMessage = strings.Replace(newMessage, ":values", fe.Param(), 1)
	getParamValue := strings.Split(fe.Param(), " ")
	var value string
	var param string
	if len(getParamValue) > 1 {
		value = getParamValue[1]
		param = getParamValue[0]
	}
	newMessage = strings.Replace(newMessage, ":other", param, 1)
	newMessage = strings.Replace(newMessage, ":value", value, 1)
	return newMessage
}

func GetErrorMsgValidation(err error, language string) map[string]any {
	errorMessages := map[string]any{}
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, fe := range ve {
			var errorMsg = []string{messageReplacer(getErrorMsg(fe, language), fe)}
			errorMessages[utils.ToSnakeCase(fe.Field())] = errorMsg
		}
	}

	return errorMessages
}
