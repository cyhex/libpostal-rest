# Libpostal REST

## install deps

    export PKG_CONFIG_PATH=/usr/local/lib/pkgconfig
    go get github.com/openvenues/gopostal/expand
    go get github.com/gorilla/mux


## start Service
    go build main.go
    ./main -port=5000


## API Example

Replace <host> with your host

### Parser
`curl -X POST -d '{"query": "100 main st buffalo ny"}' <host>:8080/parser`

** Response **
```
[
  {
    "label": "house_number",
    "value": "100"
  },
  {
    "label": "road",
    "value": "main st"
  },
  {
    "label": "city",
    "value": "buffalo"
  },
  {
    "label": "state",
    "value": "ny"
  }
]
```

### Expand
`curl -X POST -d '{"query": "100 main st buffalo ny"}' <host>:8080/expand`

** Response **
```
[
  "100 main saint buffalo new york",
  "100 main saint buffalo ny",
  "100 main street buffalo new york",
  "100 main street buffalo ny"
]
```
