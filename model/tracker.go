package model

type CostTracker struct {
	costs          []float64
	dayTotal       float64
	weekTotal      float64
	monthTotal     float64
	currentDay     int
	currentWeekDay int
}

func NewCostTracker() *CostTracker {
	return &CostTracker{
		costs:          make([]float64, 30),
		dayTotal:       0,
		weekTotal:      0,
		monthTotal:     0,
		currentDay:     0,
		currentWeekDay: 0,
	}
}

func (et *CostTracker) AddCost(amount float64) {
	et.dayTotal += amount
	et.weekTotal += amount
	et.monthTotal += amount

	et.monthTotal -= et.costs[et.currentDay]
	if et.currentWeekDay < 7 {
		et.weekTotal -= et.costs[et.currentDay]
	}

	et.costs[et.currentDay] = amount

	et.currentDay = (et.currentDay + 1) % 30
	et.currentWeekDay = (et.currentWeekDay + 1) % 7
}

func (et *CostTracker) GetDayTotal() float64 {
	return et.dayTotal
}

func (et *CostTracker) GetWeekTotal() float64 {
	return et.weekTotal
}

func (et *CostTracker) GetMonthTotal() float64 {
	return et.monthTotal
}
