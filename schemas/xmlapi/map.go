package xmlapi

type (
	FacWarSystems struct {
		EveApi
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

	Jumps struct {
		EveApi
		DataTime
		SolarSystems []struct {
			SolarSystemID int `xml:"solarSystemID,attr"`
			ShipJumps     int `xml:"shipJumps,attr"`
		} `xml:"result>rowset>row"`
	}

	Kills struct {
		EveApi
		DataTime
		SolarSystems []struct {
			SolarSystemID int `xml:"solarSystemID,attr"`
			ShipKills     int `xml:"shipKills,attr"`
			FactionKills  int `xml:"factionKills,attr"`
			PodKills      int `xml:"podKills,attr"`
		} `xml:"result>rowset>row"`
	}

	Sovereignty struct {
		EveApi
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
