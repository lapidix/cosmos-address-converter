# cosmos-address-converter

A Go-based CLI tool for converting Cosmos blockchain AccAddress into its corresponding ValAddress and ConsAddress formats.

---

## Installation

```bash
git clone https://github.com/mingi3442/cosmos-address-converter.git
cd cosmos-address-converter
go build -o cosmos-address-converter
```

## Configuration

Create a `config.toml` file in the project root directory based on the provided `config.toml.example`:

```toml
grpc_url = "localhost:9090"
account_address = "cosmos1abcd..."
```

## Usage

```bash
./cosmos-address-converter
```

Example output:

```
Account address:             cosmos1abcd...
Validator address:           cosmosvaloper1abcd...
Consensus address:           cosmosvalcons1abcd...
```
