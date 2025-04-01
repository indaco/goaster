package viewmodel

import "strings"

func GetToastEntranceClass(position string) string {
	switch strings.ToLower(position) {
	case "top-left", "top-right", "top-center":
		return "gttShowFromTop"
	case "bottom-left", "bottom-right", "bottom-center":
		return "gttShowFromBottom"
	default:
		return "gttShowFromBottom"
	}
}
