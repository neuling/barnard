package main

import (
	"crypto/tls"

	"github.com/neuling/barnard/uiterm"
	"github.com/neuling/gumble"
	"github.com/neuling/gumble/gumbleopenal"
)

type Barnard struct {
	Config *gumble.Config
	Client *gumble.Client

	Address   string
	TLSConfig tls.Config

	Stream *gumbleopenal.Stream

	Ui            *uiterm.Ui
	UiOutput      uiterm.Textview
	UiInput       uiterm.Textbox
	UiStatus      uiterm.Label
	UiTree        uiterm.Tree
	UiInputStatus uiterm.Label
}
