<h1 align="center" style="font-size: 2.5rem;">
  goaster
</h1>
<h2 align="center" style="font-size: 1.5rem;">
A configurable, themeable and non-intrusive server-rendered toast notification component for Go web applications.
</h2>
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
</p>

Built with [templ](https://github.com/a-h/templ) library for seamless integration with Go-based web frontends.

### Features

- **No External Dependencies**: Built with native Go and the `templ` library, requiring no external frontend dependencies.
- **Multiple Toasts**: Support to display multiple toast notifications.
- **Configurable**: Customize appearance, behavior, and position.
- **Variants**: Support for different toast variants including `Colorful` (default), `Accent`, and `AccentDark`.
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
go get github.com/indaco/goaster
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

## Theming

Toast notifications are themeable using CSS variables (prefix `gtt`) to customize the appearance according to your design.

Here below is the list of all CSS variables defined and their default values:

| Name                                   | Default Value                                                          | Description                                                      |
|----------------------------------------|------------------------------------------------------------------------|------------------------------------------------------------------|
| `--gtt-py`                             | 0                                                                      | Padding on the y-axis for the toast element                      |
| `--gtt-px`                             | 0.75rem                                                                | Padding on the x-axis for the toast element                      |
| `--gtt-font-family`                    | inherit                                                                | Font family for the toast text                                   |
| `--gtt-font-size`                      | 1rem                                                                   | Font size for the toast text                                     |
| `--gtt-line-height`                    | 1rem                                                                   | Line height for the toast text                                   |
| `--gtt-border-radius`                  | 0.375rem                                                               | Border radius for the toast element                              |
| `--gtt-border-style`                   | solid                                                                  | Border style for the toast                                       |
| `--gtt-border-width`                   | 1px                                                                    | Border width for the toast                                       |
| `--gtt-default-border-color`           | ![Color Preview](https://via.placeholder.com/20/f3f4f6?text=+) #f3f4f6 | Default border color for the toast                               |
| `--gtt-default-bg`                     | ![Color Preview](https://via.placeholder.com/20/f9fafb?text=+) #f9fafb | Default background color for the toast                           |
| `--gtt-default-color`                  | ![Color Preview](https://via.placeholder.com/20/1f2937?text=+) #1f2937 | Default text color for the toast                                 |
| `--gtt-success-border-color`           | ![Color Preview](https://via.placeholder.com/20/dcfce7?text=+) #dcfce7 | Success border color for the toast                               |
| `--gtt-success-bg`                     | ![Color Preview](https://via.placeholder.com/20/f0fdf4?text=+) #f0fdf4 | Success background color for the toast                           |
| `--gtt-success-color`                  | ![Color Preview](https://via.placeholder.com/20/166534?text=+) #166534 | Success text color for the toast                                 |
| `--gtt-error-border-color`             | ![Color Preview](https://via.placeholder.com/20/fee2e2?text=+) #fee2e2 | Error border color for the toast                                 |
| `--gtt-error-bg`                       | ![Color Preview](https://via.placeholder.com/20/fef2f2?text=+) #fef2f2 | Error background color for the toast                             |
| `--gtt-error-color`                    | ![Color Preview](https://via.placeholder.com/20/991b1b?text=+) #991b1b | Error text color for the toast                                   |
| `--gtt-warning-border-color`           | ![Color Preview](https://via.placeholder.com/20/ffedd5?text=+) #ffedd5 | Warning border color for the toast                               |
| `--gtt-warning-bg`                     | ![Color Preview](https://via.placeholder.com/20/fff7ed?text=+) #fff7ed | Warning background color for the toast                           |
| `--gtt-warning-color`                  | ![Color Preview](https://via.placeholder.com/20/9a3412?text=+) #9a3412 | Warning text color for the toast                                 |
| `--gtt-info-border-color`              | ![Color Preview](https://via.placeholder.com/20/dbeafe?text=+) #dbeafe | Info border color for the toast                                  |
| `--gtt-info-bg`                        | ![Color Preview](https://via.placeholder.com/20/eff6ff?text=+) #eff6ff | Info background color for the toast                              |
| `--gtt-info-color`                     | ![Color Preview](https://via.placeholder.com/20/1e40af?text=+) #1e40af | Info text color for the toast                                    |
| `--gtt-animation-entrance-duration`    | 0.5s                                                                   | Default duration for entrance animations                         |
| `--gtt-animation-entrance-direction`   | normal                                                                 | Default direction for entrance animations                        |
| `--gtt-animation-entrance-timing-func` | ease                                                                   | Default timing function for entrance animations                  |
| `--gtt-animation-entrance-delay`       | 0s                                                                     | Default delay for entrance animations                            |
| `--gtt-animation-exit-duration`        | 0.5s                                                                   | Default duration for exit animations                             |
| `--gtt-animation-exit-direction`       | normal                                                                 | Default direction for exit animations                            |
| `--gtt-animation-exit-timing-function` | ease                                                                   | Default timing function for exit animations                      |
| `--gtt-animation-exit-delay`           | 0s                                                                     | Default delay for exit animations                                |
| `--gtt-animation-entrance-name-top`    | gttFadeInDown                                                          | Name of the entrance animation when direction is from the top    |
| `--gtt-animation-exit-name-top`        | gttFadeOutUp                                                           | Name of the exit animation when direction is from the top        |
| `--gtt-animation-entrance-name-bottom` | gttFadeInUp                                                            | Name of the entrance animation when direction is from the bottom |
| `--gtt-animation-exit-name-bottom`     | gttFadeOutDown                                                         | Name of the exit animation when direction is from the bottom     |

## Custom Icons

Specify custom SVG icons for each toast level:

```go
toaster := goaster.NewToaster(
   goaster.WithIcon(toast.SuccessLevel, "<svg>...</svg>"),
   goaster.WithIcon(toast.ErrorLevel, "<svg>...</svg>"),
)
```

## Examples

- [use with `a-h/templ` template](_examples/a-h-templ-single-toast)
- [multiple messages with `a-h/templ`](_examples/a-h-templ-multiple-toasts)
- [use with `template/html`](_examples/go-html-template-single-toast)
- [multiple  messages with `template/html`](_examples/go-html-template-multiple-toasts)
- [theming](_examples/theming)
- [variants](_examples/variants)
- [custom icons](_examples/custom-icons)
- [custom animations](_examples/custom-animations)

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
