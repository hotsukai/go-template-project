# go-template-project

### Setup
```sh
make init # install tools
docker compose up # start database
cd ./cmd & go run main.go # start server
```

### Test
```sh
make gen-test # generate test
make test # test
```

### Architecture
```sh
├── cmd # Entry points
│   └── main.go 
└── pkg
    ├── domain
    │   ├── model # domain model
    │   └── repository
    │       └── user_repository.go # repository interface
    ├── infrastructure # adapters
    │   ├── database
    │       └── persistence # repository implement
    ├── presentation # ui
    │   └── web # web api
    ├── usecase # business logic
    └── util    # utility
```