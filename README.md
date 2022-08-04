
Проект назван в честь Юрия Гагарина - первого человека, полетевшего в космос.

## Веб сервер обеспечивает работу:
- gRPC API доступен на https://localhost:6060
- REST API доступен на https://localhost:7070

REST API используется для наглядности

gRPC API должен использоваться в Студии чтобы иметь преимущества в производительности при сервисном взаимодействии

## Общая информация о бэкенде:
- язык написания сервера go 1.15.5 darwin/amd64
- для генерации интерфейса клиентского и серверного gRPC API и REST API использовался https://github.com/grpc-ecosystem/grpc-gateway
- google.golang.org/protobuf v1.25.0 - компилятор Protobuf. Модуль является реализацией APIv2 протокола Protobuf https://github.com/golang/protobuf
- gRPC v1.3.0 транпортный протокол для proto сообщений https://github.com/grpc/grpc
- формат передачи данных по gRPC описан в прото файле, доступном в разделе Дополнительная информация/common.proto

## Общая информация об административном клиенте:
- для отображения информации о составе API используется swagger-ui-3.40.0 https://swagger.io/tools/swagger-ui/
- для генерации *.swagger.json используется https://github.com/grpc-ecosystem/grpc-gateway
- https://localhost:7070/common.swagger.json для тестирования REST API

##Для запуска сервера:
1. Если не установлен докер-клиент, то установить https://www.docker.com/products/docker-desktop

2. Установите скаченный docker-gagarin-api.tar (~15Мб) в локальный реестр докера командой из командной строки
   `docker load docker-gagarin-api.tar`

Для создания образа docker-gagarin-api надо выполнить команду
```
docker build -t gagarin/backend -f ./Dockerfile.build .
docker save gagarin/backend:latest -o ./bin
```

4. Запустите образ командной
   docker-compose up

5. Добавьте в баузер самоподписанный сертификат localhost.cer
   Пример инструкции https://wiki.webmoney.ru/projects/webmoney/wiki/%D0%A3%D1%81%D1%82%D0%B0%D0%BD%D0%BE%D0%B2%D0%BA%D0%B0_%D0%BA%D0%BE%D1%80%D0%BD%D0%B5%D0%B2%D0%BE%D0%B3%D0%BE_%D1%81%D0%B5%D1%80%D1%82%D0%B8%D1%84%D0%B8%D0%BA%D0%B0%D1%82%D0%B0_%D0%B2_%D0%B1%D1%80%D0%B0%D1%83%D0%B7%D0%B5%D1%80%D0%B5_Google_Chrome

6. Откройте в браузере ссылку https://localhost:7070/ чтобы увидеть интерфейс swagger-ui

## ToDo
Elasticsearch clients olivere/elastic aka HTTP transport of Elasticsearch

В docker-compose.yml добавлена конфигурация локально поднятого сервиса ES 
из образа docker.elastic.co/elasticsearch/elasticsearch:7.10.2

Опробован способ бэкапа и восстановления данных ES .\utils\create-backup.yml и .\utils\restore-backup.yml

Команды
```
docker-compose -f create-backup.yml run --rm create_backup
docker-compose -f restore-backup.yml run --rm restore_backup
```

## Запуск докера локально
1. из папки utils команду для создания тома с самоподписанным сертификатом ```docker-compose -f create-certs.yml run --rm create_certs```
2. из папки gagarin команду для поднятия ноды ES и ноды GRPC сервера ```docker-compose up -d```