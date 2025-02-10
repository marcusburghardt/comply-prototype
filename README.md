# Comply-Prototype

NOTE: This repository was an experiment that is now finished and therefore is archived.

The outcomes of this experiment were incorporated at https://github.com/complytime/complytime

## Overview

**Comply-Prototype** is a CLI tool used to manage OSCAL artifacts and policies generated from these OSCAL artifacts. The supported policies can be extended with plugins for specific technologies used both to scan and remediate a system. **Comply-Prototype** makes the underlying technology used by plugins transparent for the users, allowing a consistent experience while keeping a very flexible adoption of different technologies. **Comply-Prototype** communicates with plugins via gRPC, providing a standard and consistent communication mechanism while also allowing plugins developers to choose their preferred languages to contribute. This project is structured to allow modular development, ease of packaging, and maintainability.

## Project Structure

```
comply-prototype/
├── cmd/                # Main entry points for CLI and plugin binaries
│ └── main.go           # Main file for CLI
├── proto/              # gRPC related code
│ ├── scan_grpc.pb.go   # Auto generated content from scan.proto
│ ├── scan.pb.go        # Auto generated content from scan.proto
│ └── scan.proto        # Protocol Buffer definition for scan
├── .gitignore
├── go.mod              # Go module file
└── README.md           # This file
```

## Installation

### Prerequisites

- **Go** version 1.20 or higher
- **Protocol Buffers** compiler (for gRPC)
- **Make** (optional, for using the `Makefile` if included)

### Clone the repository

```bash
git clone https://github.com/marcusburghardt/comply-prototype.git
cd comply-prototype
```

## Build Instructions
To compile the CLI and plugin:

```bash
go build -o comply-prototype .
```

## Running
Start the plugin server:

```bash
./openscap-prototype
```

In another terminal, run the CLI to connect to the server and send the scan action:

```bash
./comply-prototype scan
```

## Development
### gRPC Protocol Buffers
1. Edit the .proto file(s) in proto as needed.
2. Regenerate the gRPC code:

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative scan.proto
```

### Testing - NOT YET
Tests are organized within each package. Run tests using:

```bash
go test ./...
```

### Packaging as RPM - NOT YET
To build an RPM package, use the spec file in build/rpm:

```bash
# Example using rpmbuild
rpmbuild -ba build/rpm/specfile.spec
```

## Contributing
Please open an issue or submit a pull request for any contributions or improvements.

## License
MIT License
