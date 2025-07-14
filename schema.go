package main

type AuthTicketRequestPayload struct {
	Key      string `json:"key"`
	AppId    uint32 `json:"appid"`
	Ticket   string `json:"ticket"`
	Identity string `json:"identity"`
}

type AuthTicketResponsePayload struct {
	Result          string `json:"result"`
	SteamId         string `json:"steamid"`
	OwnerSteamId    string `json:"ownersteamid"`
	VacBanned       bool   `json:"vacbanned"`
	PublisherBanned bool   `json:"publisherbanned"`
}
