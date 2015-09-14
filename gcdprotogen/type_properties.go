package main

type TypeProperties struct {
	ParentType     *Type // the type ref this property belongs to
	protoProperty  *ProtoProperty
	Name           string // property name
	Description    string // property description
	UnderlyingType string
	GoType         string
	Optional       bool   // is this property optional?
	EnumVals       string // possible enum values as a string
	IsRef          bool   // is a reference to another type
}

func NewTypeProperties(parentType *Type, props *ProtoProperty) *TypeProperties {
	tp := &TypeProperties{}
	tp.ParentType = parentType
	tp.protoProperty = props
	tp.Name = props.Name
	tp.Description = props.Description
	tp.Optional = props.Optional
	tp.UnderlyingType = props.Type
	return tp
}

func (p *TypeProperties) IsNonPropertiesObject() bool {
	return (p.UnderlyingType == "object" && len(p.protoProperty.Properties) == 0)
}

func (p *TypeProperties) GetUnderlyingType() string {
	return p.UnderlyingType
}

func (p *TypeProperties) IsArray() bool {
	return p.UnderlyingType == "array"
}

func (p *TypeProperties) GetArrayType() string {
	if p.protoProperty.Items.Type != "" {
		return p.protoProperty.Items.Type
	}

	if p.protoProperty.Items.Ref != "" {
		return p.protoProperty.Items.Ref
	}
	return "object"
}
