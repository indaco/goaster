# goaster Examples

All examples are unified under a single web application.

To explore them:

1. Run the example server:

   ```bash
   go run examples/main.go
   ```

2. Open your browser and go to [http://localhost:3000](http://localhost:3000) to view the homepage listing all available examples.

Each example includes:

- A **handler** file under `examples/handlers/` showing how to **configure and use goaster** for that scenario.
- A **page** file under `examples/pages/` rendering the corresponding UI using Templ.

Examples available:

- Single message toast
- Multiple toasts
- Single message using `template/html`
- Multiple messages using `template/html`
- Theming
- Variants
- Custom icons
- Custom animations
