package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

// ComputerAutomation handles mouse, keyboard, and window automation
type ComputerAutomation struct {
	app              *App
	isRecording      bool
	recordedActions  []Action
	maxRecordingSize int
}

// Action represents an automated action
type Action struct {
	Type        string                 `json:"type"`
	Timestamp   time.Time              `json:"timestamp"`
	Parameters  map[string]interface{} `json:"parameters"`
	Description string                 `json:"description"`
}

// MouseAction represents mouse-related actions
type MouseAction struct {
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Button string `json:"button"` // "left", "right", "middle"
}

// KeyboardAction represents keyboard-related actions
type KeyboardAction struct {
	Key   string `json:"key"`
	Text  string `json:"text,omitempty"`
	Shift bool   `json:"shift,omitempty"`
	Ctrl  bool   `json:"ctrl,omitempty"`
	Alt   bool   `json:"alt,omitempty"`
	Meta  bool   `json:"meta,omitempty"`
}

// WindowInfo represents window information
type WindowInfo struct {
	PID         int32  `json:"pid"`
	Title       string `json:"title"`
	X           int    `json:"x"`
	Y           int    `json:"y"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	IsActive    bool   `json:"isActive"`
	ProcessName string `json:"processName"`
}

// NewComputerAutomation creates a new ComputerAutomation instance
func NewComputerAutomation(app *App) *ComputerAutomation {
	return &ComputerAutomation{
		app:              app,
		maxRecordingSize: 1000,
		recordedActions:  make([]Action, 0),
	}
}

// Mouse Operations

// ClickAt performs a mouse click at the specified coordinates
func (ca *ComputerAutomation) ClickAt(x, y int, button string) error {
	wailsRuntime.LogDebugf(ca.app.ctx, "Clicking at (%d, %d) with %s button", x, y, button)

	// Move mouse to position
	robotgo.Move(x, y)
	time.Sleep(50 * time.Millisecond) // Small delay for natural movement

	// Perform click based on button type
	switch strings.ToLower(button) {
	case "left", "":
		robotgo.Click("left", false)
	case "right":
		robotgo.Click("right", false)
	case "middle":
		robotgo.Click("center", false)
	case "double":
		robotgo.Click("left", true)
	default:
		return fmt.Errorf("unsupported mouse button: %s", button)
	}

	// Record action if recording is enabled
	if ca.isRecording {
		ca.recordAction("mouse_click", map[string]interface{}{
			"x": x, "y": y, "button": button,
		}, fmt.Sprintf("Click %s at (%d, %d)", button, x, y))
	}

	wailsRuntime.LogInfof(ca.app.ctx, "Mouse click completed at (%d, %d)", x, y)
	return nil
}

// DragTo performs a mouse drag operation
func (ca *ComputerAutomation) DragTo(fromX, fromY, toX, toY int) error {
	wailsRuntime.LogDebugf(ca.app.ctx, "Dragging from (%d, %d) to (%d, %d)", fromX, fromY, toX, toY)

	robotgo.Move(fromX, fromY)
	time.Sleep(50 * time.Millisecond)

	robotgo.MouseToggle("down", "left")
	time.Sleep(100 * time.Millisecond)

	robotgo.Move(toX, toY)
	time.Sleep(100 * time.Millisecond)

	robotgo.MouseToggle("up", "left")

	// Record action if recording is enabled
	if ca.isRecording {
		ca.recordAction("mouse_drag", map[string]interface{}{
			"fromX": fromX, "fromY": fromY, "toX": toX, "toY": toY,
		}, fmt.Sprintf("Drag from (%d, %d) to (%d, %d)", fromX, fromY, toX, toY))
	}

	wailsRuntime.LogInfof(ca.app.ctx, "Mouse drag completed")
	return nil
}

// Scroll performs mouse scroll operation
func (ca *ComputerAutomation) Scroll(x, y int, direction string, amount int) error {
	wailsRuntime.LogDebugf(ca.app.ctx, "Scrolling at (%d, %d) %s by %d", x, y, direction, amount)

	robotgo.Move(x, y)
	time.Sleep(50 * time.Millisecond)

	switch strings.ToLower(direction) {
	case "up":
		robotgo.Scroll(0, amount)
	case "down":
		robotgo.Scroll(0, -amount)
	case "left":
		robotgo.Scroll(-amount, 0)
	case "right":
		robotgo.Scroll(amount, 0)
	default:
		return fmt.Errorf("unsupported scroll direction: %s", direction)
	}

	// Record action if recording is enabled
	if ca.isRecording {
		ca.recordAction("mouse_scroll", map[string]interface{}{
			"x": x, "y": y, "direction": direction, "amount": amount,
		}, fmt.Sprintf("Scroll %s at (%d, %d)", direction, x, y))
	}

	return nil
}

// GetMousePosition returns the current mouse position
func (ca *ComputerAutomation) GetMousePosition() (int, int) {
	x, y := robotgo.GetMousePos()
	return x, y
}

// Keyboard Operations

// TypeText types the specified text
func (ca *ComputerAutomation) TypeText(text string, delay int) error {
	wailsRuntime.LogDebugf(ca.app.ctx, "Typing text: %s", text)

	if delay <= 0 {
		delay = 50 // Default delay in milliseconds
	}

	for _, char := range text {
		robotgo.TypeStr(string(char))
		if delay > 0 {
			time.Sleep(time.Duration(delay) * time.Millisecond)
		}
	}

	// Record action if recording is enabled
	if ca.isRecording {
		ca.recordAction("keyboard_type", map[string]interface{}{
			"text": text, "delay": delay,
		}, fmt.Sprintf("Type: %s", text))
	}

	wailsRuntime.LogInfof(ca.app.ctx, "Text typing completed")
	return nil
}

// PressKey presses a specific key or key combination
func (ca *ComputerAutomation) PressKey(key string, modifiers []string) error {
	wailsRuntime.LogDebugf(ca.app.ctx, "Pressing key: %s with modifiers: %v", key, modifiers)

	// Build key combination
	var keyCombo []string
	for _, modifier := range modifiers {
		switch strings.ToLower(modifier) {
		case "ctrl", "control":
			keyCombo = append(keyCombo, "ctrl")
		case "shift":
			keyCombo = append(keyCombo, "shift")
		case "alt":
			keyCombo = append(keyCombo, "alt")
		case "meta", "cmd", "super":
			keyCombo = append(keyCombo, "cmd")
		}
	}
	keyCombo = append(keyCombo, key)

	// Press key combination
	if len(keyCombo) == 1 {
		robotgo.KeyTap(key)
	} else {
		robotgo.KeyTap(key, keyCombo[:len(keyCombo)-1])
	}

	// Record action if recording is enabled
	if ca.isRecording {
		ca.recordAction("keyboard_key", map[string]interface{}{
			"key": key, "modifiers": modifiers,
		}, fmt.Sprintf("Press: %s+%s", strings.Join(modifiers, "+"), key))
	}

	return nil
}

// Window Operations

// GetActiveWindow returns information about the currently active window
func (ca *ComputerAutomation) GetActiveWindow() (*WindowInfo, error) {
	pid := robotgo.GetPID()
	title := robotgo.GetTitle()

	// Get window bounds (this is a simplified implementation)
	// In a real implementation, you'd use platform-specific APIs
	windowInfo := &WindowInfo{
		PID:         pid,
		Title:       title,
		IsActive:    true,
		ProcessName: "unknown", // Would be populated by platform-specific code
	}

	wailsRuntime.LogDebugf(ca.app.ctx, "Active window: %s (PID: %d)", title, pid)
	return windowInfo, nil
}

// GetAllWindows returns information about all open windows
func (ca *ComputerAutomation) GetAllWindows() ([]WindowInfo, error) {
	// This is a placeholder implementation
	// In a real implementation, you'd enumerate all windows using platform-specific APIs
	activeWindow, err := ca.GetActiveWindow()
	if err != nil {
		return nil, err
	}

	windows := []WindowInfo{*activeWindow}
	return windows, nil
}

// FocusWindow brings a window to focus by its title or PID
func (ca *ComputerAutomation) FocusWindow(identifier string, identifierType string) error {
	wailsRuntime.LogDebugf(ca.app.ctx, "Focusing window: %s (%s)", identifier, identifierType)

	switch strings.ToLower(identifierType) {
	case "title":
		return robotgo.ActiveName(identifier)
	case "pid":
		// Implementation would depend on platform-specific APIs
		return fmt.Errorf("focusing by PID not implemented yet")
	default:
		return fmt.Errorf("unsupported identifier type: %s", identifierType)
	}
}

// Application Operations

// LaunchApplication launches an application by name or path
func (ca *ComputerAutomation) LaunchApplication(appPath string, args []string) error {
	wailsRuntime.LogDebugf(ca.app.ctx, "Launching application: %s with args: %v", appPath, args)

	var cmd *exec.Cmd
	if len(args) > 0 {
		cmd = exec.Command(appPath, args...)
	} else {
		cmd = exec.Command(appPath)
	}

	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to launch application: %w", err)
	}

	// Record action if recording is enabled
	if ca.isRecording {
		ca.recordAction("app_launch", map[string]interface{}{
			"path": appPath, "args": args,
		}, fmt.Sprintf("Launch: %s", appPath))
	}

	wailsRuntime.LogInfof(ca.app.ctx, "Application launched successfully: %s", appPath)
	return nil
}

// CloseApplication closes an application by name or PID
func (ca *ComputerAutomation) CloseApplication(identifier string, identifierType string) error {
	wailsRuntime.LogDebugf(ca.app.ctx, "Closing application: %s (%s)", identifier, identifierType)

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		if identifierType == "name" {
			cmd = exec.Command("taskkill", "/F", "/IM", identifier)
		} else {
			cmd = exec.Command("taskkill", "/F", "/PID", identifier)
		}
	case "darwin":
		if identifierType == "name" {
			cmd = exec.Command("pkill", "-f", identifier)
		} else {
			cmd = exec.Command("kill", "-9", identifier)
		}
	case "linux":
		if identifierType == "name" {
			cmd = exec.Command("pkill", "-f", identifier)
		} else {
			cmd = exec.Command("kill", "-9", identifier)
		}
	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to close application: %w", err)
	}

	return nil
}

// Recording Operations

// StartRecording starts recording user actions
func (ca *ComputerAutomation) StartRecording() error {
	if ca.isRecording {
		return fmt.Errorf("recording already in progress")
	}

	ca.isRecording = true
	ca.recordedActions = make([]Action, 0)

	wailsRuntime.LogInfo(ca.app.ctx, "Started recording user actions")
	return nil
}

// StopRecording stops recording user actions
func (ca *ComputerAutomation) StopRecording() []Action {
	ca.isRecording = false
	actions := ca.recordedActions

	wailsRuntime.LogInfof(ca.app.ctx, "Stopped recording. Captured %d actions", len(actions))
	return actions
}

// PlaybackActions replays recorded actions
func (ca *ComputerAutomation) PlaybackActions(actions []Action, speed float64) error {
	if speed <= 0 {
		speed = 1.0
	}

	wailsRuntime.LogInfof(ca.app.ctx, "Playing back %d actions at %.2fx speed", len(actions), speed)

	for i, action := range actions {
		// Add delay between actions based on original timing and speed
		if i > 0 {
			originalDelay := action.Timestamp.Sub(actions[i-1].Timestamp)
			adjustedDelay := time.Duration(float64(originalDelay) / speed)
			time.Sleep(adjustedDelay)
		}

		err := ca.executeAction(action)
		if err != nil {
			wailsRuntime.LogErrorf(ca.app.ctx, "Error executing action %d: %v", i, err)
			return fmt.Errorf("playback failed at action %d: %w", i, err)
		}

		// Emit progress event
		wailsRuntime.EventsEmit(ca.app.ctx, "actionPlaybackProgress", map[string]interface{}{
			"current": i + 1,
			"total":   len(actions),
			"action":  action,
		})
	}

	wailsRuntime.LogInfo(ca.app.ctx, "Action playback completed")
	return nil
}

// Helper methods

func (ca *ComputerAutomation) recordAction(actionType string, parameters map[string]interface{}, description string) {
	action := Action{
		Type:        actionType,
		Timestamp:   time.Now(),
		Parameters:  parameters,
		Description: description,
	}

	ca.recordedActions = append(ca.recordedActions, action)

	// Keep recording size under limit
	if len(ca.recordedActions) > ca.maxRecordingSize {
		ca.recordedActions = ca.recordedActions[1:]
	}
}

func (ca *ComputerAutomation) executeAction(action Action) error {
	params := action.Parameters

	switch action.Type {
	case "mouse_click":
		x := int(params["x"].(float64))
		y := int(params["y"].(float64))
		button := params["button"].(string)
		return ca.ClickAt(x, y, button)

	case "mouse_drag":
		fromX := int(params["fromX"].(float64))
		fromY := int(params["fromY"].(float64))
		toX := int(params["toX"].(float64))
		toY := int(params["toY"].(float64))
		return ca.DragTo(fromX, fromY, toX, toY)

	case "mouse_scroll":
		x := int(params["x"].(float64))
		y := int(params["y"].(float64))
		direction := params["direction"].(string)
		amount := int(params["amount"].(float64))
		return ca.Scroll(x, y, direction, amount)

	case "keyboard_type":
		text := params["text"].(string)
		delay := int(params["delay"].(float64))
		return ca.TypeText(text, delay)

	case "keyboard_key":
		key := params["key"].(string)
		modifiers := make([]string, 0)
		if mods, ok := params["modifiers"].([]interface{}); ok {
			for _, mod := range mods {
				modifiers = append(modifiers, mod.(string))
			}
		}
		return ca.PressKey(key, modifiers)

	default:
		return fmt.Errorf("unknown action type: %s", action.Type)
	}
}

// Wails-bound methods for frontend

// PerformMouseClick performs a mouse click (Wails method)
func (a *App) PerformMouseClick(x, y int, button string) error {
	if a.computerAutomation == nil {
		return fmt.Errorf("computer automation not initialized")
	}
	return a.computerAutomation.ClickAt(x, y, button)
}

// PerformMouseDrag performs a mouse drag (Wails method)
func (a *App) PerformMouseDrag(fromX, fromY, toX, toY int) error {
	if a.computerAutomation == nil {
		return fmt.Errorf("computer automation not initialized")
	}
	return a.computerAutomation.DragTo(fromX, fromY, toX, toY)
}

// PerformKeyPress performs a key press (Wails method)
func (a *App) PerformKeyPress(key string, modifiers []string) error {
	if a.computerAutomation == nil {
		return fmt.Errorf("computer automation not initialized")
	}
	return a.computerAutomation.PressKey(key, modifiers)
}

// TypeTextOnScreen types text on screen (Wails method)
func (a *App) TypeTextOnScreen(text string, delay int) error {
	if a.computerAutomation == nil {
		return fmt.Errorf("computer automation not initialized")
	}
	return a.computerAutomation.TypeText(text, delay)
}

// LaunchApp launches an application (Wails method)
func (a *App) LaunchApp(appPath string, args []string) error {
	if a.computerAutomation == nil {
		return fmt.Errorf("computer automation not initialized")
	}
	return a.computerAutomation.LaunchApplication(appPath, args)
}

// GetActiveWindowInfo returns active window information (Wails method)
func (a *App) GetActiveWindowInfo() (*WindowInfo, error) {
	if a.computerAutomation == nil {
		return nil, fmt.Errorf("computer automation not initialized")
	}
	return a.computerAutomation.GetActiveWindow()
}

// StartActionRecording starts recording actions (Wails method)
func (a *App) StartActionRecording() error {
	if a.computerAutomation == nil {
		return fmt.Errorf("computer automation not initialized")
	}
	return a.computerAutomation.StartRecording()
}

// StopActionRecording stops recording actions (Wails method)
func (a *App) StopActionRecording() []Action {
	if a.computerAutomation == nil {
		return nil
	}
	return a.computerAutomation.StopRecording()
}

// PlaybackRecordedActions replays recorded actions (Wails method)
func (a *App) PlaybackRecordedActions(actions []Action, speed float64) error {
	if a.computerAutomation == nil {
		return fmt.Errorf("computer automation not initialized")
	}
	return a.computerAutomation.PlaybackActions(actions, speed)
}
