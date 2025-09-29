package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
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

// QuestCategory represents different types of quests
type QuestCategory int

const (
	HackerQuest QuestCategory = iota
	EngineeringQuest
	AstronomicalQuest
	BiologicalQuest
	PhysicalQuest
)

// Quest represents a quest in the game
type Quest struct {
	ID           int
	Name         string
	Description  string
	Category     QuestCategory
	Difficulty   int // 1-5
	Solved       bool
	TimeLimit    time.Duration
	Reward       string
	Requirements []string
	Solution     string
	ASCII        string
	Hints        []string // Подсказки для квеста
	Example      string   // Пример решения
}

// PlayerStats represents player characteristics
type PlayerStats struct {
	Hacking     int // 0-100
	Engineering int // 0-100
	Astronomy   int // 0-100
	Biology     int // 0-100
	Physics     int // 0-100
	Energy      int // 0-100
	TimeLeft    time.Duration
}

// Item represents an object in the game
type Item struct {
	Name        string
	Description string
	Usable      bool
	ASCII       string
	QuestID     int // Associated quest ID
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
	Stats       *PlayerStats
	Quests      []*Quest
	Completed   int
}

// Game represents the main game state
type Game struct {
	Player    *Player
	Rooms     map[string]*Room
	AllQuests []*Quest
	GameStart time.Time
	GameMode  string // "tutorial", "normal", "hardcore"
}

// UI Helper functions
func clearScreen() {
	fmt.Print("\033[2J\033[H")
}

func printBanner() {
	banner := `
╔══════════════════════════════════════════════════════════════╗
║                                                              ║
║  ██████╗ ██████╗ ███████╗███╗   ███╗██╗   ██╗██████╗ ██████╗ ║
║ ██╔════╝██╔═══██╗██╔════╝████╗ ████║██║   ██║██╔══██╗██╔══██╗║
║ ██║     ██║   ██║███████╗██╔████╔██║██║   ██║██████╔╝██████╔╝║
║ ██║     ██║   ██║╚════██║██║╚██╔╝██║██║   ██║██╔═══╝ ██╔══██╗║
║ ╚██████╗╚██████╔╝███████║██║ ╚═╝ ██║╚██████╔╝██║     ██║  ██║║
║  ╚═════╝ ╚═════╝ ╚══════╝╚═╝     ╚═╝ ╚═════╝ ╚═╝     ╚═╝  ╚═╝║
║                                                              ║
║              🌌 COSMIC CYBERPUNK ROOM ESCAPE 🌌             ║
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

// Quest creation functions
func createAllQuests() []*Quest {
	quests := []*Quest{
		// Хакерские и кибернетические задачи (1-20)
		{1, "Взлом голограммы", "Расшифровать двоичный код, проецируемый голографическим интерфейсом", HackerQuest, 2, false, 5 * time.Minute, "Кибер-ключ", []string{"Голографический терминал"}, "01001000 01100001 01100011 01101011", `
    ╔══════════════════════════════╗
    ║  🔮 HOLOGRAM INTERFACE 🔮    ║
    ║  01001000 01100001 01100011  ║
    ║  01101011 01100101 01110010  ║
    ╚══════════════════════════════╝`,
			[]string{
				"💡 Подсказка 1: Это двоичный код. Каждая группа из 8 цифр представляет одну букву.",
				"💡 Подсказка 2: Переведите двоичный код в текст. 01001000 = H, 01100001 = a, 01100011 = c, 01101011 = k",
				"💡 Подсказка 3: Слово 'Hacker' на английском языке.",
			}, "Пример: 01001000 = H, 01100001 = a → 'Ha'"},

		{2, "Нейроинтерфейс", "Подключиться к мозговому чипу и решить математическую последовательность", HackerQuest, 3, false, 7 * time.Minute, "Нейро-имплант", []string{"Нейро-шлем"}, "2, 4, 8, 16, 32, 64", `
    ╔══════════════════════════════╗
    ║  🧠 NEURO INTERFACE 🧠        ║
    ║  [2] [4] [8] [16] [32] [64]  ║
    ║  Find the next number...     ║
    ╚══════════════════════════════╝`,
			[]string{
				"💡 Подсказка 1: Каждое число в 2 раза больше предыдущего.",
				"💡 Подсказка 2: 2×2=4, 4×2=8, 8×2=16, 16×2=32, 32×2=64",
				"💡 Подсказка 3: Следующее число: 64×2 = ?",
			}, "Пример: 2 → 4 → 8 → 16 → 32 → 64 → 128"},

		{3, "Квантовый пароль", "Одновременно активировать несколько терминалов в правильной последовательности", HackerQuest, 4, false, 10 * time.Minute, "Квантовый ключ", []string{"Терминал 1", "Терминал 2", "Терминал 3"}, "1-3-2-1-3", `
    ╔══════════════════════════════╗
    ║  ⚛️ QUANTUM TERMINALS ⚛️      ║
    ║  [1] [2] [3]                 ║
    ║  Activate in sequence...     ║
    ╚══════════════════════════════╝`,
			[]string{
				"💡 Подсказка 1: Последовательность: 1, затем 3, затем 2, затем 1, затем 3.",
				"💡 Подсказка 2: Начните с терминала 1, затем перейдите к терминалу 3.",
				"💡 Подсказка 3: Полная последовательность: 1-3-2-1-3",
			}, "Пример: Активируйте терминалы в порядке 1, 3, 2, 1, 3"},

		// Инженерные и технические головоломки (21-40)
		{21, "Энергетические узлы", "Перенаправить поток энергии через сложную схему", EngineeringQuest, 3, false, 8 * time.Minute, "Энерго-модуль", []string{"Энергетическая сеть"}, "A→B→C→D→E", `
    ╔══════════════════════════════╗
    ║  ⚡ ENERGY GRID ⚡            ║
    ║  A ── B ── C ── D ── E       ║
    ║  │    │    │    │    │       ║
    ║  F ── G ── H ── I ── J       ║
    ╚══════════════════════════════╝`,
			[]string{
				"💡 Подсказка 1: Следуйте по верхней линии: A → B → C → D → E",
				"💡 Подсказка 2: Не переходите на нижнюю линию (F-G-H-I-J)",
				"💡 Подсказка 3: Простое линейное соединение: A→B→C→D→E",
			}, "Пример: Начните с A, затем B, затем C, затем D, затем E"},

		{22, "Гравитационный генератор", "Настроить искусственную гравитацию в нужных зонах", EngineeringQuest, 4, false, 12 * time.Minute, "Грави-контроллер", []string{"Генератор гравитации"}, "Zone1: 0.5g, Zone2: 1.0g, Zone3: 1.5g", `
    ╔══════════════════════════════╗
    ║  🌍 GRAVITY GENERATOR 🌍     ║
    ║  Zone 1: [0.5g]              ║
    ║  Zone 2: [1.0g]              ║
    ║  Zone 3: [1.5g]              ║
    ╚══════════════════════════════╝`,
			[]string{
				"💡 Подсказка 1: Установите гравитацию: Зона 1 = 0.5g, Зона 2 = 1.0g, Зона 3 = 1.5g",
				"💡 Подсказка 2: Формат ответа: 'Zone1: 0.5g, Zone2: 1.0g, Zone3: 1.5g'",
				"💡 Подсказка 3: Постепенное увеличение гравитации от 0.5 до 1.5",
			}, "Пример: Zone1: 0.5g, Zone2: 1.0g, Zone3: 1.5g"},

		// Астрономические и космические загадки (41-60)
		{41, "Звездная карта", "Найти правильное созвездие для навигации", AstronomicalQuest, 2, false, 6 * time.Minute, "Навигационный чип", []string{"Звездная карта"}, "Орион", `
    ╔══════════════════════════════╗
    ║  ⭐ STAR MAP ⭐              ║
    ║  • • • • • • • • • • • • •   ║
    ║  • • • • • • • • • • • • •   ║
    ║  • • • • • • • • • • • • •   ║
    ╚══════════════════════════════╝`,
			[]string{
				"💡 Подсказка 1: Это одно из самых известных созвездий в северном полушарии.",
				"💡 Подсказка 2: Созвездие с тремя яркими звездами в ряд (пояс охотника).",
				"💡 Подсказка 3: Название созвездия: 'Орион' (Orion)",
			}, "Пример: Орион - одно из самых узнаваемых созвездий"},

		{42, "Планетарное выравнивание", "Дождаться, когда планеты займут нужные позиции", AstronomicalQuest, 3, false, 15 * time.Minute, "Планетарный сканер", []string{"Планетарный симулятор"}, "Mercury-Venus-Earth-Mars", `
    ╔══════════════════════════════╗
    ║  🪐 PLANETARY ALIGNMENT 🪐   ║
    ║  ☿️  ♀️  🌍  ♂️  ♃️  ♄️  ║
    ║  Wait for alignment...       ║
    ╚══════════════════════════════╝`,
			[]string{
				"💡 Подсказка 1: Планеты должны выстроиться в порядке от Солнца: Меркурий, Венера, Земля, Марс",
				"💡 Подсказка 2: Формат ответа: 'Mercury-Venus-Earth-Mars'",
				"💡 Подсказка 3: Четыре ближайшие к Солнцу планеты в правильном порядке",
			}, "Пример: Mercury-Venus-Earth-Mars - планеты в порядке от Солнца"},

		// Биологические и медицинские задачи (61-80)
		{61, "Генетический замок", "Модифицировать ДНК для доступа к биосейфу", BiologicalQuest, 4, false, 10 * time.Minute, "Генетический ключ", []string{"ДНК-анализатор"}, "ATCGATCGATCG", `
    ╔══════════════════════════════╗
    ║  🧬 DNA LOCK 🧬              ║
    ║  A T C G A T C G A T C G     ║
    ║  Modify sequence...          ║
    ╚══════════════════════════════╝`,
			[]string{
				"💡 Подсказка 1: Повторите последовательность A-T-C-G три раза подряд",
				"💡 Подсказка 2: ATCG + ATCG + ATCG = ATCGATCGATCG",
				"💡 Подсказка 3: Полная последовательность: ATCGATCGATCG",
			}, "Пример: ATCGATCGATCG - повторение базовой последовательности"},

		{62, "Синтетические органы", "Подключить искусственные органы к пациенту", BiologicalQuest, 5, false, 15 * time.Minute, "Био-имплант", []string{"Синтетическое сердце", "Нейронные связи"}, "Heart→Brain→Lungs→Liver", `
    ╔══════════════════════════════╗
    ║  🫀 SYNTHETIC ORGANS 🫀      ║
    ║  Heart → Brain → Lungs       ║
    ║  Connect in sequence...      ║
    ╚══════════════════════════════╝`,
			[]string{
				"💡 Подсказка 1: Подключите органы в порядке: Сердце → Мозг → Легкие → Печень",
				"💡 Подсказка 2: Формат ответа: 'Heart→Brain→Lungs→Liver'",
				"💡 Подсказка 3: Начните с сердца, затем мозг, затем легкие, затем печень",
			}, "Пример: Heart→Brain→Lungs→Liver - последовательность подключения органов"},

		// Физические и механические головоломки (81-100)
		{81, "Левитирующие платформы", "Управлять парящими в воздухе поверхностями", PhysicalQuest, 3, false, 8 * time.Minute, "Антиграви-модуль", []string{"Платформа-контроллер"}, "Platform1→Platform2→Platform3", `
    ╔══════════════════════════════╗
    ║  🏗️ FLOATING PLATFORMS 🏗️   ║
    ║  [1]     [2]     [3]         ║
    ║  ╱╲     ╱╲     ╱╲            ║
    ║  Navigate sequence...        ║
    ╚══════════════════════════════╝`,
			[]string{
				"💡 Подсказка 1: Перемещайтесь по платформам в порядке: 1 → 2 → 3",
				"💡 Подсказка 2: Формат ответа: 'Platform1→Platform2→Platform3'",
				"💡 Подсказка 3: Простая последовательность: Платформа 1, затем 2, затем 3",
			}, "Пример: Platform1→Platform2→Platform3 - последовательность навигации"},

		{82, "Голографические стены", "Отличить настоящие препятствия от иллюзий", PhysicalQuest, 4, false, 12 * time.Minute, "Голо-детектор", []string{"Голографический проектор"}, "Real: 1,3,5,7,9", `
    ╔══════════════════════════════╗
    ║  🎭 HOLOGRAPHIC WALLS 🎭     ║
    ║  [1][2][3][4][5][6][7][8][9] ║
    ║  Find the real ones...       ║
    ╚══════════════════════════════╝`,
			[]string{
				"💡 Подсказка 1: Настоящие стены имеют нечетные номера: 1, 3, 5, 7, 9",
				"💡 Подсказка 2: Формат ответа: 'Real: 1,3,5,7,9'",
				"💡 Подсказка 3: Все нечетные числа от 1 до 9 являются настоящими стенами",
			}, "Пример: Real: 1,3,5,7,9 - нечетные номера стен"},
	}

	return quests
}

func getCategoryName(category QuestCategory) string {
	switch category {
	case HackerQuest:
		return "Хакерские"
	case EngineeringQuest:
		return "Инженерные"
	case AstronomicalQuest:
		return "Астрономические"
	case BiologicalQuest:
		return "Биологические"
	case PhysicalQuest:
		return "Физические"
	default:
		return "Неизвестные"
	}
}

func getCategoryEmoji(category QuestCategory) string {
	switch category {
	case HackerQuest:
		return "💻"
	case EngineeringQuest:
		return "⚙️"
	case AstronomicalQuest:
		return "⭐"
	case BiologicalQuest:
		return "🧬"
	case PhysicalQuest:
		return "⚡"
	default:
		return "❓"
	}
}

// NewGame creates a new game instance
func NewGame() *Game {
	// Create all quests
	allQuests := createAllQuests()

	// Create items
	key := &Item{
		Name:        "key",
		Description: "A rusty old key that might fit somewhere",
		Usable:      true,
		ASCII: `
    ╔══════╗
    ║  🔑  ║
    ╚══════╝`,
		QuestID: 0,
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
		QuestID: 0,
	}

	// Create cyberpunk rooms
	cyberRoom := &Room{
		Name:        "Cyber Control Room",
		Description: "A high-tech control room filled with holographic displays and quantum computers. The air hums with energy.",
		Items:       []*Item{note},
		Exits:       make(map[string]*Room),
		Solved:      false,
		ASCII: `
    ┌─────────────────────────────────┐
    │  💻    🔮    🧠    ⚛️         │
    │                                 │
    │        🚪        🚪             │
    │                                 │
    │  📡    📊    📈    📉          │
    └─────────────────────────────────┘`,
	}

	engineeringBay := &Room{
		Name:        "Engineering Bay",
		Description: "A massive engineering facility with gravity generators, energy nodes, and plasma resonators.",
		Items:       []*Item{key},
		Exits:       make(map[string]*Room),
		Solved:      false,
		ASCII: `
    ┌─────────────────────────────────┐
    │  ⚡    🌍    ⚙️    🔧           │
    │                                 │
    │        🚪        🚪             │
    │                                 │
    │  ⚙️    🔩    ⚡    🌍          │
    └─────────────────────────────────┘`,
	}

	observatory := &Room{
		Name:        "Space Observatory",
		Description: "A domed observatory with star maps, planetary simulators, and cosmic navigation equipment.",
		Items:       []*Item{},
		Exits:       make(map[string]*Room),
		Solved:      false,
		ASCII: `
    ┌─────────────────────────────────┐
    │  ⭐    🪐    🌟    🌌           │
    │                                 │
    │        🚪        🚪             │
    │                                 │
    │  🔭    📡    🛰️    🚀          │
    └─────────────────────────────────┘`,
	}

	// Set up room connections
	cyberRoom.Exits["north"] = engineeringBay
	cyberRoom.Exits["east"] = observatory
	engineeringBay.Exits["south"] = cyberRoom
	engineeringBay.Exits["east"] = observatory
	observatory.Exits["west"] = cyberRoom
	observatory.Exits["south"] = engineeringBay

	rooms := map[string]*Room{
		"cyber control room": cyberRoom,
		"engineering bay":    engineeringBay,
		"observatory":        observatory,
	}

	// Create player with stats
	playerStats := &PlayerStats{
		Hacking:     50,
		Engineering: 50,
		Astronomy:   50,
		Biology:     50,
		Physics:     50,
		Energy:      100,
		TimeLeft:    60 * time.Minute, // 1 hour game time
	}

	// Select random quests for the player
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	playerQuests := make([]*Quest, 0)
	questIndices := r.Perm(len(allQuests))

	// Select 5 random quests
	for i := 0; i < 5 && i < len(questIndices); i++ {
		quest := allQuests[questIndices[i]]
		playerQuests = append(playerQuests, quest)
	}

	player := &Player{
		CurrentRoom: cyberRoom,
		Inventory:   []*Item{},
		Stats:       playerStats,
		Quests:      playerQuests,
		Completed:   0,
	}

	return &Game{
		Player:    player,
		Rooms:     rooms,
		AllQuests: allQuests,
		GameStart: time.Now(),
		GameMode:  "normal",
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

	// Show player stats
	fmt.Println()
	printColored("📊 YOUR STATS:", ColorBold+ColorCyan)
	fmt.Printf("  💻 Hacking: %d/100    ⚙️ Engineering: %d/100\n", g.Player.Stats.Hacking, g.Player.Stats.Engineering)
	fmt.Printf("  ⭐ Astronomy: %d/100  🧬 Biology: %d/100\n", g.Player.Stats.Astronomy, g.Player.Stats.Biology)
	fmt.Printf("  ⚡ Physics: %d/100    🔋 Energy: %d/100\n", g.Player.Stats.Physics, g.Player.Stats.Energy)
	fmt.Printf("  ⏰ Time Left: %s\n", g.Player.Stats.TimeLeft.Round(time.Second))

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
		if strings.EqualFold(item.Name, itemName) {
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
		if strings.EqualFold(invItem.Name, itemName) {
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

// Quest methods
func (g *Game) ShowQuests() {
	clearScreen()
	printColored("🎯 YOUR QUESTS", ColorBold+ColorYellow)
	printSeparator()

	if len(g.Player.Quests) == 0 {
		printWarning("No active quests!")
		printInfo("Debug: Player has 0 quests")
		return
	}

	printInfo(fmt.Sprintf("Debug: Player has %d quests", len(g.Player.Quests)))

	availableQuests := 0
	for i, quest := range g.Player.Quests {
		status := "❌"
		if quest.Solved {
			status = "✅"
		} else {
			availableQuests++
		}

		emoji := getCategoryEmoji(quest.Category)
		categoryName := getCategoryName(quest.Category)

		fmt.Printf("%s %s %s (ID: %d) - %s\n", status, emoji, quest.Name, quest.ID, categoryName)
		fmt.Printf("   Difficulty: %d/5 ⭐\n", quest.Difficulty)
		fmt.Printf("   Time Limit: %s\n", quest.TimeLimit.Round(time.Second))
		fmt.Printf("   Reward: %s\n", quest.Reward)
		fmt.Printf("   Description: %s\n", quest.Description)

		if !quest.Solved {
			printASCII(quest.ASCII)
		}

		if i < len(g.Player.Quests)-1 {
			fmt.Println()
		}
	}

	if availableQuests > 0 {
		printInfo(fmt.Sprintf("Available quest IDs: "))
		for _, quest := range g.Player.Quests {
			if !quest.Solved {
				fmt.Printf("%d ", quest.ID)
			}
		}
		fmt.Println()
	}

	printSeparator()
	printInfo("Press Enter to continue...")
	fmt.Scanln()
	g.Look()
}

func (g *Game) StartQuest(questID int) {
	clearScreen()

	printInfo(fmt.Sprintf("Debug: Looking for quest ID %d", questID))
	printInfo(fmt.Sprintf("Debug: Player has %d quests", len(g.Player.Quests)))

	var quest *Quest
	for i, q := range g.Player.Quests {
		printInfo(fmt.Sprintf("Debug: Quest %d: ID=%d, Name=%s, Solved=%t", i, q.ID, q.Name, q.Solved))
		if q.ID == questID && !q.Solved {
			quest = q
			break
		}
	}

	if quest == nil {
		printError("Quest not found or already completed!")
		printInfo("Available quest IDs:")
		availableCount := 0
		for _, q := range g.Player.Quests {
			if !q.Solved {
				fmt.Printf("  - %d: %s\n", q.ID, q.Name)
				availableCount++
			}
		}
		if availableCount == 0 {
			printWarning("No available quests! All quests are completed.")
		}
		time.Sleep(3 * time.Second)
		g.Look()
		return
	}

	printColored(fmt.Sprintf("🎯 STARTING QUEST: %s", quest.Name), ColorBold+ColorYellow)
	printSeparator()

	emoji := getCategoryEmoji(quest.Category)
	categoryName := getCategoryName(quest.Category)

	fmt.Printf("%s Category: %s\n", emoji, categoryName)
	fmt.Printf("⭐ Difficulty: %d/5\n", quest.Difficulty)
	fmt.Printf("⏰ Time Limit: %s\n", quest.TimeLimit.Round(time.Second))
	fmt.Printf("🎁 Reward: %s\n", quest.Reward)
	fmt.Println()
	printColored("Description:", ColorCyan)
	fmt.Println(quest.Description)
	fmt.Println()

	printASCII(quest.ASCII)

	fmt.Println()
	printColored("Enter your solution:", ColorGreen)
	fmt.Print("> ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	solution := strings.TrimSpace(scanner.Text())

	if strings.EqualFold(solution, quest.Solution) {
		quest.Solved = true
		g.Player.Completed++

		// Award experience based on category
		switch quest.Category {
		case HackerQuest:
			g.Player.Stats.Hacking += quest.Difficulty * 5
		case EngineeringQuest:
			g.Player.Stats.Engineering += quest.Difficulty * 5
		case AstronomicalQuest:
			g.Player.Stats.Astronomy += quest.Difficulty * 5
		case BiologicalQuest:
			g.Player.Stats.Biology += quest.Difficulty * 5
		case PhysicalQuest:
			g.Player.Stats.Physics += quest.Difficulty * 5
		}

		// Add reward to inventory
		rewardItem := &Item{
			Name:        quest.Reward,
			Description: fmt.Sprintf("Reward for completing: %s", quest.Name),
			Usable:      true,
			ASCII:       fmt.Sprintf("    🎁 %s", quest.Reward),
			QuestID:     quest.ID,
		}
		g.Player.Inventory = append(g.Player.Inventory, rewardItem)

		printSuccess(fmt.Sprintf("🎉 QUEST COMPLETED! You earned: %s", quest.Reward))
		printSuccess(fmt.Sprintf("Experience gained in %s!", categoryName))

		// Check if all quests completed
		allCompleted := true
		for _, q := range g.Player.Quests {
			if !q.Solved {
				allCompleted = false
				break
			}
		}

		if allCompleted {
			printSuccess("🏆 CONGRATULATIONS! You completed all quests!")
			printSuccess("You have successfully escaped the Cosmic Cyberpunk Room!")
			time.Sleep(5 * time.Second)
			os.Exit(0)
		}
	} else {
		printError("❌ Incorrect solution! Try again.")
		g.Player.Stats.Energy -= 10 // Lose energy for wrong answer
		if g.Player.Stats.Energy < 0 {
			g.Player.Stats.Energy = 0
		}
	}

	time.Sleep(3 * time.Second)
	g.Look()
}

func (g *Game) ShowStats() {
	clearScreen()
	printColored("📊 DETAILED STATS", ColorBold+ColorYellow)
	printSeparator()

	fmt.Printf("💻 Hacking: %d/100\n", g.Player.Stats.Hacking)
	fmt.Printf("⚙️ Engineering: %d/100\n", g.Player.Stats.Engineering)
	fmt.Printf("⭐ Astronomy: %d/100\n", g.Player.Stats.Astronomy)
	fmt.Printf("🧬 Biology: %d/100\n", g.Player.Stats.Biology)
	fmt.Printf("⚡ Physics: %d/100\n", g.Player.Stats.Physics)
	fmt.Printf("🔋 Energy: %d/100\n", g.Player.Stats.Energy)
	fmt.Printf("⏰ Time Left: %s\n", g.Player.Stats.TimeLeft.Round(time.Second))
	fmt.Printf("✅ Quests Completed: %d/%d\n", g.Player.Completed, len(g.Player.Quests))

	printSeparator()
	printInfo("Press Enter to continue...")
	fmt.Scanln()
	g.Look()
}

func (g *Game) ShowHints(questID int) {
	clearScreen()

	var quest *Quest
	for _, q := range g.Player.Quests {
		if q.ID == questID {
			quest = q
			break
		}
	}

	if quest == nil {
		printError("Quest not found!")
		time.Sleep(2 * time.Second)
		g.Look()
		return
	}

	printColored(fmt.Sprintf("💡 HINTS FOR: %s", quest.Name), ColorBold+ColorYellow)
	printSeparator()

	emoji := getCategoryEmoji(quest.Category)
	categoryName := getCategoryName(quest.Category)

	fmt.Printf("%s Category: %s\n", emoji, categoryName)
	fmt.Printf("⭐ Difficulty: %d/5\n", quest.Difficulty)
	fmt.Println()

	printColored("Description:", ColorCyan)
	fmt.Println(quest.Description)
	fmt.Println()

	printASCII(quest.ASCII)
	fmt.Println()

	printColored("💡 HINTS:", ColorGreen)
	for i, hint := range quest.Hints {
		fmt.Printf("%s\n", hint)
		if i < len(quest.Hints)-1 {
			fmt.Println()
		}
	}

	fmt.Println()
	printColored("Example:", ColorCyan)
	fmt.Println(quest.Example)

	printSeparator()
	printInfo("Press Enter to continue...")
	fmt.Scanln()
	g.Look()
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
	fmt.Printf("  %s - Show your quests\n", ColorCyan+"quests/q"+ColorReset)
	fmt.Printf("  %s - Start a quest\n", ColorCyan+"start <quest_id>"+ColorReset)
	fmt.Printf("  %s - Show hints for a quest\n", ColorCyan+"hints <quest_id>"+ColorReset)
	fmt.Printf("  %s - Show detailed stats\n", ColorCyan+"stats/s"+ColorReset)
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
	case "quests", "q":
		g.ShowQuests()
	case "start":
		if len(parts) > 1 {
			if questID, err := strconv.Atoi(parts[1]); err == nil {
				g.StartQuest(questID)
			} else {
				printError("Invalid quest ID. Use a number.")
			}
		} else {
			fmt.Println("Start which quest? Use quest ID number.")
		}
	case "hints":
		if len(parts) > 1 {
			if questID, err := strconv.Atoi(parts[1]); err == nil {
				g.ShowHints(questID)
			} else {
				printError("Invalid quest ID. Use a number.")
			}
		} else {
			fmt.Println("Show hints for which quest? Use quest ID number.")
		}
	case "stats", "s":
		g.ShowStats()
	case "help", "h":
		g.Help()
	case "quit", "exit":
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
	printInfo("🌌 Welcome to the Cosmic Cyberpunk Room Escape!")
	printInfo("You are trapped in a high-tech facility filled with quantum puzzles and cybernetic challenges.")
	printInfo("Complete quests to gain experience and escape!")
	fmt.Println()
	printInfo("Type 'help' for commands or 'quit' to exit.")
	fmt.Println()
	printInfo("Press Enter to start your cyberpunk adventure...")
	fmt.Scanln()

	game := NewGame()
	scanner := bufio.NewScanner(os.Stdin)

	game.Look()

	for {
		// Check if time is up
		if game.Player.Stats.TimeLeft <= 0 {
			clearScreen()
			printError("⏰ TIME'S UP! You failed to escape in time!")
			printError("The facility's security systems have locked you in permanently!")
			time.Sleep(3 * time.Second)
			os.Exit(0)
		}

		// Check if energy is depleted
		if game.Player.Stats.Energy <= 0 {
			clearScreen()
			printError("🔋 ENERGY DEPLETED! You collapsed from exhaustion!")
			printError("You need to rest to regain energy!")
			time.Sleep(3 * time.Second)
			os.Exit(0)
		}

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
