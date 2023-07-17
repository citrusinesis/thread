package config

import (
	"context"
	"github.com/citrusinesis/thread/pkg/ent"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	once   sync.Once
	client *ent.Client
)

func GetMySQLClient(config DBConfig) *ent.Client {
	var err error

	once.Do(func() {
		client, err = ent.Open("mysql", config.CreateURL())
		if err != nil {
			log.Fatalf("failed opening connection to mysql: %v", err)
		}

		if err := client.Schema.Create(context.Background()); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
	})

	return client
}
