<h1 align="center" style="font-size: 2.5rem;">
  goaster
</h1>
<p align="center">
  <a href="https://github.com/indaco/goaster/blob/main/LICENSE" target="_blank">
    <img src="https://img.shields.io/badge/license-mit-blue?style=flat-square&logo=none" alt="license" />
  </a>
  &nbsp;
  <a href="https://goreportcard.com/report/github.com/indaco/goaster/" target="_blank">
    <img src="https://goreportcard.com/badge/github.com/indaco/goaster" alt="go report card" />
  </a>
  &nbsp;
  <a href="https://pkg.go.dev/github.com/indaco/goaster/" target="_blank">
      <img src="https://pkg.go.dev/badge/github.com/indaco/goaster/.svg" alt="go reference" />
  </a>
  &nbsp;
  <a href="https://www.jetify.com/devbox/docs/contributor-quickstart/">
    <img src="https://www.jetify.com/img/devbox/shield_moon.svg" alt="Built with Devbox" />
  </a>
</p>

<p align="center">
  <a href="#features">Features</a> •
  <a href="#installation">Installation</a> •
  <a href="#usage">Usage</a> •
  <a href="#custom-icons">Custom Icons</a> •
  <a href="#theming">Theming</a> •
  <a href="#examples">Examples</a> •
  <a href="#contributing">Contributing</a>
</p>

A configurable, themeable and non-intrusive server-rendered toast notification component for Go web applications. Built with [templ](https://github.com/a-h/templ) library for seamless integration with Go-based web frontends.

### Features

- **No External Dependencies**: Built with native Go and the `templ` library, requiring no external frontend dependencies.
- **Multiple Toasts**: Support to display multiple toast notifications.
- **Configurable**: Customize appearance (bordered, rounded), behavior, and position.
- **Variants**: Provide toast style variants like `Accent`, `AccentLight` and `AccentDark`.
- **Themeable**: Use CSS variables to theme your toasts to match your application's design.
- **Icon Support**: Include default SVG icons for various toast levels (such as success, error, info, etc.), allowing you to use your preferred icons.
- **Positioning**: Flexible positioning of toast messages (top-right, bottom-left, etc.).
- **Progress Bar for Auto-Dismiss**: A visual progress bar indicates the countdown until the toast will automatically dismiss (with auto-dismiss enabled only).
- **Animations**: Entrance and exit animations for toast notifications.
  - _Default Animations_: built-in animations for toasts appearing and disappearing.
  - _Custom Animations_: Customize animations via CSS variables or integrate with external CSS animation libraries.
  - _Animation Control_: Full control over animation timing, easing, delay and effects, via CSS variables.
  - _Disable Animations Option_: Ability to completely disable animations, providing flexibility for different application needs and user preferences.

<div style="display: flex; justify-content: center;">
   <img src="statics/demo.gif" alt="Image" >
</div>

## Installation

Ensure your project is using Go Modules.

To install the module, use the `go get` command:

```sh
go get github.com/indaco/goaster@latest
```

## Usage

Import `goaster` module into your project:

```go
import "github.com/indaco/goaster"
```

### Creating a Toaster

Initialize a `Toaster` with default settings:

```go
toaster := goaster.NewToaster()
```

or customize the toaster with options:

```go
toaster := goaster.NewToaster(
   goaster.WithBorder(false),
   goaster.WithPosition(goaster.TopRight),
   goaster.WithAutoDismiss(false)
   // ...
)
```

### Add goester CSS and Javascript

`goaster` leverages the `templ` library's features, including CSS Components and JavaScript Templates, to encapsulate all necessary styling and functionality without relying on external dependencies.

- `GoasterCSS`: it supplies the required CSS, encapsulating the visual design and layout specifics of the toast notifications.
- `GoasterJS`: it provides the JavaScript logic essential for dynamic behaviors such as displaying, hiding, and managing toast notifications.

To facilitate integration with Go's `template/html` standard library, `goaster` includes a dedicated `HTMLGenerator` type to seamlessly integrate toast notifications into web applications built with Go's `html/template` standard library.

3 methods, acting as wrappers to the templ's `templ.ToGoHTML`, generate the necessary HTML to be embedded them into server-rendered pages:

- `GoasterCSSToGoHTML`: render the `GoasterCSS` component into a `template.HTML` value.
- `GoasterJSToGoHTML`: render the `GoasterJS`component into a `template.HTML` value.

> **Note**: refer to the [Examples](#examples) section to see how to use `goaster` with `templ` and `html/template`.

### Displaying Toast Messages

Display different levels of toast messages:

```go
toaster.Default("Sample message.")
toaster.Success("Operation completed successfully.")
toaster.Error("An error occurred.")
toaster.Info("Here's some information.")
toaster.Warning("This is a warning message.")
```

### Add Toast Messages to the queue and display them

Multiple toast messages in the queue:

```go
toaster.PushDefault("Sample message.")
toaster.PushSuccess("Operation completed successfully.")
toaster.PushError("An error occurred.")
toaster.PushInfo("Here's some information.")
toaster.PushWarning("This is a warning message.")

toaster.RenderAll()
```

## Custom Icons

Specify custom SVG icons for each toast level:

```go
toaster := goaster.NewToaster(
   goaster.WithIcon(toast.SuccessLevel, "<svg>...</svg>"),
   goaster.WithIcon(toast.ErrorLevel, "<svg>...</svg>"),
)
```

## Theming

Customizing the appearance of `goaster` notifications to align with your design preferences is both straightforward and flexible, accomplished by using CSS custom properties (CSS variables) prefixed with `gtt`. For a comprehensive list of CSS custom properties, along with their default values and descriptions, please consult the tabber [CSS custom Props](./docs/css-props.md) document.

## Examples

- [use with `a-h/templ` template](_examples/a-h-templ-single-toast)
- [multiple messages with `a-h/templ`](_examples/a-h-templ-multiple-toasts)
- [use with `template/html`](_examples/go-html-template-single-toast)
- [multiple messages with `template/html`](_examples/go-html-template-multiple-toasts)
- [theming](_examples/theming)
- [variants](_examples/variants)
- [custom icons](_examples/custom-icons)
- [custom animations](_examples/custom-animations)

## Contributing

Contributions are welcome!

See the [Contributing Guide](/CONTRIBUTING.md) for setup instructions.

## License

This project is licensed under the MIT License – see the [LICENSE](./LICENSE) file for details.
