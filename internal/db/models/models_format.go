package models

type IFormat interface {
	YAML() (string, error)
}
