package constants

const (
	// Central AC DefaultTempareture
	DefaultCoolingTemp = 22.0
	DefaultHeatingTemp = 28.0

	// CentralAC Status
	CentralStatusOff     = 0
	CentralStatusStandBy = 1
	CentralStatusActive  = 2

	// CentralAC Mode
	CoolMode = 1
	HeatMode = 2

	// CentralAC Default RefreshRate
	DefaultRefreshRate = 1

	// Room Status
	RoomStatusOn  = 1
	RoomStatusOff = 0

	// FanSpeed
	FanSpeedLow    = 1
	FanSpeedMedium = 2
	FanSpeedHigh   = 3

	// RequestStatus
	RequestStatusPending = 1
	RequestStatusDoing   = 2
	RequestStatusDone    = 3

	// FanSpeed Consumed
	LowSpeedConsumedPerSecond    = 0.8
	MediumSpeedConsumedPerSecond = 1
	HighSpeedConsumedPerSecond   = 1.2

	CostPerEnergy = 5
)

var RoomStatusToString = map[int]string{
	RoomStatusOn:  "on",
	RoomStatusOff: "off",
}

var FanSpeedToString = map[int]string{
	FanSpeedLow:    "low",
	FanSpeedMedium: "medium",
	FanSpeedHigh:   "high",
}

var FanSpeedToInt = map[string]int{
	"low":    FanSpeedLow,
	"medium": FanSpeedMedium,
	"high":   FanSpeedHigh,
}

var RequestStatusToString = map[int]string{
	RequestStatusPending: "Pending",
	RequestStatusDoing:   "Doing",
	RequestStatusDone:    "Done",
}

var CentralACModeToString = map[int]string{
	CoolMode: "Cool",
	HeatMode: "Heat",
}

var CentralACStatusToString = map[int]string{
	CentralStatusOff:     "Off",
	CentralStatusStandBy: "StandBy",
	CentralStatusActive:  "Active",
}
