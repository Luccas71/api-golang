package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var (
	cfg *conf
)

// variaveis de configuração da api
type conf struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`       //mapeando a struct para viper fazer a troca
	DBHost        string `mapstructure:"DB_HOST"`         //mapeando a struct para viper fazer a troca
	DBPort        string `mapstructure:"DB_PORT"`         //mapeando a struct para viper fazer a troca
	DBUser        string `mapstructure:"DB_USER"`         //mapeando a struct para viper fazer a troca
	DBPassword    string `mapstructure:"DB_PASSWORD"`     //mapeando a struct para viper fazer a troca
	DBName        string `mapstructure:"DB_NAME"`         //mapeando a struct para viper fazer a troca
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"` //mapeando a struct para viper fazer a troca
	JWTSecret     string `mapstructure:"JWT_SECRET"`      //mapeando a struct para viper fazer a troca
	JWTExperesIn  int    `mapstructure:"JWT_EXPERES_IN"`  //mapeando a struct para viper fazer a troca
	TokenAuth     *jwtauth.JWTAuth
}

func LoadConfig(path string) (*conf, error) {
	//lendo as configs
	viper.SetConfigName("config") //nome do arquivo de configuração
	viper.SetConfigType("env")    //tipo de configuração
	viper.AddConfigPath(path)     //caminho do arquivo
	viper.SetConfigFile(".env")   // nome do arquivo onde está a variaveis
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)
	return cfg, nil
}
