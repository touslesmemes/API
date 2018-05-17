package grifts

import (
	"git.aprentout.com/touslesmemes/api2/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
