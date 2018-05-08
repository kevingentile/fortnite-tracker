package tracker

import "time"

// TestResourcePath holds the path to test resources
const testResourcePath = "test-resource/"

// TestKey is a LIVE KEY. Should limit requests
const testKey = ".tracker-key"

// requestLimitDuration is the amout of time required between requests. 0.5 per second
const requestLimitDuration = time.Second * 2

// headerKeyName is the header fild name used to supply an API token to fortnite tracker
const headerKeyName = "TRN-Api-Key"

// GET https://api.fortnitetracker.com/v1/profile/{platform}/{epic-nickname}

//Version contains the current API version for fornite tracker
const Version = "v1"

//ProfileRoute contains the core route for requesting profile resources
const ProfileRoute = "https://api.fortnitetracker.com/v1/profile/"

//TODO validate JSON fields

// StatField is an abstract JSON object for a stats field
type StatField struct {
	Label        string  `json:"label"`
	Field        string  `json:"field"`
	Category     string  `json:"category"`
	ValueInt     int     `json:"valueInt"`
	Value        float64 `json:"value"`
	Rank         int     `json:"rank"`
	Percentile   float64 `json:"percentile"`
	DisplayValue string  `json:"displayValue"`
}

// Profile json
type Profile struct {
	AccountID      string        `json:"accountId"`
	Platform       int           `json:"platformId"`
	PlatformName   string        `json:"platformNameLong"`
	EpicUserHandle string        `json:"epicUserHandle"`
	Stats          PersonalStats `json:"stats"`
}

// Stat json
type Stat struct {
	TRNRating     StatField `json:"trnRating"`
	Score         StatField `json:"score"`
	Top1          StatField `json:"top1"`
	Top3          StatField `json:"top3"`
	Top5          StatField `json:"top5"`
	Top6          StatField `json:"top6"`
	Top10         StatField `json:"top10"`
	Top12         StatField `json:"top12"`
	Top25         StatField `json:"top25"`
	KD            StatField `json:"kd"`
	WinRatio      StatField `json:"winRatio"`
	Matches       StatField `json:"matches"`
	Kills         StatField `json:"kills"`
	KPG           StatField `json:"kpg"`
	ScorePerMatch StatField `json:"scorePerMatch"`
}

// Personal
type PersonalStats struct {
	P2       Stat
	P10      Stat
	P9       Stat
	CurrP10  Stat
	CurrP9   Stat
	PriorP2  Stat
	PriorP10 Stat
	PriorP9  Stat
}

// LifeTimeStats holds an array of lifetime stats
type LifeTimeStats []LifeTimeKey

// LifeTimeKey holds the key and value for lifetime stats map
type LifeTimeKey struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// RecentMatches
type RecentMatches []Match

// Match
type Match struct {
	ID              int     `json:"id"`
	AcccountID      string  `json:"accountId"`
	Playlist        string  `json:"playlist"`
	Kills           int     `json:"kills"`
	MinutesPlayed   int     `json:"minutesPlayed"`
	Top1            int     `json:"top1"`
	Top5            int     `json:"top5"`
	Top6            int     `json:"top6"`
	Top10           int     `json:"top10"`
	Top12           int     `json:"top12"`
	Top25           int     `json:"top25"`
	Matches         int     `json:"matches"`
	Top3            int     `json:"top3"`
	DateCollected   string  `json:"dateCollected"`
	Score           int     `json:"score"`
	Platform        int     `json:"platform"`
	TRNRating       float64 `json:"trnRating"`
	TRNRatingChange float64 `json:"trnRatingChange"`
}
