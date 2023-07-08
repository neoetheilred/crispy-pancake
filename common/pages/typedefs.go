package pages

import "io"

type page struct {
	TemplName string
	Data      any
}

type Page func(io.Writer, any)
