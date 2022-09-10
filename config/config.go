package config

type EnvStruct struct {
	Debug               bool   `env:"DEBUG" env-description:"dev or prod" env-default:"false"`
	Port                int32  `env:"PORT" env-description:"server port" env-default:"50051"`
	Firebase_Project_Id string `env:"FIREBASE_PROJECT_ID" env-description:"firebase project id" env-default:"NO_PROJECT"`
}

var Env EnvStruct
