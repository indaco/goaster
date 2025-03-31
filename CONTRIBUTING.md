# Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

## Development Environment Setup

To set up a development environment for this repository, you can use [devbox](https://www.jetify.com/devbox) along with the provided `devbox.json` configuration file.

1. Install devbox by following the instructions in the [devbox documentation](https://www.jetify.com/devbox/docs/installing_devbox/).
2. Clone this repository to your local machine.
3. Navigate to the root directory of the cloned repository.
4. Run `devbox install` to install all packages mentioned in the `devbox.json` file.
5. Run `devbox shell` to start a new shell with access to the environment.
6. Once the devbox environment is set up, you can start developing, testing, and contributing to the repository.

### Using Makefile

Additionally, you can make use of the provided `Makefile` to run various tasks:

```bash
make build       # The main build target
make examples    # Process templ files in the _examples folder
make templ       # Process TEMPL files
make test        # Run go tests
make help        # Print this help message
```
