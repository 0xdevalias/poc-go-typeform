# poc-go-typeform

PoC reference for accessing the modern (aka 2018+) [Typeform API](https://developer.typeform.com/).

Definitely incomplete, but nothing else seems to even exist in Golang-world for the latest API.

## Usage

```go
import (
	"log"

	"github.com/0xdevalias/poc-typeform/api"
)

def main() {
    c := api.DefaultClient("MY-PERSONAL-ACCESS-TOKEN")

    r, err := c.RetrieveForm("MY-FORM-ID")
    if err != nil {
        log.Fatalf("error: %v", err)
    }

    log.Printf("%#v", r)
}
```

## Alternative Library Options

Looking around.. there just didn't really seem to be any good/stand out options..

* https://golanglibs.com/top?q=typeform
* https://github.com/xsb/typeform-go
    * Last updated 2016, unmaintained
    * Doesn't work with the new API ([Ref](https://github.com/xsb/typeform-go/blob/master/tfio/api.go#L11))
* https://github.com/gagliardetto/go-ask-awesomely
    * Last updated 2017
    * Looked good, but doesn't work with the new API ([Issue](https://github.com/gagliardetto/go-ask-awesomely/issues/1))
* https://github.com/zachgoldstein/gotypeformio
    * Last updated 2015
    * Uses the old v0.3 API ([Ref](https://github.com/zachgoldstein/gotypeformio/blob/fc06f60f14134ae4452cdda73528902d6a0c45f4/api.go#L5))
* https://github.com/levenlabs/go-typeform
    * Last updated 2016
    * Uses the old v0.4 API ([Ref](https://github.com/levenlabs/go-typeform/blob/master/tyapi/api.go#L12))
* https://github.com/Bowbaq/typeform
    * Last updated 2015
    * Uses selenium on the website to make forms.. not the API.
