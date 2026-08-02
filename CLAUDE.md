# CLAUDE.md

## AI Skills

Follow the practices defined in `~/Projects/SiteNetSoft/ai-skills/`:
- `dev-practices/golang/` — Go style, error handling, functions, testing, linting
- `dev-practices/git/` — Git authorship rules, multi-repo workspace patterns

## Project Overview

Judge is a CLI tool for validation and auditing within the Amadla ecosystem. It discovers `judge-*` plugins on PATH and delegates validation tasks to them. Each plugin validates a specific aspect of system configuration (network, applications, security, etc.).

**Ecosystem context:** Part of the Amadla tool pipeline (`raise` -> `lay` -> `enjoin` -> `weaver` -> `waiter`). Judge validates entities and system state independently of the main pipeline. Entity schemas (no "Entity" prefix) live in `Entities/<Type>/`.

## Build Commands

```bash
make build    # Build for current platform
make test     # Run tests
make clean    # Remove build artifacts
```

## Architecture

**UNIX Plugin Protocol:**
```
judge run --from <plugin> -f <data>    →  judge-<plugin> judge -f <data>
judge plugins                          →  scans PATH for judge-* binaries
```

**Plugin Protocol (judge-* binaries):**
- `info` subcommand → JSON metadata (name, version, engine, description, supports)
- `judge -f <data>` subcommand → validates input, outputs results
- Exit codes: 0=pass, 1=fail, 2=usage error
- Data to stdout, diagnostics to stderr
