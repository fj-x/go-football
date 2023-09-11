# go-football



Live match monitoring may be scheduled in 2 cases:

1.
When user subscribe to a team - go to API and get next match start time for the selected team.
Setup a schedule for starting live match monitoring goroutine at given time for given match.
If the live match monitoring for the given match already seted up do not create the new one.

2.
Once per day call API to get the next match start time for the all subscribed teams.
Setup a schedule for starting live match monitoring goroutines at given times for given matches.


Each live match monitoring goroutine should check if a score has been changed and notify all subsribed users.
If API returns that match ended - we should exit from goroutine.


next steps:
1. interface
2. service loading
3. error handling
4. channels, waitgroup, mutex
5. test