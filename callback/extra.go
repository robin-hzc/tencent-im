package callback

type OfflinePushInfo struct {
	PushFlag int64  `json:"PushFlag"`
	Desc     string `json:"Desc"`
	Ext      string `json:"Ext"`
	Title    string `json:"Title"`
	ApnsInfo struct {
		Sound            string `json:"Sound"`
		BadgeMode        int64  `json:"BadgeMode"`
		ContentAvailable int64  `json:"ContentAvailable"`
		IsVoipPush       int64  `json:"IsVoipPush"`
	} `json:"ApnsInfo"`
	AndroidInfo struct {
		Sound              string `json:"Sound"`
		OPPOChannelID      string `json:"OPPOChannelID"`
		VIVOClassification int64  `json:"VIVOClassification"`
		HuaWeiCategory     string `json:"HuaWeiCategory"`
		VIVOCategory       string `json:"VIVOCategory"`
		OPPONotifyLevel    int64  `json:"OPPONotifyLevel"`
	} `json:"AndroidInfo"`
}
