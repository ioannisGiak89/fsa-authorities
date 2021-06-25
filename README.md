
# compare-fsa-ratings

A cli tool to compare the food hygiene rating distributions
of two local authorities written in GO.



## Installation

Clone the project

```bash
  git clone git@github.com:ioannisGiak89/fsa-authorities.git
```

Go to the project directory

```bash
  cd fsa-authorities
```

Get dependencies

```bash
  go get -d -v ./...
```

Install

```bash
  go install -v ./...
```

Run

```bash
fsa-authorities help
```

## Available commands


#### list

Displays a table with all the available authorities alongside with their ID and the FSA scheme they belong to.

#### compare

Compare the food hygiene rating distributions
of two or more authorities

| Flags        | Shorthand   | Description
| ----------- | ----------- | ----------
| --authorityIds      | -a       | A comma separated list with the local authority IDs to compare
| --schemeType   | -s        | The scheme type the authorities belong to

## Running Tests

To run tests, run the following command

```bash
  go test ./...
```

  
## Usage/Examples

```bash
fsa-authorities list
fsa-authorities compare -s fhrs -a 358,359
fsa-authorities compare -s fhis -a 206,227
```
## Future improvements

1. Add some integration tests to actually test the commands
1. Add some more test cases for the distribution calculators
1. Add support for configuration, e.g baseUrl, endpoints etc can be configurable via environment variables or a config command
1. Locally cache responses to avoid calling the API too often  

