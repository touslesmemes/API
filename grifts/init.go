package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/touslesmemes/api/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
