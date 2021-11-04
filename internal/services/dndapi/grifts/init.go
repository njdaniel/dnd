package grifts

import (
  "github.com/gobuffalo/buffalo"
	"dndapi/actions"
)

func init() {
  buffalo.Grifts(actions.App())
}
