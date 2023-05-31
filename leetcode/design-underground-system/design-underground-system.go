// 1396. Design Underground System
// https://leetcode.com/problems/design-underground-system/

package designundergroundsystem

type checkin struct {
	stationName string
	t           int
}

type averageTime struct {
	travelTime int
	count      int
}

type UndergroundSystem struct {
	checkins     map[int]checkin
	averageTimes map[string]averageTime
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		checkins:     make(map[int]checkin),
		averageTimes: make(map[string]averageTime),
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.checkins[id] = checkin{stationName, t}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	checkinStation, checkinTime := this.checkins[id].stationName, this.checkins[id].t
	travelTime := t - checkinTime
	stations := checkinStation + "-" + stationName
	if avT, ok := this.averageTimes[stations]; !ok {
		this.averageTimes[stations] = averageTime{travelTime, 1}
	} else {
		avT.travelTime += travelTime
		avT.count++
		this.averageTimes[stations] = avT
	}
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	stations := startStation + "-" + endStation
	avT := float64(this.averageTimes[stations].travelTime) / float64(this.averageTimes[stations].count)
	return avT
}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
