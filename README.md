# Fortnite Tracker API

This project allows the user to make requests to the [FortniteTracker API](https://fortnitetracker.com/site-api) with a valid API token.


### Author: Kevin Gentile
### Contact: kevin@kevingentile.com

## Example Usage
```
	type Data struct {
		Wins  int     `json:"wins"`
		KDR   float64 `json:"kdr"`
		Kills int     `json:"kills"`
	}

	platform := "pc"
	username := "laughingcabbage"
	apiToken := os.Getenv("KEY")

	profile, err := tracker.GetProfile(platform, username, apiToken)
	if err != nil {
		handleError(err, w)
		return
	}

	kills, err := tracker.GetKills(profile)
	if err != nil {
		handleError(err, w)
		return
	}
	data.Kills = kills

	wins, err := tracker.GetWins(profile)
	if err != nil {
		handleError(err, w)
		return
	}
	data.Wins = wins

	kdr, err := tracker.GetCurrentKDR(profile)
	if err != nil {
		handleError(err, w)
		return
	}
	data.KDR = kdr

	fmt.Println("My Stats: ", data)
}
```