package xmlapi

type (
	AllianceList struct {
		EveApi
		Alliances []struct {
			Name               string `xml:"name,attr"`
			ShortName          string `xml:"shortName,attr"`
			AllianceID         int    `xml:"allianceID,attr"`
			ExecutorCorpID     int    `xml:"executorCorpID,attr"`
			MemberCount        int    `xml:"memberCount,attr"`
			StartDate          Time   `xml:"startDate,attr"`
			MemberCorporations []struct {
				CorporationID int  `xml:"corporationID,attr"`
				StartDate     Time `xml:"startDate,attr"`
			} `xml:"rowset>row,omitemty"`
		} `xml:"result>rowset>row"`
	}

	CharacterAffiliation struct {
		EveApi
		Characters []struct {
			CharacterID     int    `xml:"characterID,attr"`
			CharacterName   string `xml:"characterName,attr"`
			CorporationID   int    `xml:"corporationID,attr"`
			CorporationName string `xml:"corporationName,attr"`
			AllianceID      int    `xml:"allianceID,attr"`
			AllianceName    string `xml:"allianceName,attr"`
			FactionID       int    `xml:"factionID,attr"`
			FactionName     string `xml:"factionName,attr"`
		} `xml:"result>rowset>row"`
	}

	CharacterID struct {
		EveApi
		Characters []struct {
			Name        string `xml:"name,attr"`
			CharacterID int    `xml:"characterID,attr"`
		} `xml:"result>rowset>row"`
	}

	CharacterInfo struct {
		EveApi
		CharacterID       int     `xml:"result>characterID"`
		CharacterName     string  `xml:"result>characterName"`
		Race              string  `xml:"result>race"`
		BloodlineID       int     `xml:"result>bloodlineID"`
		Bloodline         string  `xml:"result>bloodline"`
		AncestryID        int     `xml:"result>ancestryID"`
		Ancestry          string  `xml:"result>ancestry"`
		AccountBallance   float32 `xml:"result>accountBallance,omitempty"`
		SkillPoints       int     `xml:"result>skillPoints,omitempty"`
		NextTrainingEnds  Time    `xml:"result>nextTrainingEnds,omitempty"`
		ShipName          string  `xml:"result>shipName,omitempty"`
		ShipTypeID        int     `xml:"result>shipTypeID,omitempty"`
		ShipTypeName      string  `xml:"result>shipTypeName,omitempty"`
		CorporationID     int     `xml:"result>corporationID"`
		Corporation       string  `xml:"result>corporation"`
		CorporationDate   Time    `xml:"result>corporationDate"`
		AllianceID        int     `xml:"result>allianceID"`
		Alliance          string  `xml:"result>alliacne"`
		AllianceDate      Time    `xml:"result>allianceDate"`
		LastKnownLocation string  `xml:"result>lastKnownLocation,omitempty"`
		SecurityStatus    float32 `xml:"result>securityStatus,omitempty"`
		EmploymentHistory []struct {
			RecordID        int    `xml:"recordID,attr"`
			CorporationID   int    `xml:"corporationID,attr"`
			CorporationName string `xml:"corporationName,attr"`
			StartDate       Time   `xml:"startDate"`
		} `xml:"result>rowset>row"`
	}

	CharacterName struct {
		EveApi
		Characters []struct {
			Name        string `xml:"name,attr"`
			CharacterID int    `xml:"characterID,attr"`
		} `xml:"result>rowset>row"`
	}

	ConquerableStationList struct {
		EveApi
		Outposts []struct {
			StationID       int    `xml:"stationID,attr"`
			StationName     string `xml:"stationName,attr"`
			StationTypeID   int    `xml:"stationTypeID,attr"`
			SolarSystemID   int    `xml:"solarSystemID,attr"`
			CorporationID   int    `xml:"corporationID,attr"`
			CorporationName string `xml:"corporationName,attr"`
			X               int64  `xml:"x,attr"`
			Y               int64  `xml:"y,attr"`
			Z               int64  `xml:"z,attr"`
		} `xml:"result>rowset>row"`
	}

	ErrorList struct {
		EveApi
		Errors []struct {
			ErrorCode int    `xml:"errorCode,attr"`
			ErrorText string `xml:"errorText,attr"`
		} `xml:"result>rowset>row"`
	}

	RefTypes struct {
		EveApi
		RefTypes []struct {
			RefTypeID   int    `xml:"refTypeID,attr"`
			RefTypeName string `xml:"refTypeName,attr"`
		} `xml:"result>rowset>row"`
	}

	TypeName struct {
		EveApi
		Types []struct {
			TypeID   int    `xml:"typeID,attr"`
			TypeName string `xml:"typeName,attr"`
		} `xml:"result>rowset>row"`
	}
)
