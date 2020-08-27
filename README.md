# Unique ID Generator

## Motivation
I had an experience with UUID, and it's perfect for various use cases. For some number of projects I have a few additional requirements which UUID doesn't meet:
1. A unique string should achieve security requirements
1. Should be readable for code generations
1. String representation should be shorter than UUID

How requirements achieved:
1. By default, used `crypto/rand` implementation, but it extendable and can be simply changed to any other implementation (`math/rand` or any custom one)
1. By default, used base32, which is readable (because of the list of supported symbols), but can be simply replaced to any another implementation
1. By default, used base32, as a result, 20 random bytes (160 bits) are encoded to 32 symbols (UUID has 122 random bits in 32-36 symbols)

## Quick Architecture Review
The main endpoint which should be used from external applications or libraries is located in the file `/facade.go`.
The file can be found functions to create default or custom providers.
`Provider` is a struct that can generate new ID (random number of bytes) or parse from a string, decode from ints, initialize from bytes ID.
`ID` is a struct that can encode or dump the id (random number of bytes) to a string, bytes, ints.
`Encoder`, `reader`, `randomizer` are internal entities which use for:
* `Encoder` - interface to encode and decode an array of bytes to a string
* `Randomizer` - interface to generate a random array of bytes, currently implemented `encoder` only from `crypto/rand` which has a secure random implementation
* `Reader` - struct implements `io.Reader` to generate rand from packages `crypto/rand` and can be implemented from `math/rand` 

## Example

### Default Provider
In the file `/facade.go` can be found a list of functions for the project.
The default provider has random bytes size `20`, randomizer from `crypto/rand`, the encoder is `base32` with disabled paddings.

#### Generate random string
```go
p := NewProvider()
id, _ := p.Generate()
id.String() // DQSU2FIFIFLDEOOWKJJ66JT5JUXYJDAH
```

#### Parse ID from previously generated string
```go
p := NewProvider()
id, _ := p.Parse("DQSU2FIFIFLDEOOWKJJ66JT5JUXYJDAH")
id.Uint32() // []uint32{357377308, 844513541, 1397937721, 1300047599, 126649391},
```
