package internal

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/nndergunov/tgBot/bot/pkg/db/internal/models"
	"github.com/nndergunov/tgBot/bot/pkg/domain/entities"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type DB struct {
	db  *sql.DB
	ctx context.Context //nolint:containedctx
}

func NewDB(dbURL string) (*DB, error) {
	database, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("NewDB: %w", err)
	}

	ctx := context.TODO()

	return &DB{
		db:  database,
		ctx: ctx,
	}, nil
}

func (d DB) DeleteAllInfos() error {
	_, err := models.Infos().DeleteAll(d.ctx, d.db)
	if err != nil {
		return fmt.Errorf("DeleteAllInfos: %w", err)
	}

	return nil
}

func (d DB) InsertInfo(startTime time.Time, timeLayout string) error {
	var info models.Info

	info.Starttime = null.StringFrom(startTime.Format(timeLayout))

	err := info.Insert(d.ctx, d.db, boil.Infer())
	if err != nil {
		return fmt.Errorf("InsertInfo: %w", err)
	}

	return nil
}

func (d DB) GetInfo() (*models.Info, error) {
	infos, err := models.Infos().All(d.ctx, d.db)
	if err != nil {
		return nil, fmt.Errorf("GetInfo: %w", err)
	}

	lastInfo := len(infos) - 1

	return infos[lastInfo], nil
}

func (d DB) AddUserIfNotExists(userID int, userName string) error {
	allUsers, err := d.GetUsers()
	if err != nil {
		return fmt.Errorf("AddUserIfNotExists: %w", err)
	}

	for _, user := range allUsers {
		if user.Telegramid.Int == userID {
			return nil
		}
	}

	err = d.InsertUser(userID, userName)
	if err != nil {
		return fmt.Errorf("AddUserIfNotExists: %w", err)
	}

	return nil
}

func (d DB) GetUsers() (models.UserSlice, error) {
	users, err := models.Users().All(d.ctx, d.db)
	if err != nil {
		return nil, fmt.Errorf("GetUsers: %w", err)
	}

	return users, nil
}

func (d DB) InsertUser(userID int, userName string) error {
	var user models.User

	user.Telegramid = null.IntFrom(userID)
	user.Username = null.StringFrom(userName)

	err := user.Insert(d.ctx, d.db, boil.Infer())
	if err != nil {
		return fmt.Errorf("InsertUser: %w", err)
	}

	return nil
}

func (d DB) GetUserInTable(userID int) (int, error) {
	user, err := models.Users(qm.Where("telegramid=?", userID)).One(d.ctx, d.db)
	if err != nil {
		return 0, fmt.Errorf("GetUserInTable: %w", err)
	}

	return user.ID, nil
}

func (d DB) AddLocationIfNotExists(locationName string, userID int) error {
	allUserLocations, err := d.GetUserLocations(userID)
	if err != nil {
		return fmt.Errorf("AddLocationIfNotExists: %w", err)
	}

	for _, location := range allUserLocations {
		if location.Name.String == locationName {
			return nil
		}
	}

	err = d.InsertLocation(locationName, userID)
	if err != nil {
		return fmt.Errorf("AddLocationIfNotExists: %w", err)
	}

	return nil
}

func (d DB) GetUserLocations(userID int) (models.LocationSlice, error) {
	userInTable, err := d.GetUserInTable(userID)
	if err != nil {
		return nil, fmt.Errorf("GetUserLocations: %w", err)
	}

	locations, err := models.Locations(qm.Where("user_id=?", userInTable)).All(d.ctx, d.db)
	if err != nil {
		return nil, fmt.Errorf("GetUserLocations: %w", err)
	}

	return locations, nil
}

func (d DB) InsertLocation(locationName string, userID int) error {
	userInTable, err := d.GetUserInTable(userID)
	if err != nil {
		return fmt.Errorf("InsertLocation: %w", err)
	}

	var location models.Location

	location.Name = null.StringFrom(locationName)
	location.UserID = null.IntFrom(userInTable)

	err = location.Insert(d.ctx, d.db, boil.Infer())
	if err != nil {
		return fmt.Errorf("InsertLocation: %w", err)
	}

	return nil
}

func (d DB) GetLocationInTable(locationName string, userID int) (int, error) {
	userInTable, err := d.GetUserInTable(userID)
	if err != nil {
		return 0, fmt.Errorf("GetLocationInTable: %w", err)
	}

	location, err := models.Locations(qm.Where("user_id=? and name=?", userInTable, locationName)).One(
		d.ctx,
		d.db,
	)
	if err != nil {
		return 0, fmt.Errorf("GetLocationInTable: %w", err)
	}

	return location.ID, nil
}

func (d DB) AddAlbumToCollection(album entities.Album, locationName string, userID int) error {
	locationInTable, err := d.GetLocationInTable(locationName, userID)
	if err != nil {
		return fmt.Errorf("AddAlbumToCollection: %w", err)
	}

	err = d.AddArtistIfNotExists(album.Artist.Name)

	err = d.InsertAlbum(album)
	if err != nil {
		return fmt.Errorf("AddAlbumToCollection: %w", err)
	}

	albumInTable, err := d.GetAlbumInTable(album)
	if err != nil {
		return fmt.Errorf("AddAlbumToCollection: %w", err)
	}

	err = d.InsertToCollection(albumInTable, locationInTable)
	if err != nil {
		return fmt.Errorf("AddAlbumToCollection: %w", err)
	}

	return nil
}

func (d DB) AddArtistIfNotExists(artistName string) error {
	allArtists, err := d.GetArtists()
	if err != nil {
		return fmt.Errorf("AddArtistIfNotExists: %w", err)
	}

	for _, artist := range allArtists {
		if artist.Name.String == artistName {
			return nil
		}
	}

	err = d.InsertArtist(artistName)
	if err != nil {
		return fmt.Errorf("AddArtistIfNotExists: %w", err)
	}

	return nil
}

func (d DB) GetArtists() (models.ArtistSlice, error) {
	artists, err := models.Artists().All(d.ctx, d.db)
	if err != nil {
		return nil, fmt.Errorf("GetArtists: %w", err)
	}

	return artists, nil
}

func (d DB) InsertArtist(artistName string) error {
	var artist models.Artist

	artist.Name = null.StringFrom(artistName)

	err := artist.Insert(d.ctx, d.db, boil.Infer())
	if err != nil {
		return fmt.Errorf("InsertUser: %w", err)
	}

	return nil
}

func (d DB) GetArtistInTable(artistName string) (int, error) {
	artist, err := models.Artists(qm.Where("name=?", artistName)).One(d.ctx, d.db)
	if err != nil {
		return 0, fmt.Errorf("GetUserInTable: %w", err)
	}

	return artist.ID, nil
}

func (d DB) InsertAlbum(albumData entities.Album) error {
	var album models.Album

	artistInTable, err := d.GetArtistInTable(albumData.Artist.Name)
	if err != nil {
		return fmt.Errorf("InsertAlbum: %w", err)
	}

	album.ArtistID = null.IntFrom(artistInTable)
	album.AlbumName = null.StringFrom(albumData.Name)
	album.Genre = null.StringFrom(albumData.Genre)
	album.ReleaseYear = null.IntFrom(albumData.ReleaseYear)
	album.ReissueYear = null.IntFrom(albumData.ReissueYear)
	album.Label = null.StringFrom(albumData.Label)
	album.Coloured = null.BoolFrom(albumData.Coloured)
	album.CoverID = null.StringFrom(albumData.CoverID)

	err = album.Insert(d.ctx, d.db, boil.Infer())
	if err != nil {
		return fmt.Errorf("InsertAlbum: %w", err)
	}

	return nil
}

func (d DB) GetAlbumInTable(albumData entities.Album) (int, error) {
	album, err := models.Albums(qm.Where(
		"album_name=? and genre=? and release_year=? and reissue_year=? and label=?",
		albumData.Name,
		albumData.Genre,
		albumData.ReleaseYear,
		albumData.ReissueYear,
		albumData.Label,
	)).One(d.ctx, d.db)
	if err != nil {
		return 0, fmt.Errorf("GetUserInTable: %w", err)
	}

	return album.ID, nil
}

func (d DB) InsertToCollection(albumInTable, locationInTable int) error {
	var collectionElement models.Collection

	collectionElement.AlbumID = null.IntFrom(albumInTable)
	collectionElement.LocationID = null.IntFrom(locationInTable)

	err := collectionElement.Insert(d.ctx, d.db, boil.Infer())
	if err != nil {
		return fmt.Errorf("InsertToCollection: %w", err)
	}

	return nil
}
