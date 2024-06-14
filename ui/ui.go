package ui

import (
	"center-air-conditioning-interactive/constants"
	"center-air-conditioning-interactive/model"
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var ac *model.CentralAC
var uiUpdate chan func()

// 定义绑定变量
var mode binding.String
var status binding.String
var refreshRate binding.String

func RunUI() {
	// 初始化中央空调实例
	ac = model.GetCentralACInstance()

	// 初始化Fyne应用和窗口
	a := app.New()
	w := a.NewWindow("中央空调")

	// 初始化UI更新通道
	uiUpdate = make(chan func())

	// 设置主界面为初始内容
	w.SetContent(buildMainScreen(w))
	w.Resize(fyne.NewSize(600, 400))

	// 开启一个goroutine用于处理UI更新
	go func() {
		for update := range uiUpdate {
			update()
		}
	}()

	// 显示窗口并运行事件循环
	w.ShowAndRun()
}

func buildMainScreen(w fyne.Window) fyne.CanvasObject {
	// 初始化绑定变量
	mode = binding.NewString()
	status = binding.NewString()
	refreshRate = binding.BindString(&ac.RefreshRate)

	updateModeString()
	updateStatusString()

	modeLabel := widget.NewLabelWithData(mode)
	statusLabel := widget.NewLabelWithData(status)
	refreshRateLabel := widget.NewLabelWithData(refreshRate)

	// 开关按钮
	toggleButton := widget.NewButton("开关 On/Off", func() {
		if ac.IsTurnOff() {
			ac.SetStatus(constants.CentralStatusStandBy)
		} else {
			ac.SetStatus(constants.CentralStatusOff)
		}
		uiUpdate <- func() {
			updateStatusString()
		}
	})

	// 工作模式按钮
	modeButton := widget.NewButton("模式 Mode (Cool/Heat)", func() {
		if ac.Mode == constants.CoolMode {
			ac.SetMode(constants.HeatMode)
		} else {
			ac.SetMode(constants.CoolMode)
		}
		uiUpdate <- func() {
			updateModeString()
		}
	})

	// 刷新频率按钮
	refreshRateEntry := widget.NewEntry()
	refreshRateEntry.SetPlaceHolder("请输入刷新频率(rate/s)")
	setRefreshRateButton := widget.NewButton("Set", func() {
		rate, err := strconv.Atoi(refreshRateEntry.Text)
		if err == nil {
			ac.RefreshRate = rate
			uiUpdate <- func() {
				refreshRate.Set(strconv.Itoa(ac.RefreshRate))
			}
		} else {
			uiUpdate <- func() {
				dialog.ShowError(err, w)
			}
		}
	})

	refreshRateBox := container.NewHBox(widget.NewLabel("设置刷新频率: "), container.New(layout.NewGridWrapLayout(fyne.NewSize(200, refreshRateEntry.MinSize().Height)), refreshRateEntry), setRefreshRateButton)

	// 静态数据部分
	staticData := container.NewVBox(
		widget.NewForm(
			widget.NewFormItem("Mode", modeLabel),
			widget.NewFormItem("Status", statusLabel),
			widget.NewFormItem("Refresh Rate", refreshRateLabel),
		),
	)
	staticDataBox := container.NewVBox(
		widget.NewCard("", "", staticData),
	)

	controlPanel := container.NewVBox(
		toggleButton,
		modeButton,
		refreshRateBox,
	)

	return container.NewBorder(nil, nil, nil, nil,
		container.NewVBox(
			staticDataBox,
			controlPanel,
		))
}

func updateModeString() {
	modeString := constants.CentralACModeToString[ac.Mode]
	mode.Set(modeString)
}

func updateStatusString() {
	statusString := constants.CentralACStatusToString[ac.Status]
	status.Set(statusString)
}
