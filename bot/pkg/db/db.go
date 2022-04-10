package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	// postgres driver.
	_ "github.com/lib/pq"

	// other packages.
	"github.com/nndergunov/tgBot/bot/pkg/db/models"
	"github.com/nndergunov/tgBot/bot/pkg/domain/entities"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

const timeLayout = "02 Jan 06 15:04 MST"

type Database struct {
	db  *sql.DB
	ctx context.Context
}

func NewDatabase(dbURL string) (*Database, error) {
	database, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("NewDatabase: %w", err)
	}

	ctx := context.TODO()

	return &Database{
		db:  database,
		ctx: ctx,
	}, nil
}

func (d Database) PutInfo(startTime time.Time) error {
	_, err := models.Infos().DeleteAll(d.ctx, d.db)
	if err != nil {
		return fmt.Errorf("PutInfo: %w", err)
	}

	var info models.Info

	info.Starttime = null.StringFrom(startTime.Format(timeLayout))

	err = info.Insert(d.ctx, d.db, boil.Infer())
	if err != nil {
		return fmt.Errorf("PutInfo: %w", err)
	}

	return nil
}

func (d Database) GetInfo() (*entities.Info, error) {
	infos, err := models.Infos().All(d.ctx, d.db)
	if err != nil {
		return nil, fmt.Errorf("GetInfo: %w", err)
	}

	lastInfo := len(infos) - 1

	startTime, err := time.Parse(timeLayout, infos[lastInfo].Starttime.String)
	if err != nil {
		return nil, fmt.Errorf("GetInfo: %w", err)
	}

	return &entities.Info{Starttime: startTime}, nil
}
