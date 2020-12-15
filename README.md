# gagarin is framework to manage metadata, dictionary data related to specific projects

The project is named after Yuri Gagarin, the first person to fly the space.

## CoreStack
### Frontend
Vue.js 3.0.0-rc.5 - Vue 3 для создания интерфейса
Go client for the NATS messaging system

### Backend
1. gRPC v1.3.0 and greater aka transport layer for protobuf messages
gogoprotobuf instead of the default protobuf because gogo allows generation of the data model without too many pointers, e.g. []model.KeyValue instead of []*model.KeyValue (gist)

2. liftbridge message streams for the NATS messaging system
3. Elasticsearch clients olivere/elastic aka HTTP transport of Elasticsearch, e.g. for authentication
4. Protoc v3.12.4 - компилятор Protobuf, используемый для создания заглушек.
5. Node.js v14.2.0 - используется для сборки серверной части и Vue.
6. Docker v19.03.12 - Используется для запуска envoy.
7. envoy v1.14 - прокси для преобразования HTTP / 1.1 из обычного Интернета в HTTP / 2.
8. Docker Compose v1.26.2 - Используйте, чтобы сделать все Docker, вы можете работать без него.
