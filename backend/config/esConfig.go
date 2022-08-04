package config

import (
	elastic "github.com/elastic/go-elasticsearch/v7"
)

/* const (
	//Конфигурация elasticsearch проекта, потом вынесу в отдельный файл
	yamlElasticConfig = `
	es:
	#Хост
		host: https://localhost:9200
	#API ключ для добавления в заголовок запроса
		apiKey: ckY5cmVIY0JfS2RDUUxZMF9MX2Q6bGxfSmJvUlVRLVdsWWh3VHBKTW5Pdw==
	#Пользователь
		user: elastic
	#Пароль пользователя
		password: 123
	`
) */

//InitConfig() создает новый ElasticConfig из yaml файла.
func InitConfig() *elastic.Config {
	config := elastic.Config{
			Addresses: []string{
			  "https://localhost:9200",
			},
			Username: "elastic",
			Password: "123",
			APIKey: "ckY5cmVIY0JfS2RDUUxZMF9MX2Q6bGxfSmJvUlVRLVdsWWh3VHBKTW5Pdw",
	}
	return &config
}
