package xmlindex

type CompoundKind string

type CompoundType struct {
	Name   string       `xml:" name"`
	Member []MemberType `xml:" member,omitempty"`
	Refid  string       `xml:"refid,attr"`
	Kind   CompoundKind `xml:"kind,attr"`
}

type DoxygenType struct {
	Compound []CompoundType `xml:" compound,omitempty"`
	Version  string         `xml:"version,attr"`
}

type MemberKind string

type MemberType struct {
	Name  string     `xml:" name"`
	Refid string     `xml:"refid,attr"`
	Kind  MemberKind `xml:"kind,attr"`
}

type Self struct {
	Doxygenindex DoxygenType `xml:" doxygenindex"`
}
