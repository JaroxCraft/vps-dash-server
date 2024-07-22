package colors

import "github.com/fatih/color"

func S(c color.Attribute, text string) string {
	return color.New(c).Sprint(text)
}
