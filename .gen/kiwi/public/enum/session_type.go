//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package enum

import "github.com/go-jet/jet/v2/postgres"

var SessionType = &struct {
	FillProfile   postgres.StringExpression
	FillBlacklist postgres.StringExpression
	None          postgres.StringExpression
}{
	FillProfile:   postgres.NewEnumValue("fill_profile"),
	FillBlacklist: postgres.NewEnumValue("fill_blacklist"),
	None:          postgres.NewEnumValue("none"),
}