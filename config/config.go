package config

import (
	"database/sql"
	"github.com/go-redis/redis/v7"
	"github.com/joho/godotenv"
	"gofiber/pkg/jwe"
	"gofiber/pkg/jwt"
	postgresqlPkg "gofiber/pkg/postgresql"
	redisPkg "gofiber/pkg/redis"
	"gofiber/pkg/str"
	"log"
	"os"
	"time"
)

type Configs struct {
	DB          *sql.DB
	RedisClient redisPkg.RedisClient
	JweCred     jwe.Credential
	JwtCred     jwt.JwtCredential
}

func LoadConfigs() (res Configs, err error) {
	err = godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading ..env file")
	}

	//postgresql conn
	dbConn := postgresqlPkg.Connection{
		Host:                    os.Getenv("DB_HOST"),
		DbName:                  os.Getenv("DB_NAME"),
		User:                    os.Getenv("DB_USER"),
		Password:                os.Getenv("DB_PASSWORD"),
		Port:                    os.Getenv("DB_PORT"),
		SslMode:                 os.Getenv("DB_SSL_MODE"),
		DBMaxConnection:         str.StringToInt(os.Getenv("DB_MAX_CONNECTION")),
		DBMAxIdleConnection:     str.StringToInt(os.Getenv("DB_MAX_IDLE_CONNECTION")),
		DBMaxLifeTimeConnection: str.StringToInt(os.Getenv("DB_MAX_LIFETIME_CONNECTION")),
	}
	res.DB, err = dbConn.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	res.DB.SetMaxOpenConns(dbConn.DBMaxConnection)
	res.DB.SetMaxIdleConns(dbConn.DBMAxIdleConnection)
	res.DB.SetConnMaxLifetime(time.Duration(dbConn.DBMaxLifeTimeConnection) * time.Second)

	//redis conn
	redisOption := &redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	}
	res.RedisClient = redisPkg.RedisClient{Client: redis.NewClient(redisOption)}

	//jwe
	res.JweCred = jwe.Credential{
		KeyLocation: os.Getenv("PRIVATE_KEY"),
		Passphrase:  os.Getenv("PASSPHRASE"),
	}

	//jwt
	res.JwtCred = jwt.JwtCredential{
		TokenSecret:         os.Getenv("SECRET"),
		ExpiredToken:        str.StringToInt(os.Getenv("TOKEN_EXP_TIME")),
		RefreshTokenSecret:  os.Getenv("SECRET_REFRESH_TOKEN"),
		ExpiredRefreshToken: str.StringToInt(os.Getenv("REFRESH_TOKEN_EXP_TIME")),
	}

	return res,err
}
