# Contributing Guide

Thank you for your interest in the project! We welcome your contributions.

## How to Contribute

### Reporting Bugs

If you found a bug:

1. Check if it hasn't been reported in [Issues](https://github.com/Fsh10/GO-SELL-BOT/issues)
2. If not found, create a new issue with description:
   - What happened
   - What was expected
   - Steps to reproduce
   - Go version and OS
   - Logs (if applicable)

### Suggesting Enhancements

1. Create an issue describing the proposed enhancement
2. Discuss the proposal with maintainers
3. After approval, create a pull request

### Pull Requests

1. **Fork the repository**
2. **Create a branch** for your feature:
   ```bash
   git checkout -b feature/amazing-feature
   ```
3. **Follow code style**:
   - Use `go fmt` for formatting
   - Run linter: `make lint`
   - Write clear comments
   - Follow existing naming conventions

4. **Write tests** for new functionality

5. **Make sure all tests pass**:
   ```bash
   make test
   ```

6. **Commit changes**:
   ```bash
   git commit -m "Add amazing feature"
   ```
   Use clear commit messages in English.

7. **Push to branch**:
   ```bash
   git push origin feature/amazing-feature
   ```

8. **Open Pull Request**

## Code Standards

### Go Style Guide

- Follow [Effective Go](https://go.dev/doc/effective_go)
- Use `golangci-lint` for code checking
- Format code with `go fmt`
- Use `goimports` for imports

### Comments

- Comment exported functions and types
- Use English for comments
- Write clear and informative comments

### Testing

- Cover new code with tests
- Use table tests where appropriate
- Test edge cases

### Commits

- Use clear commit messages
- One commit = one logical change
- Use prefixes:
  - `feat:` - new feature
  - `fix:` - bug fix
  - `docs:` - documentation changes
  - `refactor:` - refactoring
  - `test:` - adding tests
  - `chore:` - dependency updates, configuration

## Review Process

1. All PRs go through code review
2. Maintainers may request changes
3. After approval, PR will be merged

## Questions?

If you have questions, create an issue or contact maintainers.

Thank you for your contribution! 🎉
