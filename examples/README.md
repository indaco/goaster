# goaster Examples

All examples are unified under a single web application.

To explore them locally:

1. **Clone the repository**

   ```bash
   git clone https://github.com/indaco/goaster.git
   ```

2. **Run the example server**

   You can either run it from the project root:

   ```bash
   cd goaster
   make dev # or: task dev
   ```

   Or directly from the `examples` directory:

   ```bash
   cd goaster/examples
   go run .
   ```

3. **Open your browser**

   Visit [http://localhost:3000](http://localhost:3000) browse the homepage and explore all available examples.

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
