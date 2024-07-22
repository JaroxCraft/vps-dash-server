package banner

import (
	"github.com/fatih/color"
	"github.com/jaroxcraft/vps-dash-server/colors"
	"github.com/jaroxcraft/vps-dash-server/info"
	"io"
	"os"
	"strings"
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
	banner = strings.Replace(banner, "${name}", colors.S(color.ReverseVideo, info.GetInfo().Name), -1)
	banner = strings.Replace(banner, "${version}", colors.S(color.Italic, info.GetInfo().Version), -1)

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
