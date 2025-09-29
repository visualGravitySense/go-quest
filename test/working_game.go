package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Working quest structure
type WorkingQuest struct {
	ID          int
	Name        string
	Description string
	Solution    string
	Solved      bool
	Hints       []string
}

func main() {
	fmt.Println("🌌 Cosmic Cyberpunk Room Escape - Working Version")
	fmt.Println("================================================")

	// Create working quests with simple IDs
	quests := []*WorkingQuest{
		{
			ID:          1,
			Name:        "Binary Hack",
			Description: "Decode this binary code: 01001000 01100001 01100011 01101011",
			Solution:    "Hacker",
			Solved:      false,
			Hints: []string{
				"💡 Hint 1: This is binary code. Each group of 8 digits represents one letter.",
				"💡 Hint 2: 01001000 = H, 01100001 = a, 01100011 = c, 01101011 = k",
				"💡 Hint 3: The word is 'Hacker' in English.",
			},
		},
		{
			ID:          2,
			Name:        "Math Sequence",
			Description: "What's the next number: 2, 4, 8, 16, 32, 64",
			Solution:    "128",
			Solved:      false,
			Hints: []string{
				"💡 Hint 1: Each number is 2 times the previous one.",
				"💡 Hint 2: 2×2=4, 4×2=8, 8×2=16, 16×2=32, 32×2=64",
				"💡 Hint 3: Next number: 64×2 = 128",
			},
		},
		{
			ID:          3,
			Name:        "Terminal Sequence",
			Description: "Activate terminals in the correct order: 1, 3, 2, 1, 3",
			Solution:    "1-3-2-1-3",
			Solved:      false,
			Hints: []string{
				"💡 Hint 1: Activate terminal 1, then 3, then 2, then 1, then 3.",
				"💡 Hint 2: Start with terminal 1, then go to terminal 3.",
				"💡 Hint 3: Full sequence: 1-3-2-1-3",
			},
		},
	}

	fmt.Println("\n🎯 Available Quests:")
	for _, quest := range quests {
		status := "❌"
		if quest.Solved {
			status = "✅"
		}
		fmt.Printf("%s %d. %s\n", status, quest.ID, quest.Name)
	}

	fmt.Println("\nCommands: quests, start <id>, hints <id>, quit")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\n🎮 > ")
		if !scanner.Scan() {
			break
		}

		command := strings.TrimSpace(scanner.Text())
		parts := strings.Fields(strings.ToLower(command))

		if len(parts) == 0 {
			continue
		}

		switch parts[0] {
		case "quests":
			fmt.Println("\n🎯 Available Quests:")
			for _, quest := range quests {
				status := "❌"
				if quest.Solved {
					status = "✅"
				}
				fmt.Printf("%s %d. %s\n", status, quest.ID, quest.Name)
			}
			fmt.Println("\nAvailable quest IDs: 1 2 3")

		case "start":
			if len(parts) > 1 {
				if questID, err := strconv.Atoi(parts[1]); err == nil {
					startQuest(questID, quests, scanner)
				} else {
					fmt.Println("Invalid quest ID. Use: start 1, start 2, or start 3")
				}
			} else {
				fmt.Println("Start which quest? Use: start 1, start 2, or start 3")
			}

		case "hints":
			if len(parts) > 1 {
				if questID, err := strconv.Atoi(parts[1]); err == nil {
					showHints(questID, quests)
				} else {
					fmt.Println("Invalid quest ID. Use: hints 1, hints 2, or hints 3")
				}
			} else {
				fmt.Println("Show hints for which quest? Use: hints 1, hints 2, or hints 3")
			}

		case "quit", "exit":
			fmt.Println("Thanks for playing!")
			return

		default:
			fmt.Println("Commands: quests, start <id>, hints <id>, quit")
		}
	}
}

func startQuest(questID int, quests []*WorkingQuest, scanner *bufio.Scanner) {
	var quest *WorkingQuest
	for _, q := range quests {
		if q.ID == questID && !q.Solved {
			quest = q
			break
		}
	}

	if quest == nil {
		fmt.Println("❌ Quest not found or already completed!")
		fmt.Println("Available quest IDs: 1 2 3")
		return
	}

	fmt.Printf("\n🎯 Starting Quest: %s\n", quest.Name)
	fmt.Printf("Description: %s\n", quest.Description)
	fmt.Print("Enter your solution: ")

	scanner.Scan()
	solution := strings.TrimSpace(scanner.Text())

	if strings.EqualFold(solution, quest.Solution) {
		quest.Solved = true
		fmt.Println("🎉 Correct! Quest completed!")
	} else {
		fmt.Printf("❌ Incorrect. The answer was: %s\n", quest.Solution)
	}
}

func showHints(questID int, quests []*WorkingQuest) {
	var quest *WorkingQuest
	for _, q := range quests {
		if q.ID == questID {
			quest = q
			break
		}
	}

	if quest == nil {
		fmt.Println("❌ Quest not found!")
		fmt.Println("Available quest IDs: 1 2 3")
		return
	}

	fmt.Printf("\n💡 Hints for: %s\n", quest.Name)
	fmt.Printf("Description: %s\n", quest.Description)
	fmt.Println("\nHints:")
	for _, hint := range quest.Hints {
		fmt.Println(hint)
	}
}
