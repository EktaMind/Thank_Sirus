package sirius

import (
	"github.com/EktaMind/Thank_Sirus/hash"
)

// Block is a part of an ordered chain of batches of events.
type Block struct {
	Atropos  hash.Event
	Cheaters Cheaters
}
