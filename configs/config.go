package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var cfg *conf

type conf struct {
	//mapstructure => mapeia as config para viper fazer as substituições

	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
	JwtExpiresIn  int    `mapstructure:"JWT_EXPIRESIN"`
	TokenAuth     *jwtauth.JWTAuth
}

// função que recebe o caminho de arq de config e retorna as configs ou um erro
func LoadConfig(path string) (*conf, error) {

	//nome do arquivo de configuração
	viper.SetConfigName("app_config")

	//tipo de arq de configuração
	viper.SetConfigType("env")

	//caminho do arquivo
	viper.AddConfigPath(path)

	//nome do arq onde estão as variáveis
	viper.SetConfigFile(".env")

	//permite que as configurações do seu aplicativo sejam substituídas por variáveis de ambiente sem a necessidade de modificar os arquivos de configuração.
	viper.AutomaticEnv()

	// le o arquivo de configuração
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// função viper.Unmarshal mapeia as configurações lidas para os campos correspondentes na estrutura fornecida
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	// criando uma instancia para gerar token jwt
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)

	return cfg, nil
}
