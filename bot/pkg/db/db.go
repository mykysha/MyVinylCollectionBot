package db

import (
	"fmt"
	"time"

	// postgres driver.
	_ "github.com/lib/pq"
	"github.com/nndergunov/tgBot/bot/pkg/db/internal"
	"github.com/nndergunov/tgBot/bot/pkg/domain/entities"
)

const timeLayout = "02 Jan 06 15:04 MST"

type Database struct {
	db *internal.DB
}

func NewDatabase(dbURL string) (*Database, error) {
	database, err := internal.NewDB(dbURL)
	if err != nil {
		return nil, fmt.Errorf("NewDatabase: %w", err)
	}

	return &Database{
		db: database,
	}, nil
}

func (d Database) PutInfo(info entities.Info) error {
	err := d.db.DeleteAllInfos()
	if err != nil {
		return fmt.Errorf("PutInfo: %w", err)
	}

	err = d.db.InsertInfo(info.Starttime, timeLayout)
	if err != nil {
		return fmt.Errorf("PutInfo: %w", err)
	}

	return nil
}

func (d Database) GetInfo() (*entities.Info, error) {
	info, err := d.db.GetInfo()
	if err != nil {
		return nil, fmt.Errorf("GetInfo: %w", err)
	}

	startTime, err := time.Parse(timeLayout, info.Starttime.String)
	if err != nil {
		return nil, fmt.Errorf("GetInfo: %w", err)
	}

	return &entities.Info{Starttime: startTime}, nil
}

func (d Database) AddAlbumToCollection(album entities.Album, location entities.Location) error {
	err := d.db.AddUserIfNotExists(location.Owner.ChatID, location.Owner.UserName)
	if err != nil {
		return fmt.Errorf("AddAlbumToCollection: %w", err)
	}

	err = d.db.AddLocationIfNotExists(location.Name, location.Owner.ChatID)
	if err != nil {
		return fmt.Errorf("AddAlbumToCollection: %w", err)
	}

	err = d.db.AddAlbumToCollection(album, location.Name, location.Owner.ChatID)
	if err != nil {
		return fmt.Errorf("AddAlbumToCollection: %w", err)
	}

	return nil
}
