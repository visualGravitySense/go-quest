package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// ANSI Color codes
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
	ColorBold   = "\033[1m"
)

// Item represents an object in the game
type Item struct {
	Name        string
	Description string
	Usable      bool
	ASCII       string
}

// Room represents a location in the game
type Room struct {
	Name        string
	Description string
	Items       []*Item
	Exits       map[string]*Room
	Solved      bool
	ASCII       string
}

// Player represents the game player
type Player struct {
	CurrentRoom *Room
	Inventory   []*Item
}

// Game represents the main game state
type Game struct {
	Player *Player
	Rooms  map[string]*Room
}

// UI Helper functions
func clearScreen() {
	fmt.Print("\033[2J\033[H")
}

func printBanner() {
	banner := `
╔══════════════════════════════════════════════════════════════╗
║                                                              ║
║    ██████╗  ██████╗        ██████╗ ██╗   ██╗███████╗███████╗████████╗
║   ██╔════╝ ██╔═══██╗      ██╔═══██╗██║   ██║██╔════╝██╔════╝╚══██╔══╝
║   ██║  ███╗██║   ██║      ██║   ██║██║   ██║█████╗  ███████╗   ██║   
║   ██║   ██║██║   ██║      ██║   ██║╚██╗ ██╔╝██╔══╝  ╚════██║   ██║   
║   ╚██████╔╝╚██████╔╝      ╚██████╔╝ ╚████╔╝ ███████╗███████║   ██║   
║    ╚═════╝  ╚═════╝        ╚═════╝   ╚═══╝  ╚══════╝╚══════╝   ╚═╝   
║                                                              ║
║                    🏠 ROOM ESCAPE ADVENTURE 🏠                ║
║                                                              ║
╚══════════════════════════════════════════════════════════════╝`
	fmt.Printf("%s%s\n", ColorCyan, banner)
}

func printSeparator() {
	fmt.Printf("%s%s%s\n", ColorBlue, strings.Repeat("═", 70), ColorReset)
}

func printColored(text, color string) {
	fmt.Printf("%s%s%s", color, text, ColorReset)
}

func printSuccess(message string) {
	fmt.Printf("%s✅ %s%s\n", ColorGreen, message, ColorReset)
}

func printWarning(message string) {
	fmt.Printf("%s⚠️  %s%s\n", ColorYellow, message, ColorReset)
}

func printError(message string) {
	fmt.Printf("%s❌ %s%s\n", ColorRed, message, ColorReset)
}

func printInfo(message string) {
	fmt.Printf("%sℹ️  %s%s\n", ColorCyan, message, ColorReset)
}

func printASCII(ascii string) {
	fmt.Printf("%s%s%s\n", ColorPurple, ascii, ColorReset)
}

// NewGame creates a new game instance
func NewGame() *Game {
	// Create items
	key := &Item{
		Name:        "key",
		Description: "A rusty old key that might fit somewhere",
		Usable:      true,
		ASCII: `
    ╔══════╗
    ║  🔑  ║
    ╚══════╝`,
	}

	note := &Item{
		Name:        "note",
		Description: "A crumpled note with numbers: 1234",
		Usable:      false,
		ASCII: `
    ╔══════════╗
    ║  📄 NOTE ║
    ║    1234  ║
    ╚══════════╝`,
	}

	// Create rooms
	livingRoom := &Room{
		Name:        "Living Room",
		Description: "You are in a dimly lit living room. Dust covers the furniture. There's a locked door to the north and a window to the east.",
		Items:       []*Item{note},
		Exits:       make(map[string]*Room),
		Solved:      false,
		ASCII: `
    ┌─────────────────┐
    │  🪑    🪑       │
    │                 │
    │        🚪       │
    │                 │
    │  📺             │
    └─────────────────┘`,
	}

	kitchen := &Room{
		Name:        "Kitchen",
		Description: "A small kitchen with old appliances. There's a cabinet that might contain something useful.",
		Items:       []*Item{key},
		Exits:       make(map[string]*Room),
		Solved:      false,
		ASCII: `
    ┌─────────────────┐
    │  🍳    🥘       │
    │                 │
    │  🗄️             │
    │                 │
    │  🚰             │
    └─────────────────┘`,
	}

	// Set up room connections
	livingRoom.Exits["north"] = kitchen
	livingRoom.Exits["kitchen"] = kitchen
	kitchen.Exits["south"] = livingRoom
	kitchen.Exits["living room"] = livingRoom

	rooms := map[string]*Room{
		"living room": livingRoom,
		"kitchen":     kitchen,
	}

	player := &Player{
		CurrentRoom: livingRoom,
		Inventory:   []*Item{},
	}

	return &Game{
		Player: player,
		Rooms:  rooms,
	}
}

// Look displays the current room description and items
func (g *Game) Look() {
	clearScreen()

	// Print room ASCII art
	printASCII(g.Player.CurrentRoom.ASCII)

	printSeparator()
	printColored(fmt.Sprintf("📍 %s", g.Player.CurrentRoom.Name), ColorBold+ColorYellow)
	fmt.Println()
	printSeparator()

	printColored(g.Player.CurrentRoom.Description, ColorWhite)
	fmt.Println()

	if len(g.Player.CurrentRoom.Items) > 0 {
		fmt.Println()
		printColored("🔍 You can see:", ColorGreen)
		for _, item := range g.Player.CurrentRoom.Items {
			fmt.Printf("  • ")
			printColored(item.Name, ColorCyan)
			fmt.Printf(" - %s\n", item.Description)
			printASCII(item.ASCII)
		}
	}

	fmt.Println()
	printColored("🚪 Exits:", ColorBlue)
	for direction := range g.Player.CurrentRoom.Exits {
		fmt.Printf("  • ")
		printColored(direction, ColorPurple)
		fmt.Println()
	}
	printSeparator()
}

// Take adds an item to player's inventory
func (g *Game) Take(itemName string) {
	for i, item := range g.Player.CurrentRoom.Items {
		if strings.ToLower(item.Name) == strings.ToLower(itemName) {
			g.Player.Inventory = append(g.Player.Inventory, item)
			g.Player.CurrentRoom.Items = append(g.Player.CurrentRoom.Items[:i], g.Player.CurrentRoom.Items[i+1:]...)
			printSuccess(fmt.Sprintf("You take the %s.", item.Name))
			time.Sleep(1 * time.Second)
			g.Look()
			return
		}
	}
	printError(fmt.Sprintf("There's no %s here.", itemName))
}

// Inventory displays player's current inventory
func (g *Game) Inventory() {
	clearScreen()
	printColored("🎒 YOUR INVENTORY", ColorBold+ColorYellow)
	printSeparator()

	if len(g.Player.Inventory) == 0 {
		printWarning("Your inventory is empty.")
		return
	}

	for _, item := range g.Player.Inventory {
		fmt.Printf("  • ")
		printColored(item.Name, ColorCyan)
		fmt.Printf(" - %s\n", item.Description)
		printASCII(item.ASCII)
	}
	printSeparator()
}

// Move changes the player's current room
func (g *Game) Move(direction string) {
	if room, exists := g.Player.CurrentRoom.Exits[strings.ToLower(direction)]; exists {
		g.Player.CurrentRoom = room
		printInfo(fmt.Sprintf("You go %s...", direction))
		time.Sleep(1 * time.Second)
		g.Look()
	} else {
		printError(fmt.Sprintf("You can't go %s from here.", direction))
	}
}

// Use attempts to use an item
func (g *Game) Use(itemName string) {
	// Check if player has the item
	var item *Item
	for _, invItem := range g.Player.Inventory {
		if strings.ToLower(invItem.Name) == strings.ToLower(itemName) {
			item = invItem
			break
		}
	}

	if item == nil {
		printError(fmt.Sprintf("You don't have a %s.", itemName))
		return
	}

	// Simple puzzle: using the key in the living room
	if strings.ToLower(item.Name) == "key" && strings.ToLower(g.Player.CurrentRoom.Name) == "living room" {
		clearScreen()
		printInfo("You try the key in the locked door...")
		time.Sleep(2 * time.Second)
		clearScreen()
		printSuccess("🎉 SUCCESS! The door unlocks!")
		printSuccess("You have escaped the room! Congratulations!")
		time.Sleep(3 * time.Second)
		os.Exit(0)
	} else {
		printWarning(fmt.Sprintf("You can't use the %s here.", item.Name))
	}
}

// Help displays available commands
func (g *Game) Help() {
	clearScreen()
	printColored("❓ GAME HELP", ColorBold+ColorYellow)
	printSeparator()
	fmt.Println("Available commands:")
	fmt.Printf("  %s - Look around the current room\n", ColorCyan+"look/l"+ColorReset)
	fmt.Printf("  %s - Pick up an item\n", ColorCyan+"take <item>"+ColorReset)
	fmt.Printf("  %s - Check your inventory\n", ColorCyan+"inventory/i"+ColorReset)
	fmt.Printf("  %s - Use an item\n", ColorCyan+"use <item>"+ColorReset)
	fmt.Printf("  %s - Move in a direction\n", ColorCyan+"go <direction>"+ColorReset)
	fmt.Printf("  %s - Show this help\n", ColorCyan+"help/h"+ColorReset)
	fmt.Printf("  %s - Exit the game\n", ColorCyan+"quit/q"+ColorReset)
	printSeparator()
	printInfo("Press Enter to continue...")
	fmt.Scanln()
	g.Look()
}

// ProcessCommand handles user input
func (g *Game) ProcessCommand(input string) {
	parts := strings.Fields(strings.ToLower(input))
	if len(parts) == 0 {
		return
	}

	command := parts[0]

	switch command {
	case "look", "l":
		g.Look()
	case "take":
		if len(parts) > 1 {
			itemName := strings.Join(parts[1:], " ")
			g.Take(itemName)
		} else {
			fmt.Println("Take what?")
		}
	case "inventory", "i":
		g.Inventory()
	case "use":
		if len(parts) > 1 {
			itemName := strings.Join(parts[1:], " ")
			g.Use(itemName)
		} else {
			fmt.Println("Use what?")
		}
	case "go":
		if len(parts) > 1 {
			direction := strings.Join(parts[1:], " ")
			g.Move(direction)
		} else {
			fmt.Println("Go where?")
		}
	case "help", "h":
		g.Help()
	case "quit", "q":
		clearScreen()
		printSuccess("Thanks for playing! Goodbye!")
		time.Sleep(2 * time.Second)
		os.Exit(0)
	default:
		printError("I don't understand that command. Type 'help' for available commands.")
	}
}

func main() {
	clearScreen()
	printBanner()

	fmt.Println()
	printInfo("You find yourself trapped in a mysterious house...")
	printInfo("Type 'help' for commands or 'quit' to exit.")
	fmt.Println()
	printInfo("Press Enter to start your adventure...")
	fmt.Scanln()

	game := NewGame()
	scanner := bufio.NewScanner(os.Stdin)

	game.Look()

	for {
		fmt.Print("\n🎮 > ")
		if !scanner.Scan() {
			break
		}

		command := strings.TrimSpace(scanner.Text())
		if command != "" {
			game.ProcessCommand(command)
		}
	}
}
