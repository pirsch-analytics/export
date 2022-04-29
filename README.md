# Pirsch Analytics Export

Export statistics from Pirsch by using the API.

## Usage

1. create a client ID + secret on the [Pirsch dashboard developer settings page](https://pirsch.io)
2. create a `config.toml` right next to the `pirsch-export`/`pirsch-export.exe` executable
3. paste the configuration below and adjust the settings:

```toml
client_id = "your-client-id"
client_secret = "your-client-secret"
hostname = "example.com"

export = [
    "conversion_goals_day"
]
from = 2022-01-01
to = 2022-04-29
```

4. run the program by double-clicking it or execute it in a terminal

## License

MIT
