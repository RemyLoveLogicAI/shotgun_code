package main

import (
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"time"

	"github.com/kbinani/screenshot"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// ComputerVision handles screen capture and vision-related operations
type ComputerVision struct {
	app               *App
	isCapturing       bool
	captureInterval   time.Duration
	lastScreenshot    *image.RGBA
	screenshotHistory []ScreenshotData
	maxHistorySize    int
}

// ScreenshotData represents a captured screenshot with metadata
type ScreenshotData struct {
	Timestamp   time.Time `json:"timestamp"`
	ImageBase64 string    `json:"imageBase64"`
	Width       int       `json:"width"`
	Height      int       `json:"height"`
	DisplayID   int       `json:"displayId"`
	FilePath    string    `json:"filePath,omitempty"`
}

// OCRResult represents optical character recognition results
type OCRResult struct {
	Text        string      `json:"text"`
	Confidence  float64     `json:"confidence"`
	BoundingBox BoundingBox `json:"boundingBox"`
	Words       []OCRWord   `json:"words"`
}

// OCRWord represents a single word detected by OCR
type OCRWord struct {
	Text        string      `json:"text"`
	Confidence  float64     `json:"confidence"`
	BoundingBox BoundingBox `json:"boundingBox"`
}

// BoundingBox represents a rectangular area on screen
type BoundingBox struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

// UIElement represents a detected UI element
type UIElement struct {
	Type        string            `json:"type"`
	Text        string            `json:"text,omitempty"`
	BoundingBox BoundingBox       `json:"boundingBox"`
	Attributes  map[string]string `json:"attributes,omitempty"`
	Clickable   bool              `json:"clickable"`
	Visible     bool              `json:"visible"`
}

// NewComputerVision creates a new ComputerVision instance
func NewComputerVision(app *App) *ComputerVision {
	return &ComputerVision{
		app:               app,
		captureInterval:   500 * time.Millisecond,
		maxHistorySize:    100,
		screenshotHistory: make([]ScreenshotData, 0),
	}
}

// CaptureScreen captures the current screen and returns base64 encoded image
func (cv *ComputerVision) CaptureScreen(displayID int) (*ScreenshotData, error) {
	runtime.LogDebugf(cv.app.ctx, "Capturing screen display %d", displayID)

	// Get screen bounds
	bounds := screenshot.GetDisplayBounds(displayID)
	if bounds.Dx() == 0 || bounds.Dy() == 0 {
		return nil, fmt.Errorf("invalid display bounds for display %d", displayID)
	}

	// Capture screenshot
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		return nil, fmt.Errorf("failed to capture screen: %w", err)
	}

	cv.lastScreenshot = img

	// Convert to base64
	imageBase64, err := cv.imageToBase64(img, "png")
	if err != nil {
		return nil, fmt.Errorf("failed to encode image: %w", err)
	}

	screenshotData := &ScreenshotData{
		Timestamp:   time.Now(),
		ImageBase64: imageBase64,
		Width:       img.Bounds().Dx(),
		Height:      img.Bounds().Dy(),
		DisplayID:   displayID,
	}

	// Add to history
	cv.addToHistory(*screenshotData)

	runtime.LogInfof(cv.app.ctx, "Screen captured successfully: %dx%d", screenshotData.Width, screenshotData.Height)
	return screenshotData, nil
}

// CaptureAllScreens captures all available displays
func (cv *ComputerVision) CaptureAllScreens() ([]ScreenshotData, error) {
	numDisplays := screenshot.NumActiveDisplays()
	if numDisplays == 0 {
		return nil, fmt.Errorf("no active displays found")
	}

	screenshots := make([]ScreenshotData, 0, numDisplays)

	for i := 0; i < numDisplays; i++ {
		screenshot, err := cv.CaptureScreen(i)
		if err != nil {
			runtime.LogWarningf(cv.app.ctx, "Failed to capture display %d: %v", i, err)
			continue
		}
		screenshots = append(screenshots, *screenshot)
	}

	return screenshots, nil
}

// StartContinuousCapture starts capturing screenshots at regular intervals
func (cv *ComputerVision) StartContinuousCapture(intervalMs int, displayID int) error {
	if cv.isCapturing {
		return fmt.Errorf("capture already in progress")
	}

	cv.isCapturing = true
	cv.captureInterval = time.Duration(intervalMs) * time.Millisecond

	go func() {
		ticker := time.NewTicker(cv.captureInterval)
		defer ticker.Stop()

		for cv.isCapturing {
			select {
			case <-ticker.C:
				screenshot, err := cv.CaptureScreen(displayID)
				if err != nil {
					runtime.LogErrorf(cv.app.ctx, "Continuous capture error: %v", err)
					continue
				}

				// Emit screenshot event to frontend
				runtime.EventsEmit(cv.app.ctx, "screenshotCaptured", screenshot)

			case <-cv.app.ctx.Done():
				cv.isCapturing = false
				return
			}
		}
	}()

	runtime.LogInfo(cv.app.ctx, "Started continuous screen capture")
	return nil
}

// StopContinuousCapture stops the continuous screen capture
func (cv *ComputerVision) StopContinuousCapture() {
	cv.isCapturing = false
	runtime.LogInfo(cv.app.ctx, "Stopped continuous screen capture")
}

// SaveScreenshot saves a screenshot to disk
func (cv *ComputerVision) SaveScreenshot(screenshot *ScreenshotData, filename string) error {
	if filename == "" {
		timestamp := screenshot.Timestamp.Format("20060102_150405")
		filename = fmt.Sprintf("screenshot_%s.png", timestamp)
	}

	// Create screenshots directory if it doesn't exist
	screenshotsDir := "data/screenshots"
	if err := os.MkdirAll(screenshotsDir, 0755); err != nil {
		return fmt.Errorf("failed to create screenshots directory: %w", err)
	}

	filepath := filepath.Join(screenshotsDir, filename)

	// Decode base64 image
	imgData, err := base64.StdEncoding.DecodeString(screenshot.ImageBase64)
	if err != nil {
		return fmt.Errorf("failed to decode base64 image: %w", err)
	}

	// Save to file
	if err := os.WriteFile(filepath, imgData, 0644); err != nil {
		return fmt.Errorf("failed to save screenshot: %w", err)
	}

	screenshot.FilePath = filepath
	runtime.LogInfof(cv.app.ctx, "Screenshot saved to: %s", filepath)
	return nil
}

// GetScreenshotHistory returns the screenshot history
func (cv *ComputerVision) GetScreenshotHistory() []ScreenshotData {
	return cv.screenshotHistory
}

// ClearScreenshotHistory clears the screenshot history
func (cv *ComputerVision) ClearScreenshotHistory() {
	cv.screenshotHistory = make([]ScreenshotData, 0)
	runtime.LogInfo(cv.app.ctx, "Screenshot history cleared")
}

// DetectUIElements analyzes the current screen for UI elements
func (cv *ComputerVision) DetectUIElements() ([]UIElement, error) {
	if cv.lastScreenshot == nil {
		// Capture current screen first
		_, err := cv.CaptureScreen(0)
		if err != nil {
			return nil, fmt.Errorf("failed to capture screen for UI detection: %w", err)
		}
	}

	// Placeholder for actual UI detection logic
	// In a real implementation, this would use computer vision algorithms
	// to detect buttons, text fields, menus, etc.
	elements := []UIElement{
		{
			Type:        "button",
			Text:        "Example Button",
			BoundingBox: BoundingBox{X: 100, Y: 100, Width: 120, Height: 30},
			Clickable:   true,
			Visible:     true,
		},
		{
			Type:        "text_input",
			BoundingBox: BoundingBox{X: 50, Y: 200, Width: 200, Height: 25},
			Clickable:   true,
			Visible:     true,
		},
	}

	runtime.LogInfof(cv.app.ctx, "Detected %d UI elements", len(elements))
	return elements, nil
}

// PerformOCR performs optical character recognition on the current screen
func (cv *ComputerVision) PerformOCR(region *BoundingBox) (*OCRResult, error) {
	if cv.lastScreenshot == nil {
		// Capture current screen first
		_, err := cv.CaptureScreen(0)
		if err != nil {
			return nil, fmt.Errorf("failed to capture screen for OCR: %w", err)
		}
	}

	// Extract region if specified (for future OCR implementation)
	// var targetImage image.Image = cv.lastScreenshot
	// if region != nil {
	// 	bounds := image.Rect(region.X, region.Y, region.X+region.Width, region.Y+region.Height)
	// 	targetImage = cv.lastScreenshot.SubImage(bounds)
	// }

	// Placeholder for actual OCR implementation
	// In a real implementation, this would use libraries like tesseract-go
	result := &OCRResult{
		Text:       "Example OCR text detected",
		Confidence: 0.95,
		BoundingBox: BoundingBox{
			X:      region.X,
			Y:      region.Y,
			Width:  region.Width,
			Height: region.Height,
		},
		Words: []OCRWord{
			{
				Text:        "Example",
				Confidence:  0.98,
				BoundingBox: BoundingBox{X: region.X, Y: region.Y, Width: 50, Height: 20},
			},
			{
				Text:        "OCR",
				Confidence:  0.92,
				BoundingBox: BoundingBox{X: region.X + 55, Y: region.Y, Width: 30, Height: 20},
			},
		},
	}

	runtime.LogInfof(cv.app.ctx, "OCR completed with confidence: %.2f", result.Confidence)
	return result, nil
}

// GetDisplayInfo returns information about available displays
func (cv *ComputerVision) GetDisplayInfo() ([]map[string]interface{}, error) {
	numDisplays := screenshot.NumActiveDisplays()
	displays := make([]map[string]interface{}, numDisplays)

	for i := 0; i < numDisplays; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		displays[i] = map[string]interface{}{
			"id":      i,
			"x":       bounds.Min.X,
			"y":       bounds.Min.Y,
			"width":   bounds.Dx(),
			"height":  bounds.Dy(),
			"primary": i == 0, // Assume first display is primary
		}
	}

	return displays, nil
}

// Helper methods

func (cv *ComputerVision) imageToBase64(img image.Image, format string) (string, error) {
	// Create a temporary file to encode the image
	tmpFile, err := os.CreateTemp("", "screenshot."+format)
	if err != nil {
		return "", err
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	// Encode image to file
	switch format {
	case "png":
		err = png.Encode(tmpFile, img)
	case "jpeg", "jpg":
		err = jpeg.Encode(tmpFile, img, &jpeg.Options{Quality: 90})
	default:
		return "", fmt.Errorf("unsupported format: %s", format)
	}

	if err != nil {
		return "", err
	}

	// Read file content
	tmpFile.Seek(0, 0)
	content, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(content), nil
}

func (cv *ComputerVision) addToHistory(screenshot ScreenshotData) {
	cv.screenshotHistory = append(cv.screenshotHistory, screenshot)

	// Keep history size under limit
	if len(cv.screenshotHistory) > cv.maxHistorySize {
		cv.screenshotHistory = cv.screenshotHistory[1:]
	}
}

// Wails-bound methods for frontend

// CaptureScreenForDisplay captures screen for specified display (Wails method)
func (a *App) CaptureScreenForDisplay(displayID int) (*ScreenshotData, error) {
	if a.computerVision == nil {
		return nil, fmt.Errorf("computer vision not initialized")
	}
	return a.computerVision.CaptureScreen(displayID)
}

// StartScreenCapture starts continuous screen capture (Wails method)
func (a *App) StartScreenCapture(intervalMs int, displayID int) error {
	if a.computerVision == nil {
		return fmt.Errorf("computer vision not initialized")
	}
	return a.computerVision.StartContinuousCapture(intervalMs, displayID)
}

// StopScreenCapture stops continuous screen capture (Wails method)
func (a *App) StopScreenCapture() {
	if a.computerVision != nil {
		a.computerVision.StopContinuousCapture()
	}
}

// GetDisplayInfo returns display information (Wails method)
func (a *App) GetDisplayInfo() ([]map[string]interface{}, error) {
	if a.computerVision == nil {
		return nil, fmt.Errorf("computer vision not initialized")
	}
	return a.computerVision.GetDisplayInfo()
}

// DetectUIElements detects UI elements on screen (Wails method)
func (a *App) DetectUIElements() ([]UIElement, error) {
	if a.computerVision == nil {
		return nil, fmt.Errorf("computer vision not initialized")
	}
	return a.computerVision.DetectUIElements()
}

// PerformScreenOCR performs OCR on screen region (Wails method)
func (a *App) PerformScreenOCR(region *BoundingBox) (*OCRResult, error) {
	if a.computerVision == nil {
		return nil, fmt.Errorf("computer vision not initialized")
	}
	return a.computerVision.PerformOCR(region)
}
