package ebase

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/jilin7105/ebase/config"
	"github.com/jilin7105/ebase/logger"
)

func newes(config config.EsConfig) (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{}
	if config.Type == "host" {
		cfg = elasticsearch.Config{
			Addresses: config.Hosts,
		}
		if config.User != "" {
			cfg.Username = config.User

		}

		if config.Pass != "" {
			cfg.Password = config.Pass
		}
	}

	if config.Type == "cloud" {
		cfg.CloudID = config.CloudId
		cfg.APIKey = config.ApiKey
	}

	return elasticsearch.NewClient(cfg)

}

func (e *Eb) Inites() {

	for _, config := range e.Config.Es {
		esclient, err := newes(config)
		if err != nil {
			logger.Error("es  初始化失败 %s %s", config.Name, err.Error())
		}
		e.ES[config.Name] = esclient
	}

}