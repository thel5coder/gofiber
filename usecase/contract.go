package usecase

import (
	"database/sql"
	"errors"
	"gofiber/pkg/aes"
	queue "gofiber/pkg/amqp"
	"gofiber/pkg/aws"
	"gofiber/pkg/fcm"
	"gofiber/pkg/jwe"
	"gofiber/pkg/jwt"
	"gofiber/pkg/logruslogger"
	"gofiber/pkg/mailing"
	"gofiber/pkg/mandrill"
	"gofiber/pkg/messages"
	"gofiber/pkg/pusher"
	twilioHelper "gofiber/pkg/twilio"
	"gofiber/usecase/viewmodel"
	"math/rand"
	"os"
	"strings"
	"time"

	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	jwtFiber "github.com/gofiber/jwt/v2"
	"github.com/streadway/amqp"
	"gofiber/pkg/redis"
)

const (
	defaultLimit   = 10
	maxLimit       = 50
	defaultSort    = "asc"
	PasswordLength  = 6
	defaultLastPage = 0
)

var (
	AmqpConnection *amqp.Connection
	AmqpChannel *amqp.Channel
)

// UcContract ...
type UcContract struct {
	ReqID        string
	UserID       string
	DB           *sql.DB
	TX           *sql.Tx
	AES          aes.Credential
	AmqpConn     *amqp.Connection
	AmqpChannel  *amqp.Channel
	RedisClient  redis.RedisClient
	JweCred      jwe.Credential
	Validate     *validator.Validate
	Translator   ut.Translator
	JwtCred      jwt.JwtCredential
	JwtConfig    jwtFiber.Config
	AWSS3        aws.AWSS3
	Pusher       pusher.Credential
	GoMailConfig mailing.GoMailConfig
	Fcm          fcm.Connection
	TwilioClient *twilioHelper.Client
	Mandrill     mandrill.Credential
}

func (uc UcContract) setPaginationParameter(page, limit int, orderBy, sort string, orderByWhiteLists []string) (int, int, int, string, string) {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 || limit > maxLimit {
		limit = defaultLimit
	}

	orderBy = uc.checkWhiteList(orderBy, orderByWhiteLists)

	if sort == "" {
		sort = defaultSort
	}
	offset := (page - 1) * limit

	return offset, limit, page, orderBy, sort
}

func (uc UcContract) checkWhiteList(orderBy string, whiteLists []string) string {
	for _, whiteList := range whiteLists {
		if orderBy == whiteList {
			return orderBy
		}
	}

	return whiteLists[0]
}

func (uc UcContract) setPaginationResponse(page, limit, total int) (paginationResponse viewmodel.PaginationVm) {
	var lastPage int

	if total > 0 {
		lastPage = total / limit

		if total%limit != 0 {
			lastPage = lastPage + 1
		}
	} else {
		lastPage = defaultLastPage
	}

	paginationResponse = viewmodel.PaginationVm{
		CurrentPage: page,
		LastPage:    lastPage,
		Total:       total,
		PerPage:     limit,
	}

	return paginationResponse
}

// GetRandomString ...
func (uc UcContract) GetRandomString(length int) string {
	if length == 0 {
		length = PasswordLength
	}

	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
		"abcdefghijklmnopqrstuvwxyzåäö" +
		"0123456789")

	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	password := b.String()

	return password
}

// LimitRetryByKey ...
func (uc UcContract) LimitRetryByKey(key string, limit float64) (err error) {
	var count float64
	res := map[string]interface{}{}

	err = uc.RedisClient.GetFromRedis(key, &res)
	if err != nil {
		err = nil
		res = map[string]interface{}{
			"counter": count,
		}
	}
	count = res["counter"].(float64) + 1
	if count > limit {
		uc.RedisClient.RemoveFromRedis(key)

		return errors.New(messages.MaxRetryKey)
	}

	res["counter"] = count
	uc.RedisClient.StoreToRedistWithExpired(key, res, "24h")

	return err
}

// PushToQueue ...
func (uc UcContract) PushToQueue(queueBody map[string]interface{}, queueType, deadLetterType string) (err error) {
	mqueue := queue.NewQueue(AmqpConnection, AmqpChannel)

	_, _, err = mqueue.PushQueueReconnect(os.Getenv("AMQP_URL"), queueBody, queueType, deadLetterType)
	if err != nil {
		return err
	}

	return err
}

func (uc UcContract) ErrorHandler(ctx, scope string, err error) error {
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, scope, uc.ReqID)
		return err
	}

	return nil
}
