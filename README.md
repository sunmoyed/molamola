# Mola Mola

## Dev Setup

### Client

- Install npm and node
  - Recommendation: use nvm to install npm and node.
    - [Install nvm](https://github.com/creationix/nvm#installation)
    - [Install node with nvm](https://github.com/creationix/nvm#usage)
- Go into the web directory: `cd web`
- Install dependencies of this project: `npm i`
- Return to the project root: `cd ..`
- Run development build: `make web-dev`

### Server

- Install golang
  - Do appropriate `GOPATH` setup
- `go get github.com/sunmoyed/molamola`
- `make compile-server`

### Run
- `make server-dev`
  - Your project is serving at [http://localhost:4477](), check it out.
