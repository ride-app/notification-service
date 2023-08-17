package config

type EnvStruct struct {
	Production            bool   `env:"PRODUCTION" env-description:"dev or prod" env-default:"true"`
	LogDebug              bool   `env:"LOG_DEBUG" env-description:"should log at debug level" env-default:"false"`
	Port                  int32  `env:"PORT" env-description:"server port" env-default:"50051"`
	Project_Id            string `env:"PROJECT_ID" env-description:"firebase project id" env-default:"NO_PROJECT"`
	Firebase_Database_url string `env:"FIREBASE_DATABASE_URL" env-description:"firebase database url" env-default:"NO_DATABASE_URL"`
}

var Env EnvStruct
