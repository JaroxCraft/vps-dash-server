package banner

import (
	"io"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/jaroxcraft/vps-dash-server/colors"
	"github.com/jaroxcraft/vps-dash-server/info"
)

const defaultBanner = ""

func PrintBanner() {
	banner := getBanner()
	banner = colorize(banner)
	banner = replaceInfo(banner)
	print(banner)
}

func colorize(banner string) string {
	return colors.S(color.FgYellow, banner)
}

func replaceInfo(banner string) string {
	banner = strings.ReplaceAll(banner, "${name}", colors.S(color.ReverseVideo, info.GetInfo().Name))
	banner = strings.ReplaceAll(banner, "${version}", colors.S(color.Italic, info.GetInfo().Version))

	return banner
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
