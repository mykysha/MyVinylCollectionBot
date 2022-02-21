package db

type AlbumsCriteria struct {
	ID          *int
	ArtistID    *int
	AlbumName   *string
	Genre       *string
	ReleaseYear *int
	ReissueYear *int
	Label       *int
	Coloured    *bool
	Cover       *string
}
