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

	startTime, err := time.Parse(timeLayout, info.Starttime)
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

func (d Database) GetCollection(userID string) ([]entities.Album, error) {
	var albums []entities.Album

	dbAlbums, err := d.db.GetCollection(userID)
	if err != nil {
		return nil, fmt.Errorf("GetCollection: %w", err)
	}

	for _, dbAlbum := range dbAlbums {
		artist, err := d.db.GetArtistByID(dbAlbum.ArtistID)
		if err != nil {
			return nil, fmt.Errorf("GetCollection: %w", err)
		}

		album := entities.Album{
			Artist: entities.Artist{
				Name: artist.Name,
			},
			Name:        dbAlbum.AlbumName,
			Genre:       dbAlbum.Genre,
			ReleaseYear: dbAlbum.ReleaseYear,
			ReissueYear: dbAlbum.ReissueYear,
			Label:       dbAlbum.Label,
			Coloured:    dbAlbum.Coloured,
			CoverID:     dbAlbum.CoverID,
		}

		albums = append(albums, album)
	}

	return albums, nil
}

func (d Database) GetGenres(userID string) ([]string, error) {
	var genres []string

	dbAlbums, err := d.db.GetCollection(userID)
	if err != nil {
		return nil, fmt.Errorf("GetCollection: %w", err)
	}

	for _, dbAlbum := range dbAlbums {
		genres = append(genres, dbAlbum.Genre)
	}

	return genres, nil
}

func (d Database) GetArtists(userID string) ([]entities.Artist, error) {
	var artists []entities.Artist

	dbAlbums, err := d.db.GetCollection(userID)
	if err != nil {
		return nil, fmt.Errorf("GetCollection: %w", err)
	}

	for _, dbAlbum := range dbAlbums {
		artist, err := d.db.GetArtistByID(dbAlbum.ArtistID)
		if err != nil {
			return nil, fmt.Errorf("GetCollection: %w", err)
		}

		artists = append(artists, entities.Artist{Name: artist.Name})
	}

	return artists, nil
}

func (d Database) GetLocationByName(locationName, userID, userName string) (entities.Location, error) {
	loc, err := d.db.GetLocationByName(locationName, userID)
	if err != nil {
		return entities.Location{}, fmt.Errorf("GetLocationByName: %w", err)
	}

	return entities.Location{
		Owner: entities.User{
			ChatID:   userID,
			UserName: userName,
		},
		Name: loc.Name,
	}, nil
}

func (d Database) DeleteAlbum(albumNumber int, userID string) error {
	err := d.db.DeleteAlbum(userID, albumNumber-1)
	if err != nil {
		return fmt.Errorf("GetCollection: %w", err)
	}

	return nil
}
