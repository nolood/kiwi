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

var Profiles = newProfilesTable("public", "profiles", "")

type profilesTable struct {
	postgres.Table

	// Columns
	ID        postgres.ColumnInteger
	UserID    postgres.ColumnInteger
	UserTgID  postgres.ColumnInteger
	Name      postgres.ColumnString
	Age       postgres.ColumnInteger
	Gender    postgres.ColumnString
	PhotoID   postgres.ColumnString
	About     postgres.ColumnString
	IsActive  postgres.ColumnBool
	Latitude  postgres.ColumnFloat
	Longitude postgres.ColumnFloat

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type ProfilesTable struct {
	profilesTable

	EXCLUDED profilesTable
}

// AS creates new ProfilesTable with assigned alias
func (a ProfilesTable) AS(alias string) *ProfilesTable {
	return newProfilesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ProfilesTable with assigned schema name
func (a ProfilesTable) FromSchema(schemaName string) *ProfilesTable {
	return newProfilesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ProfilesTable with assigned table prefix
func (a ProfilesTable) WithPrefix(prefix string) *ProfilesTable {
	return newProfilesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ProfilesTable with assigned table suffix
func (a ProfilesTable) WithSuffix(suffix string) *ProfilesTable {
	return newProfilesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newProfilesTable(schemaName, tableName, alias string) *ProfilesTable {
	return &ProfilesTable{
		profilesTable: newProfilesTableImpl(schemaName, tableName, alias),
		EXCLUDED:      newProfilesTableImpl("", "excluded", ""),
	}
}

func newProfilesTableImpl(schemaName, tableName, alias string) profilesTable {
	var (
		IDColumn        = postgres.IntegerColumn("id")
		UserIDColumn    = postgres.IntegerColumn("user_id")
		UserTgIDColumn  = postgres.IntegerColumn("user_tg_id")
		NameColumn      = postgres.StringColumn("name")
		AgeColumn       = postgres.IntegerColumn("age")
		GenderColumn    = postgres.StringColumn("gender")
		PhotoIDColumn   = postgres.StringColumn("photo_id")
		AboutColumn     = postgres.StringColumn("about")
		IsActiveColumn  = postgres.BoolColumn("is_active")
		LatitudeColumn  = postgres.FloatColumn("latitude")
		LongitudeColumn = postgres.FloatColumn("longitude")
		allColumns      = postgres.ColumnList{IDColumn, UserIDColumn, UserTgIDColumn, NameColumn, AgeColumn, GenderColumn, PhotoIDColumn, AboutColumn, IsActiveColumn, LatitudeColumn, LongitudeColumn}
		mutableColumns  = postgres.ColumnList{UserIDColumn, UserTgIDColumn, NameColumn, AgeColumn, GenderColumn, PhotoIDColumn, AboutColumn, IsActiveColumn, LatitudeColumn, LongitudeColumn}
	)

	return profilesTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		UserID:    UserIDColumn,
		UserTgID:  UserTgIDColumn,
		Name:      NameColumn,
		Age:       AgeColumn,
		Gender:    GenderColumn,
		PhotoID:   PhotoIDColumn,
		About:     AboutColumn,
		IsActive:  IsActiveColumn,
		Latitude:  LatitudeColumn,
		Longitude: LongitudeColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
