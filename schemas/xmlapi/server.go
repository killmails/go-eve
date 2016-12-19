package xmlapi

type ServerStatus struct {
	EveApi
	ServerOpen    bool `xml:"result>serverOpen"`
	OnlinePlayers int  `xml:"result>onlinePlayers"`
}
