package goaster

// Predefined toast levels
const (
	DefaultLevel Level = "default"
	SuccessLevel Level = "success"
	ErrorLevel   Level = "error"
	WarningLevel Level = "warning"
	InfoLevel    Level = "info"
)

// Predefined toast positions
const (
	TopRight     = Position("top-right")
	TopLeft      = Position("top-left")
	TopCenter    = Position("top-center")
	BottomRight  = Position("bottom-right")
	BottomLeft   = Position("bottom-left")
	BottomCenter = Position("bottom-center")
)

// Predefined style variants
const (
	Accent      = Variant("accent")
	AccentLight = Variant("accent_light")
	AccentDark  = Variant("accent_dark")
)
