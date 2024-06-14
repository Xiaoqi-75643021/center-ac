package model

import (
	"center-air-conditioning-interactive/constants"
	"sync"
)

type CentralAC struct {
	Mode        int // cooling/1 or heating/2
	DefaultTemp float64
	RefreshRate int
	Status      int // off/0 or standby/1 or active/2
}

var CentralACInstance *CentralAC
var CentralACOnce sync.Once

func GetCentralACInstance() *CentralAC {
	CentralACOnce.Do(func() {
		CentralACInstance = &CentralAC{
			Mode:        constants.CoolMode,
			DefaultTemp: constants.DefaultCoolingTemp,
			Status:      constants.CentralStatusStandBy,
			RefreshRate: constants.DefaultRefreshRate,
		}
	})
	return CentralACInstance
}

func (ac *CentralAC) SetMode(mode int) {
	if mode == constants.CoolMode || mode == constants.HeatMode {
		ac.Mode = mode
		if mode == constants.CoolMode {
			ac.DefaultTemp = constants.DefaultCoolingTemp
		} else {
			ac.DefaultTemp = constants.DefaultHeatingTemp
		}
	}
}

func (ac *CentralAC) SetStatus(status int) {
	if status == constants.CentralStatusOff || status == constants.CentralStatusActive || status == constants.CentralStatusStandBy {
		ac.Status = status
	}
}

func (ac *CentralAC) IsTurnOff() bool {
	return ac.Status == constants.CentralStatusOff
}
