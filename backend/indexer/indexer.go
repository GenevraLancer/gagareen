package indexer
import (  
	"context"
	"encoding/json"
	"log"
  
	"github.com/olivere/elastic/v7"
	 myapi "github.com/GenevraLancer/gagarin/backend/gen/api"
	 "github.com/GenevraLancer/gagarin/backend/config"
  )

  //Обёртка для клиента es 
  type ElasticClient struct {  
	client *elastic.Client
  }
  
  //Инициализация клиента с параметрами из .env
  func New(url string) (*ElasticIndex, error) {  
	config := config.GagarinConfig
	client, err := elastic.NewClient(
	  elastic.SetURL(viper.Get"https://localhost:9200"),
	  elastic.SetSniff(false),
	)
	if err != nil {
	  return nil, err
	}

	fmt.Println("ElasticSearch initialized in gagarin-api...")
	return &ElasticIndex{client}, nil
  }
  
  func (r *ElasticIndex) Close() {  
  }
  
  func (r *ElasticIndex) InsertField(ctx context.Context, myapi ) error { 
  }