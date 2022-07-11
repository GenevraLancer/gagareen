package config

import (
	"bytes"

	"github.com/spf13/viper"
)

const (
	// Структура индекса проверок в эластике
	jsonChecksIndexSettings = `
	{
	  "settings": {
		"index": {
		  "number_of_shards": 1,
		  "number_of_replicas": "0",
		  "refresh_interval": "-1",
		  "requests": {
			"cache": {
			  "enable": "false"
			}
		  },
		  "blocks": {
			"read_only_allow_delete": "false"
		  }
		}
	  },
	  "mappings": {
		"dynamic": false,
		"properties": {
		  "isBinary": {
			"type": "keyword"
		  },
		  "keyCDO": {
			"type": "keyword"
		  },
		  "checkID": {
			"type":"text",
			"store": true
		  },
		  "checkDesc": {
			"type":"text",
			"store": true
		  },
		  "formularSysName": {
			"type":"text",
			"store": true
		  }
		}
	  }
	}
	`
)

//Метод config.SetIndexCheckStruct() для инициализации структуры индекса проверок
func (config *ElasticConfig) Init() error {
	viper.SetConfigType("yaml")
	c := []byte(yamlElasticConfig)

	var err error
	err = viper.ReadConfig(bytes.NewBuffer(c))

	//Приведение к типу
	config.Scheme = viper.Get("es.scheme").(string)
	config.Host = viper.Get("es.host").(string)
	config.Sniff = viper.Get("es.sniff").(bool)
	config.Gzip = viper.Get("es.gzip").(bool)
	config.User = viper.Get("es.user").(string)
	config.Password = viper.Get("es.password").(string)
	config.ApiKey = viper.Get("es.apiKey").(string)

	return err
}
