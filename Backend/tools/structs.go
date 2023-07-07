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

type WarLogItems struct {
	Tag             string `json:"tag"`
	Name			string `json:"name"`
	JoinDate        string `json:"joinDate"`
	ClanStatus      int `json:"clanStatus"`
	Role            string `json:"role"`
	Trophies        int `json:"trophies"`
	ClanRank        int `json:"clanRank"`
	Fame            int `json:"fame"`
	DecksUsed       int `json:"decksUsed"`
	MissedDecks     int `json:"missedDecks"`
	RepairPoints    int `json:"repairPoints"`
	BoatAttacks     int `json:"boatAttacks"`
}

type WarLog struct {
	Items []WarLogItems `json:"items"`
	Identifier string `json:"identifier"`
	MaxOffset string `json:"maxOffset"`
}


type ClanWeekReport struct {
	Id       int `json:"id"`
	Fame     int `json:"fame"`
	FameGain int `json:"fame_gain"`
}

// DailyReport is a struct for a daily report
type DailyReport struct {
	ID             int    `json:"id"`
	DecksUsedToday int    `json:"decksUsedToday"`
	Fame           int    `json:"fame"`
	DayIdentifier  string `json:"dayIdentifier"`
	Date           string `json:"date"`
	FkPerson       string `json:"fkPerson"`
}



type ClanWeekReports []ClanWeekReport
