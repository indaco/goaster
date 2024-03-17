package goaster

import (
	"github.com/a-h/templ"
	"strings"
)

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

func isAnyAccentVariant(variant Variant) bool {
	return strings.Contains(variant.String(), "accent")
}

// mapToastLevelToCSSClass returns a map of CSS classes corresponding to each toast level
// based on the specified variant.
func mapToastLevelToCSSClass(variant Variant) map[Level]templ.CSSClass {
	var variantClasses map[Level]templ.CSSClass

	switch variant {
	case Colorful:
		variantClasses = map[Level]templ.CSSClass{
			DefaultLevel: gttDefaultLevel(),
			SuccessLevel: gttSuccessLevel(),
			ErrorLevel:   gttErrorLevel(),
			WarningLevel: gttWarningLevel(),
			InfoLevel:    gttInfoLevel(),
		}
	case Accent:
		variantClasses = map[Level]templ.CSSClass{
			DefaultLevel: gttAccentDefaultLevel(),
			SuccessLevel: gttAccentSuccessLevel(),
			ErrorLevel:   gttAccentErrorLevel(),
			WarningLevel: gttAccentWarningLevel(),
			InfoLevel:    gttAccentInfoLevel(),
		}
	case AccentDark:
		variantClasses = map[Level]templ.CSSClass{
			DefaultLevel: gttAccentDarkDefaultLevel(),
			SuccessLevel: gttAccentDarkSuccessLevel(),
			ErrorLevel:   gttAccentDarkErrorLevel(),
			WarningLevel: gttAccentDarkWarningLevel(),
			InfoLevel:    gttAccentDarkInfoLevel(),
		}
	default:
		// Default variant
		variantClasses = map[Level]templ.CSSClass{
			DefaultLevel: gttDefaultLevel(),
			SuccessLevel: gttSuccessLevel(),
			ErrorLevel:   gttErrorLevel(),
			WarningLevel: gttWarningLevel(),
			InfoLevel:    gttInfoLevel(),
		}
	}

	return variantClasses
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
