package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)
//Config Struct
type Config struct {
	BASE_URL   string `mapstructure:"BASE_URL"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBPassword string `mapstructure:"DB_PASSWORD"`

	AUTHTOKEN   string `mapstructure:"TWILIO_AUTHTOKEN"`
	ACCOUNTSID  string `mapstructure:"TWILIO_ACCOUNTSID"`
	SERVICESSID string `mapstructure:"TWILIO_SEVICESID"`

	KEY       string `mapstructure:"KEY"`
	KEY_ADMIN string `mapstructure:"KEY_ADMIN"`

	KEY_ID_FOR_PAY     string `mapstructure:"KEY_ID_FOR_PAY"`
	SECRET_KEY_FOR_PAY string `mapstructure:"SECRET_KEY_FOR_PAY"`
}
//Creates a slice named envs containing the names of environment
// variables expected to be used in the configuration
var envs = []string{
	"BASE_URL", "DB_HOST", "DB_NAME", "DB_PORT", "DB_PASSWORD", "TWILIO_AUTHTOKEN", "TWILIO_ACCOUNTSID", "TWILIO_SERVICESID", "KEY", "KEY_ADMIN", "KEY_ID_FOR_PAY", "SECRET_KEY_FOR_PAY",
}

func LoadConfig() (Config, error) {
	//Declaration and Initialization
	var confg Config
	//Set configuration file and path
	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	//Binding environment vaiables
	for _,env := range envs{
		if err := viper.BindEnv(env); err != nil{
			return confg,err
		}
	}
	//Unmarshalling configuration
	if err := viper.Unmarshal(&confg);err !=nil{
		return confg,err
	}
	//Struct validation
	if err := validator.New().Struct(&confg);err != nil{
		return confg,err
	}
	//return config and error
	return confg,nil
}
