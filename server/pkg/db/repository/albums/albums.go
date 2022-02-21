package albums

import (
	"context"

	db "github.com/nndergunov/tgBot/server/pkg/db/tgbot/public"
	"github.com/nndergunov/tgBot/server/pkg/db/tgbot/public/model"
)

type AlbumsRepository interface {
	CreateAlbums(ctx context.Context, name string) (id int, err error)
	ReadAlbums(ctx context.Context, id int) (*model.Albums, error)
	UpdateAlbums(ctx context.Context, room *model.Albums) error
	DeleteAlbums(ctx context.Context, id int) error

	ListAlbums(ctx context.Context, list *db.ListOptions, criteria *db.AlbumsCriteria) ([]*model.Albums, error)
}
