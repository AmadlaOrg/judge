# judge
🧑‍⚖️ Judge verifies that system settings meet required configurations and resource specifications 🧑‍⚖️

## How it works
Judge communicates with different applications and scripts to validate. For this reason there is an entity named judge
that indicate what script/application to load and how to execute it (similar to a plugin system). That does its own 
validation of the environment.

It communicates to `hery` to get the entities. Once it hits the `judge` entity in the `_meta` then it pulls the script 
or application to the plugins directory and then runs it. And prints the result in a standard cli table.

### Help
```bash
judge --help|-h
```

### List evaluators
```bash
judge --evaluators|-e
```

### Run
```bash
judge run <evaluator name>
```