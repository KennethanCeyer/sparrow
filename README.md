<h1 align="center">tit</h1>
<p align="center">:bird: Golang env based configure management module which is supported yaml, toml, json</p>

## Getting Started

### What is tit

tit is configure managing tool on Golang. When you need to separate configure file as file (toml, yml, json). This module easily read and parse them to pass the content to Golang code.

And tit supports environment isolation, When you use multiple stage for production, alpha and local. tit will overwrite specific env configuration to general configuration. 

### Quick Start

```go
import (
	"log"
	"https://github.com/KennethanCeyer/tit"
)

type Config struct {
	AppName     string
	AppVersion  string
	Debug       string
}

func main() {
	config := new(Config)
    tit.ReadFile("./config/config.yml", config)
	log.Println(config)
}
```

**OUTPUT**

```go
&{tit 1.0.0 verbose}
```

## Roadmap

### Configuration parser

- [ ] YAML
- [ ] JSON
- [ ] TOML

### Documentation

- [ ] Example
- [ ] API Guide
