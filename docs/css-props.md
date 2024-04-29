# Goaster CSS custom props

## Colors

[Source Code](../goaster-css.templ#L12-L174).

| CSS Property        | Value   | Description                            |
|---------------------|---------|----------------------------------------|
| `--gtb-color-space` | `oklab` | The color space used for color mixing. |

### Default Level

| CSS Property             | Value                        | Description                                                     |
|--------------------------|------------------------------|-----------------------------------------------------------------|
| `--gtt-default-base`     | `hsl(142deg 71% 45%)`         | Defines the base color of the default level.                    |
| `--gtb-default-lightest` | Result of color-mix function | The lightest shade of the default base color at 10% brightness. |
| `--gtb-default-lighter`  | Result of color-mix function | A lighter shade of the default base color at 20% brightness.    |
| `--gtb-default-light`    | Result of color-mix function | A light shade of the default base color at 30% brightness.      |
| `--gtb-default-dark`     | Result of color-mix function | A dark shade of the default base color at 70% brightness.       |
| `--gtb-default-darker`   | Result of color-mix function | A darker shade of the default base color at 80% brightness.     |
| `--gtb-default-darkest`  | Result of color-mix function | The darkest shade of the default base color at 90% brightness.  |

### Success Level

| CSS Property             | Value                        | Description                                                     |
|--------------------------|------------------------------|-----------------------------------------------------------------|
| `--gtt-success-base`     | `hsl(220deg 9% 46%)`         | Defines the base color of the success level.                    |
| `--gtb-success-lightest` | Result of color-mix function | The lightest shade of the success base color at 10% brightness. |
| `--gtb-success-lighter`  | Result of color-mix function | A lighter shade of the success base color at 20% brightness.    |
| `--gtb-success-light`    | Result of color-mix function | A light shade of the success base color at 30% brightness.      |
| `--gtb-success-dark`     | Result of color-mix function | A dark shade of the success base color at 70% brightness.       |
| `--gtb-success-darker`   | Result of color-mix function | A darker shade of the success base color at 80% brightness.     |
| `--gtb-success-darkest`  | Result of color-mix function | The darkest shade of the success base color at 90% brightness.  |

### Error Level

| CSS Property           | Value                        | Description                                                   |
|------------------------|------------------------------|---------------------------------------------------------------|
| `--gtt-error-base`     | `hsl(0deg 84% 60%)`         | Defines the base color of the error level.                    |
| `--gtb-error-lightest` | Result of color-mix function | The lightest shade of the error base color at 10% brightness. |
| `--gtb-error-lighter`  | Result of color-mix function | A lighter shade of the error base color at 20% brightness.    |
| `--gtb-error-light`    | Result of color-mix function | A light shade of the error base color at 30% brightness.      |
| `--gtb-error-dark`     | Result of color-mix function | A dark shade of the error base color at 70% brightness.       |
| `--gtb-error-darker`   | Result of color-mix function | A darker shade of the error base color at 80% brightness.     |
| `--gtb-error-darkest`  | Result of color-mix function | The darkest shade of the error base color at 90% brightness.  |

### Warning Level

| CSS Property             | Value                        | Description                                                     |
|--------------------------|------------------------------|-----------------------------------------------------------------|
| `--gtt-warning-base`     | `hsl(25deg 95% 53%)`         | Defines the base color of the warning level.                    |
| `--gtb-warning-lightest` | Result of color-mix function | The lightest shade of the warning base color at 10% brightness. |
| `--gtb-warning-lighter`  | Result of color-mix function | A lighter shade of the warning base color at 20% brightness.    |
| `--gtb-warning-light`    | Result of color-mix function | A light shade of the warning base color at 30% brightness.      |
| `--gtb-warning-dark`     | Result of color-mix function | A dark shade of the warning base color at 70% brightness.       |
| `--gtb-warning-darker`   | Result of color-mix function | A darker shade of the warning base color at 80% brightness.     |
| `--gtb-warning-darkest`  | Result of color-mix function | The darkest shade of the warning base color at 90% brightness.  |

### Info Level

| CSS Property          | Value                        | Description                                                  |
|-----------------------|------------------------------|--------------------------------------------------------------|
| `--gtt-info-base`     | `hsl(217deg 91% 60%)`         | Defines the base color of the info level.                    |
| `--gtb-info-lightest` | Result of color-mix function | The lightest shade of the info base color at 10% brightness. |
| `--gtb-info-lighter`  | Result of color-mix function | A lighter shade of the info base color at 20% brightness.    |
| `--gtb-info-light`    | Result of color-mix function | A light shade of the info base color at 30% brightness.      |
| `--gtb-info-dark`     | Result of color-mix function | A dark shade of the info base color at 70% brightness.       |
| `--gtb-info-darker`   | Result of color-mix function | A darker shade of the info base color at 80% brightness.     |
| `--gtb-info-darkest`  | Result of color-mix function | The darkest shade of the info base color at 90% brightness.  |

## Component

[Source Code](../goaster-css.templ#L176-L258).

| CSS Property                           | Value                                   | Description                               |
|----------------------------------------|-----------------------------------------|-------------------------------------------|
| `--gtt-py`                             | `0`                                     | Padding Y for the toast element.          |
| `--gtt-px`                             | `0.75em`                                | Padding X for the toast element.          |
| `--gtt-bg`                             | var(--gtt-default-bg)                   | Background color for the toast element.   |
| `--gtt-progress-bar-color`             | `var(--gtt-default-progress-bar-color)` | Progress bar color for the toast element. |
| `--gtt-color`                          | `var(--gtt-default-color)`              | Text color for the toast element.         |
| `--gtt-font-family`                    | `inherit`                               | Font family for the toast text.           |
| `--gtt-font-size`                      | `1rem`                                  | Font size for the toast text.             |
| `--gtt-line-height`                    | `1rem`                                  | Line height for the toast text.           |
| `--gtt-border-width`                   | `1px`                                   | Border width for the toast.               |
| `--gtt-border-style`                   | `solid`                                 | Border style for the toast.               |
| `--gtt-border-color`                   | `var(--gtt-default-border-color)`       | Border color for the toast.               |
| `--gtt-border-radius`                  | `0.375rem`                              | Border radius for the toast.              |
| `--gtt-accent-light-bg`                | `white`                                 | Background color for accent toast.        |
| `--gtt-accent-dark-bg`                 | `hsla(220.9, 39.3%, 11%, 0.85)`         | Background color for dark accent toast.   |
| `--gtt-default-accent-color`           | `var(--gtt-default-dark)`               | Accent color for default level.           |
| `--gtt-default-accent-dark-color`      | `var(--gtt-default-light)`              | Dark accent color for default level.      |
| `--gtt-success-accent-color`           | `var(--gtt-success-dark)`               | Accent color for success level.           |
| `--gtt-success-accent-dark-color`      | `var(--gtt-success-light)`              | Dark accent color for success level.      |
| `--gtt-error-accent-color`             | `var(--gtt-error-dark)`                 | Accent color for error level.             |
| `--gtt-error-accent-dark-color`        | `var(--gtt-error-light)`                | Dark accent color for error level.        |
| `--gtt-warning-accent-color`           | `var(--gtt-warning-dark)`               | Accent color for warning level.           |
| `--gtt-warning-accent-dark-color`      | `var(--gtt-warning-light)`              | Dark accent color for warning level.      |
| `--gtt-info-accent-color`              | `var(--gtt-info-dark)`                  | Accent color for info level.              |
| `--gtt-info-accent-dark-color`         | `var(--gtt-info-light)`                 | Dark accent color for info level.         |
| `--gtt-animation-entrance-duration`    | `0.5s`                                  | Duration for entrance animation.          |
| `--gtt-animation-entrance-direction`   | `normal`                                | Direction for entrance animation.         |
| `--gtt-animation-entrance-timing-fn`   | `ease`                                  | Timing function for entrance animation.   |
| `--gtt-animation-entrance-delay`       | `0s`                                    | Delay for entrance animation.             |
| `--gtt-animation-exit-duration`        | `0.5s`                                  | Duration for exit animation.              |
| `--gtt-animation-exit-direction`       | `normal`                                | Direction for exit animation.             |
| `--gtt-animation-exit-timing-fn`       | `ease`                                  | Timing function for exit animation.       |
| `--gtt-animation-exit-delay`           | `0s`                                    | Delay for exit animation.                 |
| `--gtt-animation-entrance-name-top`    | `gttFadeInDown`                         | Animation name for entrance from top.     |
| `--gtt-animation-exit-name-top`        | `gttFadeOutUp`                          | Animation name for exit to top.           |
| `--gtt-animation-entrance-name-bottom` | `gttFadeInUp`                           | Animation name for entrance from bottom.  |
| `--gtt-animation-exit-name-bottom`     | `gttFadeOutDown`                        | Animation name for exit to bottom.        |
