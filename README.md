<img src=".assets/judge.jpg" alt="Justice image" style="width: 400px;" align="right">

# `judge`

Judge is a CLI tool for validation and auditing. It discovers `judge-*` plugins on PATH and delegates validation to the appropriate plugin.

## Plugins

| Plugin | Purpose |
|--------|---------|
| `judge-application` | Validates applications/binaries are installed and accessible |
| `judge-network` | Validates network connectivity, DNS, port accessibility |

## Usage

```bash
# List discovered plugins
judge plugins

# Run validation using a specific plugin
judge run --from network -f checks.yaml

# Run application validation
judge run --from application -f requirements.yaml
```

## License

The license for the code and documentation can be found in the [LICENSE](./LICENSE) file.

---

Made in Quebec, Canada!
