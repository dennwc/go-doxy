package xmlfile

type ChildnodeType struct {
	Edgelabel []string         `xml:" edgelabel,omitempty" json:",omitempty"`
	Refid     string           `xml:"refid,attr,omitempty" json:",omitempty"`
	Relation  DoxGraphRelation `xml:"relation,attr,omitempty" json:",omitempty"`
}

type CodelineType struct {
	Highlight []HighlightType `xml:" highlight,omitempty" json:",omitempty"`
	Lineno    int             `xml:"lineno,attr,omitempty" json:",omitempty"`
	Refid     string          `xml:"refid,attr,omitempty" json:",omitempty"`
	Refkind   DoxRefKind      `xml:"refkind,attr,omitempty" json:",omitempty"`
	External  DoxBool         `xml:"external,attr,omitempty" json:",omitempty"`
}

type CompoundRefType struct {
	Value string            `xml:",chardata"`
	Refid string            `xml:"refid,attr,omitempty" json:",omitempty"`
	Prot  DoxProtectionKind `xml:"prot,attr,omitempty" json:",omitempty"`
	Virt  DoxVirtualKind    `xml:"virt,attr,omitempty" json:",omitempty"`
}

type CompounddefType struct {
	Compoundname        string                 `xml:" compoundname"`
	Title               string                 `xml:" title,omitempty" json:",omitempty"`
	Basecompoundref     []CompoundRefType      `xml:" basecompoundref,omitempty" json:",omitempty"`
	Derivedcompoundref  []CompoundRefType      `xml:" derivedcompoundref,omitempty" json:",omitempty"`
	Includes            []IncType              `xml:" includes,omitempty" json:",omitempty"`
	Includedby          []IncType              `xml:" includedby,omitempty" json:",omitempty"`
	Incdepgraph         GraphType              `xml:" incdepgraph,omitempty" json:",omitempty"`
	Invincdepgraph      GraphType              `xml:" invincdepgraph,omitempty" json:",omitempty"`
	Innerdir            []RefType              `xml:" innerdir,omitempty" json:",omitempty"`
	Innerfile           []RefType              `xml:" innerfile,omitempty" json:",omitempty"`
	Innerclass          []RefType              `xml:" innerclass,omitempty" json:",omitempty"`
	Innernamespace      []RefType              `xml:" innernamespace,omitempty" json:",omitempty"`
	Innerpage           []RefType              `xml:" innerpage,omitempty" json:",omitempty"`
	Innergroup          []RefType              `xml:" innergroup,omitempty" json:",omitempty"`
	Templateparamlist   *TemplateparamlistType `xml:" templateparamlist,omitempty" json:",omitempty"`
	Sectiondef          []SectiondefType       `xml:" sectiondef,omitempty" json:",omitempty"`
	Briefdescription    *DescriptionType       `xml:" briefdescription,omitempty" json:",omitempty"`
	Detaileddescription *DescriptionType       `xml:" detaileddescription,omitempty" json:",omitempty"`
	Inheritancegraph    GraphType              `xml:" inheritancegraph,omitempty" json:",omitempty"`
	Collaborationgraph  GraphType              `xml:" collaborationgraph,omitempty" json:",omitempty"`
	Programlisting      ListingType            `xml:" programlisting,omitempty" json:",omitempty"`
	Location            LocationType           `xml:" location,omitempty" json:",omitempty"`
	Listofallmembers    ListofallmembersType   `xml:" listofallmembers,omitempty" json:",omitempty"`
	Kind                DoxCompoundKind        `xml:"kind,attr,omitempty" json:",omitempty"`
	Language            DoxLanguage            `xml:"language,attr,omitempty" json:",omitempty"`
	Prot                DoxProtectionKind      `xml:"prot,attr,omitempty" json:",omitempty"`
	Final               DoxBool                `xml:"final,attr,omitempty" json:",omitempty"`
	Sealed              DoxBool                `xml:"sealed,attr,omitempty" json:",omitempty"`
	Abstract            DoxBool                `xml:"abstract,attr,omitempty" json:",omitempty"`
}

type DescriptionType struct {
	Title    string           `xml:" title,omitempty" json:",omitempty"`
	Para     []DocParaType    `xml:" para,omitempty" json:",omitempty"`
	Sect1    []DocSect1Type   `xml:" sect1,omitempty" json:",omitempty"`
	Internal *DocInternalType `xml:" internal,omitempty" json:",omitempty"`
}

type DocAnchorType struct {
}

type DocBlockQuoteType struct {
	Para []DocParaType `xml:" para,omitempty" json:",omitempty"`
}

type DocCaptionType struct {
	Ulink          *DocURLLink     `xml:" ulink"`
	Bold           *DocMarkupType  `xml:" bold"`
	Emphasis       *DocMarkupType  `xml:" emphasis"`
	Computeroutput *DocMarkupType  `xml:" computeroutput"`
	Subscript      *DocMarkupType  `xml:" subscript"`
	Superscript    *DocMarkupType  `xml:" superscript"`
	Center         *DocMarkupType  `xml:" center"`
	Small          *DocMarkupType  `xml:" small"`
	Htmlonly       string          `xml:" htmlonly"`
	Manonly        string          `xml:" manonly"`
	Xmlonly        string          `xml:" xmlonly"`
	Rtfonly        string          `xml:" rtfonly"`
	Latexonly      string          `xml:" latexonly"`
	Dot            string          `xml:" dot"`
	Plantuml       string          `xml:" plantuml"`
	Anchor         DocAnchorType   `xml:" anchor"`
	Formula        DocFormulaType  `xml:" formula"`
	Ref            *DocRefTextType `xml:" ref"`
}

type DocCopyType struct {
	Para     []DocParaType   `xml:" para,omitempty" json:",omitempty"`
	Sect1    []DocSect1Type  `xml:" sect1,omitempty" json:",omitempty"`
	Internal DocInternalType `xml:" internal,omitempty" json:",omitempty"`
	Link     string          `xml:"link,attr,omitempty" json:",omitempty"`
}

type DocEntryType struct {
	Para  []DocParaType `xml:" para,omitempty" json:",omitempty"`
	Thead DoxBool       `xml:"thead,attr,omitempty" json:",omitempty"`
}

type DocFileType struct {
	Ulink          *DocURLLink     `xml:" ulink"`
	Bold           *DocMarkupType  `xml:" bold"`
	Emphasis       *DocMarkupType  `xml:" emphasis"`
	Computeroutput *DocMarkupType  `xml:" computeroutput"`
	Subscript      *DocMarkupType  `xml:" subscript"`
	Superscript    *DocMarkupType  `xml:" superscript"`
	Center         *DocMarkupType  `xml:" center"`
	Small          *DocMarkupType  `xml:" small"`
	Htmlonly       string          `xml:" htmlonly"`
	Manonly        string          `xml:" manonly"`
	Xmlonly        string          `xml:" xmlonly"`
	Rtfonly        string          `xml:" rtfonly"`
	Latexonly      string          `xml:" latexonly"`
	Dot            string          `xml:" dot"`
	Plantuml       string          `xml:" plantuml"`
	Anchor         DocAnchorType   `xml:" anchor"`
	Formula        DocFormulaType  `xml:" formula"`
	Ref            *DocRefTextType `xml:" ref"`

	Name string `xml:"name,attr,omitempty" json:",omitempty"`
}

type DocFormulaType struct {
}

type DocHeadingType struct {
	Ulink          *DocURLLink     `xml:" ulink"`
	Bold           *DocMarkupType  `xml:" bold"`
	Emphasis       *DocMarkupType  `xml:" emphasis"`
	Computeroutput *DocMarkupType  `xml:" computeroutput"`
	Subscript      *DocMarkupType  `xml:" subscript"`
	Superscript    *DocMarkupType  `xml:" superscript"`
	Center         *DocMarkupType  `xml:" center"`
	Small          *DocMarkupType  `xml:" small"`
	Htmlonly       string          `xml:" htmlonly"`
	Manonly        string          `xml:" manonly"`
	Xmlonly        string          `xml:" xmlonly"`
	Rtfonly        string          `xml:" rtfonly"`
	Latexonly      string          `xml:" latexonly"`
	Dot            string          `xml:" dot"`
	Plantuml       string          `xml:" plantuml"`
	Anchor         DocAnchorType   `xml:" anchor"`
	Formula        DocFormulaType  `xml:" formula"`
	Ref            *DocRefTextType `xml:" ref"`

	Level int `xml:"level,attr,omitempty" json:",omitempty"`
}

type DocImageType struct {
	Ulink          *DocURLLink     `xml:" ulink"`
	Bold           *DocMarkupType  `xml:" bold"`
	Emphasis       *DocMarkupType  `xml:" emphasis"`
	Computeroutput *DocMarkupType  `xml:" computeroutput"`
	Subscript      *DocMarkupType  `xml:" subscript"`
	Superscript    *DocMarkupType  `xml:" superscript"`
	Center         *DocMarkupType  `xml:" center"`
	Small          *DocMarkupType  `xml:" small"`
	Htmlonly       string          `xml:" htmlonly"`
	Manonly        string          `xml:" manonly"`
	Xmlonly        string          `xml:" xmlonly"`
	Rtfonly        string          `xml:" rtfonly"`
	Latexonly      string          `xml:" latexonly"`
	Dot            string          `xml:" dot"`
	Plantuml       string          `xml:" plantuml"`
	Anchor         DocAnchorType   `xml:" anchor"`
	Formula        DocFormulaType  `xml:" formula"`
	Ref            *DocRefTextType `xml:" ref"`

	Type   DoxImageKind `xml:"type,attr,omitempty" json:",omitempty"`
	Name   string       `xml:"name,attr,omitempty" json:",omitempty"`
	Width  string       `xml:"width,attr,omitempty" json:",omitempty"`
	Height string       `xml:"height,attr,omitempty" json:",omitempty"`
}

type DocIndexEntryType struct {
	Primaryie   string `xml:" primaryie"`
	Secondaryie string `xml:" secondaryie"`
}

type DocInternalS1Type struct {
	Para  []DocParaType  `xml:" para,omitempty" json:",omitempty"`
	Sect2 []DocSect2Type `xml:" sect2,omitempty" json:",omitempty"`
}

type DocInternalS2Type struct {
	Para  []DocParaType  `xml:" para,omitempty" json:",omitempty"`
	Sect3 []DocSect3Type `xml:" sect3,omitempty" json:",omitempty"`
}

type DocInternalS3Type struct {
	Para  []DocParaType  `xml:" para,omitempty" json:",omitempty"`
	Sect3 []DocSect4Type `xml:" sect3,omitempty" json:",omitempty"`
}

type DocInternalS4Type struct {
	Para []DocParaType `xml:" para,omitempty" json:",omitempty"`
}

type DocInternalType struct {
	Para  []DocParaType  `xml:" para,omitempty" json:",omitempty"`
	Sect1 []DocSect1Type `xml:" sect1,omitempty" json:",omitempty"`
}

type DocLanguageType struct {
	Para   []DocParaType `xml:" para,omitempty" json:",omitempty"`
	Langid string        `xml:"langid,attr,omitempty" json:",omitempty"`
}

type DocListItemType struct {
	Para []DocParaType `xml:" para,omitempty" json:",omitempty"`
}

type DocListType struct {
	Listitem []DocListItemType `xml:" listitem"`
}

type DocMarkupType struct {
	Ulink          *DocURLLink     `xml:" ulink"`
	Bold           *DocMarkupType  `xml:" bold"`
	Emphasis       *DocMarkupType  `xml:" emphasis"`
	Computeroutput *DocMarkupType  `xml:" computeroutput"`
	Subscript      *DocMarkupType  `xml:" subscript"`
	Superscript    *DocMarkupType  `xml:" superscript"`
	Center         *DocMarkupType  `xml:" center"`
	Small          *DocMarkupType  `xml:" small"`
	Htmlonly       string          `xml:" htmlonly"`
	Manonly        string          `xml:" manonly"`
	Xmlonly        string          `xml:" xmlonly"`
	Rtfonly        string          `xml:" rtfonly"`
	Latexonly      string          `xml:" latexonly"`
	Dot            string          `xml:" dot"`
	Plantuml       string          `xml:" plantuml"`
	Anchor         DocAnchorType   `xml:" anchor"`
	Formula        DocFormulaType  `xml:" formula"`
	Ref            *DocRefTextType `xml:" ref"`

	Preformatted   *DocMarkupType      `xml:" preformatted"`
	Programlisting ListingType         `xml:" programlisting"`
	Verbatim       string              `xml:" verbatim"`
	Indexentry     DocIndexEntryType   `xml:" indexentry"`
	Orderedlist    DocListType         `xml:" orderedlist"`
	Itemizedlist   DocListType         `xml:" itemizedlist"`
	Simplesect     DocSimpleSectType   `xml:" simplesect"`
	Title          DocTitleType        `xml:" title"`
	Variablelist   DocVariableListType `xml:" variablelist"`
	Table          DocTableType        `xml:" table"`
	Heading        DocHeadingType      `xml:" heading"`
	Image          DocImageType        `xml:" image"`
	Dotfile        DocFileType         `xml:" dotfile"`
	Mscfile        DocFileType         `xml:" mscfile"`
	Diafile        DocFileType         `xml:" diafile"`
	Toclist        DocTocListType      `xml:" toclist"`
	Language       DocLanguageType     `xml:" language"`
	Parameterlist  DocParamListType    `xml:" parameterlist"`
	Xrefsect       DocXRefSectType     `xml:" xrefsect"`
	Copydoc        DocCopyType         `xml:" copydoc"`
	Blockquote     DocBlockQuoteType   `xml:" blockquote"`
	Parblock       DocParBlockType     `xml:" parblock"`
}

type DocParBlockType struct {
	Para []DocParaType `xml:" para,omitempty" json:",omitempty"`
}

type DocParaType struct {
	Ulink          *DocURLLink     `xml:" ulink"`
	Bold           *DocMarkupType  `xml:" bold"`
	Emphasis       *DocMarkupType  `xml:" emphasis"`
	Computeroutput *DocMarkupType  `xml:" computeroutput"`
	Subscript      *DocMarkupType  `xml:" subscript"`
	Superscript    *DocMarkupType  `xml:" superscript"`
	Center         *DocMarkupType  `xml:" center"`
	Small          *DocMarkupType  `xml:" small"`
	Htmlonly       string          `xml:" htmlonly"`
	Manonly        string          `xml:" manonly"`
	Xmlonly        string          `xml:" xmlonly"`
	Rtfonly        string          `xml:" rtfonly"`
	Latexonly      string          `xml:" latexonly"`
	Dot            string          `xml:" dot"`
	Plantuml       string          `xml:" plantuml"`
	Anchor         DocAnchorType   `xml:" anchor"`
	Formula        DocFormulaType  `xml:" formula"`
	Ref            *DocRefTextType `xml:" ref"`

	Preformatted   *DocMarkupType      `xml:" preformatted"`
	Programlisting ListingType         `xml:" programlisting"`
	Verbatim       string              `xml:" verbatim"`
	Indexentry     DocIndexEntryType   `xml:" indexentry"`
	Orderedlist    DocListType         `xml:" orderedlist"`
	Itemizedlist   DocListType         `xml:" itemizedlist"`
	Simplesect     DocSimpleSectType   `xml:" simplesect"`
	Title          DocTitleType        `xml:" title"`
	Variablelist   DocVariableListType `xml:" variablelist"`
	Table          DocTableType        `xml:" table"`
	Heading        DocHeadingType      `xml:" heading"`
	Image          DocImageType        `xml:" image"`
	Dotfile        DocFileType         `xml:" dotfile"`
	Mscfile        DocFileType         `xml:" mscfile"`
	Diafile        DocFileType         `xml:" diafile"`
	Toclist        DocTocListType      `xml:" toclist"`
	Language       DocLanguageType     `xml:" language"`
	Parameterlist  DocParamListType    `xml:" parameterlist"`
	Xrefsect       DocXRefSectType     `xml:" xrefsect"`
	Copydoc        DocCopyType         `xml:" copydoc"`
	Blockquote     DocBlockQuoteType   `xml:" blockquote"`
	Parblock       DocParBlockType     `xml:" parblock"`
}

type DocParamListItem struct {
	Parameternamelist    []DocParamNameList `xml:" parameternamelist,omitempty" json:",omitempty"`
	Parameterdescription *DescriptionType   `xml:" parameterdescription"`
}

type DocParamListType struct {
	Parameteritem []DocParamListItem `xml:" parameteritem,omitempty" json:",omitempty"`
	Kind          DoxParamListKind   `xml:"kind,attr,omitempty" json:",omitempty"`
}

type DocParamName struct {
	Ref       RefTextType `xml:" ref,omitempty" json:",omitempty"`
	Direction DoxParamDir `xml:"direction,attr,omitempty" json:",omitempty"`
}

type DocParamNameList struct {
	Parametertype []DocParamType `xml:" parametertype,omitempty" json:",omitempty"`
	Parametername []DocParamName `xml:" parametername,omitempty" json:",omitempty"`
}

type DocParamType struct {
	Ref RefTextType `xml:" ref,omitempty" json:",omitempty"`
}

type DocRefTextType struct {
	Ulink          *DocURLLink     `xml:" ulink"`
	Bold           *DocMarkupType  `xml:" bold"`
	Emphasis       *DocMarkupType  `xml:" emphasis"`
	Computeroutput *DocMarkupType  `xml:" computeroutput"`
	Subscript      *DocMarkupType  `xml:" subscript"`
	Superscript    *DocMarkupType  `xml:" superscript"`
	Center         *DocMarkupType  `xml:" center"`
	Small          *DocMarkupType  `xml:" small"`
	Htmlonly       string          `xml:" htmlonly"`
	Manonly        string          `xml:" manonly"`
	Xmlonly        string          `xml:" xmlonly"`
	Rtfonly        string          `xml:" rtfonly"`
	Latexonly      string          `xml:" latexonly"`
	Dot            string          `xml:" dot"`
	Plantuml       string          `xml:" plantuml"`
	Anchor         DocAnchorType   `xml:" anchor"`
	Formula        DocFormulaType  `xml:" formula"`
	Ref            *DocRefTextType `xml:" ref"`

	Refid    string     `xml:"refid,attr,omitempty" json:",omitempty"`
	Kindref  DoxRefKind `xml:"kindref,attr,omitempty" json:",omitempty"`
	External string     `xml:"external,attr,omitempty" json:",omitempty"`
}

type DocRowType struct {
	Entry []DocEntryType `xml:" entry,omitempty" json:",omitempty"`
}

type DocSect1Type struct {
	Title    string            `xml:" title"`
	Para     []DocParaType     `xml:" para,omitempty" json:",omitempty"`
	Sect2    []DocSect2Type    `xml:" sect2,omitempty" json:",omitempty"`
	Internal DocInternalS1Type `xml:" internal,omitempty" json:",omitempty"`
}

type DocSect2Type struct {
	Title    string            `xml:" title"`
	Para     []DocParaType     `xml:" para,omitempty" json:",omitempty"`
	Sect3    []DocSect3Type    `xml:" sect3,omitempty" json:",omitempty"`
	Internal DocInternalS2Type `xml:" internal,omitempty" json:",omitempty"`
}

type DocSect3Type struct {
	Title    string            `xml:" title"`
	Para     []DocParaType     `xml:" para,omitempty" json:",omitempty"`
	Sect4    []DocSect4Type    `xml:" sect4,omitempty" json:",omitempty"`
	Internal DocInternalS3Type `xml:" internal,omitempty" json:",omitempty"`
}

type DocSect4Type struct {
	Title    string            `xml:" title"`
	Para     []DocParaType     `xml:" para,omitempty" json:",omitempty"`
	Internal DocInternalS4Type `xml:" internal,omitempty" json:",omitempty"`
}

type DocSimpleSectType struct {
	Title DocTitleType      `xml:" title,omitempty" json:",omitempty"`
	Para  []DocParaType     `xml:" para"`
	Kind  DoxSimpleSectKind `xml:"kind,attr,omitempty" json:",omitempty"`
}

type DocTableType struct {
	Row     []DocRowType   `xml:" row,omitempty" json:",omitempty"`
	Caption DocCaptionType `xml:" caption,omitempty" json:",omitempty"`
	Rows    int            `xml:"rows,attr,omitempty" json:",omitempty"`
	Cols    int            `xml:"cols,attr,omitempty" json:",omitempty"`
}

type DocTitleType struct {
	Ulink          *DocURLLink     `xml:" ulink"`
	Bold           *DocMarkupType  `xml:" bold"`
	Emphasis       *DocMarkupType  `xml:" emphasis"`
	Computeroutput *DocMarkupType  `xml:" computeroutput"`
	Subscript      *DocMarkupType  `xml:" subscript"`
	Superscript    *DocMarkupType  `xml:" superscript"`
	Center         *DocMarkupType  `xml:" center"`
	Small          *DocMarkupType  `xml:" small"`
	Htmlonly       string          `xml:" htmlonly"`
	Manonly        string          `xml:" manonly"`
	Xmlonly        string          `xml:" xmlonly"`
	Rtfonly        string          `xml:" rtfonly"`
	Latexonly      string          `xml:" latexonly"`
	Dot            string          `xml:" dot"`
	Plantuml       string          `xml:" plantuml"`
	Anchor         DocAnchorType   `xml:" anchor"`
	Formula        DocFormulaType  `xml:" formula"`
	Ref            *DocRefTextType `xml:" ref"`
}

type DocTocItemType struct {
	Ulink          *DocURLLink     `xml:" ulink"`
	Bold           *DocMarkupType  `xml:" bold"`
	Emphasis       *DocMarkupType  `xml:" emphasis"`
	Computeroutput *DocMarkupType  `xml:" computeroutput"`
	Subscript      *DocMarkupType  `xml:" subscript"`
	Superscript    *DocMarkupType  `xml:" superscript"`
	Center         *DocMarkupType  `xml:" center"`
	Small          *DocMarkupType  `xml:" small"`
	Htmlonly       string          `xml:" htmlonly"`
	Manonly        string          `xml:" manonly"`
	Xmlonly        string          `xml:" xmlonly"`
	Rtfonly        string          `xml:" rtfonly"`
	Latexonly      string          `xml:" latexonly"`
	Dot            string          `xml:" dot"`
	Plantuml       string          `xml:" plantuml"`
	Anchor         DocAnchorType   `xml:" anchor"`
	Formula        DocFormulaType  `xml:" formula"`
	Ref            *DocRefTextType `xml:" ref"`
}

type DocTocListType struct {
	Tocitem []DocTocItemType `xml:" tocitem,omitempty" json:",omitempty"`
}

type DocURLLink struct {
	Ulink          *DocURLLink     `xml:" ulink"`
	Bold           *DocMarkupType  `xml:" bold"`
	Emphasis       *DocMarkupType  `xml:" emphasis"`
	Computeroutput *DocMarkupType  `xml:" computeroutput"`
	Subscript      *DocMarkupType  `xml:" subscript"`
	Superscript    *DocMarkupType  `xml:" superscript"`
	Center         *DocMarkupType  `xml:" center"`
	Small          *DocMarkupType  `xml:" small"`
	Htmlonly       string          `xml:" htmlonly"`
	Manonly        string          `xml:" manonly"`
	Xmlonly        string          `xml:" xmlonly"`
	Rtfonly        string          `xml:" rtfonly"`
	Latexonly      string          `xml:" latexonly"`
	Dot            string          `xml:" dot"`
	Plantuml       string          `xml:" plantuml"`
	Anchor         DocAnchorType   `xml:" anchor"`
	Formula        DocFormulaType  `xml:" formula"`
	Ref            *DocRefTextType `xml:" ref"`

	Url string `xml:"url,attr,omitempty" json:",omitempty"`
}

type DocVarListEntryType struct {
	Term DocTitleType `xml:" term"`
}

type DocVariableListType struct {
	Varlistentry DocVarListEntryType `xml:" varlistentry"`
	Listitem     DocListItemType     `xml:" listitem"`
}

type DocXRefSectType struct {
	Xreftitle       []string         `xml:" xreftitle,omitempty" json:",omitempty"`
	Xrefdescription *DescriptionType `xml:" xrefdescription"`
}

type DoxAccessor string

type DoxCharRange string

type DoxCompoundKind string

type DoxGraphRelation string

type DoxHighlightClass string

type DoxImageKind string

type DoxLanguage string

type DoxMemberKind string

type DoxParamDir string

type DoxParamListKind string

type DoxProtectionKind string

type DoxRefKind string

type DoxRefQualifierKind string

type DoxSectionKind string

type DoxSimpleSectKind string

type DoxVersionNumber string

type DoxVirtualKind string

type DoxygenType struct {
	Compounddef CompounddefType  `xml:" compounddef,omitempty" json:",omitempty"`
	Version     DoxVersionNumber `xml:"version,attr"`
}

type EnumvalueType struct {
	Name                string            `xml:" name"`
	Initializer         *LinkedTextType   `xml:" initializer,omitempty" json:",omitempty"`
	Briefdescription    *DescriptionType  `xml:" briefdescription,omitempty" json:",omitempty"`
	Detaileddescription *DescriptionType  `xml:" detaileddescription,omitempty" json:",omitempty"`
	Prot                DoxProtectionKind `xml:"prot,attr,omitempty" json:",omitempty"`
}

type GraphType struct {
	Node []NodeType `xml:" node" json:",omitempty"`
}

type HighlightType struct {
	Sp    string            `xml:" sp"`
	Ref   RefTextType       `xml:" ref"`
	Class DoxHighlightClass `xml:"class,attr,omitempty" json:",omitempty"`
}

type IncType struct {
	Value string  `xml:",chardata"`
	Refid string  `xml:"refid,attr,omitempty" json:",omitempty"`
	Local DoxBool `xml:"local,attr,omitempty" json:",omitempty"`
}

type LinkType struct {
	Refid    string `xml:"refid,attr,omitempty" json:",omitempty"`
	External string `xml:"external,attr,omitempty" json:",omitempty"`
}

type LinkedTextType struct {
	Type string        `xml:",chardata" json:",omitempty"`
	Ref  []RefTextType `xml:" ref,omitempty" json:",omitempty"`
}

type ListingType struct {
	Codeline []CodelineType `xml:" codeline,omitempty" json:",omitempty"`
}

type ListofallmembersType struct {
	Member []MemberRefType `xml:" member,omitempty" json:",omitempty"`
}

type LocationType struct {
	File      string `xml:"file,attr,omitempty" json:",omitempty"`
	Line      int    `xml:"line,attr,omitempty" json:",omitempty"`
	Column    int    `xml:"column,attr,omitempty" json:",omitempty"`
	Bodyfile  string `xml:"bodyfile,attr,omitempty" json:",omitempty"`
	Bodystart int    `xml:"bodystart,attr,omitempty" json:",omitempty"`
	Bodyend   int    `xml:"bodyend,attr,omitempty" json:",omitempty"`
}

type MemberRefType struct {
	Scope          string            `xml:" scope"`
	Name           string            `xml:" name"`
	Refid          string            `xml:"refid,attr,omitempty" json:",omitempty"`
	Prot           DoxProtectionKind `xml:"prot,attr,omitempty" json:",omitempty"`
	Virt           DoxVirtualKind    `xml:"virt,attr,omitempty" json:",omitempty"`
	Ambiguityscope string            `xml:"ambiguityscope,attr,omitempty" json:",omitempty"`
}

type MemberdefType struct {
	Templateparamlist   *TemplateparamlistType `xml:" templateparamlist,omitempty" json:",omitempty"`
	Type                *LinkedTextType        `xml:" type,omitempty" json:",omitempty"`
	Definition          string                 `xml:" definition,omitempty" json:",omitempty"`
	Argsstring          string                 `xml:" argsstring,omitempty" json:",omitempty"`
	Name                string                 `xml:" name"`
	Read                string                 `xml:" read,omitempty" json:",omitempty"`
	Write               string                 `xml:" write,omitempty" json:",omitempty"`
	Bitfield            string                 `xml:" bitfield,omitempty" json:",omitempty"`
	Reimplements        []ReimplementType      `xml:" reimplements,omitempty" json:",omitempty"`
	Reimplementedby     []ReimplementType      `xml:" reimplementedby,omitempty" json:",omitempty"`
	Param               []ParamType            `xml:" param,omitempty" json:",omitempty"`
	Enumvalue           []EnumvalueType        `xml:" enumvalue,omitempty" json:",omitempty"`
	Initializer         *LinkedTextType        `xml:" initializer,omitempty" json:",omitempty"`
	Exceptions          *LinkedTextType        `xml:" exceptions,omitempty" json:",omitempty"`
	Briefdescription    *DescriptionType       `xml:" briefdescription,omitempty" json:",omitempty"`
	Detaileddescription *DescriptionType       `xml:" detaileddescription,omitempty" json:",omitempty"`
	Inbodydescription   *DescriptionType       `xml:" inbodydescription,omitempty" json:",omitempty"`
	Location            LocationType           `xml:" location"`
	References          []ReferenceType        `xml:" references,omitempty" json:",omitempty"`
	Referencedby        []ReferenceType        `xml:" referencedby,omitempty" json:",omitempty"`
	Kind                DoxMemberKind          `xml:"kind,attr,omitempty" json:",omitempty"`
	Prot                DoxProtectionKind      `xml:"prot,attr,omitempty" json:",omitempty"`
	Static              DoxBool                `xml:"static,attr,omitempty" json:",omitempty"`
	Const               DoxBool                `xml:"const,attr,omitempty" json:",omitempty"`
	Explicit            DoxBool                `xml:"explicit,attr,omitempty" json:",omitempty"`
	Inline              DoxBool                `xml:"inline,attr,omitempty" json:",omitempty"`
	Refqual             DoxRefQualifierKind    `xml:"refqual,attr,omitempty" json:",omitempty"`
	Virt                DoxVirtualKind         `xml:"virt,attr,omitempty" json:",omitempty"`
	Volatile            DoxBool                `xml:"volatile,attr,omitempty" json:",omitempty"`
	Mutable             DoxBool                `xml:"mutable,attr,omitempty" json:",omitempty"`
	Readable            DoxBool                `xml:"readable,attr,omitempty" json:",omitempty"`
	Writable            DoxBool                `xml:"writable,attr,omitempty" json:",omitempty"`
	Initonly            DoxBool                `xml:"initonly,attr,omitempty" json:",omitempty"`
	Settable            DoxBool                `xml:"settable,attr,omitempty" json:",omitempty"`
	Gettable            DoxBool                `xml:"gettable,attr,omitempty" json:",omitempty"`
	Final               DoxBool                `xml:"final,attr,omitempty" json:",omitempty"`
	Sealed              DoxBool                `xml:"sealed,attr,omitempty" json:",omitempty"`
	New                 DoxBool                `xml:"new,attr,omitempty" json:",omitempty"`
	Add                 DoxBool                `xml:"add,attr,omitempty" json:",omitempty"`
	Remove              DoxBool                `xml:"remove,attr,omitempty" json:",omitempty"`
	Raise               DoxBool                `xml:"raise,attr,omitempty" json:",omitempty"`
	Optional            DoxBool                `xml:"optional,attr,omitempty" json:",omitempty"`
	Required            DoxBool                `xml:"required,attr,omitempty" json:",omitempty"`
	Accessor            DoxAccessor            `xml:"accessor,attr,omitempty" json:",omitempty"`
	Attribute           DoxBool                `xml:"attribute,attr,omitempty" json:",omitempty"`
	Property            DoxBool                `xml:"property,attr,omitempty" json:",omitempty"`
	Readonly            DoxBool                `xml:"readonly,attr,omitempty" json:",omitempty"`
	Bound               DoxBool                `xml:"bound,attr,omitempty" json:",omitempty"`
	Removable           DoxBool                `xml:"removable,attr,omitempty" json:",omitempty"`
	Contrained          DoxBool                `xml:"contrained,attr,omitempty" json:",omitempty"`
	Transient           DoxBool                `xml:"transient,attr,omitempty" json:",omitempty"`
	Maybevoid           DoxBool                `xml:"maybevoid,attr,omitempty" json:",omitempty"`
	Maybedefault        DoxBool                `xml:"maybedefault,attr,omitempty" json:",omitempty"`
	Maybeambiguous      DoxBool                `xml:"maybeambiguous,attr,omitempty" json:",omitempty"`
}

type NodeType struct {
	Label     string          `xml:" label"`
	Link      *LinkType       `xml:" link,omitempty" json:",omitempty"`
	Childnode []ChildnodeType `xml:" childnode,omitempty" json:",omitempty"`
}

type ParamType struct {
	Type             *LinkedTextType  `xml:" type,omitempty" json:",omitempty"`
	Declname         string           `xml:" declname,omitempty" json:",omitempty"`
	Defname          string           `xml:" defname,omitempty" json:",omitempty"`
	Array            string           `xml:" array,omitempty" json:",omitempty"`
	Defval           *LinkedTextType  `xml:" defval,omitempty" json:",omitempty"`
	Typeconstraint   *LinkedTextType  `xml:" typeconstraint,omitempty" json:",omitempty"`
	Briefdescription *DescriptionType `xml:" briefdescription,omitempty" json:",omitempty"`
}

type RefTextType struct {
	Value    string     `xml:",chardata"`
	Refid    string     `xml:"refid,attr,omitempty" json:",omitempty"`
	Kindref  DoxRefKind `xml:"kindref,attr,omitempty" json:",omitempty"`
	External string     `xml:"external,attr,omitempty" json:",omitempty"`
	Tooltip  string     `xml:"tooltip,attr,omitempty" json:",omitempty"`
}

type RefType struct {
	Value string            `xml:",chardata"`
	Refid string            `xml:"refid,attr,omitempty" json:",omitempty"`
	Prot  DoxProtectionKind `xml:"prot,attr,omitempty" json:",omitempty"`
}

type ReferenceType struct {
	Refid       string `xml:"refid,attr,omitempty" json:",omitempty"`
	Compoundref string `xml:"compoundref,attr,omitempty" json:",omitempty"`
	Startline   int    `xml:"startline,attr,omitempty" json:",omitempty"`
	Endline     int    `xml:"endline,attr,omitempty" json:",omitempty"`
}

type ReimplementType struct {
	Value string `xml:",chardata"`
	Refid string `xml:"refid,attr,omitempty" json:",omitempty"`
}

type SectiondefType struct {
	Header      string           `xml:" header,omitempty" json:",omitempty"`
	Description *DescriptionType `xml:" description,omitempty" json:",omitempty"`
	Memberdef   []MemberdefType  `xml:" memberdef"`
	Kind        DoxSectionKind   `xml:"kind,attr,omitempty" json:",omitempty"`
}

type Self struct {
	Doxygen DoxygenType `xml:" doxygen"`
}

type TemplateparamlistType struct {
	Param []ParamType `xml:" param,omitempty" json:",omitempty"`
}
