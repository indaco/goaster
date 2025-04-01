package goaster

// Inline SVGs for default toast icons
const (
	defaultIconDefault = `<svg xmlns="http://www.w3.org/2000/svg" style="stroke: currentColor;" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>`
	defaultIconSuccess = `<svg xmlns="http://www.w3.org/2000/svg" style="stroke: currentColor;" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>`
	defaultIconError   = `<svg xmlns="http://www.w3.org/2000/svg" style="stroke: currentColor;" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>`
	defaultIconWarning = `<svg xmlns="http://www.w3.org/2000/svg" style="stroke: currentColor;" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"></path></svg>`
	defaultIconInfo    = `<svg xmlns="http://www.w3.org/2000/svg" style="stroke: currentColor;" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>`
)

// defaultIcons returns a map of default SVG icons for each toast level.
func defaultIcons() map[Level]string {
	return map[Level]string{
		DefaultLevel: defaultIconDefault,
		SuccessLevel: defaultIconSuccess,
		ErrorLevel:   defaultIconError,
		WarningLevel: defaultIconWarning,
		InfoLevel:    defaultIconInfo,
	}
}
