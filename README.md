# Sorting Algorithms in Go

[![Build Status](https://github.com/adityasanka/gosort/actions/workflows/build.yml/badge.svg?branch=master)](https://github.com/adityasanka/gosort/actions/workflows/build.yml)
[![Tests](https://github.com/adityasanka/gosort/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/adityasanka/gosort/actions/workflows/tests.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This repository contains implementations of standard algorithms like Bubble Sort, Merge Sort, Quick Sort, and others, with clear and concise explanations. This project focuses on learning and applying best practices rather than production-grade performance or optimization.

## Roadmap

- [X] Project Setup
  - [X] README file
  - [X] Git commit message template
  - [X] Ignoring files
  - [X] Makefile
- [ ] GitHub Actions
  - [X] Continuous integration
  - [ ] Code Coverage
  - [ ] Publishing a module
- [ ] Badges
  - [ ] Go Report Card
  - [X] Build Status
  - [X] Test Status
  - [ ] Test Coverage
  - [X] License
  - [ ] Go Doc
- [ ] Implement Sorting Algorithms
  - [ ] Bubble Sort
  - [ ] Selection Sort
  - [ ] Quick Sort
  - [ ] Merge Sort
  - [ ] Heap Sort
- [ ] Benchmarks
- [ ] Cmd App
  - [ ] Specify options with cmd-line flags
  - [ ] Print help text
  - [ ] Print app version
  - [ ] Choose sorting algorithm
  - [ ] Sort in reverse order

## Requirements

Go version 1.18 or higher.

## Make Commands

This project uses a Makefile to automate common development tasks. Here are the available commands:

- `make install`: Install project dependencies
- `make fmt`: Format all source files
- `make lint`: Lint all source files
- `make test`: Run tests
- `make run`: Run the application
- `make build`: Build the application
- `make help`: Display help information about available commands

To enable verbose output for any command, you can use the VERBOSE=1 flag. For example:

```sh
make run VERBOSE=1
```

This will display all the commands being executed.

## Commit Message Template

Run this command in the repository root folder to use the project's commit message template:

```sh
git config commit.template $PWD/.gitmessage
```

This can be useful for ensuring consistent commit message format across the project.

## Contributing

Pull requests are welcome. Please make sure to update tests as appropriate.

Open an issue to discuss feedback.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
