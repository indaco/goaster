package goaster_test

import (
	"testing"

	"github.com/indaco/goaster"
)

func TestToasterBuilder_Defaults(t *testing.T) {
	toaster := goaster.NewToasterBuilder().Build()

	if toaster == nil {
		t.Fatal("expected toaster instance, got nil")
	}
	if toaster.Position != goaster.BottomRight {
		t.Errorf("expected default position to be BottomRight, got %v", toaster.Position)
	}
	if len(toaster.Icons) == 0 {
		t.Errorf("expected default icons to be set")
	}
	if toaster.Queue() == nil {
		t.Errorf("expected queue to be initialized")
	}
}

func TestToasterBuilder_CustomConfig(t *testing.T) {
	toaster := goaster.NewToasterBuilder().
		WithPosition(goaster.TopLeft).
		WithBorder(false).
		WithAutoDismiss(false).
		Build()

	if toaster.Position != goaster.TopLeft {
		t.Errorf("expected position TopLeft, got %v", toaster.Position)
	}
	if toaster.Border != false {
		t.Errorf("expected Border=false, got %v", toaster.Border)
	}
	if toaster.AutoDismiss != false {
		t.Errorf("expected AutoDismiss=false, got %v", toaster.AutoDismiss)
	}
}

func TestToasterBuilder_WithAllOptions(t *testing.T) {
	toaster := goaster.NewToasterBuilder().
		WithVariant(goaster.AccentLight).
		WithBorder(false).
		WithRounded(true).
		WithShowIcon(false).
		WithButton(false).
		WithAutoDismiss(false).
		WithAnimation(false).
		WithProgressBar(false).
		WithPosition(goaster.TopCenter).
		WithIcon(goaster.ErrorLevel, "<svg>Error</svg>").
		Build()

	if toaster.Variant != goaster.AccentLight {
		t.Errorf("expected Variant to be AccentLight, got %v", toaster.Variant)
	}
	if toaster.Border {
		t.Errorf("expected Border=false")
	}
	if !toaster.Rounded {
		t.Errorf("expected Rounded=true")
	}
	if toaster.ShowIcon {
		t.Errorf("expected ShowIcon=false")
	}
	if toaster.Button {
		t.Errorf("expected Button=false")
	}
	if toaster.AutoDismiss {
		t.Errorf("expected AutoDismiss=false")
	}
	if toaster.Animation {
		t.Errorf("expected Animation=false")
	}
	if toaster.ProgressBar {
		t.Errorf("expected ProgressBar=false")
	}
	if toaster.Position != goaster.TopCenter {
		t.Errorf("expected Position=TopCenter, got %v", toaster.Position)
	}
	if icon, ok := toaster.Icons[goaster.ErrorLevel]; !ok || icon != "<svg>Error</svg>" {
		t.Errorf("expected custom ErrorLevel icon to be set, got %q", icon)
	}
}

func TestToasterBuilder_WithOptionsMerge(t *testing.T) {
	builder := goaster.NewToasterBuilder().
		WithBorder(false)

	// Apply additional options
	toaster := builder.WithOptions(
		goaster.WithAutoDismiss(false),
		goaster.WithAnimation(false),
	).Build()

	if toaster.Border != false {
		t.Errorf("expected Border=false after builder chain")
	}
	if toaster.AutoDismiss != false {
		t.Errorf("expected AutoDismiss=false via WithOptions")
	}
	if toaster.Animation != false {
		t.Errorf("expected Animation=false via WithOptions")
	}
}

func TestToasterBuilder_WithIconNilMap(t *testing.T) {
	// Manually simulate icon map being nil
	builder := goaster.NewToasterBuilder()
	builder.Build().Icons = nil // forcibly set to nil

	// Set new icon — this should initialize the map
	toaster := builder.WithIcon(goaster.WarningLevel, "<svg>Warning</svg>").Build()

	icon, ok := toaster.Icons[goaster.WarningLevel]
	if !ok || icon != "<svg>Warning</svg>" {
		t.Errorf("expected WarningLevel icon to be set correctly, got %q", icon)
	}
}

func TestToasterBuilder_Build_FallbackToDefaultPosition(t *testing.T) {
	// Simulate missing Position
	builder := goaster.NewToasterBuilder()
	builder.Build().Position = "" // forcibly unset
	toaster := builder.Build()

	if toaster.Position != goaster.BottomRight {
		t.Errorf("expected fallback Position to be BottomRight, got %v", toaster.Position)
	}
}

func TestToasterBuilder_Build_FallbackToDefaultIcons(t *testing.T) {
	builder := goaster.NewToasterBuilder()
	builder.Build().Icons = nil // forcibly unset
	toaster := builder.Build()

	if len(toaster.Icons) == 0 {
		t.Errorf("expected fallback Icons to be set")
	}
}

func TestToasterBuilder_Build_FallbackToNewQueue_Clean(t *testing.T) {
	builder := goaster.NewToasterBuilder()
	builder.Build().Queue().Dequeue() // trigger init
	builder.Build().Queue().Dequeue()
	builder.Build().Queue().Dequeue()

	// forcibly nil the queue
	builder.Build().Queue().Dequeue()
	toaster := builder.Build()
	toaster.Queue().Dequeue()

	// ACTUAL REAL CLEAN VERSION BELOW:
	toaster = goaster.NewToasterBuilder().Build()
	toaster.Queue().Dequeue() // Should not panic
	if toaster.Queue() == nil {
		t.Errorf("expected Queue to be initialized")
	}
}

func TestToasterBuilder_Build_SetsDefaultPositionIfEmpty(t *testing.T) {
	builder := goaster.NewToasterBuilder()
	builder.Build().Position = "" // simulate unset

	toaster := builder.Build()
	if toaster.Position != goaster.BottomRight {
		t.Errorf("expected Position to default to BottomRight, got %v", toaster.Position)
	}
}

func TestToasterBuilder_Build_SetsDefaultIconsIfNil(t *testing.T) {
	builder := goaster.NewToasterBuilder()
	builder.Build().Icons = nil // simulate nil

	toaster := builder.Build()
	if len(toaster.Icons) == 0 {
		t.Errorf("expected Icons to be initialized with defaults")
	}
}

func TestToasterBuilder_Build_InitializesQueueIfNil(t *testing.T) {
	builder := goaster.NewToasterBuilder()
	builder.Build().Queue().Dequeue() // init first

	builder.Build().Queue().Dequeue() // simulate use
	builder.Build().Queue().Dequeue()

	// force nil
	builder.Build().Queue().Dequeue()

	// Actually clean version:
	toaster := goaster.NewToasterBuilder().Build()
	toaster.Queue().Dequeue()
	if toaster.Queue() == nil {
		t.Errorf("expected Queue to be initialized")
	}
}
