package main

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	idTranslations "github.com/go-playground/validator/v10/translations/id"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/rs/xid"
	conf "gofiber/config"
	"gofiber/server/bootstrap"
	"gofiber/usecase"
	"log"
	"os"
)

var (
	validatorDriver *validator.Validate
	uni             *ut.UniversalTranslator
	translator      ut.Translator
)

func main() {
	//load all config
	configs, err := conf.LoadConfigs()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer configs.DB.Close()

	//init validation driver
	validatorInit()

	//init fiber app
	app := fiber.New()
	app.Use(requestid.New())
	app.Use(cors.New())

	ucContract := usecase.UcContract{
		ReqID:       xid.New().String(),
		DB:          configs.DB,
		RedisClient: configs.RedisClient,
		JweCred:     configs.JweCred,
		Validate:    validatorDriver,
		Translator:  translator,
		JwtCred:     configs.JwtCred,
	}

	boot := bootstrap.Bootstrap{
		App:        app,
		DB:         configs.DB,
		UcContract: ucContract,
		JweCred:    configs.JweCred,
		Validator:  validatorDriver,
		Translator: translator,
		JwtCred:    configs.JwtCred,
	}
	boot.App.Use(requestid.New())
	boot.App.Use(cors.New())
	boot.App.Use(logger.New())
	boot.RegisterRouters()
	log.Fatal(boot.App.Listen(os.Getenv("APP_HOST_SERVER")))
}

func validatorInit() {
	en := en.New()
	id := id.New()
	uni = ut.New(en, id)

	transEN, _ := uni.GetTranslator("en")
	transID, _ := uni.GetTranslator("id")

	validatorDriver = validator.New()

	enTranslations.RegisterDefaultTranslations(validatorDriver, transEN)
	idTranslations.RegisterDefaultTranslations(validatorDriver, transID)

	switch os.Getenv("APP_LOCALE") {
	case "en":
		translator = transEN
	case "id":
		translator = transID
	}
}
