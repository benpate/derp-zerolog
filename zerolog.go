package plugins

import (
	"encoding/json"

	"github.com/benpate/derp"
	"github.com/rs/zerolog/log"
)

// Zerolog outputs errors to the zerolog logger.
type Zerolog struct{}

// Report implements the `derp.Plugin` interface, which allows the Console
// to be called by the derp.Report() method.
func (zerolog Zerolog) Report(err error) {

	switch typed := err.(type) {
	case derp.Error:
		buffer, _ := json.MarshalIndent(typed, "", "\t")
		log.Error().Str("loc", typed.Location).Str("message", typed.Message).RawJSON("details", buffer).Send()
	default:
		log.Error().Msg(err.Error())
	}
}
