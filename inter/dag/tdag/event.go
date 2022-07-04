package tdag

import (
	"github.com/EktaMind/Thank_Sirus/hash"
	"github.com/EktaMind/Thank_Sirus/inter/dag"
)

type TestEvent struct {
	dag.MutableBaseEvent
	Name string
}

func (e *TestEvent) AddParent(id hash.Event) {
	parents := e.Parents()
	parents.Add(id)
	e.SetParents(parents)
}
