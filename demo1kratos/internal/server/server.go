package server

import (
	"github.com/google/wire"
	"go.einride.tech/aip/fieldbehavior"
	"google.golang.org/protobuf/proto"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer)

// validateRequiredFields rejects a request whose fields marked (google.api.field_behavior) = REQUIRED are unset
func validateRequiredFields(req any) error {
	if msg, ok := req.(proto.Message); ok {
		return fieldbehavior.ValidateRequiredFields(msg)
	}
	return nil
}
