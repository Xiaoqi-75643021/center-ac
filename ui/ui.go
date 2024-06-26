package ui

import (
	"center-air-conditioning-interactive/constants"
	"center-air-conditioning-interactive/model"
	"center-air-conditioning-interactive/service"
	"errors"
	"path/filepath"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
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
	w := a.NewWindow("Central Air Conditioner Controller")

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

	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for range ticker.C {
			uiUpdate <- func() {
				updateStatusString()
			}
		}
	}()

	// 显示窗口并运行事件循环
	w.ShowAndRun()
}

func buildMainScreen(w fyne.Window) fyne.CanvasObject {
	// 初始化绑定变量
	mode = binding.NewString()
	status = binding.NewString()
	refreshRate = binding.NewString()

	updateModeString()
	updateStatusString()
	updateRefreshRateString()

	modeLabel := widget.NewLabelWithData(mode)
	statusLabel := widget.NewLabelWithData(status)
	refreshRateLabel := widget.NewLabelWithData(refreshRate)

	// 开关按钮
	toggleButton := widget.NewButton("Toggle Status (On/Off)", func() {
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
	modeButton := widget.NewButton("Toggle Mode (Cool/Heat)", func() {
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
	refreshRateEntry.SetPlaceHolder("Enter Refresh Rate")
	setRefreshRateButton := widget.NewButton("Set", func() {
		rate, err := strconv.Atoi(refreshRateEntry.Text)
		if (err == nil) && (rate > 0) {
			ac.SetRefreshRate(rate)
			uiUpdate <- func() {
				updateRefreshRateString()
			}
		} else {
			uiUpdate <- func() {
				dialog.ShowError(errors.New("Invalid refresh rate"), w)
			}
		}
	})

	refreshRateBox := container.NewHBox(widget.NewLabel("Set Refresh Rate: "), container.New(layout.NewGridWrapLayout(fyne.NewSize(200, refreshRateEntry.MinSize().Height)), refreshRateEntry), setRefreshRateButton)

	// Room ID和Period输入框
	roomIdEntry := widget.NewEntry()
	roomIdEntry.SetPlaceHolder("Enter Room ID")
	periodEntry := widget.NewEntry()
	periodEntry.SetPlaceHolder("Enter Period (daily/weekly/monthly)")

	// 导出报表按钮
	exportReportButton := widget.NewButton("Export Report", func() {
		roomId := roomIdEntry.Text
		period := periodEntry.Text
		if roomId == "" || period == "" {
			dialog.ShowError(errors.New("Room ID and Period must be filled"), w)
			return
		}
		if !(period == "daily" || period == "weekly" || period == "monthly") {
			dialog.ShowError(errors.New("Invalid period"), w)
			return
		}
		if err := ExportRoomReportWithDialog(roomId, period, w); err != nil {
			dialog.ShowError(err, w)
		}
	})

	// Room ID和Period输入框放在按钮下面
	exportReportBox := container.NewVBox(
		widget.NewLabel("Enter Room ID and Period:"),
		roomIdEntry,
		periodEntry,
		exportReportButton,
	)

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
		exportReportBox,
	)

	return container.NewBorder(nil, nil, nil, nil,
		container.NewVBox(
			staticDataBox,
			controlPanel,
		))
}


func updateModeString() {
	modeString := ac.GetModeString()
	mode.Set(modeString)
}

func updateStatusString() {
	statusString := ac.GetStatusString()
	status.Set(statusString)
}

func updateRefreshRateString() {
	refreshRateString := ac.GetRefreshRateString()
	refreshRate.Set(refreshRateString)
}

func ExportRoomReportWithDialog(roomId string, period string, w fyne.Window) error {
	dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
		if err != nil || uri == nil {
			dialog.ShowError(errors.New("未选择文件夹或选择失败"), w)
			return
		}

		savePath := filepath.Join(uri.Path(), "RoomReportLog.log")
		err = service.ExportRoomReport(roomId, period, savePath)
		if err != nil {
			dialog.ShowError(err, w)
		} else {
			dialog.ShowInformation("Success", "Report exported successfully!", w)
		}
	}, w)

	return nil
}