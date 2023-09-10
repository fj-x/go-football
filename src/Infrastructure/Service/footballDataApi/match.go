package footballdataapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
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
	StartDate   string      `json:"utcDate"`
	Status      string      `json:"status"`
	Group       string      `json:"group"`
	HomeTeam    Team        `json:"homeTeam"`
	AwayTeam    Team        `json:"awayTeam"`
	Goals       []Goal      `json:"goals"`
}

func (rcv *httpClient) GetMatchesList() (*MatchesList, error) {
	response, err := rcv.Get("competitions/PL/matches?matchday=4")
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

func (rcv *httpClient) FetchMatchInfo(matchId int32) (*Match, error) {
	srrr := "matches/" + strconv.Itoa(int(matchId))
	response, err := rcv.Get(srrr)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	result := &Match{}

	switch response.StatusCode {
	case http.StatusOK:
		if err = json.NewDecoder(response.Body).Decode(&result); err != nil {
			fmt.Println("eRRRR")
			fmt.Println(err)
			return nil, err
		}

		return result, nil
	case http.StatusNotFound:
		return nil, nil
	default:
	}

	return nil, errors.New(fmt.Sprintf("Unexpected status code: %d", response.StatusCode))
}

// {
//     "area": {
//         "id": 2072,
//         "name": "England",
//         "code": "ENG",
//         "flag": "https://crests.football-data.org/770.svg"
//     },
//     "competition": {
//         "id": 2021,
//         "name": "Premier League",
//         "code": "PL",
//         "type": "LEAGUE",
//         "emblem": "https://crests.football-data.org/PL.png"
//     },
//     "season": {
//         "id": 733,
//         "startDate": "2021-08-13",
//         "endDate": "2022-05-22",
//         "currentMatchday": 37,
//         "winner": null,
//         "stages": [
//             "REGULAR_SEASON"
//         ]
//     },
//     "id": 327117,
//     "utcDate": "2022-02-12T12:30:00Z",
//     "status": "FINISHED",
//     "minute": 90,
//     "injuryTime": 6,
//     "attendance": 73084,
//     "venue": "Old Trafford",
//     "matchday": 25,
//     "stage": "REGULAR_SEASON",
//     "group": null,
//     "lastUpdated": "2022-05-17T17:33:08Z",
//     "homeTeam": {
//         "id": 66,
//         "name": "Manchester United FC",
//         "shortName": "Man United",
//         "tla": "MUN",
//         "crest": "https://crests.football-data.org/66.png",
//         "coach": {
//             "id": 59070,
//             "name": "Ralf Rangnick",
//             "nationality": "Germany"
//         },
//         "leagueRank": null,
//         "formation": "4-2-3-1"
//     },
//     "awayTeam": {
//         "id": 340,
//         "name": "Southampton FC",
//         "shortName": "Southampton",
//         "tla": "SOU",
//         "crest": "https://crests.football-data.org/340.png",
//         "coach": {
//             "id": 43924,
//             "name": "Ralph Hasenhüttl",
//             "nationality": "Austria"
//         },
//         "leagueRank": null,
//         "formation": "4-4-2"
//     },
//     "score": {
//         "winner": "DRAW",
//         "duration": "REGULAR",
//         "fullTime": {
//             "home": 1,
//             "away": 1
//         },
//         "halfTime": {
//             "home": 1,
//             "away": 0
//         }
//     },
//     "goals": [
//         {
//             "minute": 21,
//             "injuryTime": null,
//             "type": "REGULAR",
//             "team": {
//                 "id": 66,
//                 "name": "Manchester United FC"
//             },
//             "scorer": {
//                 "id": 146,
//                 "name": "Jadon Sancho"
//             },
//             "assist": {
//                 "id": 3331,
//                 "name": "Marcus Rashford"
//             },
//             "score": {
//                 "home": 1,
//                 "away": 0
//             }
//         },
//         {
//             "minute": 48,
//             "injuryTime": null,
//             "type": "REGULAR",
//             "team": {
//                 "id": 340,
//                 "name": "Southampton FC"
//             },
//             "scorer": {
//                 "id": 4119,
//                 "name": "Che Adams"
//             },
//             "assist": {
//                 "id": 16060,
//                 "name": "Mohamed Elyounoussi"
//             },
//             "score": {
//                 "home": 1,
//                 "away": 1
//             }
//         }
//     ],
//     "substitutions": [
//         {
//             "minute": 46,
//             "team": {
//                 "id": 340,
//                 "name": "Southampton FC"
//             },
//             "playerOut": {
//                 "id": 8086,
//                 "name": "Jan Bednarek"
//             },
//             "playerIn": {
//                 "id": 8085,
//                 "name": "Jack Stephens"
//             }
//         },
//         {
//             "minute": 71,
//             "team": {
//                 "id": 340,
//                 "name": "Southampton FC"
//             },
//             "playerOut": {
//                 "id": 16060,
//                 "name": "Mohamed Elyounoussi"
//             },
//             "playerIn": {
//                 "id": 168712,
//                 "name": "Valentino Livramento"
//             }
//         }
//     ]
//}
