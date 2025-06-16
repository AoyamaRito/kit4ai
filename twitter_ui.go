package main

import (
	"fmt"
	"os"
	"kit4ai/pkg/canvas"
)

func generateTwitterUI() string {
	// Create Twitter-style UI
	twitterCanvas := canvas.NewByteCanvas()
	
	currentWidth := canvas.GetCurrentWidth()
	maxX := currentWidth - 1
	
	// Header bar with Twitter branding
	twitterCanvas.DrawBox(0, 0, maxX, 2)
	twitterCanvas.WriteBytesASCII(2, 1, "X (Twitter)")
	if currentWidth >= 80 {
		twitterCanvas.WriteBytesASCII(currentWidth-18, 1, "2024-06-15 14:32")
	}
	
	// Navigation bar
	twitterCanvas.DrawBox(0, 3, maxX, 5)
	if currentWidth >= 80 {
		twitterCanvas.WriteBytesASCII(2, 4, "[Home] [Search] [Notifications] [Messages] [Bookmarks] [Profile]")
		twitterCanvas.WriteBytesASCII(currentWidth-20, 4, "[Settings] [@username]")
	} else {
		twitterCanvas.WriteBytesASCII(2, 4, "[Home] [Search] [Notifications] [Messages]")
	}
	
	// Tweet compose box
	twitterCanvas.DrawBox(0, 6, maxX, 11)
	twitterCanvas.WriteBytesASCII(2, 7, "What's happening?")
	twitterCanvas.WriteBytesASCII(2, 9, "")
	twitterCanvas.WriteBytesASCII(2, 10, "[Photo] [Poll] [Emoji] [Location]")
	if currentWidth >= 70 {
		twitterCanvas.WriteBytesASCII(currentWidth-15, 10, "[Tweet]")
	}
	
	// Tweet feed
	twitterCanvas.DrawBox(0, 12, maxX, 40)
	
	// Tweet 1
	twitterCanvas.WriteBytesASCII(2, 13, "@user1")
	if currentWidth >= 60 {
		twitterCanvas.WriteBytesASCII(currentWidth-10, 13, "2h")
	}
	twitterCanvas.WriteBytesASCII(2, 14, "This is a sample tweet with some interesting content...")
	twitterCanvas.WriteBytesASCII(2, 15, "[Reply 12] [Retweet 5] [Like 23] [Share]")
	
	// Separator
	twitterCanvas.DrawBox(0, 16, maxX, 16)
	
	// Tweet 2
	twitterCanvas.WriteBytesASCII(2, 17, "@user2")
	if currentWidth >= 60 {
		twitterCanvas.WriteBytesASCII(currentWidth-10, 17, "4h")
	}
	twitterCanvas.WriteBytesASCII(2, 18, "Another tweet with a longer message that spans multiple")
	twitterCanvas.WriteBytesASCII(2, 19, "lines and includes some hashtags #trending #news")
	twitterCanvas.WriteBytesASCII(2, 20, "[Reply 45] [Retweet 12] [Like 89] [Share]")
	
	// Separator
	twitterCanvas.DrawBox(0, 21, maxX, 21)
	
	// Tweet 3 (Retweet)
	twitterCanvas.WriteBytesASCII(2, 22, "@user3")
	if currentWidth >= 60 {
		twitterCanvas.WriteBytesASCII(currentWidth-10, 22, "6h")
	}
	twitterCanvas.WriteBytesASCII(2, 23, "RT @someone: Great article about technology trends")
	twitterCanvas.WriteBytesASCII(2, 24, "[Reply 8] [Retweet 15] [Like 34] [Share]")
	
	// Separator
	twitterCanvas.DrawBox(0, 25, maxX, 25)
	
	// Tweet 4
	twitterCanvas.WriteBytesASCII(2, 26, "@user4")
	if currentWidth >= 60 {
		twitterCanvas.WriteBytesASCII(currentWidth-10, 26, "8h")
	}
	twitterCanvas.WriteBytesASCII(2, 27, "Just had the best coffee! #mondaymorning #coffee")
	twitterCanvas.WriteBytesASCII(2, 28, "[Reply 3] [Retweet 2] [Like 15] [Share]")
	
	// Separator
	twitterCanvas.DrawBox(0, 29, maxX, 29)
	
	// Tweet 5
	twitterCanvas.WriteBytesASCII(2, 30, "@user5")
	if currentWidth >= 60 {
		twitterCanvas.WriteBytesASCII(currentWidth-10, 30, "12h")
	}
	twitterCanvas.WriteBytesASCII(2, 31, "Working on a new project. Excited to share updates soon!")
	twitterCanvas.WriteBytesASCII(2, 32, "[Reply 7] [Retweet 4] [Like 28] [Share]")
	
	// Trending sidebar (if width allows)
	if currentWidth >= 100 {
		sidebarStart := currentWidth - 25
		twitterCanvas.DrawBox(sidebarStart, 6, maxX, 25)
		twitterCanvas.WriteBytesASCII(sidebarStart+2, 7, "Trending")
		twitterCanvas.WriteBytesASCII(sidebarStart+2, 9, "• #WorldNews")
		twitterCanvas.WriteBytesASCII(sidebarStart+2, 10, "• #Technology")  
		twitterCanvas.WriteBytesASCII(sidebarStart+2, 11, "• #Sports")
		twitterCanvas.WriteBytesASCII(sidebarStart+2, 12, "• JavaScript")
		twitterCanvas.WriteBytesASCII(sidebarStart+2, 13, "• #AI")
		twitterCanvas.WriteBytesASCII(sidebarStart+2, 14, "• #OpenSource")
		
		twitterCanvas.WriteBytesASCII(sidebarStart+2, 16, "Who to follow")
		twitterCanvas.WriteBytesASCII(sidebarStart+2, 18, "@developer1")
		twitterCanvas.WriteBytesASCII(sidebarStart+2, 19, "@techwriter2")
		twitterCanvas.WriteBytesASCII(sidebarStart+2, 20, "@designer3")
	}
	
	// Footer/Status bar
	twitterCanvas.DrawBox(0, 41, maxX, 43)
	twitterCanvas.WriteBytesASCII(2, 42, "Connected | Timeline updated")
	if currentWidth >= 80 {
		twitterCanvas.WriteBytesASCII(currentWidth-20, 42, "280 characters max")
	}
	
	// Generate output
	configName := canvas.GetConfigName()
	output := fmt.Sprintf("X (Twitter) UI Design\n")
	output += fmt.Sprintf("====================\n\n")
	output += fmt.Sprintf("Configuration: %s\n", configName)
	output += fmt.Sprintf("Features: Tweet feed, compose box, navigation, trending\n")
	output += fmt.Sprintf("ASCII Filter: Enabled (all full-width characters removed)\n\n")
	output += fmt.Sprintf("Layout:\n")
	output += twitterCanvas.String()
	
	return output
}

func main() {
	// Set canvas configuration for standard width
	canvas.SetConfig(canvas.StandardConfig)
	
	// Generate Twitter UI
	uiContent := generateTwitterUI()
	
	// Write to file
	filename := "twitter_ui_design.txt"
	err := os.WriteFile(filename, []byte(uiContent), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}
	
	fmt.Printf("X (Twitter) UI created: %s\n", filename)
}