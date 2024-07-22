package banner

import (
	"github.com/fatih/color"
	"io"
	"os"
	"strings"
)

const defaultBanner = "                                  .___             .__\n___  ________  ______           __| _/____    _____|  |__\n\\  \\/ /\\____ \\/  ___/  ______  / __ |\\__  \\  /  ___/  |  \\\n \\   / |  |_> >___ \\  /_____/ / /_/ | / __ \\_\\___ \\|   Y  \\\n  \\_/  |   __/____  >         \\____ |(____  /____  >___|  /\n       |__|       \\/               \\/     \\/     \\/     \\/\n\n"

func PrintBanner() {
	banner := getBanner()
	c := color.New(color.FgYellow, color.Bold)

	_, _ = c.Print(banner)
}

func getBanner() string {
	open, err := os.Open("banner.txt")
	if err != nil {
		return defaultBanner
	}

	bannerB, err := io.ReadAll(open)
	if err != nil {
		return defaultBanner
	}
	bannerS := string(bannerB)

	if !strings.HasSuffix(bannerS, "\n") {
		bannerS += "\n"
	}

	return bannerS
}
