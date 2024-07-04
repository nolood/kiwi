//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import "errors"

type SessionType string

const (
	SessionType_FillProfile   SessionType = "fill_profile"
	SessionType_FillBlacklist SessionType = "fill_blacklist"
	SessionType_None          SessionType = "none"
)

func (e *SessionType) Scan(value interface{}) error {
	var enumValue string
	switch val := value.(type) {
	case string:
		enumValue = val
	case []byte:
		enumValue = string(val)
	default:
		return errors.New("jet: Invalid scan value for AllTypesEnum enum. Enum value has to be of type string or []byte")
	}

	switch enumValue {
	case "fill_profile":
		*e = SessionType_FillProfile
	case "fill_blacklist":
		*e = SessionType_FillBlacklist
	case "none":
		*e = SessionType_None
	default:
		return errors.New("jet: Invalid scan value '" + enumValue + "' for SessionType enum")
	}

	return nil
}

func (e SessionType) String() string {
	return string(e)
}
