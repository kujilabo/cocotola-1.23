package config

import (
	"embed"
	"os"

	"gopkg.in/yaml.v2"

	rslibconfig "github.com/kujilabo/cocotola-1.23/redstart/lib/config"
	rslibdomain "github.com/kujilabo/cocotola-1.23/redstart/lib/domain"
	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"

	libconfig "github.com/kujilabo/cocotola-1.23/lib/config"
)

type AppConfig struct {
	Name                 string `yaml:"name" validate:"required"`
	HTTPPort             int    `yaml:"httpPort" validate:"required"`
	MetricsPort          int    `yaml:"metricsPort" validate:"required"`
	ReadHeaderTimeoutSec int    `yaml:"readHeaderTimeoutSec" validate:"gte=1"`
}

type InternalAuthConfig struct {
	Type     string `yaml:"type" validate:"required"`
	Username string `yaml:"username" validate:"required"`
	Password string `yaml:"password" validate:"required"`
}

type Config struct {
	App          *AppConfig                 `yaml:"app" validate:"required"`
	DB           *rslibconfig.DBConfig      `yaml:"db" validate:"required"`
	InternalAuth *InternalAuthConfig        `yaml:"internalAuth" validate:"required"`
	Trace        *rslibconfig.TraceConfig   `yaml:"trace" validate:"required"`
	CORS         *rslibconfig.CORSConfig    `yaml:"cors" validate:"required"`
	Shutdown     *libconfig.ShutdownConfig  `yaml:"shutdown" validate:"required"`
	Log          *rslibconfig.LogConfig     `yaml:"log" validate:"required"`
	Swagger      *rslibconfig.SwaggerConfig `yaml:"swagger" validate:"required"`
	Debug        *libconfig.DebugConfig     `yaml:"debug"`
}

//go:embed local.yml
//go:embed production.yml
var config embed.FS

func LoadConfig(env string) (*Config, error) {
	filename := env + ".yml"
	confContent, err := config.ReadFile(filename)
	if err != nil {
		return nil, rsliberrors.Errorf("config.ReadFile. filename: %s, err: %w", filename, err)
	}

	confContent = []byte(os.ExpandEnv(string(confContent)))
	conf := &Config{}
	if err := yaml.Unmarshal(confContent, conf); err != nil {
		return nil, rsliberrors.Errorf("yaml.Unmarshal. filename: %s, err: %w", filename, err)
	}

	if err := rslibdomain.Validator.Struct(conf); err != nil {
		return nil, rsliberrors.Errorf("Validator.Struct. filename: %s, err: %w", filename, err)
	}

	return conf, nil
}