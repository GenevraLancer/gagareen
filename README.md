# gagarin is framework to manage metadata, dictionary data related to specific projects

The project is named after Yuri Gagarin, the first person to fly the space.

## CoreStack
### Frontend
Vue.js 3.0.0-rc.5 - Vue 3 для создания интерфейса
Node.js v14.2.0 - используется для сборки клиентской части и Vue.
Go client for the NATS messaging system

### Backend
1. Язык написания сервера go 1.15.5 darwin/amd64
2. google.golang.org/protobuf v1.25.0 - компилятор Protobuf, используемый для создания заглушек *.pb.go. Модуль является реализацией APIv2 протокола Protobuf 

1. gRPC v1.3.0 and greater aka transport layer for protobuf messages
2. liftbridge message streams for the NATS messaging system
3. Elasticsearch clients olivere/elastic aka HTTP transport of Elasticsearch, e.g. for authentication
4. Docker v19.03.12 - Используется для запуска envoy.
5. envoy v1.14 - прокси для преобразования HTTP / 1.1 из обычного Интернета в HTTP / 2.
6. Docker Compose v1.26.2 - Используйте, чтобы сделать все Docker, вы можете работать без него.
