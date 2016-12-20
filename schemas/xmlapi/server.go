package xmlapi

// ServerStatus is the structure used to represent the
// /server/ServerStatus.xml.aspx response
type ServerStatus struct {
	EveAPI
	ServerOpen    bool `xml:"result>serverOpen"`
	OnlinePlayers int  `xml:"result>onlinePlayers"`
}
