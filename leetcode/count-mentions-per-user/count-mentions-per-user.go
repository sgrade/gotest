// 3433. Count Mentions Per User
// https://leetcode.com/problems/count-mentions-per-user/

package countmentionsperuser

import (
	"sort"
	"strconv"
	"strings"
)

func countMentions(numberOfUsers int, events [][]string) []int {
	const (
		eventMessage = "MESSAGE"
		targetAll    = "ALL"
		targetHere   = "HERE"
	)

	// Parse timestamps once during sorting
	sort.Slice(events, func(i, j int) bool {
		timeI, _ := strconv.Atoi(events[i][1])
		timeJ, _ := strconv.Atoi(events[j][1])

		if timeI != timeJ {
			return timeI < timeJ
		}
		// OFFLINE comes before MESSAGE at same timestamp
		return events[i][0] != eventMessage && events[j][0] == eventMessage
	})

	counter := make([]int, numberOfUsers)
	nextOnline := make([]int, numberOfUsers)

	for _, event := range events {
		curTime, _ := strconv.Atoi(event[1])

		if event[0] == eventMessage {
			switch event[2] {
			case targetAll:
				for userID := range numberOfUsers {
					counter[userID]++
				}
			case targetHere:
				for userID, nextOnlineTime := range nextOnline {
					if nextOnlineTime <= curTime {
						counter[userID]++
					}
				}
			default:
				// Specific user mentions
				parts := strings.Fields(event[2])
				for _, strID := range parts {
					userID, _ := strconv.Atoi(strings.TrimPrefix(strID, "id"))
					counter[userID]++
				}
			}
		} else {
			// OFFLINE event
			userID, _ := strconv.Atoi(event[2])
			nextOnline[userID] = curTime + 60
		}
	}

	return counter
}
