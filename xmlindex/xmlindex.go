package xmlindex

//go:generate xsdgen -ns "" -pkg xmlindex ../index.xsd

const (
	CompoundKindClass     = CompoundKind("class")
	CompoundKindStruct    = CompoundKind("struct")
	CompoundKindUnion     = CompoundKind("union")
	CompoundKindInterface = CompoundKind("interface")
	CompoundKindProtocol  = CompoundKind("protocol")
	CompoundKindCategory  = CompoundKind("category")
	CompoundKindException = CompoundKind("exception")
	CompoundKindFile      = CompoundKind("file")
	CompoundKindNamespace = CompoundKind("namespace")
	CompoundKindGroup     = CompoundKind("group")
	CompoundKindPage      = CompoundKind("page")
	CompoundKindExample   = CompoundKind("example")
	CompoundKindDir       = CompoundKind("dir")
)

const (
	MemberKindDefine    = MemberKind("define")
	MemberKindProperty  = MemberKind("property")
	MemberKindEvent     = MemberKind("event")
	MemberKindVariable  = MemberKind("variable")
	MemberKindTypedef   = MemberKind("typedef")
	MemberKindEnum      = MemberKind("enum")
	MemberKindEnumvalue = MemberKind("enumvalue")
	MemberKindFunction  = MemberKind("function")
	MemberKindSignal    = MemberKind("signal")
	MemberKindPrototype = MemberKind("prototype")
	MemberKindFriend    = MemberKind("friend")
	MemberKindDcop      = MemberKind("dcop")
	MemberKindSlot      = MemberKind("slot")
)
