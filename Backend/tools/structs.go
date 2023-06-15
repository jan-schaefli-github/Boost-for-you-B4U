package tools

// Clan is a struct for a clan
type Clan struct {
	Tag string `json:"tag"`
}

// Person is a struct for a person
type Person struct {
	Tag        string `json:"tag"`
	WholeFame  string `json:"wholeFame"`
	ClanStatus string `json:"clanStatus"`
	JoinDate   string `json:"joinDate"`
	FkClan     string `json:"fk_clan"`
}
