package models

import (
	"time"

	"github.com/mathiasb/snippetbox/pkg/utils"
)

var (
	ErrNoRecord           = utils.Error("models: no matching record found")
	ErrInvalidCredentials = utils.Error("models: invalid credentials")
	ErrDuplicateEmail     = utils.Error("models: duplicate email")
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
	Active         bool
}
