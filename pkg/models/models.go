package models

import (
	"time"

	"github.com/mathiasb/snippetbox/pkg/utils"
)

var ErrNoRecord = utils.Error("models: no matching record found")

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
