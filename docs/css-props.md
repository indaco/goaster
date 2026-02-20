# Goaster CSS Custom Properties

> Variables prefixed with `--_gtt-` are internal. Only `--gtt-*` variables are intended for user override.

## Semantic Color Levels

[Source Code](../assets/css/colors.css)

All semantic color levels use [`color-mix()`](<https://developer.mozilla.org/en-US/docs/Web/CSS/color-mix()>) in the `oklab` color space. Override the color space with:

```css
:root {
  --gtt-color-space: oklab; /* default */
}
```

| Level   | Base Value                   |
| :------ | :--------------------------- |
| default | `oklch(0.554 0.046 257.417)` |
| success | `oklch(0.723 0.219 149.579)` |
| error   | `oklch(0.637 0.237 25.331)`  |
| warning | `oklch(0.769 0.188 70.08)`   |
| info    | `oklch(0.623 0.214 259.815)` |

Four derived variables are generated per level from the base color:

- `--gtt-color-<level>-lightest`: 10% base, 90% white
- `--gtt-color-<level>-lighter`: 30% base, 70% white
- `--gtt-color-<level>-light`: 70% base, 30% white
- `--gtt-color-<level>-dark`: 70% base, 30% black

## Variants

[Source Code](../assets/css/variants/)

All variants (except `accent-light` and `accent-dark`) automatically adapt to light/dark mode via `@media (prefers-color-scheme)`.

| Variant        | Light Mode                               | Dark Mode                                       | Key Visual                 |
| :------------- | :--------------------------------------- | :---------------------------------------------- | :------------------------- |
| `accent`       | White bg, dark text, colored left border | Dark bg, light text, colored left border        | Auto-adaptive left accent  |
| `accent-light` | White bg, dark text, colored left border | - (forced light)                                | Explicit light mode        |
| `accent-dark`  | - (forced dark)                          | Dark bg, light text, colored left border        | Explicit dark mode         |
| `filled`       | Base-color bg, white text                | Dark-color bg, lightest text                    | Bold solid background      |
| `outlined`     | Transparent bg, base border, dark text   | Transparent bg, light border, light text        | Full colored border        |
| `soft`         | Lightest bg, dark text, no border        | Subtle tinted dark bg, lightest text, no border | Subtle tinted background   |
| `minimal`      | Transparent bg, dark text, no border     | Transparent bg, light text, no border           | Just colored text          |
| `brutalist`    | White bg, thick border, offset shadow    | Near-black bg, thick border, offset shadow      | Raw, heavy, blocky         |
| `retro`        | Cream bg, warm tones, inset shadow       | Warm dark bg, warm tones                        | Vintage, nostalgic         |
| `neon`         | Falls back to default                    | Dark bg, neon glow borders                      | Cyberpunk glow (dark only) |

### Accent

Shared background tokens:

| Variant        | CSS Variable            | Light Default | Dark Default                    |
| :------------- | :---------------------- | :------------ | :------------------------------ |
| `accent`       | `--gtt-accent-bg`       | `white`       | `hsla(220.9, 39.3%, 11%, 0.85)` |
| `accent-light` | `--gtt-accent-light-bg` | `white`       | -                               |
| `accent-dark`  | `--gtt-accent-dark-bg`  | -             | `hsla(220.9, 39.3%, 11%, 0.85)` |

Per-level overrides for `accent` and `accent-light`:

- `--gtt-<level>-accent-color` - text color
- `--gtt-<level>-accent-border-color` - left border color

Per-level overrides for `accent-dark`:

- `--gtt-<level>-accent-dark-color` - text color
- `--gtt-<level>-accent-border-color` - left border color (shared with accent/accent-light)

### Filled

Per-level overrides:

- `--gtt-filled-<level>-bg` - background color
- `--gtt-filled-<level>-color` - text color

### Outlined

Per-level overrides:

- `--gtt-outlined-<level>-color` - text color
- `--gtt-outlined-<level>-border` - border color

Shared:

- `--gtt-outlined-bg` - background (default: `transparent`)

### Soft

Per-level overrides:

- `--gtt-soft-<level>-bg` - background color (light: lightest tint; dark: 30% dark shade mixed with neutral)
- `--gtt-soft-<level>-color` - text color

### Minimal

Per-level overrides:

- `--gtt-minimal-<level>-color` - text color

Background and border are always `transparent`.

### Brutalist

Per-level overrides:

- `--gtt-brutalist-<level>-color` - text color
- `--gtt-brutalist-<level>-border` - border color

Shared:

- `--gtt-brutalist-bg` - background (light: `white`; dark: `hsl(0 0% 8%)`)

Structural: `border-width: 3px`, `border-radius: 0`, `box-shadow: 4px 4px 0 0` with level base color, monospace font stack.

### Retro

Per-level overrides:

- `--gtt-retro-<level>-color` - text color
- `--gtt-retro-<level>-border` - border color

Shared:

- `--gtt-retro-bg` - background (light: `hsl(45 40% 96%)`; dark: `hsl(35 25% 18%)`)

Structural: `border-width: 1px`, `box-shadow: inset 0 1px 2px` for paper-like depth, serif font stack (Georgia, Palatino).

### Neon (dark mode only)

No custom override properties. Uses level colors directly.

Structural: `border-width: 1px`, dual `box-shadow` glow via `color-mix()` at 50% and 20% opacity, monospace font stack. In light mode, toasts fall back to default appearance.

## Theming

### Override a base color

To change a level's color (e.g., success), override its base variable. All derived shades update automatically:

```css
:root {
  --gtt-color-success-base: oklch(0.65 0.3 155);
}
```

### Override a specific level's appearance

Each level exposes semantic theme aliases that sit between the color primitives and the component. Override these to change one level without touching the base color:

- `--gtt-<level>-bg` - background
- `--gtt-<level>-color` - text color
- `--gtt-<level>-border-color` - border color
- `--gtt-<level>-progress-bar-color` - progress bar color

```css
:root {
  --gtt-success-bg: #d4edda;
  --gtt-success-color: #155724;
}
```

## Component Styles

[Source Code](../assets/css/styles.css)

### Layout

| Property              | Default    | Description        |
| :-------------------- | :--------- | :----------------- |
| `--gtt-py`            | `0`        | Vertical padding   |
| `--gtt-px`            | `0.75em`   | Horizontal padding |
| `--gtt-bg`            | per-level  | Background color   |
| `--gtt-border-color`  | per-level  | Border color       |
| `--gtt-border-width`  | `1px`      | Border width       |
| `--gtt-border-style`  | `solid`    | Border style       |
| `--gtt-border-radius` | `0.375rem` | Border radius      |
| `--gtt-z-index`       | `1000`     | Stack order        |

### Typography

| Property            | Default   | Description |
| :------------------ | :-------- | :---------- |
| `--gtt-color`       | per-level | Text color  |
| `--gtt-font-family` | `inherit` | Font family |
| `--gtt-font-size`   | `1rem`    | Font size   |
| `--gtt-line-height` | `1.5`     | Line height |

### Progress Bar

| Property               | Default   | Description        |
| :--------------------- | :-------- | :----------------- |
| `--gtt-progress-color` | per-level | Progress bar color |

### Animation

| Property                     | Default          | Description                    |
| :--------------------------- | :--------------- | :----------------------------- |
| `--gtt-anim-name-in-top`     | `gttFadeInDown`  | Entrance animation from top    |
| `--gtt-anim-name-in-bottom`  | `gttFadeInUp`    | Entrance animation from bottom |
| `--gtt-anim-name-out-top`    | `gttFadeOutUp`   | Exit animation to top          |
| `--gtt-anim-name-out-bottom` | `gttFadeOutDown` | Exit animation to bottom       |
| `--gtt-anim-duration-in`     | `0.5s`           | Entrance duration              |
| `--gtt-anim-delay-in`        | `0s`             | Entrance delay                 |
| `--gtt-anim-direction-in`    | `normal`         | Entrance direction             |
| `--gtt-anim-timing-in`       | `ease`           | Entrance timing function       |
| `--gtt-anim-duration-out`    | `0.5s`           | Exit duration                  |
| `--gtt-anim-delay-out`       | `0s`             | Exit delay                     |
| `--gtt-anim-direction-out`   | `normal`         | Exit direction                 |
| `--gtt-anim-timing-out`      | `ease`           | Exit timing function           |
