package pkg_a

import (
	"github.com/jadekler/module-testing/pkg_b"
	"github.com/jadekler/module-testing/internal"
)

const FooA = "FooA"

var _ = pkg_b.FooB
var _= internal.InternalThing