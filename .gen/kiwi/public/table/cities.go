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

var Cities = newCitiesTable("public", "cities", "")

type citiesTable struct {
	postgres.Table

	// Columns
	ID               postgres.ColumnInteger
	Name             postgres.ColumnString
	Asciiname        postgres.ColumnString
	Alternatenames   postgres.ColumnString
	Latitude         postgres.ColumnFloat
	Longitude        postgres.ColumnFloat
	FeatureClass     postgres.ColumnString
	FeatureCode      postgres.ColumnString
	CountryCode      postgres.ColumnString
	Cc2              postgres.ColumnString
	Admin1Code       postgres.ColumnString
	Admin2Code       postgres.ColumnString
	Admin3Code       postgres.ColumnString
	Admin4Code       postgres.ColumnString
	Population       postgres.ColumnInteger
	Elevation        postgres.ColumnInteger
	Dem              postgres.ColumnInteger
	Timezone         postgres.ColumnString
	ModificationDate postgres.ColumnDate

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type CitiesTable struct {
	citiesTable

	EXCLUDED citiesTable
}

// AS creates new CitiesTable with assigned alias
func (a CitiesTable) AS(alias string) *CitiesTable {
	return newCitiesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new CitiesTable with assigned schema name
func (a CitiesTable) FromSchema(schemaName string) *CitiesTable {
	return newCitiesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new CitiesTable with assigned table prefix
func (a CitiesTable) WithPrefix(prefix string) *CitiesTable {
	return newCitiesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new CitiesTable with assigned table suffix
func (a CitiesTable) WithSuffix(suffix string) *CitiesTable {
	return newCitiesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newCitiesTable(schemaName, tableName, alias string) *CitiesTable {
	return &CitiesTable{
		citiesTable: newCitiesTableImpl(schemaName, tableName, alias),
		EXCLUDED:    newCitiesTableImpl("", "excluded", ""),
	}
}

func newCitiesTableImpl(schemaName, tableName, alias string) citiesTable {
	var (
		IDColumn               = postgres.IntegerColumn("id")
		NameColumn             = postgres.StringColumn("name")
		AsciinameColumn        = postgres.StringColumn("asciiname")
		AlternatenamesColumn   = postgres.StringColumn("alternatenames")
		LatitudeColumn         = postgres.FloatColumn("latitude")
		LongitudeColumn        = postgres.FloatColumn("longitude")
		FeatureClassColumn     = postgres.StringColumn("feature_class")
		FeatureCodeColumn      = postgres.StringColumn("feature_code")
		CountryCodeColumn      = postgres.StringColumn("country_code")
		Cc2Column              = postgres.StringColumn("cc2")
		Admin1CodeColumn       = postgres.StringColumn("admin1_code")
		Admin2CodeColumn       = postgres.StringColumn("admin2_code")
		Admin3CodeColumn       = postgres.StringColumn("admin3_code")
		Admin4CodeColumn       = postgres.StringColumn("admin4_code")
		PopulationColumn       = postgres.IntegerColumn("population")
		ElevationColumn        = postgres.IntegerColumn("elevation")
		DemColumn              = postgres.IntegerColumn("dem")
		TimezoneColumn         = postgres.StringColumn("timezone")
		ModificationDateColumn = postgres.DateColumn("modification_date")
		allColumns             = postgres.ColumnList{IDColumn, NameColumn, AsciinameColumn, AlternatenamesColumn, LatitudeColumn, LongitudeColumn, FeatureClassColumn, FeatureCodeColumn, CountryCodeColumn, Cc2Column, Admin1CodeColumn, Admin2CodeColumn, Admin3CodeColumn, Admin4CodeColumn, PopulationColumn, ElevationColumn, DemColumn, TimezoneColumn, ModificationDateColumn}
		mutableColumns         = postgres.ColumnList{NameColumn, AsciinameColumn, AlternatenamesColumn, LatitudeColumn, LongitudeColumn, FeatureClassColumn, FeatureCodeColumn, CountryCodeColumn, Cc2Column, Admin1CodeColumn, Admin2CodeColumn, Admin3CodeColumn, Admin4CodeColumn, PopulationColumn, ElevationColumn, DemColumn, TimezoneColumn, ModificationDateColumn}
	)

	return citiesTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:               IDColumn,
		Name:             NameColumn,
		Asciiname:        AsciinameColumn,
		Alternatenames:   AlternatenamesColumn,
		Latitude:         LatitudeColumn,
		Longitude:        LongitudeColumn,
		FeatureClass:     FeatureClassColumn,
		FeatureCode:      FeatureCodeColumn,
		CountryCode:      CountryCodeColumn,
		Cc2:              Cc2Column,
		Admin1Code:       Admin1CodeColumn,
		Admin2Code:       Admin2CodeColumn,
		Admin3Code:       Admin3CodeColumn,
		Admin4Code:       Admin4CodeColumn,
		Population:       PopulationColumn,
		Elevation:        ElevationColumn,
		Dem:              DemColumn,
		Timezone:         TimezoneColumn,
		ModificationDate: ModificationDateColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}