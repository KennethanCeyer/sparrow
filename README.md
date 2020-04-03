<h1 align="center">sparrow</h1>
<p align="center">:bird: Golang env based configure management module which is supported yaml, toml, json</p>

## Getting Started

### What is sparrow

sparrow is configure managing tool on Golang. When you need to separate configure file as file (toml, yml, json). This module easily read and parse them to pass the content to Golang code.

And sparrow supports environment isolation, When you use multiple stage for production, alpha and local. sparrow will overwrite specific env configuration to general configuration. 

## Quick Start

### Installation

In your terminal

```bash
go get github.com/KennethanCeyer/sparrow
```

In your code

```go
import "github.com/KennethanCeyer/sparrow"

var config map[string]interface{}

err := sparrow.ReadFile("./your_config_path.yml", &config)
```

### Example


> YAML file
```yaml
appName: 'sparrow'
appVersion: '1.0.0'
debug: true

```

> JSON file
```json
{
  "appName": "sparrow",
  "appVersion": "1.0.0",
  "debug": true
}
```

> TOML file
```toml
appName = "sparrow"
appVersion = "1.0.0"
debug = true
```

```go
import (
	"github.com/KennethanCeyer/sparrow/sparrow"
	"log"
)

type Config struct {
	AppName    string
	AppVersion string
	Debug      bool
}

func main() {
var ymlConfig Config
	var jsonConfig Config
	var tomlConfig Config
	var err error

	err = sparrow.ReadFile("./config.yml", &ymlConfig)
	if err != nil {
		return err
	}

	err = sparrow.ReadFile("./config.json", &jsonConfig)
	if err != nil {
		return err
	}

	err = sparrow.ReadFile("./config.toml", &tomlConfig)
	if err != nil {
		return err
	}

	log.Println("ymlConfig", ymlConfig)
	log.Println("jsonConfig", jsonConfig)
	log.Println("tomlConfig", jsonConfig)

	return nil
}
```

**OUTPUT**

```bash
2019/08/12 16:54:02 ymlConfig {sparrow 1.0.0 true}
2019/08/12 16:54:02 jsonConfig {sparrow 1.0.0 true}
2019/08/12 16:54:02 tomlConfig {sparrow 1.0.0 true}
```

For your more information, Please check [the example](./example)

## Roadmap

### Configuration parser

- [x] YAML
- [x] JSON
- [x] TOML

### Documentation

- [x] Example
- [ ] API Guide
