package footballdataapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type MatchesList struct {
	Competition Competition `json:"competition"`
	ResultSet   ResultSet   `json:"resultSet"`
	Filters     Filter      `json:"filters"`
	Matches     []Match     `json:"matches"`
}

type Match struct {
	Area        Area        `json:"area"`
	Competition Competition `json:"competition"`
	Season      Season      `json:"season"`
	Id          int32       `json:"id"`
	UtcDate     string      `json:"utcDate"`
	Status      string      `json:"status"`
	Group       string      `json:"group"`
	HomeTeam    Team        `json:"homeTeam"`
	AwayTeam    Team        `json:"awayTeam"`
}

func (rcv *httpClient) GetMatchesList(league string) (*MatchesList, error) {
	response, err := rcv.Get("competitions/" + league + "/matches?matchday=1")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	result := &MatchesList{}

	switch response.StatusCode {
	case http.StatusOK:
		if err = json.NewDecoder(response.Body).Decode(&result); err != nil {
			fmt.Println("eRRRR")
			fmt.Println(err)
			return nil, err
		}

		return result, nil
	case http.StatusNotFound:
		return result, nil
	default:
	}

	return nil, errors.New(fmt.Sprintf("Unexpected status code: %d", response.StatusCode))
}

// {
//     "filters": {
//         "season": "2023",
//         "matchday": "1"
//     },
//     "resultSet": {
//         "count": 10,
//         "first": "2023-08-11",
//         "last": "2023-08-14",
//         "played": 10
//     },
//     "competition": {
//         "id": 2021,
//         "name": "Premier League",
//         "code": "PL",
//         "type": "LEAGUE",
//         "emblem": "https://crests.football-data.org/PL.png"
//     },
//     "matches": [
//         {
//             "area": {
//                 "id": 2072,
//                 "name": "England",
//                 "code": "ENG",
//                 "flag": "https://crests.football-data.org/770.svg"
//             },
//             "competition": {
//                 "id": 2021,
//                 "name": "Premier League",
//                 "code": "PL",
//                 "type": "LEAGUE",
//                 "emblem": "https://crests.football-data.org/PL.png"
//             },
//             "season": {
//                 "id": 1564,
//                 "startDate": "2023-08-11",
//                 "endDate": "2024-05-19",
//                 "currentMatchday": 3,
//                 "winner": null
//             },
//             "id": 435943,
//             "utcDate": "2023-08-11T19:00:00Z",
//             "status": "FINISHED",
//             "matchday": 1,
//             "stage": "REGULAR_SEASON",
//             "group": null,
//             "lastUpdated": "2023-08-12T01:04:51Z",
//             "homeTeam": {
//                 "id": 328,
//                 "name": "Burnley FC",
//                 "shortName": "Burnley",
//                 "tla": "BUR",
//                 "crest": "https://crests.football-data.org/328.png"
//             },
//             "awayTeam": {
//                 "id": 65,
//                 "name": "Manchester City FC",
//                 "shortName": "Man City",
//                 "tla": "MCI",
//                 "crest": "https://crests.football-data.org/65.png"
//             },
//             "score": {
//                 "winner": "AWAY_TEAM",
//                 "duration": "REGULAR",
//                 "fullTime": {
//                     "home": 0,
//                     "away": 3
//                 },
//                 "halfTime": {
//                     "home": 0,
//                     "away": 2
//                 }
//             },
//             "odds": {
//                 "msg": "Activate Odds-Package in User-Panel to retrieve odds."
//             },
//             "referees": [
//                 {
//                     "id": 11585,
//                     "name": "Craig Pawson",
//                     "type": "REFEREE",
//                     "nationality": "England"
//                 }
//             ]
//         }
// 	]
// }
