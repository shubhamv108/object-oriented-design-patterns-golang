package typedobject

import (
	"fmt"
)

type TypedObject struct {
	attribute1 string
	attribute2 string
}

func Builder() *TypeObjectBuilder {
	return &TypeObjectBuilder{};
}

func (t *TypedObject) String() string {
	return fmt.Sprintf("{attribute1: %s, attribute2: %s}", t.attribute1, t.attribute2);
}