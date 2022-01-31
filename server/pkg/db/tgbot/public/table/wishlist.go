//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var Wishlist = newWishlistTable("public", "wishlist", "")

type wishlistTable struct {
	postgres.Table

	//Columns
	ID      postgres.ColumnInteger
	OwnerID postgres.ColumnInteger
	AlbumID postgres.ColumnInteger
	Store   postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type WishlistTable struct {
	wishlistTable

	EXCLUDED wishlistTable
}

// AS creates new WishlistTable with assigned alias
func (a WishlistTable) AS(alias string) *WishlistTable {
	return newWishlistTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new WishlistTable with assigned schema name
func (a WishlistTable) FromSchema(schemaName string) *WishlistTable {
	return newWishlistTable(schemaName, a.TableName(), a.Alias())
}

func newWishlistTable(schemaName, tableName, alias string) *WishlistTable {
	return &WishlistTable{
		wishlistTable: newWishlistTableImpl(schemaName, tableName, alias),
		EXCLUDED:      newWishlistTableImpl("", "excluded", ""),
	}
}

func newWishlistTableImpl(schemaName, tableName, alias string) wishlistTable {
	var (
		IDColumn       = postgres.IntegerColumn("id")
		OwnerIDColumn  = postgres.IntegerColumn("owner_id")
		AlbumIDColumn  = postgres.IntegerColumn("album_id")
		StoreColumn    = postgres.StringColumn("store")
		allColumns     = postgres.ColumnList{IDColumn, OwnerIDColumn, AlbumIDColumn, StoreColumn}
		mutableColumns = postgres.ColumnList{IDColumn, OwnerIDColumn, AlbumIDColumn, StoreColumn}
	)

	return wishlistTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:      IDColumn,
		OwnerID: OwnerIDColumn,
		AlbumID: AlbumIDColumn,
		Store:   StoreColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}