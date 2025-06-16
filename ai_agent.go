package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// AIAgent represents the main AI agent that coordinates computer use and vision
type AIAgent struct {
	app            *App
	config         *AgentConfig
	taskManager    *TaskManager
	memoryManager  *MemoryManager
	visionAnalyzer *VisionAnalyzer
	isActive       bool
	currentTask    *Task
	contextHistory []ContextSnapshot
	learningData   []LearningEvent
	mu             sync.RWMutex
}

// AgentConfig contains configuration for the AI agent
type AgentConfig struct {
	ModelProvider      string                 `json:"modelProvider"` // "openai", "anthropic", "local"
	ModelName          string                 `json:"modelName"`
	APIKey             string                 `json:"apiKey"`
	BaseURL            string                 `json:"baseUrl,omitempty"`
	MaxContextLength   int                    `json:"maxContextLength"`
	Temperature        float64                `json:"temperature"`
	VisionEnabled      bool                   `json:"visionEnabled"`
	AutomationEnabled  bool                   `json:"automationEnabled"`
	LearningEnabled    bool                   `json:"learningEnabled"`
	SafetyChecks       bool                   `json:"safetyChecks"`
	CustomInstructions string                 `json:"customInstructions"`
	Capabilities       map[string]bool        `json:"capabilities"`
	Settings           map[string]interface{} `json:"settings"`
}

// Task represents a task that the AI agent can execute
type Task struct {
	ID          string                 `json:"id"`
	Type        string                 `json:"type"`
	Description string                 `json:"description"`
	Status      string                 `json:"status"` // "pending", "running", "completed", "failed", "paused"
	CreatedAt   time.Time              `json:"createdAt"`
	UpdatedAt   time.Time              `json:"updatedAt"`
	Priority    int                    `json:"priority"`
	Steps       []TaskStep             `json:"steps"`
	Context     map[string]interface{} `json:"context"`
	Result      *TaskResult            `json:"result,omitempty"`
	Metadata    map[string]interface{} `json:"metadata"`
}

// TaskStep represents a single step in a task
type TaskStep struct {
	ID          string                 `json:"id"`
	Type        string                 `json:"type"` // "vision", "action", "analysis", "wait"
	Description string                 `json:"description"`
	Status      string                 `json:"status"`
	Parameters  map[string]interface{} `json:"parameters"`
	Result      interface{}            `json:"result,omitempty"`
	Error       string                 `json:"error,omitempty"`
	Duration    time.Duration          `json:"duration"`
	Timestamp   time.Time              `json:"timestamp"`
}

// TaskResult contains the result of a completed task
type TaskResult struct {
	Success     bool                   `json:"success"`
	Data        interface{}            `json:"data,omitempty"`
	Error       string                 `json:"error,omitempty"`
	Screenshots []ScreenshotData       `json:"screenshots,omitempty"`
	Actions     []Action               `json:"actions,omitempty"`
	Metadata    map[string]interface{} `json:"metadata"`
}

// ContextSnapshot represents the state of the screen and system at a point in time
type ContextSnapshot struct {
	Timestamp    time.Time              `json:"timestamp"`
	Screenshot   *ScreenshotData        `json:"screenshot"`
	UIElements   []UIElement            `json:"uiElements"`
	ActiveWindow *WindowInfo            `json:"activeWindow"`
	MousePos     [2]int                 `json:"mousePos"`
	SystemState  map[string]interface{} `json:"systemState"`
}

// LearningEvent represents something the agent has learned
type LearningEvent struct {
	Timestamp  time.Time              `json:"timestamp"`
	EventType  string                 `json:"eventType"`
	Context    ContextSnapshot        `json:"context"`
	Action     interface{}            `json:"action"`
	Outcome    string                 `json:"outcome"`
	Feedback   string                 `json:"feedback,omitempty"`
	Confidence float64                `json:"confidence"`
	Metadata   map[string]interface{} `json:"metadata"`
}

// TaskManager handles task queue and execution
type TaskManager struct {
	tasks      map[string]*Task
	queue      []*Task
	history    []*Task
	maxHistory int
	mu         sync.RWMutex
}

// MemoryManager handles agent memory and learning
type MemoryManager struct {
	shortTermMemory []ContextSnapshot
	longTermMemory  []LearningEvent
	patterns        map[string]interface{}
	preferences     map[string]interface{}
	maxShortTerm    int
	maxLongTerm     int
	mu              sync.RWMutex
}

// VisionAnalyzer provides advanced computer vision analysis
type VisionAnalyzer struct {
	app *App
}

// NewAIAgent creates a new AI agent instance
func NewAIAgent(app *App) *AIAgent {
	agent := &AIAgent{
		app: app,
		config: &AgentConfig{
			ModelProvider:     "openai",
			ModelName:         "gpt-4-vision-preview",
			MaxContextLength:  128000,
			Temperature:       0.7,
			VisionEnabled:     true,
			AutomationEnabled: true,
			LearningEnabled:   true,
			SafetyChecks:      true,
			Capabilities: map[string]bool{
				"screen_capture":   true,
				"ui_interaction":   true,
				"text_recognition": true,
				"app_control":      true,
				"file_operations":  true,
				"web_automation":   true,
			},
			Settings: make(map[string]interface{}),
		},
		taskManager:    NewTaskManager(),
		memoryManager:  NewMemoryManager(),
		visionAnalyzer: &VisionAnalyzer{app: app},
		contextHistory: make([]ContextSnapshot, 0),
		learningData:   make([]LearningEvent, 0),
	}

	return agent
}

// NewTaskManager creates a new task manager
func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks:      make(map[string]*Task),
		queue:      make([]*Task, 0),
		history:    make([]*Task, 0),
		maxHistory: 1000,
	}
}

// NewMemoryManager creates a new memory manager
func NewMemoryManager() *MemoryManager {
	return &MemoryManager{
		shortTermMemory: make([]ContextSnapshot, 0),
		longTermMemory:  make([]LearningEvent, 0),
		patterns:        make(map[string]interface{}),
		preferences:     make(map[string]interface{}),
		maxShortTerm:    100,
		maxLongTerm:     10000,
	}
}

// Agent Control Methods

// Start activates the AI agent
func (agent *AIAgent) Start() error {
	agent.mu.Lock()
	defer agent.mu.Unlock()

	if agent.isActive {
		return fmt.Errorf("agent is already active")
	}

	agent.isActive = true

	// Start background monitoring
	go agent.monitorSystem()

	runtime.LogInfo(agent.app.ctx, "AI Agent started successfully")
	runtime.EventsEmit(agent.app.ctx, "agentStatusChanged", map[string]interface{}{
		"status":  "active",
		"message": "AI Agent started",
	})

	return nil
}

// Stop deactivates the AI agent
func (agent *AIAgent) Stop() {
	agent.mu.Lock()
	defer agent.mu.Unlock()

	agent.isActive = false

	// Cancel current task if any
	if agent.currentTask != nil {
		agent.currentTask.Status = "paused"
	}

	runtime.LogInfo(agent.app.ctx, "AI Agent stopped")
	runtime.EventsEmit(agent.app.ctx, "agentStatusChanged", map[string]interface{}{
		"status":  "inactive",
		"message": "AI Agent stopped",
	})
}

// ExecuteTask executes a natural language task
func (agent *AIAgent) ExecuteTask(taskDescription string, priority int) (*Task, error) {
	agent.mu.Lock()
	defer agent.mu.Unlock()

	if !agent.isActive {
		return nil, fmt.Errorf("agent is not active")
	}

	// Create new task
	task := &Task{
		ID:          fmt.Sprintf("task_%d", time.Now().Unix()),
		Type:        "user_request",
		Description: taskDescription,
		Status:      "pending",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Priority:    priority,
		Steps:       make([]TaskStep, 0),
		Context:     make(map[string]interface{}),
		Metadata:    make(map[string]interface{}),
	}

	// Add to task manager
	agent.taskManager.AddTask(task)

	// Start task execution in background
	go agent.executeTaskAsync(task)

	runtime.LogInfof(agent.app.ctx, "Task created: %s", task.Description)
	runtime.EventsEmit(agent.app.ctx, "taskCreated", task)

	return task, nil
}

// executeTaskAsync executes a task asynchronously
func (agent *AIAgent) executeTaskAsync(task *Task) {
	agent.currentTask = task
	task.Status = "running"
	task.UpdatedAt = time.Now()

	runtime.EventsEmit(agent.app.ctx, "taskStatusChanged", task)

	// Capture initial context
	context, err := agent.captureContext()
	if err != nil {
		runtime.LogErrorf(agent.app.ctx, "Failed to capture context: %v", err)
	} else {
		agent.addToContext(*context)
	}

	// Plan and execute task
	err = agent.planAndExecute(task)
	if err != nil {
		task.Status = "failed"
		task.Result = &TaskResult{
			Success: false,
			Error:   err.Error(),
		}
		runtime.LogErrorf(agent.app.ctx, "Task failed: %v", err)
	} else {
		task.Status = "completed"
		runtime.LogInfo(agent.app.ctx, "Task completed successfully")
	}

	task.UpdatedAt = time.Now()
	agent.taskManager.MoveToHistory(task)
	agent.currentTask = nil

	runtime.EventsEmit(agent.app.ctx, "taskCompleted", task)
}

// planAndExecute plans and executes a task using AI
func (agent *AIAgent) planAndExecute(task *Task) error {
	// Get current screen state
	screenshot, err := agent.app.computerVision.CaptureScreen(0)
	if err != nil {
		return fmt.Errorf("failed to capture screen: %w", err)
	}

	// Analyze UI elements
	uiElements, err := agent.app.computerVision.DetectUIElements()
	if err != nil {
		runtime.LogWarningf(agent.app.ctx, "Failed to detect UI elements: %v", err)
		uiElements = []UIElement{}
	}

	// Create task planning prompt
	prompt := agent.createTaskPlanningPrompt(task, screenshot, uiElements)

	// Get AI response for task planning
	plan, err := agent.callAIModel(prompt, screenshot)
	if err != nil {
		return fmt.Errorf("failed to get AI plan: %w", err)
	}

	// Parse and execute plan
	return agent.executePlan(task, plan)
}

// createTaskPlanningPrompt creates a prompt for AI task planning
func (agent *AIAgent) createTaskPlanningPrompt(task *Task, screenshot *ScreenshotData, uiElements []UIElement) string {
	prompt := fmt.Sprintf(`You are an AI agent that can control a computer to help users complete tasks. 

Current Task: %s

Available capabilities:
- Take screenshots and analyze the screen
- Click on UI elements (buttons, links, etc.)
- Type text
- Press keyboard shortcuts
- Scroll and navigate
- Launch applications
- Detect and read text (OCR)

Current screen analysis:
- Screenshot dimensions: %dx%d
- Detected UI elements: %d

UI Elements detected:
`, task.Description, screenshot.Width, screenshot.Height, len(uiElements))

	for i, element := range uiElements {
		prompt += fmt.Sprintf("- Element %d: %s at (%d,%d) %dx%d\n",
			i+1, element.Type, element.BoundingBox.X, element.BoundingBox.Y,
			element.BoundingBox.Width, element.BoundingBox.Height)
		if element.Text != "" {
			prompt += fmt.Sprintf("  Text: %s\n", element.Text)
		}
	}

	prompt += `
Please analyze the current screen and provide a step-by-step plan to complete the task. 
Respond with a JSON object containing:
{
  "analysis": "Your analysis of the current screen",
  "plan": [
    {
      "step": 1,
      "action": "action_type",
      "description": "What this step does",
      "parameters": {"param1": "value1"}
    }
  ]
}

Available action types:
- "click": Click at coordinates or on UI element
- "type": Type text
- "key": Press keyboard key
- "scroll": Scroll in direction
- "wait": Wait for condition
- "capture": Take screenshot
- "analyze": Analyze screen content

Be specific with coordinates and parameters.`

	return prompt
}

// callAIModel calls the configured AI model with vision
func (agent *AIAgent) callAIModel(prompt string, screenshot *ScreenshotData) (map[string]interface{}, error) {
	// This is a placeholder implementation
	// In a real implementation, this would call the actual AI model API

	runtime.LogDebugf(agent.app.ctx, "Calling AI model with prompt length: %d", len(prompt))

	// Simulate AI response
	response := map[string]interface{}{
		"analysis": "I can see the current screen and will help complete the task.",
		"plan": []interface{}{
			map[string]interface{}{
				"step":        1,
				"action":      "capture",
				"description": "Take initial screenshot for analysis",
				"parameters":  map[string]interface{}{},
			},
			map[string]interface{}{
				"step":        2,
				"action":      "analyze",
				"description": "Analyze current screen content",
				"parameters":  map[string]interface{}{},
			},
		},
	}

	// In a real implementation, you would:
	// 1. Format the request for your chosen AI provider (OpenAI, Anthropic, etc.)
	// 2. Include the screenshot image in the request
	// 3. Send HTTP request to the AI API
	// 4. Parse and return the response

	return response, nil
}

// executePlan executes the AI-generated plan
func (agent *AIAgent) executePlan(task *Task, plan map[string]interface{}) error {
	planSteps, ok := plan["plan"].([]interface{})
	if !ok {
		return fmt.Errorf("invalid plan format")
	}

	for _, stepInterface := range planSteps {
		stepMap, ok := stepInterface.(map[string]interface{})
		if !ok {
			continue
		}

		step := TaskStep{
			ID:          fmt.Sprintf("%s_step_%d", task.ID, len(task.Steps)+1),
			Type:        stepMap["action"].(string),
			Description: stepMap["description"].(string),
			Status:      "running",
			Parameters:  stepMap["parameters"].(map[string]interface{}),
			Timestamp:   time.Now(),
		}

		task.Steps = append(task.Steps, step)

		// Execute step
		err := agent.executeStep(&step)
		if err != nil {
			step.Status = "failed"
			step.Error = err.Error()
			runtime.LogErrorf(agent.app.ctx, "Step failed: %v", err)
			return err
		}

		step.Status = "completed"
		step.Duration = time.Since(step.Timestamp)

		runtime.EventsEmit(agent.app.ctx, "taskStepCompleted", step)
	}

	return nil
}

// executeStep executes a single task step
func (agent *AIAgent) executeStep(step *TaskStep) error {
	switch step.Type {
	case "capture":
		screenshot, err := agent.app.computerVision.CaptureScreen(0)
		if err != nil {
			return err
		}
		step.Result = screenshot

	case "click":
		x := int(step.Parameters["x"].(float64))
		y := int(step.Parameters["y"].(float64))
		button := "left"
		if b, ok := step.Parameters["button"].(string); ok {
			button = b
		}
		return agent.app.computerAutomation.ClickAt(x, y, button)

	case "type":
		text := step.Parameters["text"].(string)
		delay := 50
		if d, ok := step.Parameters["delay"].(float64); ok {
			delay = int(d)
		}
		return agent.app.computerAutomation.TypeText(text, delay)

	case "key":
		key := step.Parameters["key"].(string)
		modifiers := make([]string, 0)
		if mods, ok := step.Parameters["modifiers"].([]interface{}); ok {
			for _, mod := range mods {
				modifiers = append(modifiers, mod.(string))
			}
		}
		return agent.app.computerAutomation.PressKey(key, modifiers)

	case "wait":
		duration := 1000 // Default 1 second
		if d, ok := step.Parameters["duration"].(float64); ok {
			duration = int(d)
		}
		time.Sleep(time.Duration(duration) * time.Millisecond)

	case "analyze":
		// Perform screen analysis
		uiElements, err := agent.app.computerVision.DetectUIElements()
		if err != nil {
			return err
		}
		step.Result = uiElements

	default:
		return fmt.Errorf("unknown step type: %s", step.Type)
	}

	return nil
}

// Context and Memory Management

// captureContext captures the current system context
func (agent *AIAgent) captureContext() (*ContextSnapshot, error) {
	screenshot, err := agent.app.computerVision.CaptureScreen(0)
	if err != nil {
		return nil, err
	}

	uiElements, _ := agent.app.computerVision.DetectUIElements()
	activeWindow, _ := agent.app.computerAutomation.GetActiveWindow()
	mouseX, mouseY := agent.app.computerAutomation.GetMousePosition()

	context := &ContextSnapshot{
		Timestamp:    time.Now(),
		Screenshot:   screenshot,
		UIElements:   uiElements,
		ActiveWindow: activeWindow,
		MousePos:     [2]int{mouseX, mouseY},
		SystemState:  make(map[string]interface{}),
	}

	return context, nil
}

// addToContext adds a context snapshot to history
func (agent *AIAgent) addToContext(context ContextSnapshot) {
	agent.mu.Lock()
	defer agent.mu.Unlock()

	agent.contextHistory = append(agent.contextHistory, context)

	// Keep history size manageable
	if len(agent.contextHistory) > 100 {
		agent.contextHistory = agent.contextHistory[1:]
	}
}

// monitorSystem continuously monitors system state
func (agent *AIAgent) monitorSystem() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for agent.isActive {
		select {
		case <-ticker.C:
			context, err := agent.captureContext()
			if err != nil {
				runtime.LogWarningf(agent.app.ctx, "Failed to capture context: %v", err)
				continue
			}

			agent.addToContext(*context)

			// Emit context update event
			runtime.EventsEmit(agent.app.ctx, "contextUpdate", context)

		case <-agent.app.ctx.Done():
			return
		}
	}
}

// Task Manager Methods

// AddTask adds a task to the queue
func (tm *TaskManager) AddTask(task *Task) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	tm.tasks[task.ID] = task
	tm.queue = append(tm.queue, task)
}

// GetTask retrieves a task by ID
func (tm *TaskManager) GetTask(id string) (*Task, bool) {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	task, exists := tm.tasks[id]
	return task, exists
}

// MoveToHistory moves a completed task to history
func (tm *TaskManager) MoveToHistory(task *Task) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	// Remove from queue
	for i, t := range tm.queue {
		if t.ID == task.ID {
			tm.queue = append(tm.queue[:i], tm.queue[i+1:]...)
			break
		}
	}

	// Add to history
	tm.history = append(tm.history, task)

	// Keep history size under limit
	if len(tm.history) > tm.maxHistory {
		tm.history = tm.history[1:]
	}
}

// Wails-bound methods for frontend

// StartAgent starts the AI agent (Wails method)
func (a *App) StartAgent() error {
	if a.aiAgent == nil {
		return fmt.Errorf("AI agent not initialized")
	}
	return a.aiAgent.Start()
}

// StopAgent stops the AI agent (Wails method)
func (a *App) StopAgent() {
	if a.aiAgent != nil {
		a.aiAgent.Stop()
	}
}

// ExecuteAgentTask executes a task via the AI agent (Wails method)
func (a *App) ExecuteAgentTask(description string, priority int) (*Task, error) {
	if a.aiAgent == nil {
		return nil, fmt.Errorf("AI agent not initialized")
	}
	return a.aiAgent.ExecuteTask(description, priority)
}

// GetAgentStatus returns the current agent status (Wails method)
func (a *App) GetAgentStatus() map[string]interface{} {
	if a.aiAgent == nil {
		return map[string]interface{}{
			"active":  false,
			"message": "Agent not initialized",
		}
	}

	a.aiAgent.mu.RLock()
	defer a.aiAgent.mu.RUnlock()

	return map[string]interface{}{
		"active":         a.aiAgent.isActive,
		"currentTask":    a.aiAgent.currentTask,
		"contextHistory": len(a.aiAgent.contextHistory),
	}
}

// GetAgentConfig returns the agent configuration (Wails method)
func (a *App) GetAgentConfig() *AgentConfig {
	if a.aiAgent == nil {
		return nil
	}
	return a.aiAgent.config
}

// UpdateAgentConfig updates the agent configuration (Wails method)
func (a *App) UpdateAgentConfig(config *AgentConfig) error {
	if a.aiAgent == nil {
		return fmt.Errorf("AI agent not initialized")
	}

	a.aiAgent.mu.Lock()
	defer a.aiAgent.mu.Unlock()

	a.aiAgent.config = config
	runtime.LogInfo(a.app.ctx, "Agent configuration updated")

	return nil
}
