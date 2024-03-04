package goaster

// Constants for different levels of toast notifications.
const (
	DefaultLevel Level = "default"
	SuccessLevel Level = "success"
	ErrorLevel   Level = "error"
	WarningLevel Level = "warning"
	InfoLevel    Level = "info"
)

// Default SVG icons for each level of toast notifications.
// These are inline SVGs for default icons, which can be replaced by the user.
const (
	defaultDefaultIcon = `<svg xmlns="http://www.w3.org/2000/svg" style="stroke: currentColor;" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>`
	defaultSuccessIcon = `<svg xmlns="http://www.w3.org/2000/svg" style="stroke: currentColor;" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>`
	defaultErrorIcon   = `<svg xmlns="http://www.w3.org/2000/svg" style="stroke: currentColor;" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"></path></svg>`
	defaultWarningIcon = `<svg xmlns="http://www.w3.org/2000/svg" style="stroke: currentColor;" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>`
	defaultInfoIcon    = `<svg xmlns="http://www.w3.org/2000/svg" style="stroke: currentColor;" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>`
)

// Position constants define where the toast will appear on the screen.
var (
	TopRight     = Position{Name: "top-right", CSS: gttContainerTopRight()}
	TopLeft      = Position{Name: "top-left", CSS: gttContainerTopLeft()}
	TopCenter    = Position{Name: "top-center", CSS: gttContainerTopCenter()}
	BottomRight  = Position{Name: "bottom-right", CSS: gttContainerBottomRight()}
	BottomLeft   = Position{Name: "bottom-left", CSS: gttContainerBottomLeft()}
	BottomCenter = Position{Name: "bottom-center", CSS: gttContainerBottomCenter()}
)
