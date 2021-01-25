/**
 * @Author: hqd
 * @Description: translations
 * @File:  translations
 * @Version: 1.0.0
 * @Date: 2021/1/25 18:27
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		uni := ut.New(zh.New(), en.New())
		locale := c.GetHeader("locale")
		trans, _ := uni.GetTranslator(locale)
		validate, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			switch locale {
			case "zh":
				zh_translations.RegisterDefaultTranslations(validate, trans)
			case "en":
				en_translations.RegisterDefaultTranslations(validate, trans)
			default:
				zh_translations.RegisterDefaultTranslations(validate, trans)
			}
			c.Set("trans", trans)
		}
		c.Next()
	}
}
