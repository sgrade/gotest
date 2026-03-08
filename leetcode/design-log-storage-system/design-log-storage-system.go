// 635. Design Log Storage System
// design-log-storage-system

package designlogstoragesystem

type log struct {
	id        int
	timestamp string
}

type LogSystem struct {
	logs []log
}

var granLen = map[string]int{
	"Year": 4, "Month": 7, "Day": 10,
	"Hour": 13, "Minute": 16, "Second": 19,
}

func Constructor() LogSystem {
	return LogSystem{}
}

func (this *LogSystem) Put(id int, timestamp string) {
	this.logs = append(this.logs, log{id, timestamp})
}

func (this *LogSystem) Retrieve(start string, end string, granularity string) []int {
	n := granLen[granularity]
	s, e := start[:n], end[:n]
	var res []int
	for _, l := range this.logs {
		t := l.timestamp[:n]
		if t >= s && t <= e {
			res = append(res, l.id)
		}
	}
	return res
}

/**
 * Your LogSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(id,timestamp);
 * param_2 := obj.Retrieve(start,end,granularity);
 */
