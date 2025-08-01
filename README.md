![](https://github.com/user-attachments/assets/058bf4a2-9f81-406c-96ea-795cd4eaf118)

**Tired of Cursor cutting off context, missing your files and folders, and spitting out empty responses?**

Save your context with Shotgun!
→ Prepare a truly GIGANTIC prompt
→ Paste it into **Google AI Studio** and receive a massive patch for your code. 25 free queries per day!
→ Drop that patch into Cursor or Windsurf and apply the entire diff in a single request.

**That means you get 25 huge, fully coherent patches per day for your codebase—absolutely free, thanks to complete context transfer.**

**🤖 NEW: Introducing AI Agent Co-pilot!**
Now with comprehensive computer use and computer vision capabilities! Your AI assistant can see your screen, control your computer, and help with everyday tasks through natural language commands.

Perfect for dynamically-typed languages:

Python
JavaScript

# Shotgun App + AI Agent Co-pilot

*One‑click codebase "blast" for Large‑Language‑Model workflows + intelligent computer automation*

---

## 1. What Shotgun Does

### Original Shotgun Features
Shotgun is a tiny desktop tool that **explodes an entire project into a single,
well‑structured text payload** designed for AI assistants.
Think of it as a rapid‑fire alternative to copy‑pasting dozens of files by hand:

*   **Select a folder → get an instant tree + file dump**
    in a predictable delimiter format (`*#*#*...*#*#*begin … *#*#*end*#*#*`).
*   **Tick check‑boxes to exclude noise** (logs, build artifacts, `node_modules`, …).
*   **Paste the result into ChatGPT, Gemini 2.5, Cursor, etc.**
    to ask for multi‑file edits, refactors, bug fixes, reviews, or documentation.
*   **Receive a diff‑style reply** and apply changes with your favourite patch tool.

### 🤖 New AI Agent Co-pilot Features
The AI Agent Co-pilot adds comprehensive computer automation:

*   **🎯 Natural Language Commands** - "Take a screenshot and save it", "Open calculator", "Read text on screen"
*   **👁️ Computer Vision** - Real-time screen capture, UI element detection, OCR text recognition
*   **🖱️ Smart Automation** - Mouse control, keyboard input, window management, app launching
*   **🧠 Context Awareness** - Understands screen content and user intent
*   **🛡️ Safety Features** - Confirmation prompts, action logging, permission controls
*   **📊 Task Management** - Queue complex workflows, monitor progress, replay actions

---

## 2. Why You Might Need It

| Scenario                 | Traditional Shotgun Benefit | AI Agent Co-pilot Benefit |
|--------------------------|-------------------------------|---------------------------|
| **Bulk bug fixing**      | Generates complete snapshot for LLM | Automates testing and verification |
| **Large‑scale refactor** | LLM gets full context | Helps navigate and modify UI elements |
| **Documentation**        | Produce searchable text files | Screenshots and automates doc generation |
| **Accessibility**        | N/A | Voice commands and screen reading |
| **Repetitive Tasks**     | N/A | Record and replay complex workflows |
| **Screen Analysis**      | N/A | Understand UI layouts and extract information |

---

## 3. Key Features

### Original Shotgun Features
*   ⚡ **Fast tree scan** (Go + Wails backend) – thousands of files in milliseconds.
*   ✅ **Interactive exclude list** – skip folders, temporary files, or secrets.
*   📝 **Deterministic delimiters** – easy for LLMs to parse and for you to split.
*   🔄 **Re‑generate anytime** – tweak the excludes and hit *Shotgun* again.
*   🪶 **Lightweight** – no DB, no cloud; a single native executable plus a Vue UI.
*   🖥️ **Cross‑platform** – Windows, macOS, Linux.

### 🤖 New AI Agent Features
*   🎯 **Natural Language Processing** – Understand and execute spoken/typed commands
*   👁️ **Computer Vision** – Screen capture, UI detection, OCR, visual analysis
*   🖱️ **Computer Automation** – Mouse, keyboard, window control, app management
*   🧠 **AI Integration** – OpenAI, Anthropic, local models for intelligent planning
*   📊 **Task Management** – Queue, monitor, and replay complex workflows
*   🛡️ **Safety & Security** – Confirmation prompts, logging, permission controls
*   📱 **Real-time Preview** – Live screen monitoring and element highlighting
*   🎬 **Action Recording** – Record and replay manual workflows

---

## 4. Quick Start

### Traditional Shotgun Workflow
1.  **Step 1: Prepare Context** - Select project, exclude unwanted files
2.  **Step 2: Compose Prompt** - Generate context and create LLM prompts  
3.  **Step 3: Execute Prompt** - Send to AI service and get responses
4.  **Step 4: Apply Patch** - Apply the generated changes

### 🤖 New AI Agent Workflow
1.  **Start the Agent** - Click "Start Agent" in the AI Agent panel
2.  **Configure Settings** - Set up AI provider (OpenAI, Anthropic, etc.)
3.  **Give Commands** - Type natural language tasks like:
    - `"Take a screenshot and analyze the UI elements"`
    - `"Open calculator and compute 15 * 23"`
    - `"Read all text on screen and summarize it"`
    - `"Monitor this window and alert me when it changes"`
4.  **Watch Magic Happen** - The agent sees your screen and controls your computer

---

## 5. Installation

### 5.1. Prerequisites
*   **Go ≥ 1.21**   `go version`
*   **Node.js LTS**  `node -v`
*   **Wails CLI**    `go install github.com/wailsapp/wails/v2/cmd/wails@latest`
*   **🤖 AI API Key**  OpenAI, Anthropic, or other supported provider

### 5.2. Clone & Bootstrap
```bash
git clone https://github.com/glebkudr/shotgun_code
cd shotgun_code
go mod tidy           # backend deps
cd frontend
npm install           # Vue deps
cd ..

# Optional: Install Python extensions for advanced computer vision
pip install -r requirements.txt
```

### 5.3. Run in Dev Mode
```bash
wails dev
```
Hot‑reloads Vue; restart the command for Go code changes.

### 5.4. Build a Release
```bash
wails build           # binaries land in build/bin/
```

---

## 6. Quick‑Start Workflows

### Original Shotgun Workflow
1.  Run `wails dev`. The app window will open.
2.  **Step 1: Prepare Context** - Select project folder, exclude unwanted items
3.  **Step 2: Compose Prompt** - Generate context and create prompts
4.  **Step 3: Execute Prompt** - Process with AI services  
5.  **Step 4: Apply Patch** - Apply generated changes

### 🤖 AI Agent Workflow
1.  **Navigate to AI Agent Panel** - New tab in the interface
2.  **Configure Agent** - Click settings ⚙️, add API key, enable features
3.  **Start Agent** - Click "Start Agent" button
4.  **Try Simple Tasks**:
    - `"Take a screenshot"` - Captures and displays your screen
    - `"Open calculator"` - Launches calculator app
    - `"Detect UI elements"` - Analyzes current screen
    - `"Type 'Hello World'"` - Types text automatically
5.  **Advanced Commands**:
    - `"Monitor this window for changes"`
    - `"Read all text and copy to clipboard"`
    - `"Take a screenshot every 30 seconds for 5 minutes"`

---

## 7. Documentation

### Core Documentation
*   **[AI Agent Co-pilot Guide](docs/AI_AGENT_COPILOT.md)** - Complete feature documentation
*   **[Quickstart Guide](docs/QUICKSTART_AI_AGENT.md)** - Get started in minutes
*   **[Original Architecture](design/architecture.md)** - Technical details

### Example Outputs

#### Traditional Shotgun Output
```text
app/
├── main.go
├── app.go
└── frontend/
    ├── App.vue
    └── components/
        └── FileTree.vue

<file path="main.go">
package main
...
</file>
```

#### 🤖 AI Agent Capabilities
- **Screen Analysis**: Detects buttons, text fields, images, and interactive elements
- **Smart Automation**: Contextual clicking, typing, and navigation
- **Task Orchestration**: Multi-step workflows with error handling
- **Visual Feedback**: Real-time overlays showing detected elements

---

## 8. Best Practices

### Traditional Shotgun
*   **Trim the noise** – exclude lock files, vendored libs, generated assets.
*   **Ask for diffs, not whole files** – keeps responses concise.
*   **Iterate** – generate → ask → patch → re‑generate if needed.

### 🤖 AI Agent Safety
*   **Start Simple** – Begin with screenshots and basic commands
*   **Enable Safety Checks** – Always confirm destructive actions
*   **Review Logs** – Monitor what the agent does
*   **Test First** – Try automation in safe environments
*   **Use Specific Commands** – Clear instructions work better than vague requests

---

## 9. Troubleshooting

| Symptom                     | Fix                                                          |
|-----------------------------|--------------------------------------------------------------|
| `wails: command not found`  | Ensure `$GOROOT/bin` or `$HOME/go/bin` is on `PATH`.         |
| Blank window on `wails dev` | Check Node version & reinstall frontend deps.              |
| 🤖 Agent won't start       | Check API key, permissions, and dependency installation     |
| 🤖 Screen capture fails    | Grant screen recording permissions (macOS/Linux)           |
| 🤖 Automation blocked      | Enable accessibility permissions, run as admin (Windows)   |

### Platform-Specific Setup

#### macOS
```bash
# Grant permissions in System Preferences → Security & Privacy:
# - Screen Recording (for screenshots)
# - Accessibility (for automation)
```

#### Linux  
```bash
# Install automation dependencies
sudo apt-get install libx11-dev libxtst-dev xdotool
```

#### Windows
```bash
# May need to run as Administrator for automation features
# Install Visual C++ Redistributable if needed
```

---

## 10. Roadmap

### ✅ Completed Features
- ✅ **Core Shotgun functionality** - File context generation
- ✅ **AI Agent Co-pilot** - Computer vision and automation
- ✅ **Screen Capture** - Real-time monitoring and analysis  
- ✅ **UI Detection** - Element recognition and interaction
- ✅ **Task Management** - Natural language command processing
- ✅ **Safety Controls** - Permission management and logging

### 🚀 Coming Soon
- ☐ **Advanced OCR** - Layout-preserving text extraction
- ☐ **Voice Commands** - Speech-to-text task input
- ☐ **Mobile Control** - iOS/Android device automation
- ☐ **Web Automation** - Browser interaction capabilities
- ☐ **Multi-Monitor** - Enhanced multi-display support
- ☐ **Custom Actions** - User-defined automation templates
- ☐ **Learning Mode** - Pattern recognition and optimization
- ☐ **API Extensions** - Third-party service integrations

### 🔬 Research Areas
- Machine learning-based UI prediction
- Cross-platform automation consistency
- Advanced computer vision algorithms
- Natural language understanding improvements

---

## 11. Contributing

PRs and issues are welcome!

### Development Guidelines
- Format Go code with `go fmt`
- Follow Vue 3 style guidelines  
- Test AI agent features in safe environments
- Document new automation capabilities
- Include safety considerations in PRs

### 🤖 AI Agent Contributions
When contributing to AI Agent features:
- Test across multiple platforms
- Include safety checks and permission handling
- Document new computer vision capabilities
- Provide example commands and use cases

---

## 12. License

Custom MIT-like – see `LICENSE.md` file.

**Security Notice**: The AI Agent Co-pilot can control your computer. Always:
- Review what the agent plans to do
- Test in safe environments first  
- Keep safety checks enabled
- Monitor automation logs
- Use appropriate API key security

---

Shotgun – load, aim, blast your code straight into the mind of an LLM.
AI Agent Co-pilot – let AI see and control your computer to get things done.

**Iterate faster. Ship better. Automate smarter.** 🚀
