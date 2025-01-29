<img src=".assets/judge.jpg" alt="Justice image" style="width: 400px;" align="right">

# judge 🧑‍⚖️
🧑‍⚖️ Judge verifies that system settings meet required configurations and resource specifications 🧑‍⚖️

## How it works
Judge communicates with different applications and scripts to validate. For this reason there is an entity named judge
that indicate what script/application to load and how to execute it (similar to a plugin system). That does its own
validation of the environment.

It communicates to `hery` to get the entities. Once it hits the `judge` entity in the `_meta` then it pulls the script
or application to the plugins directory and then runs it. And prints the result in a standard cli table.

## Auditor
The auditor only gets the content to make the evaluation. Judge can handle the full entity to extract the entities the
different auditor need.

Example of what an auditor can receive (it can be YAML or JSON):
```yaml
- _entity: github.com/AmadlaOrg/Entity@v1.0.0
  _body:
    name: htop
    which: $PATH
    exec:
      command: htop
      version: htop --version
      help: htop --help
    lay:
      name: htop
```

OR

```json
[
  {
    "_entity": "github.com/AmadlaOrg/EntityApplication@v1.0.0",
    "_body": {
      "name": "htop",
      "which": "$PATH",
      "exec": {
        "command": "htop",
        "version": "htop --version",
        "help": "htop --help"
      },
      "lay": {
        "name": "htop"
      }
    }
  }
]
```

This format permits the streaming of the data. This makes it less heavy on memory.

### Help
```bash
judge --help|-h
```

### List auditor
```bash
judge --auditor|-e
```

### Run
```bash
judge run <evaluator name>
```

## ©️ Copyright
- "<a rel="noopener noreferrer" href="https://www.flickr.com/photos/37667416@N04/3680735931">[Iustitia]</a>" by <a rel="noopener noreferrer" href="https://www.flickr.com/photos/37667416@N04">Biblioteca Rector Machado y Nuñez</a> is marked with <a rel="noopener noreferrer" href="https://creativecommons.org/publicdomain/mark/1.0/?ref=openverse">Public Domain Mark 1.0 <img src="https://mirrors.creativecommons.org/presskit/icons/pd.svg" style="height: 1em; margin-right: 0.125em; display: inline;" /></a>.

## :scroll: License

The license for the code and documentation can be found in the [LICENSE](./LICENSE) file.

---

Made in Québec 🏴󠁣󠁡󠁱󠁣󠁿, Canada 🇨🇦!
