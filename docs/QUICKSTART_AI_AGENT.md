# 🚀 AI Agent Co-pilot Quickstart Guide

Get up and running with the Shotgun AI Agent Co-pilot in minutes! This guide will walk you through setup, configuration, and your first automated tasks.

## 📋 Prerequisites

Before you begin, ensure you have:

- **Go 1.21+** installed
- **Node.js LTS** for the frontend
- **Wails CLI** installed: `go install github.com/wailsapp/wails/v2/cmd/wails@latest`
- An **OpenAI API key** (or other supported AI provider)
- **Administrator/root permissions** (required for automation on some systems)

## 🛠️ Installation

### 1. Install Dependencies

```bash
# Install Go dependencies
go mod tidy

# Install frontend dependencies
cd frontend
npm install
cd ..

# Install Python dependencies (optional, for extensions)
pip install -r requirements.txt
```

### 2. Platform-Specific Setup

#### **macOS**
```bash
# Grant accessibility permissions
# System Preferences → Security & Privacy → Accessibility
# Add your terminal and the built application

# Install Tesseract for OCR (optional)
brew install tesseract
```

#### **Linux**
```bash
# Install X11 development packages
sudo apt-get install libx11-dev libxtst-dev libxinerama-dev libxrandr-dev libxcursor-dev

# Install Tesseract for OCR (optional)
sudo apt-get install tesseract-ocr
```

#### **Windows**
```bash
# Install Tesseract for OCR (optional)
# Download from: https://github.com/tesseract-ocr/tesseract/wiki
# Add to PATH

# May require running as Administrator for automation features
```

## 🎯 Quick Setup

### 1. Build and Run

```bash
# Development mode (hot reload)
wails dev

# Or build for production
wails build
```

### 2. Initial Configuration

When the app opens:

1. **Navigate to the AI Agent Panel** (new tab in the interface)
2. **Click the Settings ⚙️ button**
3. **Configure your AI provider:**
   - Select "OpenAI" as provider
   - Enter your API key
   - Enable Computer Vision and Automation
   - Save settings

### 3. Start the Agent

Click the **"Start Agent"** button in the AI Agent panel. You should see:
- Status indicator turns green
- "AI Agent started successfully" in the activity log
- Agent monitoring begins

## 🎮 First Tasks

### Task 1: Take a Screenshot
1. In the task input box, type: `"Take a screenshot"`
2. Click **"Execute Task"**
3. Watch the agent capture your screen and display it in the preview area

### Task 2: Computer Vision Demo
1. Click **"📸 Capture Screen"** to take a manual screenshot
2. Click **"🔍 Detect UI Elements"** to analyze the interface
3. Enable **"Show UI Elements"** to see detected components highlighted

### Task 3: Simple Automation
1. Type: `"Open calculator"`
2. Click **"Execute Task"**
3. The agent will analyze your screen and launch the calculator app

### Task 4: Text Recognition
1. Ensure you have some text visible on screen
2. Type: `"Read all text on the screen and summarize it"`
3. The agent will perform OCR and provide a summary

## 🔧 Configuration Options

### Agent Settings

| Setting | Description | Default |
|---------|-------------|---------|
| **Model Provider** | AI service (OpenAI, Anthropic, Local) | OpenAI |
| **Vision Enabled** | Enable computer vision features | ✅ |
| **Automation Enabled** | Allow computer control | ✅ |
| **Safety Checks** | Confirm before risky actions | ✅ |
| **Custom Instructions** | Additional behavior guidelines | Empty |

### Computer Vision Controls

- **📸 Capture Screen**: Single screenshot
- **▶️ Start Live Capture**: Continuous monitoring (1 sec interval)
- **🔍 Detect UI Elements**: Analyze current screen for clickable elements
- **👁️ Show UI Elements**: Overlay detected elements on preview

### Quick Actions

Pre-configured common tasks:
- **🌐 Browser**: Open web browser
- **💻 Terminal**: Open terminal/command prompt
- **📷 Screenshot**: Take and save screenshot
- **👁️ Read Screen**: OCR analysis of visible text

## 💡 Example Tasks

### Simple Commands
```
"Click on the save button"
"Type 'Hello World'"
"Press Ctrl+C"
"Scroll down"
"Open notepad"
```

### Multi-Step Tasks
```
"Take a screenshot, save it to desktop, then open it in an image viewer"
"Open browser, go to google.com, search for 'AI automation'"
"Find the calculator app and calculate 15 * 23"
```

### Advanced Workflows
```
"Monitor this window and notify me when it changes"
"Take a screenshot every 30 seconds for the next 5 minutes"
"Read all text on screen, copy it to clipboard, then paste it in notepad"
```

## 🐛 Troubleshooting

### Agent Won't Start
```bash
# Check logs in the Activity Log
# Common issues:
- Missing API key
- Network connectivity
- Permission issues
```

### Screen Capture Fails
```bash
# macOS: Grant Screen Recording permission
# System Preferences → Security & Privacy → Screen Recording

# Linux: Check X11 permissions
xhost +local:

# Windows: Run as Administrator
```

### Automation Not Working
```bash
# macOS: Enable Accessibility
# System Preferences → Security & Privacy → Accessibility

# Linux: Install automation packages
sudo apt-get install xdotool

# Windows: Check UAC settings
```

### Performance Issues
- Reduce live capture frequency
- Disable UI element detection for simple tasks
- Clear screenshot history regularly

## 🔐 Security Best Practices

### Safe Usage
1. **Start with simple tasks** to understand behavior
2. **Enable safety checks** for all automation
3. **Review activity logs** regularly
4. **Test in isolated environments** first
5. **Use specific instructions** rather than vague commands

### API Key Security
- Store API keys securely
- Use environment variables in production
- Rotate keys regularly
- Monitor usage and costs

### Permission Management
- Grant minimal required permissions
- Review automation scope regularly
- Use separate user accounts for testing
- Monitor system access logs

## 📊 Monitoring and Logs

### Activity Log
The activity log shows:
- Task execution status
- Computer vision operations
- Automation actions
- Error messages and warnings

### Performance Metrics
Monitor:
- Task completion rates
- Average execution time
- API usage and costs
- System resource usage

## 🎯 Next Steps

### Explore Advanced Features
1. **Action Recording**: Record and replay complex workflows
2. **Custom Prompts**: Create specialized task templates
3. **Context Learning**: Let the agent learn from your patterns
4. **Multi-Display**: Configure for multiple monitors

### Integration Ideas
- **Development Workflows**: Automate code building and testing
- **Content Creation**: Screenshot documentation automation
- **System Administration**: Automated monitoring and alerts
- **Accessibility**: Voice-controlled computer interaction

### Customization
- Modify AI prompts for your specific use case
- Add custom quick actions
- Integrate with your existing tools
- Create automation templates

## 📚 Learn More

- [Complete Documentation](./AI_AGENT_COPILOT.md)
- [API Reference](./AI_AGENT_COPILOT.md#api-reference)
- [Security Guide](./AI_AGENT_COPILOT.md#security-considerations)
- [Troubleshooting](./AI_AGENT_COPILOT.md#troubleshooting)

## 🆘 Getting Help

### Common Resources
- Check the Activity Log for error details
- Review the troubleshooting section
- Test with simple tasks first
- Verify permissions and dependencies

### Community Support
- Create GitHub issues for bugs
- Share automation ideas and tips
- Contribute improvements and features
- Help others with setup questions

---

**⚡ Pro Tip**: Start with simple tasks like screenshots and text reading before attempting complex automation workflows. The agent learns from context and becomes more effective with practice!

**🛡️ Safety First**: Always test automation in safe environments and review what the agent plans to do before execution in production environments.