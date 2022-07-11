package config

import (
	"bytes"

	"github.com/spf13/viper"
)

const (
	//Конфигурация elasticsearch проекта
	yamlElasticConfig = `
	es:
	#Хост
		host: localhost
	#Протокол соединения
		scheme: https
	#Запрашивать статус эластика
		sniff: false
	#Сжимать данные
		gzip: false
	#API ключ для добавления в заголовок запроса
		apiKey: ckY5cmVIY0JfS2RDUUxZMF9MX2Q6bGxfSmJvUlVRLVdsWWh3VHBKTW5Pdw==
	#Пользователь
		user: elastic
	#Пароль пользователя
		password: 123
	`
)

//Структура конфигурации elasticsearch
type ElasticConfig struct {
	Scheme   string // Протокол соединения
	Host     string // Хост
	Sniff    bool   // Запрашивать статус эластика
	Gzip     bool   // Сжимать данные
	User     string // Пользователь
	Password string // Пароль пользователя
	ApiKey   string // API ключ для добавления в заголовок запроса
}

//Метод config.Init() для инициализации базовой конфигурации
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
