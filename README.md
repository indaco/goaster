<h1 align="center">
  goaster
</h1>
<h2 align="center" style="font-size: 1.5em;">
  toast notification component for Go web applications.
</h2>
<p align="center">
  <a href="https://github.com/indaco/goaster/actions/workflows/ci.yml" target="_blank">
      <img src="https://github.com/indaco/goaster/actions/workflows/ci.yml/badge.svg" alt="CI" />
    </a>
    <a href="https://codecov.io/gh/indaco/goaster">
      <img src="https://codecov.io/gh/indaco/goaster/branch/main/graph/badge.svg" alt="Code coverage" />
    </a>
  <a href="https://goreportcard.com/report/github.com/indaco/goaster/" target="_blank">
    <img src="https://goreportcard.com/badge/github.com/indaco/goaster" alt="go report card" />
  </a>
  <a href="https://badge.fury.io/gh/indaco%2Fgoaster">
    <img src="https://badge.fury.io/gh/indaco%2Fgoaster.svg" alt="GitHub version" height="18">
  </a>
  <a href="https://pkg.go.dev/github.com/indaco/goaster/" target="_blank">
      <img src="https://pkg.go.dev/badge/github.com/indaco/goaster/.svg" alt="go reference" />
  </a>
   <a href="https://github.com/indaco/goaster/blob/main/LICENSE" target="_blank">
    <img src="https://img.shields.io/badge/license-mit-blue?style=flat-square&logo=none" alt="license" />
  </a>
  <a href="https://www.jetify.com/devbox/docs/contributor-quickstart/">
    <img src="https://www.jetify.com/img/devbox/shield_moon.svg" alt="Built with Devbox" />
  </a>
</p>

<p align="center">
  <a href="#features">Features</a> |
  <a href="#installation">Installation</a> |
  <a href="#usage">Usage</a> |
  <a href="#custom-icons">Custom Icons</a> |
  <a href="#theming">Theming</a> |
  <a href="#examples">Examples</a> |
  <a href="#contributing">Contributing</a>
</p>

![goaster demo](https://raw.githubusercontent.com/indaco/gh-assets/main/goaster/demo.gif)

### Features

- **No External Dependencies**: Built with native Go and the [`templ`](https://templ.guide) library - no JavaScript required.
- **Multiple Toasts**: Display several notifications simultaneously.
- **Configurable**: Control appearance (border, rounded), behavior, and position.
- **Variants**: `Accent`, `AccentLight`, `AccentDark`, `Filled`, `Outlined`, `Soft`, `Minimal`, `Brutalist`, `Retro`, and `Neon`.
- **Themeable**: Style with CSS variables to match your app's design.
- **Icon Support**: Default SVG icons for common toast types - or use your own.
- **Flexible Positioning**: Place toasts in any corner of the screen.
- **Auto-Dismiss Progress Bar**: Visual countdown indicator for auto-dismissed toasts.
- **Smooth Animations**: Built-in entrance/exit transitions, fully controllable via CSS variables. Can be disabled.

## Installation

```sh
go get github.com/indaco/goaster@latest
```

## Usage

```go
import "github.com/indaco/goaster"
```

### Creating a Toaster

#### 1. Default settings

```go
func HandleSingle(w http.ResponseWriter, r *http.Request) {
  toaster := goaster.ToasterDefaults()
  templ.Handler(pages.HomePage(toaster)).ServeHTTP(w, r)
}
```

#### 2. Functional options

```go
func HandleSingle(w http.ResponseWriter, r *http.Request) {
  toaster := goaster.NewToaster(
    goaster.WithVariant(goaster.Filled),
    goaster.WithBorder(false),
    goaster.WithRounded(true),
    goaster.WithShowIcon(true),
    goaster.WithButton(true),
    goaster.WithAutoDismiss(false),
    goaster.WithAnimation(true),
    goaster.WithProgressBar(false),
    goaster.WithPosition(goaster.TopRight),
  )
  templ.Handler(pages.HomePage(toaster)).ServeHTTP(w, r)
}
```

#### Options

| Option                    | Description                                      | Default        |
| ------------------------- | ------------------------------------------------ | -------------- |
| `WithVariant(Variant)`    | Style variant (`Accent`, `Filled`, `Neon`, etc.) | _(none)_       |
| `WithBorder(bool)`        | Display border around the toast                  | `true`         |
| `WithRounded(bool)`       | Use rounded corners                              | `true`         |
| `WithShowIcon(bool)`      | Show the toast icon                              | `true`         |
| `WithButton(bool)`        | Show the close button                            | `true`         |
| `WithAutoDismiss(bool)`   | Auto-dismiss after timeout                       | `true`         |
| `WithAnimation(bool)`     | Enable entrance/exit animations                  | `true`         |
| `WithProgressBar(bool)`   | Show auto-dismiss progress bar                   | `true`         |
| `WithPosition(Position)`  | Screen position (`TopRight`, `BottomLeft`, etc.) | `BottomRight`  |
| `WithIcon(Level, string)` | Custom SVG icon for a toast level                | built-in icons |

#### 3. Builder pattern

```go
func HandleSingle(w http.ResponseWriter, r *http.Request) {
  toaster := goaster.NewToasterBuilder().
    WithVariant(goaster.Neon).
    WithBorder(true).
    WithRounded(true).
    WithAutoDismiss(false).
    WithPosition(goaster.TopRight).
    Build()
  templ.Handler(pages.HomePage(toaster)).ServeHTTP(w, r)
}
```

### Displaying Toast Messages

In your templ template, call the method for the desired level:

```templ
templ UserPage(toaster *goaster.Toaster) {
  @toaster.Default("Sample message.")
  @toaster.Success("Operation completed successfully.")
  @toaster.Error("An error occurred.")
  @toaster.Info("Here's some information.")
  @toaster.Warning("This is a warning message.")
}
```

#### Queueing Multiple Toasts

To display multiple toasts, queue them in your handler:

```go
toaster.PushDefault("Sample message.")
toaster.PushSuccess("Operation completed successfully.")
toaster.PushError("An error occurred.")
toaster.PushInfo("Here's some information.")
toaster.PushWarning("This is a warning message.")
```

Then render them all in your templ page:

```templ
templ UserPage(toaster *goaster.Toaster) {
  @toaster.RenderAll()
}
```

## Custom Icons

Override the default icon for any toast level:

```go
toaster := goaster.NewToaster(
   goaster.WithIcon(goaster.SuccessLevel, "<svg>...</svg>"),
   goaster.WithIcon(goaster.ErrorLevel, "<svg>...</svg>"),
)
```

## Theming

Theme goaster using CSS custom properties prefixed with `gtt`. See the [CSS Custom Properties](./docs/css-props.md) reference for full details.

## html/template Integration

`goaster` includes `HTMLGenerator` for use with Go's `html/template`:

```go
func HandlePage(w http.ResponseWriter, r *http.Request) {
  toaster := goaster.ToasterDefaults()
  toaster.PushSuccess("Operation completed.")

  gen := goaster.NewHTMLGenerator()
  toastHTML, err := gen.DisplayAll(r.Context(), toaster)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  tmpl.Execute(w, map[string]any{
    "Toasts": toastHTML,
  })
}
```

## Examples

See the [examples](examples/) for a complete showcase with setup instructions.

## Contributing

See the [Contributing Guide](/CONTRIBUTING.md).

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.
