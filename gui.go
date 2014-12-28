package jlsamplergtk

import (
	"github.com/johnnylee/gtk"
	"github.com/johnnylee/statemachine"
)

// Gui is the main struct for our gui.
type Gui struct {
	fsm *statemachine.StateMachine

	// Widgets:
	mainWindow gtk.Widget

	// Loading widgets.
	loadSamplerBtn gtk.Widget
	jackName       gtk.Widget
	loadSpinner    gtk.Widget

	// Info-label.
	samplerInfoLabel gtk.Widget

	// Toolbar buttons.
	loadSettingsBtn gtk.Widget
	saveSettingsBtn gtk.Widget

	// Controls.
	amp          gtk.Widget
	gammaAmp     gtk.Widget
	velMult      gtk.Widget
	cropThresh   gtk.Widget
	tauFadeIn    gtk.Widget
	tau          gtk.Widget
	tauCut       gtk.Widget
	transpose    gtk.Widget
	pitchBendMax gtk.Widget
	mixLayers    gtk.Widget
	gammaLayer   gtk.Widget
	panLow       gtk.Widget
	panHigh      gtk.Widget
	rmsTime      gtk.Widget
	rmsLow       gtk.Widget
	rmsHigh      gtk.Widget
}

func NewGui() *Gui {
	gui := new(Gui)
	gui.fsm = statemachine.New("JLSampler")
	gui.fsm.AddState("NotLoaded", nil, nil, "LoadingSamples")
	gui.fsm.AddState("LoadingSamples", nil, nil)
	gui.fsm.SetInitialState("NotLoaded")
	return gui
}

// Main is the entry point for our gui application.
func (gui *Gui) Main() error {
	gtk.Init()

	// Build the gui from the glade file.
	data, err := Asset("data/gui.glade")
	if err != nil {
		return err
	}
	gtk.AddFromString(data)

	// Create widgets in Go.
	gui.connectWidgets()

	// Show the main window.
	gui.mainWindow.Show()
	gtk.Main()

	return nil
}

func (gui *Gui) connectWidgets() {
	// Main window.
	gui.mainWindow = gtk.GetWidget("mainWindow")

	// Loading widgets.
	gui.loadSamplerBtn = gtk.GetWidget("loadSamplerBtn")
	gui.loadSamplerBtn.SignalConnect("toggled", gui.onLoadSamplerBtn)

	gui.jackName = gtk.GetWidget("jackName")
	gui.loadSpinner = gtk.GetWidget("loadSpinner")

	// Toolbar buttons.
	gui.loadSettingsBtn = gtk.GetWidget("loadSettingsBtn")
	gui.loadSettingsBtn.SignalConnect("clicked", gui.onLoadSettingsBtn)
	gui.saveSettingsBtn = gtk.GetWidget("saveSettingsBtn")
	gui.saveSettingsBtn.SignalConnect("clicked", gui.onSaveSettingsBtn)

	// Control: Amplify
	gui.amp = gtk.GetWidget("amp")
	gui.amp.SpinButtonSetRange(0, 99)
	gui.amp.SpinButtonSetIncrements(0.01, 0.1)
	gui.amp.SignalConnect("value-changed", func(w gtk.Widget) {
		println(w.SpinButtonGetValue())
	})

	// Control: Amp gamma
	gui.gammaAmp = gtk.GetWidget("gammaAmp")
	gui.gammaAmp.SpinButtonSetRange(0.01, 9.99)
	gui.gammaAmp.SpinButtonSetIncrements(0.01, 0.1)
	gui.gammaAmp.SignalConnect("value-changed", func(w gtk.Widget) {
		println(w.SpinButtonGetValue())
	})

	// Control: Velocity mult
	gui.velMult = gtk.GetWidget("velMult")
	gui.velMult.SpinButtonSetRange(0.01, 9.99)
	gui.velMult.SpinButtonSetIncrements(0.01, 0.1)
	gui.velMult.SignalConnect("value-changed", func(w gtk.Widget) {
		println(w.SpinButtonGetValue())
	})

	// Control: Crop threshold
	gui.cropThresh = gtk.GetWidget("cropThresh")
	gui.cropThresh.SpinButtonSetRange(0, 0.999)
	gui.cropThresh.SpinButtonSetIncrements(0.001, 0.01)
	gui.cropThresh.SignalConnect("value-changed", func(w gtk.Widget) {
		println(w.SpinButtonGetValue())
	})

	// Control: Fade-in
	gui.tauFadeIn = gtk.GetWidget("tauFadeIn")
	gui.tauFadeIn.SpinButtonSetRange(0, 9999)
	gui.tauFadeIn.SpinButtonSetIncrements(1, 10)
	gui.tauFadeIn.SignalConnect("value-changed", func(w gtk.Widget) {
		println(w.SpinButtonGetValue())
	})

	// Control: Key-up fade
	gui.tau = gtk.GetWidget("tau")
	gui.tau.SpinButtonSetRange(0, 9999)
	gui.tau.SpinButtonSetIncrements(1, 10)
	gui.tau.SignalConnect("value-changed", func(w gtk.Widget) {
		println(w.SpinButtonGetValue())
	})

	// Control: Key-cut fade
	gui.tauCut = gtk.GetWidget("tauCut")
	gui.tauCut.SpinButtonSetRange(0, 9999)
	gui.tauCut.SpinButtonSetIncrements(1, 10)
	gui.tauCut.SignalConnect("value-changed", func(w gtk.Widget) {
		println(w.SpinButtonGetValue())
	})

	// Control: Transpose
	gui.transpose = gtk.GetWidget("transpose")
	gui.transpose.SpinButtonSetRange(-88, 88)
	gui.transpose.SpinButtonSetIncrements(1, 5)
	gui.transpose.SignalConnect("value-changed", func(w gtk.Widget) {
		println(w.SpinButtonGetValue())
	})

	// Control: Max pitch bend
	gui.pitchBendMax = gtk.GetWidget("pitchBendMax")
	gui.pitchBendMax.SpinButtonSetRange(0, 88)
	gui.pitchBendMax.SpinButtonSetIncrements(1, 5)
	gui.pitchBendMax.SignalConnect("value-changed", func(w gtk.Widget) {
		println(w.SpinButtonGetValue())
	})

	// Control: Mix layers
	gui.mixLayers = gtk.GetWidget("mixLayers")
	gui.mixLayers.SignalConnect("toggled", func(w gtk.Widget) {
		println(w.ToggleButtonGetActive())
	})

	// Control: Layer gamma
	gui.gammaLayer = gtk.GetWidget("gammaLayer")
	gui.gammaLayer.SpinButtonSetRange(0.01, 9.99)
	gui.gammaLayer.SpinButtonSetIncrements(0.01, 0.1)
	gui.gammaLayer.SignalConnect("value-changed", func(w gtk.Widget) {
		println(w.SpinButtonGetValue())
	})

	// Control: Pan low
	gui.panLow = gtk.GetWidget("panLow")
	gui.panLow.SpinButtonSetRange(-1.0, 1.0)
	gui.panLow.SpinButtonSetIncrements(0.01, 0.1)
	gui.panLow.SignalConnect("value-changed", func(w gtk.Widget) {
		println(w.SpinButtonGetValue())
	})

	// Control: Pan high
	gui.panHigh = gtk.GetWidget("panHigh")
	gui.panHigh.SpinButtonSetRange(-1.0, 1.0)
	gui.panHigh.SpinButtonSetIncrements(0.01, 0.1)
	gui.panHigh.SignalConnect("value-changed", func(w gtk.Widget) {
		println(w.SpinButtonGetValue())
	})

	// Control: Rms time
	gui.rmsTime = gtk.GetWidget("rmsTime")
	gui.rmsTime.SpinButtonSetRange(0, 9999)
	gui.rmsTime.SpinButtonSetIncrements(1, 10)
	gui.rmsTime.SignalConnect("value-changed", func(w gtk.Widget) {
		println(w.SpinButtonGetValue())
	})

	// Control: Rms low
	gui.rmsLow = gtk.GetWidget("rmsLow")
	gui.rmsLow.SpinButtonSetRange(0, 0.999)
	gui.rmsLow.SpinButtonSetIncrements(0.001, 0.010)
	gui.rmsLow.SignalConnect("value-changed", func(w gtk.Widget) {
		println(w.SpinButtonGetValue())
	})

	// Control: Rms high
	gui.rmsHigh = gtk.GetWidget("rmsHigh")
	gui.rmsHigh.SpinButtonSetRange(0, 0.999)
	gui.rmsHigh.SpinButtonSetIncrements(0.001, 0.010)
	gui.rmsHigh.SignalConnect("value-changed", func(w gtk.Widget) {
		//println(w.SpinButtonGetValue())
		gui.rmsLow.SpinButtonSetValue(w.SpinButtonGetValue())
	})
}

/*****************************************************************************
 * SIGNAL HANDLERS
 *****************************************************************************/

func (gui *Gui) onLoadSamplerBtn(w gtk.Widget) {
	if w.ToggleButtonGetActive() {
		gui.loadSpinner.SpinnerStart()
	} else {
		gui.loadSpinner.SpinnerStop()
	}

}

func (gui *Gui) onLoadSettingsBtn(w gtk.Widget) {
	println("Load")
}

func (gui *Gui) onSaveSettingsBtn(w gtk.Widget) {
	println("Save")
}
