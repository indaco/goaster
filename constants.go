package goaster

/* ------------------------------------------------------------------------- */
/* LEVELS                                                                    */
/* ------------------------------------------------------------------------- */

// Predefined toast levels
const (
	DefaultLevel Level = "default"
	SuccessLevel Level = "success"
	ErrorLevel   Level = "error"
	WarningLevel Level = "warning"
	InfoLevel    Level = "info"
)

/* ------------------------------------------------------------------------- */
/* POSITIONS                                                                 */
/* ------------------------------------------------------------------------- */

// Predefined toast positions
const (
	TopRight     = Position("top-right")
	TopLeft      = Position("top-left")
	TopCenter    = Position("top-center")
	BottomRight  = Position("bottom-right")
	BottomLeft   = Position("bottom-left")
	BottomCenter = Position("bottom-center")
)

/* ------------------------------------------------------------------------- */
/* VARIANTS                                                                  */
/* ------------------------------------------------------------------------- */

// Predefined style variants
const (
	Accent      = Variant("accent")
	AccentLight = Variant("accent-light")
	AccentDark  = Variant("accent-dark")
)

/* ------------------------------------------------------------------------- */
/* ENTRANCE ANIMATION CLASSES                                                */
/* ------------------------------------------------------------------------- */

var toastEntranceClasses = map[Position]string{
	TopRight:     "gttShowFromTop",
	TopLeft:      "gttShowFromTop",
	TopCenter:    "gttShowFromTop",
	BottomRight:  "gttShowFromBottom",
	BottomLeft:   "gttShowFromBottom",
	BottomCenter: "gttShowFromBottom",
}
