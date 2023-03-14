package bgapi2

const (
	strLen = 1024
)

const (
	EventMode_Unregistered = iota
	EventMode_Polling
	EventMode_EventHandler
)

const (
	NodePayloadType_Unknown   = "Unknown"
	NodePayloadType_Image     = "Image"
	NodePayloadType_RawData   = "RawData"
	NodePayloadType_File      = "File"
	NodePayloadType_ChunkData = "ChunkData"
	NodePayloadType_CustomId  = "CustomID_1000"
	NodePayloadType_Jpeg      = "Jpeg"
)

const (
	NodeInterface_Category    = "ICategory"
	NodeInterface_Integer     = "IInteger"
	NodeInterface_Register    = "IRegister"
	NodeInterface_Boolean     = "IBoolean"
	NodeInterface_Command     = "ICommand"
	NodeInterface_Float       = "IFloat"
	NodeInterface_Enumeration = "IEnumeration"
	NodeInterface_String      = "IString"
	NodeInterface_Port        = "IPort"
)

const (
	NodeVisibility_Beginner  = "Beginner"
	NodeVisibility_Expert    = "Expert"
	NodeVisibility_Guru      = "Guru"
	NodeVisibility_Invisible = "Invisible"
)

const (
	NodeAccess_Readwrite      = "RW"
	NodeAccess_Readonly       = "RO"
	NodeAccess_Writeonly      = "WO"
	NodeAccess_Notavailable   = "NA"
	NodeAccess_Notimplemented = "NI"
)

const (
	NodeRepresentation_Linear      = "Linear"
	NodeRepresentation_Logarithmic = "Logarithmic"
	NodeRepresentation_Purenumber  = "PureNumber"
	NodeRepresentation_Boolean     = "Boolean"
	NodeRepresentation_Hexnumber   = "HexNumber"
	NodeRepresentation_Ipv4address = "IPV4Address"
	NodeRepresentation_Macaddress  = "MACAddress"
)
