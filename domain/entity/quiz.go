package entity

import "github.com/mag30/project-backend/domain/base"

type Quiz struct {
	base.EntityWithIdKey
	Name string

	Tasks []Task
}
