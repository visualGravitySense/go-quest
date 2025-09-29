# 🌌 Cosmic Cyberpunk Room Escape

A futuristic text-based adventure game written in Go featuring cyberpunk themes, quantum puzzles, and cosmic challenges.

## 🎮 Game Features

- **5 Quest Categories**: Hacking, Engineering, Astronomy, Biology, and Physics
- **100 Unique Quests**: Based on the comprehensive quest list from `quests-1.md`
- **Player Stats System**: Track your skills in different areas
- **Time Management**: Limited time to complete all quests
- **Energy System**: Manage your energy levels
- **Interactive ASCII Art**: Beautiful visual representations
- **Multiple Rooms**: Explore different areas of the cyberpunk facility
- **Reward System**: Earn items and experience for completing quests

## 🚀 How to Play

1. **Run the game**: `go run main.go`
2. **Explore the facility**: Use `look` to examine your surroundings
3. **Check your quests**: Use `quests` to see available challenges
4. **Start a quest**: Use `start <quest_id>` to begin a quest
5. **Solve puzzles**: Enter solutions when prompted
6. **Manage resources**: Watch your energy and time
7. **Complete all quests** to escape!

## 🎯 Available Commands

- `look` or `l` - Look around the current room
- `take <item>` - Pick up an item
- `inventory` or `i` - Check your inventory
- `use <item>` - Use an item
- `go <direction>` - Move in a direction
- `quests` or `q` - Show your active quests
- `start <quest_id>` - Start a specific quest
- `hints <quest_id>` - Show hints for a quest
- `stats` or `s` - Show detailed player statistics
- `help` or `h` - Show help
- `quit` or `exit` - Exit the game

## 🏆 Quest Categories

### 💻 Hacker Quests (1-20)
- Hologram hacking
- Neural interface puzzles
- Quantum password systems
- Binary code decryption

### ⚙️ Engineering Quests (21-40)
- Energy grid management
- Gravity generator configuration
- Plasma resonator tuning
- Anti-matter containment

### ⭐ Astronomical Quests (41-60)
- Star map navigation
- Planetary alignment
- Solar flare prediction
- Black hole trajectory

### 🧬 Biological Quests (61-80)
- DNA modification
- Synthetic organ connection
- Neural impulse stimulation
- Genetic lock systems

### ⚡ Physical Quests (81-100)
- Floating platform navigation
- Holographic wall detection
- Time paradox resolution
- Quantum entanglement

## 🎮 Game Mechanics

- **Time Limit**: You have 60 minutes to complete all quests
- **Energy System**: Actions consume energy, wrong answers cost more
- **Skill Progression**: Complete quests to improve your abilities
- **Adaptive Difficulty**: Quest complexity varies
- **Multiple Solutions**: Some quests may have alternative answers
- **Hint System**: Get helpful hints for any quest using `hints <quest_id>`
- **Progressive Hints**: Each quest has 3 levels of hints from basic to specific

## 🛠️ Requirements

- Go 1.21 or later
- Terminal with ANSI color support (for best experience)

## 🎨 Visual Features

- Colorful ASCII art for rooms and items
- Status indicators for quests and stats
- Progress tracking
- Immersive cyberpunk atmosphere

## 🏁 Victory Condition

Complete all 5 randomly selected quests to escape the Cosmic Cyberpunk Room!

## 🚀 Getting Started

1. Make sure you have Go installed on your system
2. Clone or download this project
3. Run the game:
   ```bash
   go run main.go
   ```

## 🏗️ Project Structure

- `main.go` - Main game logic with quest system
- `go.mod` - Go module definition
- `quests-1.md` - Complete list of 100 quest ideas
- `README.md` - This documentation

Good luck, cyberpunk explorer! 🌌✨