package weekly

type UndergroundSystem struct {
	Travels map[int]*Travel
	Costs map[string][]float64
}

type Travel struct {
	Id int
	Station string
	T int
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{make(map[int]*Travel), make(map[string][]float64)}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int)  {
	if _, ok := this.Travels[id];!ok {
		cur := &Travel{id, stationName, t}
		this.Travels[id] = cur
	}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int)  {
	if v, ok := this.Travels[id]; ok {
		startStation := v.Station
		startTime := v.T
		key := startStation + "_" + stationName
		cost := t - startTime
		if _, ok := this.Costs[key]; ok {
			this.Costs[key] = append(this.Costs[key], float64(cost))
		} else {
			this.Costs[key] = make([]float64, 0)
			this.Costs[key] = append(this.Costs[key], float64(cost))
		}
		delete(this.Travels, id)
	}
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	key := startStation + "_" + endStation
	var res float64
	if _, ok := this.Costs[key]; ok {
		var sum float64
		for _, v := range this.Costs[key] {
			sum += v
		}
		res = sum/float64(len(this.Costs[key]))
	}
	return res
}



/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */