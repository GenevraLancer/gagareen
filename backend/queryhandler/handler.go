package queryhandler

import (
	"sync"

	myapi "github.com/GenevraLancer/gagarin/backend/gen/api"
)

// QueryService implements the protobuf interface
type QueryService struct {
	mu     *sync.RWMutex
	fields []*myapi.AddFieldParams
}

/* // New initializes a new QueryService struct.
func New() *QueryService {
	return &QueryService{
		mu: &sync.RWMutex{},
	}
}

// AddField adds a field to the in-memory store.
func (b *Backend) AddField(ctx context.Context, _ *pbField.AddFieldParams) (*pbField.FieldObject, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	field := &pbField.FieldObject{
		Id: uuid.Must(uuid.NewV4()).String(),
	}
	b.fields = append(b.fields, field)

	return field, nil
}

// ListFields lists all fields in the store.
func (b *Backend) GetField(_ *pbField.GetFieldParams, srv pbField.FieldService_SubscribeServer) error {
	b.mu.RLock()
	defer b.mu.RUnlock()

	for _, field := range b.fields {
		err := srv.Send(field)
		if err != nil {
			return err
		}
	}

	return nil
}
*/
