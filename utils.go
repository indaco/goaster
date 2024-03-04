package goaster

import "github.com/a-h/templ"

// defaultIcons returns a map of default SVG icons for each toast level.
func defaultIcons() map[Level]string {
	return map[Level]string{
		DefaultLevel: defaultDefaultIcon,
		SuccessLevel: defaultSuccessIcon,
		ErrorLevel:   defaultErrorIcon,
		WarningLevel: defaultWarningIcon,
		InfoLevel:    defaultInfoIcon,
	}
}

// getToastCSSClassesByLevel returns a map of CSS classes for toasts based on their level.
func getToastCSSClassesByLevel() map[Level]templ.CSSClass {
	return map[Level]templ.CSSClass{
		SuccessLevel: gttSuccessLevel(),
		ErrorLevel:   gttErrorLevel(),
		WarningLevel: gttWarningLevel(),
		InfoLevel:    gttInfoLevel(),
	}
}

// getToastEntranceCSSClassesByPosition returns a map of CSS classes for toast entrances based on their position.
func getToastEntranceCSSClassesByPosition() map[Position]templ.CSSClass {
	return map[Position]templ.CSSClass{
		TopRight:     gttShowFromTop(),
		TopLeft:      gttShowFromTop(),
		TopCenter:    gttShowFromTop(),
		BottomRight:  gttShowFromBottom(),
		BottomLeft:   gttShowFromBottom(),
		BottomCenter: gttShowFromBottom(),
	}
}
