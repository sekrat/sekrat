Sekrat is an embedded key/value store for secrets. These secrets are stored in a Warehouse (Memory, Filesystem, Amazon S3, etc), and they are encrypted/decrypted with a Crypter (Passthrough, AES, etc).

[![GoDoc](https://godoc.org/github.com/sekrat/serkat?status.svg)]

## Installation ##

```
go get github.com/sekrat/sekrat
```

## Usage ##

To use `sekrat`, you basically instantiate a `sekrat.Manager` with a `sekrat.Warehouse` and a `sekrat.Crypter`, then use that manager to `Put` and `Get` your secrets:

```go
package main

import (
  "fmt"
  
  "github.com/sekrat/sekrat"
)

func main() {
  // Create a sekrat.Warehouse instance
  warehouse := sekrat.NewMemoryWarehouse()

  // Create a sekrat.Crypter instance
  crypter := sekrat.NewPassthroughCrypter()

  // Create a secret manager
  confidant := sekrat.New(warehouse, crypter)

  // Store a secret
  err := confidant.Put("an identifier", "an encryption key", []byte("This is a secret."))
  if err != nil {
    panic(err.Error())
  }

  // Retrieve a secret
  secret, err := confidant.Get("an identifier", "an encryption key")
  if err != nil {
    panic(err.Error())
  }

  // Get the list of all secret IDs
  fmt.Println("Known secrets:")

  for _, id := range confidant.IDs() {
    fmt.Println("\tID:", id)
  }

  fmt.Printf("\nThe secret was '%s'\n", string(secret))

}
```

## Warehouses ##

The `sekrat.Warehouse` interface describes an object that knows how to store ID-indexed data, then retrieve that same data via its ID.

A reference implementation, `MemoryWarehouse`, is included in the base package. As the name implies, it uses an in-memory `map` to store the data handed to it. That's only really useful in initial testing and development, so it would be a very good idea to either create your own Warehouse implementation or maybe peruse [Warehouse Implementations](https://github.com/sekrat/sekrat/wiki/Warehouse-Implementations).

## Crypters ##

The `sekrat.Crypter` interface describes an object that knows how to encrypt and decrypt data with a provided key.

A reference implementation, `PassthroughCrypter`, is included in the base package. This crypter does not actually do any encrypting or decrypting, nor does it ever return an error. Rather, it just returns the exact data that it was given. That's only really useful in initial testing and development, so it would be a very good idea to either create your own Crypter implementation or mayper peruse [Crypter Implementations](https://github.com/sekrat/sekrat/wiki/Crypter-Implementations).

## Contributing ##

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request against the "develop" branch

## History ##

* v1.0.0 - Initial Release

## License ##

Sekrat is released under the Apache 2.0 license. See [LICENSE](https://github.com/sekrat/sekrat/blob/master/LICSENSE)
