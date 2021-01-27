package handlers

import (
	"context"
	"sync"

	myapi "github.com/GenevraLancer/gagarin/backend/gen/api"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CommonService implements the protobuf interface
type CommonService struct {
	mu                                     *sync.RWMutex        //mutex защищает AddFieldParams структуру
	fields                                 []*myapi.FieldObject //выделение памяти для слайса FieldObject
	myapi.UnimplementedCommonServiceServer                      // Embed the unimplemented server
}

//New initializes a new QueryService struct.
func New() *CommonService {
	return &CommonService{
		mu: &sync.RWMutex{}, //интерфейс
	}
}

// GetFields lists all fields in the store.
func (cs *CommonService) BulkGetFields(_ *myapi.BulkFields, srv myapi.CommonService_BulkGetFieldsServer) error {
	cs.mu.RLock()         //lock только на чтение данных
	defer cs.mu.RUnlock() //unlock чтения данных через defer

	for _, field := range cs.fields {
		err := srv.Send(field)
		if err != nil {
			return err
		}
	}

	return nil
}

// AddField adds a field to the in-memory store.
func (cs *CommonService) AddField(ctx context.Context, req *myapi.AddFieldParams) (*myapi.FieldObject, error) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	field := &myapi.FieldObject{
		Id:         uuid.Must(uuid.NewV4()).String(),
		SystemName: req.GetSystemName(),
	}
	cs.fields = append(cs.fields, field)

	return field, nil
}

// GetField get a field from the in-memory store by id.
func (cs *CommonService) GetField(ctx context.Context, req *myapi.GetFieldParams) (*myapi.FieldObject, error) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	for _, field := range cs.fields {
		if field.Id == req.GetId() {
			return field, nil
		}
	}

	return nil, status.Errorf(codes.NotFound, "пользователь с ID %q не найден", req.GetId())
}
