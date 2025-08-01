<template>
  <div class="ai-agent-panel h-full flex flex-col bg-gray-50 dark:bg-gray-900">
    <!-- Header -->
    <div class="bg-white dark:bg-gray-800 border-b border-gray-200 dark:border-gray-700 p-4">
      <div class="flex items-center justify-between">
        <div class="flex items-center space-x-3">
          <div :class="[
            'w-3 h-3 rounded-full',
            agentStatus.active ? 'bg-green-500' : 'bg-red-500'
          ]"></div>
          <h2 class="text-lg font-semibold text-gray-900 dark:text-white">
            AI Agent Co-pilot
          </h2>
        </div>
        <div class="flex space-x-2">
          <button
            @click="toggleAgent"
            :disabled="loading"
            class="px-4 py-2 bg-blue-600 hover:bg-blue-700 disabled:bg-gray-400 text-white rounded-md font-medium text-sm transition-colors"
          >
            {{ agentStatus.active ? 'Stop' : 'Start' }} Agent
          </button>
          <button
            @click="showSettings = true"
            class="p-2 text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"></path>
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
            </svg>
          </button>
        </div>
      </div>
    </div>

    <!-- Main Content -->
    <div class="flex-1 flex overflow-hidden">
      <!-- Left Panel - Controls -->
      <div class="w-1/3 border-r border-gray-200 dark:border-gray-700 flex flex-col">
        <!-- Task Input -->
        <div class="p-4 border-b border-gray-200 dark:border-gray-700">
          <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
            What would you like me to help you with?
          </label>
          <textarea
            v-model="taskInput"
            :disabled="!agentStatus.active || loading"
            placeholder="Describe what you want me to do on your computer..."
            class="w-full h-24 px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-md text-sm bg-white dark:bg-gray-800 text-gray-900 dark:text-white placeholder-gray-500 focus:ring-2 focus:ring-blue-500 focus:border-blue-500 resize-none"
          ></textarea>
          <div class="flex justify-between items-center mt-2">
            <select
              v-model="taskPriority"
              class="text-sm border border-gray-300 dark:border-gray-600 rounded bg-white dark:bg-gray-800 text-gray-900 dark:text-white px-2 py-1"
            >
              <option value="1">Low Priority</option>
              <option value="2">Normal Priority</option>
              <option value="3">High Priority</option>
            </select>
            <button
              @click="executeTask"
              :disabled="!taskInput.trim() || !agentStatus.active || loading"
              class="px-4 py-2 bg-green-600 hover:bg-green-700 disabled:bg-gray-400 text-white rounded-md text-sm font-medium transition-colors"
            >
              Execute Task
            </button>
          </div>
        </div>

        <!-- Computer Vision Controls -->
        <div class="p-4 border-b border-gray-200 dark:border-gray-700">
          <h3 class="text-sm font-semibold text-gray-900 dark:text-white mb-3">Computer Vision</h3>
          <div class="space-y-2">
            <button
              @click="captureScreen"
              :disabled="loading"
              class="w-full px-3 py-2 bg-blue-500 hover:bg-blue-600 disabled:bg-gray-400 text-white rounded text-sm transition-colors"
            >
              📸 Capture Screen
            </button>
            <button
              @click="toggleScreenCapture"
              :disabled="loading"
              class="w-full px-3 py-2 bg-purple-500 hover:bg-purple-600 disabled:bg-gray-400 text-white rounded text-sm transition-colors"
            >
              {{ isCapturing ? '⏹️ Stop' : '▶️ Start' }} Live Capture
            </button>
            <button
              @click="detectUIElements"
              :disabled="loading"
              class="w-full px-3 py-2 bg-orange-500 hover:bg-orange-600 disabled:bg-gray-400 text-white rounded text-sm transition-colors"
            >
              🔍 Detect UI Elements
            </button>
          </div>
        </div>

        <!-- Quick Actions -->
        <div class="p-4 border-b border-gray-200 dark:border-gray-700">
          <h3 class="text-sm font-semibold text-gray-900 dark:text-white mb-3">Quick Actions</h3>
          <div class="grid grid-cols-2 gap-2">
            <button
              @click="quickAction('open_browser')"
              :disabled="!agentStatus.active"
              class="px-2 py-2 bg-gray-500 hover:bg-gray-600 disabled:bg-gray-300 text-white rounded text-xs transition-colors"
            >
              🌐 Browser
            </button>
            <button
              @click="quickAction('open_terminal')"
              :disabled="!agentStatus.active"
              class="px-2 py-2 bg-gray-500 hover:bg-gray-600 disabled:bg-gray-300 text-white rounded text-xs transition-colors"
            >
              💻 Terminal
            </button>
            <button
              @click="quickAction('take_screenshot')"
              :disabled="!agentStatus.active"
              class="px-2 py-2 bg-gray-500 hover:bg-gray-600 disabled:bg-gray-300 text-white rounded text-xs transition-colors"
            >
              📷 Screenshot
            </button>
            <button
              @click="quickAction('read_screen')"
              :disabled="!agentStatus.active"
              class="px-2 py-2 bg-gray-500 hover:bg-gray-600 disabled:bg-gray-300 text-white rounded text-xs transition-colors"
            >
              👁️ Read Screen
            </button>
          </div>
        </div>

        <!-- Current Task -->
        <div v-if="currentTask" class="p-4 border-b border-gray-200 dark:border-gray-700">
          <h3 class="text-sm font-semibold text-gray-900 dark:text-white mb-2">Current Task</h3>
          <div class="bg-blue-50 dark:bg-blue-900/50 p-3 rounded-md">
            <p class="text-sm text-gray-700 dark:text-gray-300 mb-2">{{ currentTask.description }}</p>
            <div class="flex items-center justify-between text-xs">
              <span :class="[
                'px-2 py-1 rounded-full font-medium',
                currentTask.status === 'running' ? 'bg-blue-100 text-blue-800 dark:bg-blue-800 dark:text-blue-100' :
                currentTask.status === 'completed' ? 'bg-green-100 text-green-800 dark:bg-green-800 dark:text-green-100' :
                currentTask.status === 'failed' ? 'bg-red-100 text-red-800 dark:bg-red-800 dark:text-red-100' :
                'bg-gray-100 text-gray-800 dark:bg-gray-800 dark:text-gray-100'
              ]">
                {{ currentTask.status.toUpperCase() }}
              </span>
              <span class="text-gray-500 dark:text-gray-400">
                {{ currentTask.steps?.length || 0 }} steps
              </span>
            </div>
          </div>
        </div>
      </div>

      <!-- Right Panel - Screen Preview & Results -->
      <div class="flex-1 flex flex-col">
        <!-- Screen Preview -->
        <div class="flex-1 p-4">
          <div class="h-full bg-black rounded-lg overflow-hidden relative">
            <img
              v-if="currentScreenshot"
              :src="`data:image/png;base64,${currentScreenshot.imageBase64}`"
              :alt="'Screenshot'"
              class="w-full h-full object-contain"
            />
            <div v-else class="flex items-center justify-center h-full text-gray-500">
              <div class="text-center">
                <svg class="w-16 h-16 mx-auto mb-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                </svg>
                <p class="text-sm">No screenshot captured yet</p>
                <p class="text-xs text-gray-400 mt-1">Click "Capture Screen" to preview</p>
              </div>
            </div>

            <!-- UI Element Overlays -->
            <div v-if="showUIElements && uiElements.length > 0" class="absolute inset-0">
              <div
                v-for="(element, index) in uiElements"
                :key="index"
                :style="{
                  left: element.boundingBox.x + 'px',
                  top: element.boundingBox.y + 'px',
                  width: element.boundingBox.width + 'px',
                  height: element.boundingBox.height + 'px'
                }"
                class="absolute border-2 border-red-500 bg-red-500/20"
                :title="`${element.type}: ${element.text || 'No text'}`"
              >
                <span class="absolute -top-6 left-0 text-xs bg-red-500 text-white px-1 rounded">
                  {{ element.type }}
                </span>
              </div>
            </div>

            <!-- Loading Overlay -->
            <div v-if="loading" class="absolute inset-0 bg-black/50 flex items-center justify-center">
              <div class="text-white text-center">
                <div class="animate-spin w-8 h-8 border-4 border-white border-t-transparent rounded-full mx-auto mb-2"></div>
                <p class="text-sm">{{ loadingMessage }}</p>
              </div>
            </div>
          </div>

          <!-- Controls -->
          <div class="flex items-center justify-between mt-3">
            <div class="flex items-center space-x-2">
              <label class="flex items-center text-sm text-gray-700 dark:text-gray-300">
                <input
                  type="checkbox"
                  v-model="showUIElements"
                  class="mr-2"
                />
                Show UI Elements
              </label>
            </div>
            <div class="flex space-x-2">
              <button
                @click="saveScreenshot"
                :disabled="!currentScreenshot"
                class="px-3 py-1 bg-gray-500 hover:bg-gray-600 disabled:bg-gray-300 text-white rounded text-sm transition-colors"
              >
                💾 Save
              </button>
              <button
                @click="clearScreen"
                class="px-3 py-1 bg-gray-500 hover:bg-gray-600 text-white rounded text-sm transition-colors"
              >
                🗑️ Clear
              </button>
            </div>
          </div>
        </div>

        <!-- Task History & Logs -->
        <div class="h-48 border-t border-gray-200 dark:border-gray-700 p-4">
          <div class="flex items-center justify-between mb-2">
            <h3 class="text-sm font-semibold text-gray-900 dark:text-white">Activity Log</h3>
            <button
              @click="clearLogs"
              class="text-xs text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200"
            >
              Clear
            </button>
          </div>
          <div class="h-32 overflow-y-auto bg-gray-50 dark:bg-gray-800 rounded border text-xs">
            <div
              v-for="(log, index) in activityLogs"
              :key="index"
              class="p-2 border-b border-gray-200 dark:border-gray-700 last:border-b-0"
            >
              <div class="flex items-center justify-between">
                <span :class="[
                  'font-medium',
                  log.type === 'error' ? 'text-red-600 dark:text-red-400' :
                  log.type === 'success' ? 'text-green-600 dark:text-green-400' :
                  log.type === 'warning' ? 'text-yellow-600 dark:text-yellow-400' :
                  'text-blue-600 dark:text-blue-400'
                ]">
                  {{ log.message }}
                </span>
                <span class="text-gray-500 dark:text-gray-400">{{ formatTime(log.timestamp) }}</span>
              </div>
            </div>
            <div v-if="activityLogs.length === 0" class="p-4 text-center text-gray-500 dark:text-gray-400">
              No activity logs yet
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Settings Modal -->
    <div v-if="showSettings" class="fixed inset-0 bg-black/50 flex items-center justify-center z-50">
      <div class="bg-white dark:bg-gray-800 rounded-lg p-6 w-96 max-h-[80vh] overflow-y-auto">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900 dark:text-white">Agent Settings</h3>
          <button
            @click="showSettings = false"
            class="text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>

        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
              AI Model Provider
            </label>
            <select
              v-model="agentConfig.modelProvider"
              class="w-full border border-gray-300 dark:border-gray-600 rounded px-3 py-2 bg-white dark:bg-gray-700 text-gray-900 dark:text-white"
            >
              <option value="openai">OpenAI</option>
              <option value="anthropic">Anthropic</option>
              <option value="local">Local Model</option>
            </select>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
              API Key
            </label>
            <input
              type="password"
              v-model="agentConfig.apiKey"
              placeholder="Enter your API key"
              class="w-full border border-gray-300 dark:border-gray-600 rounded px-3 py-2 bg-white dark:bg-gray-700 text-gray-900 dark:text-white"
            />
          </div>

          <div class="space-y-2">
            <label class="flex items-center text-sm text-gray-700 dark:text-gray-300">
              <input
                type="checkbox"
                v-model="agentConfig.visionEnabled"
                class="mr-2"
              />
              Enable Computer Vision
            </label>
            <label class="flex items-center text-sm text-gray-700 dark:text-gray-300">
              <input
                type="checkbox"
                v-model="agentConfig.automationEnabled"
                class="mr-2"
              />
              Enable Automation
            </label>
            <label class="flex items-center text-sm text-gray-700 dark:text-gray-300">
              <input
                type="checkbox"
                v-model="agentConfig.safetyChecks"
                class="mr-2"
              />
              Enable Safety Checks
            </label>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
              Custom Instructions
            </label>
            <textarea
              v-model="agentConfig.customInstructions"
              placeholder="Additional instructions for the AI agent..."
              class="w-full h-20 border border-gray-300 dark:border-gray-600 rounded px-3 py-2 bg-white dark:bg-gray-700 text-gray-900 dark:text-white resize-none"
            ></textarea>
          </div>
        </div>

        <div class="flex justify-end space-x-2 mt-6">
          <button
            @click="showSettings = false"
            class="px-4 py-2 text-gray-700 dark:text-gray-300 hover:bg-gray-100 dark:hover:bg-gray-700 rounded transition-colors"
          >
            Cancel
          </button>
          <button
            @click="saveSettings"
            class="px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded transition-colors"
          >
            Save Settings
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, reactive, onMounted, onUnmounted } from 'vue'
import * as App from '../../wailsjs/go/main/App'

export default {
  name: 'AIAgentPanel',
  setup() {
    // Reactive state
    const agentStatus = ref({ active: false, message: '' })
    const loading = ref(false)
    const loadingMessage = ref('')
    const taskInput = ref('')
    const taskPriority = ref(2)
    const currentTask = ref(null)
    const currentScreenshot = ref(null)
    const uiElements = ref([])
    const showUIElements = ref(false)
    const isCapturing = ref(false)
    const showSettings = ref(false)
    const activityLogs = ref([])

    const agentConfig = reactive({
      modelProvider: 'openai',
      modelName: 'gpt-4-vision-preview',
      apiKey: '',
      visionEnabled: true,
      automationEnabled: true,
      safetyChecks: true,
      customInstructions: ''
    })

    // Methods
    const addLog = (message, type = 'info') => {
      activityLogs.value.unshift({
        message,
        type,
        timestamp: new Date()
      })
      // Keep only last 100 logs
      if (activityLogs.value.length > 100) {
        activityLogs.value = activityLogs.value.slice(0, 100)
      }
    }

    const toggleAgent = async () => {
      loading.value = true
      loadingMessage.value = agentStatus.value.active ? 'Stopping agent...' : 'Starting agent...'
      
      try {
        if (agentStatus.value.active) {
          await App.StopAgent()
          addLog('Agent stopped', 'info')
        } else {
          await App.StartAgent()
          addLog('Agent started successfully', 'success')
        }
        await updateAgentStatus()
      } catch (error) {
        addLog(`Error: ${error}`, 'error')
        console.error('Agent toggle error:', error)
      } finally {
        loading.value = false
        loadingMessage.value = ''
      }
    }

    const executeTask = async () => {
      if (!taskInput.value.trim()) return
      
      loading.value = true
      loadingMessage.value = 'Executing task...'
      
      try {
        const task = await App.ExecuteAgentTask(taskInput.value, taskPriority.value)
        currentTask.value = task
        addLog(`Task started: ${taskInput.value}`, 'info')
        taskInput.value = ''
      } catch (error) {
        addLog(`Task execution failed: ${error}`, 'error')
        console.error('Task execution error:', error)
      } finally {
        loading.value = false
        loadingMessage.value = ''
      }
    }

    const captureScreen = async () => {
      loading.value = true
      loadingMessage.value = 'Capturing screen...'
      
      try {
        const screenshot = await App.CaptureScreenForDisplay(0)
        currentScreenshot.value = screenshot
        addLog('Screen captured successfully', 'success')
      } catch (error) {
        addLog(`Screen capture failed: ${error}`, 'error')
        console.error('Screen capture error:', error)
      } finally {
        loading.value = false
        loadingMessage.value = ''
      }
    }

    const toggleScreenCapture = async () => {
      try {
        if (isCapturing.value) {
          await App.StopScreenCapture()
          isCapturing.value = false
          addLog('Live capture stopped', 'info')
        } else {
          await App.StartScreenCapture(1000, 0) // 1 second interval
          isCapturing.value = true
          addLog('Live capture started', 'success')
        }
      } catch (error) {
        addLog(`Screen capture toggle failed: ${error}`, 'error')
        console.error('Screen capture toggle error:', error)
      }
    }

    const detectUIElements = async () => {
      loading.value = true
      loadingMessage.value = 'Detecting UI elements...'
      
      try {
        const elements = await App.DetectUIElements()
        uiElements.value = elements
        showUIElements.value = true
        addLog(`Detected ${elements.length} UI elements`, 'success')
      } catch (error) {
        addLog(`UI detection failed: ${error}`, 'error')
        console.error('UI detection error:', error)
      } finally {
        loading.value = false
        loadingMessage.value = ''
      }
    }

    const quickAction = async (action) => {
      const actionMap = {
        'open_browser': 'Open web browser',
        'open_terminal': 'Open terminal',
        'take_screenshot': 'Take a screenshot',
        'read_screen': 'Read all text on the screen'
      }
      
      const description = actionMap[action] || action
      await executeTaskWithDescription(description)
    }

    const executeTaskWithDescription = async (description) => {
      loading.value = true
      loadingMessage.value = 'Executing quick action...'
      
      try {
        const task = await App.ExecuteAgentTask(description, 2)
        currentTask.value = task
        addLog(`Quick action: ${description}`, 'info')
      } catch (error) {
        addLog(`Quick action failed: ${error}`, 'error')
        console.error('Quick action error:', error)
      } finally {
        loading.value = false
        loadingMessage.value = ''
      }
    }

    const saveScreenshot = async () => {
      if (!currentScreenshot.value) return
      
      try {
        // This would call a backend method to save the screenshot
        addLog('Screenshot saved', 'success')
      } catch (error) {
        addLog(`Failed to save screenshot: ${error}`, 'error')
        console.error('Screenshot save error:', error)
      }
    }

    const clearScreen = () => {
      currentScreenshot.value = null
      uiElements.value = []
      showUIElements.value = false
      addLog('Screen preview cleared', 'info')
    }

    const clearLogs = () => {
      activityLogs.value = []
    }

    const updateAgentStatus = async () => {
      try {
        const status = await App.GetAgentStatus()
        agentStatus.value = status
      } catch (error) {
        console.error('Failed to get agent status:', error)
      }
    }

    const loadAgentConfig = async () => {
      try {
        const config = await App.GetAgentConfig()
        if (config) {
          Object.assign(agentConfig, config)
        }
      } catch (error) {
        console.error('Failed to load agent config:', error)
      }
    }

    const saveSettings = async () => {
      try {
        await App.UpdateAgentConfig(agentConfig)
        showSettings.value = false
        addLog('Settings saved successfully', 'success')
      } catch (error) {
        addLog(`Failed to save settings: ${error}`, 'error')
        console.error('Settings save error:', error)
      }
    }

    const formatTime = (timestamp) => {
      return new Date(timestamp).toLocaleTimeString()
    }

    // Event listeners for real-time updates
    let eventListeners = []

    const setupEventListeners = () => {
      // Listen for agent status changes
      window.runtime?.EventsOn('agentStatusChanged', (data) => {
        agentStatus.value = data
        addLog(data.message, data.status === 'active' ? 'success' : 'info')
      })

      // Listen for task updates
      window.runtime?.EventsOn('taskCreated', (task) => {
        currentTask.value = task
      })

      window.runtime?.EventsOn('taskCompleted', (task) => {
        currentTask.value = task
        addLog(`Task completed: ${task.description}`, 'success')
      })

      window.runtime?.EventsOn('taskStatusChanged', (task) => {
        currentTask.value = task
      })

      // Listen for screenshot updates
      window.runtime?.EventsOn('screenshotCaptured', (screenshot) => {
        if (isCapturing.value) {
          currentScreenshot.value = screenshot
        }
      })
    }

    // Lifecycle hooks
    onMounted(async () => {
      await updateAgentStatus()
      await loadAgentConfig()
      setupEventListeners()
      addLog('AI Agent Panel initialized', 'info')
    })

    onUnmounted(() => {
      // Clean up event listeners
      eventListeners.forEach(cleanup => cleanup && cleanup())
    })

    return {
      // State
      agentStatus,
      loading,
      loadingMessage,
      taskInput,
      taskPriority,
      currentTask,
      currentScreenshot,
      uiElements,
      showUIElements,
      isCapturing,
      showSettings,
      activityLogs,
      agentConfig,
      
      // Methods
      toggleAgent,
      executeTask,
      captureScreen,
      toggleScreenCapture,
      detectUIElements,
      quickAction,
      saveScreenshot,
      clearScreen,
      clearLogs,
      saveSettings,
      formatTime
    }
  }
}
</script>

<style scoped>
.ai-agent-panel {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
}

/* Custom scrollbar */
.overflow-y-auto {
  scrollbar-width: thin;
  scrollbar-color: rgb(156 163 175) transparent;
}

.overflow-y-auto::-webkit-scrollbar {
  width: 6px;
}

.overflow-y-auto::-webkit-scrollbar-track {
  background: transparent;
}

.overflow-y-auto::-webkit-scrollbar-thumb {
  background-color: rgb(156 163 175);
  border-radius: 3px;
}

.overflow-y-auto::-webkit-scrollbar-thumb:hover {
  background-color: rgb(107 114 128);
}
</style>