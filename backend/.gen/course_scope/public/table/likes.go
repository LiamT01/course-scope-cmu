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

var Likes = newLikesTable("public", "likes", "")

type likesTable struct {
	postgres.Table

	// Columns
	ID        postgres.ColumnInteger
	UserID    postgres.ColumnInteger
	RatingID  postgres.ColumnInteger
	CreatedAt postgres.ColumnTimestampz

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type LikesTable struct {
	likesTable

	EXCLUDED likesTable
}

// AS creates new LikesTable with assigned alias
func (a LikesTable) AS(alias string) *LikesTable {
	return newLikesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new LikesTable with assigned schema name
func (a LikesTable) FromSchema(schemaName string) *LikesTable {
	return newLikesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new LikesTable with assigned table prefix
func (a LikesTable) WithPrefix(prefix string) *LikesTable {
	return newLikesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new LikesTable with assigned table suffix
func (a LikesTable) WithSuffix(suffix string) *LikesTable {
	return newLikesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newLikesTable(schemaName, tableName, alias string) *LikesTable {
	return &LikesTable{
		likesTable: newLikesTableImpl(schemaName, tableName, alias),
		EXCLUDED:   newLikesTableImpl("", "excluded", ""),
	}
}

func newLikesTableImpl(schemaName, tableName, alias string) likesTable {
	var (
		IDColumn        = postgres.IntegerColumn("id")
		UserIDColumn    = postgres.IntegerColumn("user_id")
		RatingIDColumn  = postgres.IntegerColumn("rating_id")
		CreatedAtColumn = postgres.TimestampzColumn("created_at")
		allColumns      = postgres.ColumnList{IDColumn, UserIDColumn, RatingIDColumn, CreatedAtColumn}
		mutableColumns  = postgres.ColumnList{UserIDColumn, RatingIDColumn, CreatedAtColumn}
	)

	return likesTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		UserID:    UserIDColumn,
		RatingID:  RatingIDColumn,
		CreatedAt: CreatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
