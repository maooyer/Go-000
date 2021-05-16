package data

import (
	"context"

	"github.com/maooyer/Go-000/Week04/homework/internal/config"
	"github.com/maooyer/Go-000/Week04/homework/internal/data/ent"
	"github.com/maooyer/Go-000/Week04/homework/pkg/log"
)

type Data struct {
	db *ent.Client
}

func NewData(conf *config.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper("data", logger)
	client, err := ent.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	if err != nil {
		log.Errorf("failed opening connection to sqlite: %v", err)
		return nil, nil, err
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Errorf("failed creating schema resources: %v", err)
		return nil, nil, err
	}

	d := &Data{
		db: client,
	}
	return d, func() {
		log.Info("message", "closing the data resources")
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
