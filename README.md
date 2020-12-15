# gagarin is framework to manage metadata, dictionary data related to specific projects

The project is named after Yuri Gagarin, the first person to fly the space.

## CoreStack
### Frontend
Vue3.js
Go client for the NATS messaging system

### Backend
1. gRPC v1.3.0 and greater aka transport layer for protobuf messages
gogoprotobuf instead of the default protobuf because gogo allows generation of the data model without too many pointers, e.g. []model.KeyValue instead of []*model.KeyValue (gist)

2. liftbridge message streams for the NATS messaging system
3. Elasticsearch clients olivere/elastic aka HTTP transport of Elasticsearch, e.g. for authentication

