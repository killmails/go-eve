package xmlapi

type (
	AccountStatus struct {
		EveApi
		PaidUntil              Time `xml:"result>paidUntil"`
		CreateDate             Time `xml:"result>createDate"`
		LogonCount             int  `xml:"result>logonCount"`
		LogonMinutes           int  `xml:"result>logonMinutes"`
		MultiCharacterTraining []struct {
			TrainingEnd Time `xml:"trainingEnd,attr"`
		} `xml:"result>rowset>row,omitempty"`
	}

	APIKeyInfo struct {
		EveApi
		Key struct {
			AccessMask int64  `xml:"accessMask,attr"`
			Type       string `xml:"type,attr"`
			Expires    Time   `xml:"expires,attr,omitempty"`
			Characters []struct {
				CharacterID     int    `xml:"characterID,attr"`
				CharacterName   string `xml:"characterName,attr"`
				CorporationID   int    `xml:"corporationID,attr"`
				CorporationName string `xml:"corporationName,attr"`
				AllianceID      int    `xml:"allianceID,attr"`
				AllianceName    string `xml:"allianceName,attr"`
				FactionID       int    `xml:"factionID,attr"`
				FactionName     string `xml:"factionName,attr"`
			} `xml:"rowset>row"`
		} `xml:"result>key"`
	}

	Characters struct {
		EveApi
		Characters []struct {
			Name            string `xml:"name,attr"`
			CharacterID     int    `xml:"characterID,attr"`
			CorporationName string `xml:"corporationName,attr"`
			CorporationID   int    `xml:"corporationID,attr"`
			AllianceID      int    `xml:"allianceID,attr"`
			AllianceName    string `xml:"allianceName,attr"`
			FactionID       int    `xml:"factionID,attr"`
			FactionName     string `xml:"factionName,attr"`
		} `xml:"result>rowset>row"`
	}
)
