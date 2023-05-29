// 1603. Design Parking System
// https://leetcode.com/problems/design-parking-system/

package designparkingsystem

type ParkingSystem struct {
	slots map[int]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		slots: map[int]int{
			1: big,
			2: medium,
			3: small,
		},
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.slots[carType] != 0 {
		this.slots[carType] -= 1
		return true
	}
	return false
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
