package main

type Clan struct {
	Tag string `json:"tag"`
}

type Person struct {
	Tag        string `json:"tag"`
	WholeFame  string `json:"wholeFame"`
	ClanStatus string `json:"clanStatus"`
	JoinDate   string `json:"joinDate"`
	FkClan     string `json:"fk_clan"`
}
