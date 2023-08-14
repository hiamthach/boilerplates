package config

import (
	"log"
	"os"

	"github.com/golobby/dotenv"
)

var app *EnvConfig

type EnvConfig struct {
	Env string `env:"ENV"`
	App struct {
		Name    string `env:"APP_NAME"`
		Version string `env:"APP_VERSION"`
	}
	Redis struct {
		Primary         string `env:"REDIS_PRIMARY"`
		PrimaryPass     string `env:"REDIS_PRIMARY_PASS"`
		PrimaryDBNum    int    `env:"REDIS_PRIMARY_DB_NUM"`
		Reader          string `env:"REDIS_READER"`
		ReaderPass      string `env:"REDIS_READER_PASS"`
		ReaderDBNum     int    `env:"REDIS_READER_DB_NUM"`
		ServerForceSS   string `env:"REDIS_ADDRESS_FORCE_SESSION"`
		PasswordForceSS string `env:"REDIS_PASS_FORCE_SESSION"`
		DbNumberForceSS int    `env:"REDIS_DB_NUM_FORCE_SESSION"`
	}
	Log struct {
		Path        string `env:"LOG_PATH"`
		MinLevel    uint32 `env:"LOG_MIN_LEVEL"`
		Type        string `env:"LOG_TYPE"`
		LimitMsgLen uint32 `env:"LOG_LIMIT_MGS_LEN"`
	}
	MongoDb struct {
		Connection string `env:"MONGO_DB_CON"`
		DbName     string `env:"MONGO_DB_NAME"`
	}
	Server struct {
		HTTPServerUrl string `env:"HTTP_URL"`
		WSServerURL   string `env:"WS_URL"`
		GRPCServerUrl string `env:"GRPC_URL"`
		Cors          string `env:"HTTP_CORS"`
		CorsHeader    string `env:"HTTP_CORS_HEADER"`
		CorsMethod    string `env:"HTTP_CORS_METHOD"`
	}
	Auth struct {
		Issuer                       string `env:"AUTH_ISSUER"`
		PublicKeyHex                 string `env:"AUTH_TOKEN_PUBLIC_KEY_HEX"`
		PrivateKeyHex                string `env:"AUTH_TOKEN_PRIVATE_KEY_HEX"`
		AccessTokenDurationInMinute  int64  `env:"AUTH_TOKEN_DURATION"`
		RefreshTokenDurationInMinute int64  `env:"AUTH_TOKEN_REFRESH_DURATION"`
		AllowWrongPassTime           int    `env:"AUTH_WRONGPASS_TIME"`
		AllowWrongPassInMinute       int64  `env:"AUTH_WRONGPASS_IN_MIN"`
		AuthThirdPartyRequestTimeOut int64  `env:"AUTH_THIRD_PARTY_REQ_TIME"`
		SymmetricKey                 string `env:"AUTH_TOKEN_SYMMETRICKEY_KEY"`
		ConfirmCodeTimeOut           uint32 `env:"AUTH_CONFIRM_CODE_TIMEOUT"`
	}
	PubNub struct {
		PublishKey   string `env:"PUBNUB_PUB_KEY"`
		SubscribeKey string `env:"PUBNUB_SUB_KEY"`
		SecretKey    string `env:"PUBNUB_SECRET_KEY"`
	}
}

func init() {
	configApp()
	log.Println("app config loaded")
}

func Get() *EnvConfig {
	if app == nil {
		configApp()
	}
	return app
}

func configApp() {
	file, err := os.Open("app.env")
	if err != nil {
		log.Println("error when load app config")
	} else {
		app = &EnvConfig{}
		errLoadCf := dotenv.NewDecoder(file).Decode(app)
		if errLoadCf != nil {
			log.Println("error when load app config ", errLoadCf)
		}
	}
}
