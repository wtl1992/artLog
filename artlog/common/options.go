package common

import (
	"artlog/artlog/colors"
)

type Option interface {
	SetOptionValue(attrs *colors.Attrs)
}
