package config

import (
	"embed"
	"os"

	"gopkg.in/yaml.v3"

	rslibdomain "github.com/kujilabo/cocotola-1.23/redstart/lib/domain"
	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
)

type AuthConfig struct {
	Username string `yaml:"username" validate:"required"`
	Password string `yaml:"password" validate:"required"`
}

type TatoebaDataSourceConfig struct {
	Dir              string `yaml:"dir" validate:"required"`
	EngSentencesFile string `yaml:"engSentences" validate:"required"`
	JpnSentencesFile string `yaml:"jpnSentences" validate:"required"`
	LinksFile        string `yaml:"links" validate:"required"`
}

type DataSourceConfig struct {
	TatoebaDataSource *TatoebaDataSourceConfig `yaml:"tatoeba" validate:"required"`
}

type TatoebaAPIonfig struct {
	Endpoint string `yaml:"endpoint" validate:"required"`
	Username string `yaml:"username" validate:"required"`
	Password string `yaml:"password" validate:"required"`
}

type Config struct {
	TatoebaAPI *TatoebaAPIonfig  `yaml:"tatoebaApi" validate:"required"`
	DataSource *DataSourceConfig `yaml:"datasource" validate:"required"`
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
