# Sorting Algorithms in Go

This repository contains implementations of standard algorithms like Bubble Sort, Merge Sort, Quick Sort, and others, with clear and concise explanations. This project focuses on learning and applying best practices rather than production-grade performance or optimization.

## Roadmap

- [ ] Project Setup
  - [ ] README file
  - [ ] Git commit message template
  - [ ] Ignoring files
  - [ ] Git hooks for static code analysis
  - [ ] Makefile
- [ ] GitHub Actions
  - [ ] Continuous integration
  - [ ] Code Coverage
  - [ ] Publishing a module
- [ ] Badges
  - [ ] Go Report Card
  - [ ] Test Status
  - [ ] Test Coverage
  - [ ] License
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

## Commit Message Template

To use the project's commit message template:

1. Copy the prepare-commit-msg hook:
   ```
   cp hooks/prepare-commit-msg .git/hooks/
   chmod +x .git/hooks/prepare-commit-msg
   ```

2. Set the commit template:
   ```
   git config commit.template .gitmessage
   ```

These steps will ensure you're using the project's commit message format.

## Contributing

Pull requests are welcome. Please make sure to update tests as appropriate.

Open an issue to discuss feedback.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
