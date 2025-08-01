# AI Agent Co-pilot Documentation

## Overview

The **AI Agent Co-pilot** is a comprehensive computer use and computer vision system that transforms the Shotgun App into an intelligent assistant capable of understanding and interacting with your computer screen. It combines advanced computer vision, automation capabilities, and AI-powered task execution to help users complete everyday computing tasks through natural language instructions.

## Key Features

### 🤖 **Intelligent Task Execution**
- Natural language task understanding
- AI-powered action planning and execution
- Context-aware decision making
- Multi-step task automation

### 👁️ **Computer Vision Capabilities**
- Real-time screen capture and analysis
- UI element detection and recognition
- Optical Character Recognition (OCR)
- Visual context understanding
- Screen content analysis

### 🖱️ **Computer Automation**
- Mouse control (clicking, dragging, scrolling)
- Keyboard input and shortcuts
- Window management
- Application launching and control
- Action recording and playback

### 🧠 **Learning & Memory**
- Context history tracking
- Pattern recognition
- User preference learning
- Task execution optimization

### 🛡️ **Safety & Security**
- Safety checks before actions
- User confirmation for sensitive operations
- Action logging and monitoring
- Configurable permission levels

## Architecture

### Backend Components

#### 1. **ComputerVision** (`computer_vision.go`)
Handles all screen capture and visual analysis:

```go
type ComputerVision struct {
    app               *App
    isCapturing       bool
    captureInterval   time.Duration
    lastScreenshot    *image.RGBA
    screenshotHistory []ScreenshotData
    maxHistorySize    int
}
```

**Key Methods:**
- `CaptureScreen(displayID int)` - Capture specific display
- `CaptureAllScreens()` - Capture all displays
- `StartContinuousCapture()` - Begin live monitoring
- `DetectUIElements()` - Find UI components
- `PerformOCR(region *BoundingBox)` - Extract text from regions

#### 2. **ComputerAutomation** (`computer_automation.go`)
Manages computer control and automation:

```go
type ComputerAutomation struct {
    app              *App
    isRecording      bool
    recordedActions  []Action
    maxRecordingSize int
}
```

**Key Methods:**
- `ClickAt(x, y int, button string)` - Mouse clicking
- `TypeText(text string, delay int)` - Keyboard input
- `PressKey(key string, modifiers []string)` - Key combinations
- `LaunchApplication(appPath string)` - App control
- `StartRecording()` / `PlaybackActions()` - Action recording

#### 3. **AIAgent** (`ai_agent.go`)
Coordinates all capabilities with AI planning:

```go
type AIAgent struct {
    app             *App
    config          *AgentConfig
    taskManager     *TaskManager
    memoryManager   *MemoryManager
    visionAnalyzer  *VisionAnalyzer
    isActive        bool
    currentTask     *Task
    contextHistory  []ContextSnapshot
    learningData    []LearningEvent
}
```

**Key Methods:**
- `ExecuteTask(description string)` - Process natural language tasks
- `planAndExecute(task *Task)` - AI-powered task planning
- `captureContext()` - Gather system state
- `monitorSystem()` - Continuous monitoring

### Frontend Components

#### **AIAgentPanel.vue**
The main user interface for the AI Agent system:

**Features:**
- Task input and execution
- Real-time screen preview
- Computer vision controls
- Activity logging
- Settings management
- Quick action buttons

## Getting Started

### 1. **Installation**

Add the required dependencies to your Go project:

```bash
go get github.com/kbinani/screenshot
go get github.com/go-vgo/robotgo
```

### 2. **Configuration**

The AI Agent can be configured through the settings panel or programmatically:

```go
config := &AgentConfig{
    ModelProvider:     "openai",
    ModelName:        "gpt-4-vision-preview",
    APIKey:          "your-api-key",
    VisionEnabled:    true,
    AutomationEnabled: true,
    SafetyChecks:     true,
    CustomInstructions: "Be helpful and safe",
}
```

### 3. **Basic Usage**

#### Starting the Agent
```javascript
// Frontend
await App.StartAgent()
```

#### Executing Tasks
```javascript
// Natural language task execution
const task = await App.ExecuteAgentTask("Take a screenshot and save it", 2)
```

#### Capturing Screen
```javascript
// Single screenshot
const screenshot = await App.CaptureScreenForDisplay(0)

// Start live capture
await App.StartScreenCapture(1000, 0) // 1 second interval
```

## Task Examples

### Simple Tasks
- "Take a screenshot"
- "Open calculator"
- "Type 'Hello World'"
- "Click on the save button"

### Complex Tasks
- "Find all text on the screen and summarize it"
- "Open a browser, navigate to google.com, and search for 'AI agent'"
- "Take a screenshot, analyze it for UI elements, and create a report"
- "Record my actions for the next 30 seconds, then replay them"

### Automation Workflows
- "Every 5 minutes, check if there are new emails and notify me"
- "Monitor this window and alert me when it changes"
- "Create a backup of my desktop by taking periodic screenshots"

## API Reference

### Core Agent Methods

#### `StartAgent() error`
Activates the AI agent and begins monitoring.

#### `StopAgent()`
Deactivates the agent and stops all monitoring.

#### `ExecuteAgentTask(description string, priority int) (*Task, error)`
Executes a natural language task with specified priority (1-3).

#### `GetAgentStatus() map[string]interface{}`
Returns current agent status and statistics.

### Computer Vision Methods

#### `CaptureScreenForDisplay(displayID int) (*ScreenshotData, error)`
Captures a screenshot of the specified display.

#### `DetectUIElements() ([]UIElement, error)`
Analyzes the current screen for UI components.

#### `PerformScreenOCR(region *BoundingBox) (*OCRResult, error)`
Extracts text from a specific screen region.

### Automation Methods

#### `PerformMouseClick(x, y int, button string) error`
Clicks at specified coordinates with given button type.

#### `TypeTextOnScreen(text string, delay int) error`
Types text with configurable delay between characters.

#### `PerformKeyPress(key string, modifiers []string) error`
Presses key combinations (e.g., Ctrl+C).

#### `LaunchApp(appPath string, args []string) error`
Launches applications with optional arguments.

## Data Structures

### Task
```go
type Task struct {
    ID          string                 `json:"id"`
    Type        string                 `json:"type"`
    Description string                 `json:"description"`
    Status      string                 `json:"status"`
    CreatedAt   time.Time              `json:"createdAt"`
    UpdatedAt   time.Time              `json:"updatedAt"`
    Priority    int                    `json:"priority"`
    Steps       []TaskStep             `json:"steps"`
    Context     map[string]interface{} `json:"context"`
    Result      *TaskResult            `json:"result,omitempty"`
    Metadata    map[string]interface{} `json:"metadata"`
}
```

### ScreenshotData
```go
type ScreenshotData struct {
    Timestamp   time.Time `json:"timestamp"`
    ImageBase64 string    `json:"imageBase64"`
    Width       int       `json:"width"`
    Height      int       `json:"height"`
    DisplayID   int       `json:"displayId"`
    FilePath    string    `json:"filePath,omitempty"`
}
```

### UIElement
```go
type UIElement struct {
    Type        string            `json:"type"`
    Text        string            `json:"text,omitempty"`
    BoundingBox BoundingBox       `json:"boundingBox"`
    Attributes  map[string]string `json:"attributes,omitempty"`
    Clickable   bool              `json:"clickable"`
    Visible     bool              `json:"visible"`
}
```

## Security Considerations

### Safety Mechanisms
1. **Action Confirmation** - Prompts for destructive operations
2. **Permission Levels** - Configurable capability restrictions
3. **Logging** - Complete action history tracking
4. **Rate Limiting** - Prevents rapid automation abuse
5. **Scope Limiting** - Restrict to specific applications/areas

### Best Practices
- Always enable safety checks in production
- Review action logs regularly
- Use appropriate permission levels
- Test automation in safe environments
- Keep API keys secure

## Troubleshooting

### Common Issues

#### Agent Won't Start
- Check if required dependencies are installed
- Verify API key configuration
- Ensure no permission conflicts

#### Screen Capture Fails
- Check display permissions (macOS/Linux)
- Verify display ID is valid
- Test with administrator privileges

#### Automation Not Working
- Check accessibility permissions
- Verify target applications are responsive
- Test manual operations first

#### AI Planning Errors
- Verify API key and connectivity
- Check model availability
- Review custom instructions

### Debug Mode
Enable detailed logging by setting:
```go
agent.config.Settings["debug"] = true
```

## Performance Optimization

### Memory Management
- Screenshot history is limited to 100 items
- Task history limited to 1000 items
- Regular cleanup of temporary files

### CPU Usage
- Configurable capture intervals
- Lazy UI element detection
- Efficient image encoding

### Network Usage
- Compressed image uploads to AI models
- Cached model responses
- Batched API requests

## Extension Points

### Custom AI Models
Implement custom model providers by extending the `callAIModel` method:

```go
func (agent *AIAgent) callCustomModel(prompt string, screenshot *ScreenshotData) (map[string]interface{}, error) {
    // Custom model integration
}
```

### Additional Computer Vision
Add new vision capabilities:

```go
func (cv *ComputerVision) CustomVisionAnalysis() (*CustomResult, error) {
    // Custom analysis implementation
}
```

### Automation Extensions
Create custom automation actions:

```go
func (ca *ComputerAutomation) CustomAction(params map[string]interface{}) error {
    // Custom automation logic
}
```

## Future Roadmap

### Planned Features
- [ ] Multi-monitor support enhancement
- [ ] Advanced OCR with layout preservation
- [ ] Natural language file operations
- [ ] Voice command integration
- [ ] Mobile device control
- [ ] Cross-platform consistency improvements
- [ ] Machine learning-based UI prediction
- [ ] Integration with popular productivity tools

### Integration Possibilities
- Browser automation (Selenium-like)
- Mobile app control
- Cloud service integration
- Development environment automation
- System administration tasks

## License and Disclaimer

This AI Agent system is designed to assist with legitimate computing tasks. Users are responsible for ensuring compliance with:
- Application terms of service
- Local automation policies
- Privacy regulations
- Security requirements

Always test automation in safe environments before production use.

---

**Support**: For issues or questions, please refer to the main project documentation or create an issue in the project repository.