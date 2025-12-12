// 3433. Count Mentions Per User
// https://leetcode.com/problems/count-mentions-per-user/

package countmentionsperuser

import (
	"sort"
	"strconv"
	"strings"
)

func countMentions(numberOfUsers int, events [][]string) []int {
	sort.Slice(events, func(i, j int) bool {
		timeI, _ := strconv.Atoi(events[i][1])
		timeJ, _ := strconv.Atoi(events[j][1])

		// Primary: sort by timestamp
		if timeI != timeJ {
			return timeI < timeJ
		}
		// Secondary: OFFLINE (not MESSAGE) comes before MESSAGE
		// In Python: True > False, so MESSAGE sorts after
		return events[i][0] != "MESSAGE" && events[j][0] == "MESSAGE"
	})

	counter := make([]int, numberOfUsers)
	nextOnline := make([]int, numberOfUsers)
	for _, event := range events {
		curTime, _ := strconv.Atoi(event[1])
		if event[0] == "MESSAGE" {
			if event[2] == "ALL" {
				for userId := range numberOfUsers {
					counter[userId]++
				}
			} else if event[2] == "HERE" {
				for userId, nextOnlineTime := range nextOnline {
					if nextOnlineTime <= curTime {
						counter[userId]++
					}
				}
			} else {
				parts := strings.Fields(event[2])
				for _, strID := range parts {
					userID, _ := strconv.Atoi(strID[2:])
					counter[userID]++
				}
			}
		} else {
			userID, _ := strconv.Atoi(event[2])
			nextOnline[userID] = curTime + 60
		}
	}
	return counter
}
