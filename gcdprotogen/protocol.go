package main

// Top level API protocol file
type ProtoDebuggerApi struct {
	Version *ProtoApiVersion `json:"version"`
	Domains []*ProtoDomain   `json:"domains"`
}

// Version information
type ProtoApiVersion struct {
	Major string `json:"major"`
	Minor string `json:"minor"`
}

// The Domain (contains all objects, their type/commands/events)
type ProtoDomain struct {
	Domain      string          `json:"domain"`
	Description string          `json:"description,omitempty"`
	Types       []*ProtoType    `json:"types,omitempty"`
	Commands    []*ProtoCommand `json:"commands,omitempty"`
	Events      []*ProtoEvent   `json:"events,omitempty"`
	Hidden      bool            `json:"hidden,omitempty"`
	Items       *ProtoItem      `json:"items,omitempty"`
}

// A Type which represents objects specific to the API method
type ProtoType struct {
	Id          string           `json:"id"`
	Type        string           `json:"type"`
	Description string           `json:"description,omitempty"`
	Enum        []string         `json:"enum,omitempty"`
	Properties  []*ProtoProperty `json:"properties,omitempty"`
	Hidden      bool             `json:"hidden,omitempty"`
	Items       *ProtoItem       `json:"items,omitempty"`
	MinItems    int64            `json:"minItems,omitempty"`
	MaxItems    int64            `json:"maxItems,omitempty"`
}

func (p *ProtoType) IsNonPropertiesObject() bool {
	return (p.Type == "object" && len(p.Properties) == 0)
}

func (p *ProtoType) GetUnderlyingType() string {
	return p.Type
}

func (p *ProtoType) GetArrayType() string {
	if p.Type != "array" || p.Items == nil {
		return ""
	}
	if p.Items.Type != "" {
		return p.Items.Type
	}
	if p.Items.Ref != "" {
		return p.Items.Ref
	}
	return ""
}

func (p *ProtoType) IsArray() bool {
	return p.Type == "array"
}

// A property & Parameter type used by both commands & types
type ProtoProperty struct {
	Name        string           `json:"name"`
	Type        string           `json:"type,omitempty"`
	Description string           `json:"description,omitempty"`
	Ref         string           `json:"$ref,omitempty"`
	Optional    bool             `json:"optional,omitempty"`
	Hidden      bool             `json:"hidden,omitempty"`
	Enum        []string         `json:"enum,omitempty"`
	Items       *ProtoItem       `json:"items,omitempty"`
	Properties  []*ProtoProperty `json:"properties,omitempty"`
}

func (p *ProtoProperty) IsNonPropertiesObject() bool {
	return (p.Type == "object" && len(p.Properties) == 0)
}

func (p *ProtoProperty) GetUnderlyingType() string {
	return p.Type
}

func (p *ProtoProperty) GetArrayType() string {
	if p.Type != "array" || p.Items == nil {
		return ""
	}
	if p.Items.Type != "" {
		return p.Items.Type
	}
	if p.Items.Ref != "" {
		return p.Items.Ref
	}
	return ""
}

func (p *ProtoProperty) IsArray() bool {
	return p.Type == "array"
}

// An item used by types, properties and events.
type ProtoItem struct {
	Type        string           `json:"type,omitempty"`
	Ref         string           `json:"$ref,omitempty"`
	Properties  []*ProtoProperty `json:"properties,omitempty"`
	Description string           `json:"description,omitempty"`
	Enum        []string         `json:"enum,omitempty"`
}

// The API Command call.
type ProtoCommand struct {
	Name        string                 `json:"name"`
	Type        string                 `json:"type,omitempty"`
	Description string                 `json:"description,omitempty"`
	Handlers    []string               `json:"handlers,omitempty"`
	Parameters  []*ProtoProperty       `json:"parameters,omitempty"`
	Returns     []*ProtoCommandReturns `json:"returns,omitempty"`
	Hidden      bool                   `json:"hidden,omitempty"`
	Async       bool                   `json:"async,omitempty"`
	Redirect    string                 `json:"redirect,omitempty"`
}

// Parameters to the API Command call
type ProtoCommandParameters struct {
	Name        string `json:"name"`
	Type        string `json:"type,omitempty"`
	Ref         string `json:"$ref,omitempty"`
	Description string `json:"description,omitempty"`
	Optional    bool   `json:"optional,omitempty"`
}

func (p *ProtoCommandParameters) IsNonPropertiesObject() bool {
	return (p.Type == "object" && len(p.Properties) == 0)
}

func (p *ProtoCommandParameters) GetUnderlyingType() string {
	return p.Type
}

func (p *ProtoCommandParameters) GetArrayType() string {
	if p.Type != "array" || p.Items == nil {
		return ""
	}
	if p.Items.Type != "" {
		return p.Items.Type
	}
	if p.Items.Ref != "" {
		return p.Items.Ref
	}
	return ""
}

func (p *ProtoCommandParameters) IsArray() bool {
	return p.Type == "array"
}

// The return parameters for an API call
type ProtoCommandReturns struct {
	Name        string     `json:"name"`
	Type        string     `json:"type,omitempty"`
	Ref         string     `json:"$ref,omitempty"`
	Items       *ProtoItem `json:"items,omitempty"`
	Description string     `json:"description,omitempty"`
}

// An event, asynchornous events that can come in once
// enabled.
type ProtoEvent struct {
	Name        string           `json:"name"`
	Description string           `json:"description,omitempty"`
	Parameters  []*ProtoProperty `json:"parameters,omitempty"`
	Deprecated  bool             `json:"deprecated,omitempty"`
	Hidden      bool             `json:"hidden,omitempty"`
}
