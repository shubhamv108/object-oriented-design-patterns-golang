package typedobject

type TypeObjectBuilder struct {
	attribute1 string
	attribute2 string
}

func (t *TypeObjectBuilder) Build() *TypedObject {
	return &TypedObject{
		attribute1: t.attribute1,
		attribute2: t.attribute2,
	}
}

func NewTypeObjectBuilder() *TypeObjectBuilder {
	return &TypeObjectBuilder{}
}

func (t *TypeObjectBuilder) WithAttribute1(attribute1 string) *TypeObjectBuilder {
	t.attribute1 = attribute1
	return t
}

func (t *TypeObjectBuilder) WithAttribute2(attribute2 string) *TypeObjectBuilder {
	t.attribute2 = attribute2
	return t
}