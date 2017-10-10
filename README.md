# Mola Mola

## Dev Setup

### Server (Do this FIRST!)

- Install golang
  - Do appropriate `GOPATH` setup
  - Test with: `echo $GOPATH`
- If you DID NOT yet clone the repo:
    - `go get github.com/sunmoyed/molamola`
- If you DID already clone the repo:
    - `mkdir -p $GOPATH/src/github.com/sunmoyed`
    - `mv /path/to/molamola $GOPATH/src/github.com/sunmoyed/molamola`
- `make server-dev`
  - Your project is serving at [http://localhost:4477]()
  - Username/Password mola/mola
- Golang dependency management is done with [dep](https://github.com/golang/dep)

### Client

- Install npm and node
  - Recommendation: use nvm to install npm and node.
    - [Install nvm](https://github.com/creationix/nvm#installation)
    - [Install node with nvm](https://github.com/creationix/nvm#usage)
- Go into the web directory: `cd web`
- Install dependencies of this project: `npm i`
- Return to the project root: `cd ..`
- Run development build: `make web-dev`
