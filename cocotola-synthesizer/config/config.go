package config

import (
	_ "embed"
	// rslibconfig "github.com/kujilabo/redstart/lib/config"
	// rslibdomain "github.com/kujilabo/redstart/lib/domain"
	// rsliberrors "github.com/kujilabo/redstart/lib/errors"
	// libconfig "github.com/kujilabo/cocotola-1.21/lib/config"
)

// type AppConfig struct {
// 	Name                 string `yaml:"name" validate:"required"`
// 	HTTPPort             int    `yaml:"httpPort" validate:"required"`
// 	MetricsPort          int    `yaml:"metricsPort" validate:"required"`
// 	ReadHeaderTimeoutSec int    `yaml:"readHeaderTimeoutSec" validate:"gte=1"`
// }

// type InternalAuthConfig struct {
// 	Type     string `yaml:"type" validate:"required"`
// 	Username string `yaml:"username" validate:"required"`
// 	Password string `yaml:"password" validate:"required"`
// }

// type ShutdownConfig struct {
// 	TimeSec1 int `yaml:"timeSec1" validate:"gte=1"`
// 	TimeSec2 int `yaml:"timeSec2" validate:"gte=1"`
// }

// type GoogleTextToSpeechConfig struct {
// 	APIKey        string `yaml:"apiKey" validate:"required"`
// 	APITimeoutSec int    `yaml:"apiTimeoutSec" validate:"gte=5"`
// }

// type Config struct {
// 	App          *AppConfig                 `yaml:"app" validate:"required"`
// 	InternalAuth *InternalAuthConfig        `yaml:"internalAuth" validate:"required"`
// 	DB           *rslibconfig.DBConfig      `yaml:"db" validate:"required"`
// 	TTS          *GoogleTextToSpeechConfig  `yaml:"googleTextToSpeech" validate:"required"`
// 	Trace        *rslibconfig.TraceConfig   `yaml:"trace" validate:"required"`
// 	CORS         *rslibconfig.CORSConfig    `yaml:"cors" validate:"required"`
// 	Shutdown     *ShutdownConfig            `yaml:"shutdown" validate:"required"`
// 	Log          *rslibconfig.LogConfig     `yaml:"log" validate:"required"`
// 	Swagger      *rslibconfig.SwaggerConfig `yaml:"swagger" validate:"required"`
// 	Debug        *libconfig.DebugConfig     `yaml:"debug"`
// }

// //go:embed local.yml
// //go:embed production.yml
// var config embed.FS

// func LoadConfig(env string) (*Config, error) {
// 	filename := env + ".yml"
// 	confContent, err := config.ReadFile(filename)
// 	if err != nil {
// 		return nil, rsliberrors.Errorf("config.ReadFile. filename: %s, err: %w", filename, err)
// 	}

// 	confContent = []byte(os.ExpandEnv(string(confContent)))
// 	conf := &Config{}
// 	if err := yaml.Unmarshal(confContent, conf); err != nil {
// 		return nil, rsliberrors.Errorf("yaml.Unmarshal. filename: %s, err: %w", filename, err)
// 	}

// 	if err := rslibdomain.Validator.Struct(conf); err != nil {
// 		return nil, rsliberrors.Errorf("Validator.Struct. filename: %s, err: %w", filename, err)
// 	}

// 	return conf, nil
// }
