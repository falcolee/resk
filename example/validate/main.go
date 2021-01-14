package main

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	vtzh "gopkg.in/go-playground/validator.v9/translations/zh"
)
import "gopkg.in/go-playground/validator.v9"

type User struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Email     string `validate:"required,email"`
}

func main() {
	user := &User{
		FirstName: "firstName",
		LastName:  "lastName",
		Email:     "aaafl163.com",
	}
	validate := validator.New()
	cn := zh.New()
	uni := ut.New(cn, cn)
	translator, found := uni.GetTranslator("zh")
	if found {
		err := vtzh.RegisterDefaultTranslations(validate, translator)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("not found translator")
	}
	err := validate.Struct(user)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			for _, err := range errs {
				fmt.Println(err.Translate(translator))
			}
		}
	}
}
