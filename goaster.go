package goaster

var (
	// entranceCssClassesByPosition maps each Position to a corresponding templ.CSSClass representing the entrance direction for toast animations.
	entranceCssClassesByPosition map[Position]string
)

// getToastEntranceCSSClassesByPosition returns a map of CSS classes for toast entrances based on their position.
func getToastEntranceCSSClassesByPosition() map[Position]string {
	return map[Position]string{
		TopRight:     "gttShowFromTop",
		TopLeft:      "gttShowFromTop",
		TopCenter:    "gttShowFromTop",
		BottomRight:  "gttShowFromBottom",
		BottomLeft:   "gttShowFromBottom",
		BottomCenter: "gttShowFromBottom",
	}
}

func init() {
	entranceCssClassesByPosition = getToastEntranceCSSClassesByPosition()
}
