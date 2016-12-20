package xmlapi

type (
	// FacWarSystems is the structure used to represent the
	// /map/FacWarSystems.xml.aspx response
	FacWarSystems struct {
		EveAPI
		SolarSystems []struct {
			SolarSystemID         int    `xml:"solarSystemID,attr"`
			SolarSystemName       string `xml:"solarSystemName,attr"`
			OccupyingFactionID    int    `xml:"occupyingFactionID,attr"`
			OccupyingFactionName  string `xml:"occupyingFactionName,attr"`
			OwningFactionID       int    `xml:"owningFactionID,attr"`
			OwningFactionName     string `xml:"owningFactionName,attr"`
			Contested             bool   `xml:"contested,attr"`
			VictoryPoints         int    `xml:"victoryPoints,attr"`
			VictoryPointThreshold int    `xml:"victoryPointThreshold,attr"`
		} `xml:"result>rowset>row"`
	}

	// Jumps is the structure used to represent the
	// /map/Jumps.xml.aspx response
	Jumps struct {
		EveAPI
		DataTime
		SolarSystems []struct {
			SolarSystemID int `xml:"solarSystemID,attr"`
			ShipJumps     int `xml:"shipJumps,attr"`
		} `xml:"result>rowset>row"`
	}

	// Kills is the structure used to represent the
	// /map/Kills.xml.aspx response
	Kills struct {
		EveAPI
		DataTime
		SolarSystems []struct {
			SolarSystemID int `xml:"solarSystemID,attr"`
			ShipKills     int `xml:"shipKills,attr"`
			FactionKills  int `xml:"factionKills,attr"`
			PodKills      int `xml:"podKills,attr"`
		} `xml:"result>rowset>row"`
	}

	// Sovereignty is the structure used to represent the
	// /map/Sovereignty.xml.aspx response
	Sovereignty struct {
		EveAPI
		DataTime
		SolarSystems []struct {
			SolarSystemID   int    `xml:"solarSystemID,attr"`
			SolarSystemName string `xml:"solarSystemName,attr"`
			FactionID       int    `xml:"factionID,attr"`
			AllianceID      int    `xml:"allianceID,attr"`
			CorporationID   int    `xml:"corporationID,attr"`
		} `xml:"result>rowset>row"`
	}
)
