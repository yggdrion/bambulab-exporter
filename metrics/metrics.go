package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

const metricPrefix = "bambulab_"

func makeCollector[C prometheus.Collector](collector C) C {
	prometheus.MustRegister(collector)
	return collector
}

// ClearAllMetrics deletes all metric values for a given printer when it goes offline.
// This prevents stale metrics from being reported when the printer is not connected.
func ClearAllMetrics(printerName string) {
	// Online metrics
	OnlineAhb.DeleteLabelValues(printerName)
	OnlineRfid.DeleteLabelValues(printerName)
	OnlineVersion.DeleteLabelValues(printerName)

	// Print metrics
	NozzleTemperature.DeleteLabelValues(printerName)
	NozzleTargetTemperature.DeleteLabelValues(printerName)
	BedTemperature.DeleteLabelValues(printerName)
	BedTargetTemperature.DeleteLabelValues(printerName)
	ChamberTemperature.DeleteLabelValues(printerName)
	HeatbreakFanSpeed.DeleteLabelValues(printerName)
	CoolingFanSpeed.DeleteLabelValues(printerName)
	BigFan1Speed.DeleteLabelValues(printerName)
	BigFan2Speed.DeleteLabelValues(printerName)
	PrintProgress.DeleteLabelValues(printerName)
	PrintRemainingTime.DeleteLabelValues(printerName)
	PrintSpeedMagnitude.DeleteLabelValues(printerName)
	PrintSpeedLevel.DeleteLabelValues(printerName)
	PrintError.DeleteLabelValues(printerName)
	WifiSignal.DeleteLabelValues(printerName)
	LayerNumber.DeleteLabelValues(printerName)
	TotalLayers.DeleteLabelValues(printerName)
	AmsStatus.DeleteLabelValues(printerName)
	AmsRfidStatus.DeleteLabelValues(printerName)
	HardwareSwitchState.DeleteLabelValues(printerName)
	Stage.DeleteLabelValues(printerName)
	PrintStage.DeleteLabelValues(printerName)
	PrintSubStage.DeleteLabelValues(printerName)
	GcodeFilePreparePercent.DeleteLabelValues(printerName)
	QueueNumber.DeleteLabelValues(printerName)
	QueueTotal.DeleteLabelValues(printerName)
	QueueEstimatedTime.DeleteLabelValues(printerName)
	QueueStatus.DeleteLabelValues(printerName)
	McPrintLineNumber.DeleteLabelValues(printerName)
	NozzleDiameter.DeleteLabelValues(printerName)
	CalibrationK.DeleteLabelValues(printerName)
	FanGear.DeleteLabelValues(printerName)
	HomeFlag.DeleteLabelValues(printerName)
	Flag3.DeleteLabelValues(printerName)
	CalibrationVersion.DeleteLabelValues(printerName)
	SdcardPresent.DeleteLabelValues(printerName)
	ForceUpgrade.DeleteLabelValues(printerName)
	WifiSignalRaw.DeleteLabelValues(printerName)
	GcodeState.DeleteLabelValues(printerName)
	PrintType.DeleteLabelValues(printerName)
	MessProductionState.DeleteLabelValues(printerName)
	SequenceId.DeleteLabelValues(printerName)
	Msg.DeleteLabelValues(printerName)

	// AMS metrics (single printer label)
	AmsExistBits.DeleteLabelValues(printerName)
	TrayExistBits.DeleteLabelValues(printerName)
	TrayIsBblBits.DeleteLabelValues(printerName)
	TrayTarget.DeleteLabelValues(printerName)
	TrayCurrent.DeleteLabelValues(printerName)
	TrayPrevious.DeleteLabelValues(printerName)
	TrayReadDoneBits.DeleteLabelValues(printerName)
	TrayReadingBits.DeleteLabelValues(printerName)
	AmsVersion.DeleteLabelValues(printerName)
	AmsInsertFlag.DeleteLabelValues(printerName)
	AmsPowerOnFlag.DeleteLabelValues(printerName)

	// AMS metrics with multiple labels (printer, ams, tray)
	// We need to delete all possible combinations
	// Since we don't know which AMS/tray combinations exist, we'll try common ones
	for amsId := 0; amsId < 4; amsId++ {
		amsIdStr := string(rune('0' + amsId))
		AmsHumidity.DeleteLabelValues(printerName, amsIdStr)
		AmsHumidityRaw.DeleteLabelValues(printerName, amsIdStr)
		AmsTemp.DeleteLabelValues(printerName, amsIdStr)
		AmsDryTime.DeleteLabelValues(printerName, amsIdStr)
		AmsInfo.DeleteLabelValues(printerName, amsIdStr)

		for trayId := 0; trayId < 4; trayId++ {
			trayIdStr := string(rune('0' + trayId))
			AmsTrayId.DeleteLabelValues(printerName, amsIdStr, trayIdStr)
			AmsTrayState.DeleteLabelValues(printerName, amsIdStr, trayIdStr)
			AmsTrayRemain.DeleteLabelValues(printerName, amsIdStr, trayIdStr)
			AmsTrayK.DeleteLabelValues(printerName, amsIdStr, trayIdStr)
			AmsTrayN.DeleteLabelValues(printerName, amsIdStr, trayIdStr)
			AmsTrayCaliIdx.DeleteLabelValues(printerName, amsIdStr, trayIdStr)
			AmsTrayTotalLen.DeleteLabelValues(printerName, amsIdStr, trayIdStr)
			AmsTrayDiameter.DeleteLabelValues(printerName, amsIdStr, trayIdStr)
			AmsTrayTemp.DeleteLabelValues(printerName, amsIdStr, trayIdStr)
			AmsTrayTime.DeleteLabelValues(printerName, amsIdStr, trayIdStr)
			AmsTrayBedTempType.DeleteLabelValues(printerName, amsIdStr, trayIdStr)
			AmsTrayBedTemp.DeleteLabelValues(printerName, amsIdStr, trayIdStr)
			AmsTrayNozzleTempMax.DeleteLabelValues(printerName, amsIdStr, trayIdStr)
			AmsTrayNozzleTempMin.DeleteLabelValues(printerName, amsIdStr, trayIdStr)
			AmsTrayCtype.DeleteLabelValues(printerName, amsIdStr, trayIdStr)
		}
	}

	// Network metrics
	NetConf.DeleteLabelValues(printerName)

	// IP camera metrics
	IpcamDev.DeleteLabelValues(printerName)
	IpcamRecord.DeleteLabelValues(printerName)
	IpcamTimelapse.DeleteLabelValues(printerName)
	IpcamTutkServer.DeleteLabelValues(printerName)
	IpcamModeBits.DeleteLabelValues(printerName)

	// Upload metrics
	UploadProgress.DeleteLabelValues(printerName)
	UploadStatus.DeleteLabelValues(printerName)

	// Upgrade state metrics
	UpgradeStateSequenceId.DeleteLabelValues(printerName)
	UpgradeStateProgress.DeleteLabelValues(printerName)
	UpgradeStateStatus.DeleteLabelValues(printerName)
	UpgradeStateConsistencyRequest.DeleteLabelValues(printerName)
	UpgradeStateDisState.DeleteLabelValues(printerName)
	UpgradeStateErrCode.DeleteLabelValues(printerName)
	UpgradeStateForceUpgrade.DeleteLabelValues(printerName)
	UpgradeStateModule.DeleteLabelValues(printerName)
	UpgradeStateNewVersionState.DeleteLabelValues(printerName)
	UpgradeStateCurStateCode.DeleteLabelValues(printerName)
	UpgradeStateIdx2.DeleteLabelValues(printerName)

	// VT tray metrics
	VtTrayId.DeleteLabelValues(printerName)
	VtTrayRemain.DeleteLabelValues(printerName)
	VtTrayK.DeleteLabelValues(printerName)
	VtTrayN.DeleteLabelValues(printerName)
	VtTrayCaliIdx.DeleteLabelValues(printerName)
	VtTrayDiameter.DeleteLabelValues(printerName)
	VtTrayTemp.DeleteLabelValues(printerName)
	VtTrayTime.DeleteLabelValues(printerName)
	VtTrayBedTempType.DeleteLabelValues(printerName)
	VtTrayBedTemp.DeleteLabelValues(printerName)
	VtTrayNozzleTempMax.DeleteLabelValues(printerName)
	VtTrayNozzleTempMin.DeleteLabelValues(printerName)
	VtTrayCtype.DeleteLabelValues(printerName)

	// Lights metrics
	LightsReportChamberLightMode.DeleteLabelValues(printerName)
}
