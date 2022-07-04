package abft

import (
	"github.com/EktaMind/Thank_Sirus/hash"
	"github.com/EktaMind/Thank_Sirus/inter/dag"
)

// EventSource is a callback for getting events from an external storage.
type EventSource interface {
	HasEvent(hash.Event) bool
	GetEvent(hash.Event) dag.Event
}
