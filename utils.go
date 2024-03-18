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

// getMapCssClassByLevel returns a map of CSS classes for toasts based on their level.
func getMapCssClassByLevel() map[Level]templ.CSSClass {
	return map[Level]templ.CSSClass{
		SuccessLevel: gttSuccessLevel(),
		ErrorLevel:   gttErrorLevel(),
		WarningLevel: gttWarningLevel(),
		InfoLevel:    gttInfoLevel(),
	}
}

func isAccent(variant Variant) bool {
	return strings.Contains(variant.String(), "accent")
}

// getMapCssClassByVariant returns a map of CSS classes corresponding to each toast level
// based on the specified variant.
func getMapCssClassByVariant(variant Variant) map[Level]templ.CSSClass {
	var variantClasses map[Level]templ.CSSClass

	switch variant {
	case Accent:
		variantClasses = map[Level]templ.CSSClass{
			DefaultLevel: gttAccentDefaultLevel(),
			SuccessLevel: gttAccentSuccessLevel(),
			ErrorLevel:   gttAccentErrorLevel(),
			WarningLevel: gttAccentWarningLevel(),
			InfoLevel:    gttAccentInfoLevel(),
		}
	case AccentLight:
		variantClasses = map[Level]templ.CSSClass{
			DefaultLevel: gttAccentLightDefaultLevel(),
			SuccessLevel: gttAccentLightSuccessLevel(),
			ErrorLevel:   gttAccentLightErrorLevel(),
			WarningLevel: gttAccentLightWarningLevel(),
			InfoLevel:    gttAccentLightInfoLevel(),
		}
	case AccentDark:
		variantClasses = map[Level]templ.CSSClass{
			DefaultLevel: gttAccentDarkDefaultLevel(),
			SuccessLevel: gttAccentDarkSuccessLevel(),
			ErrorLevel:   gttAccentDarkErrorLevel(),
			WarningLevel: gttAccentDarkWarningLevel(),
			InfoLevel:    gttAccentDarkInfoLevel(),
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
