# Schema

Using [JSON Schema](https://json-schema.org/) it verifies

Example:
```json
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "title": "System Configuration Rules",
  "type": "object",
  "properties": {
    "cpu_cores": {
      "type": "number",
      "minimum": 4
    },
    "ram_gb": {
      "type": "number",
      "minimum": 8
    },
    "kernel_params": {
      "type": "object",
      "properties": {
        "net_ipv4_ip_forward": {
          "type": "number",
          "const": 1
        },
        "vm.swappiness": {
          "type": "number",
          "maximum": 10
        }
      }
    }
  },
  "required": ["cpu_cores", "ram_gb", "kernel_params"]
}
```