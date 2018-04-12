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

## Changing JSON Library

Since we use [Resty](https://github.com/go-resty/resty) under the hood, [it's really easy](https://github.com/go-resty/resty/issues/76#issuecomment-314015250) to switch the JSON library used for marshalling/unmarshalling. For example:

Using [json-iterator](https://github.com/json-iterator/go) for fast and efficient en/decoding:

```go
import (
    jsoniter "github.com/json-iterator/go"
)

c := api.DefaultClient("SOME-API-TOKEN")

var json = jsoniter.ConfigCompatibleWithStandardLibrary
c.RestyClient().JSONMarshal = json.Marshal
c.RestyClient().JSONUnmarshal = json.Unmarshal
```

Using [koki/json](https://github.com/koki/json) to log unhandled keys (good for debugging/API changes):

```go
import (
  "log"

    "github.com/koki/json"
    "github.com/koki/json/jsonutil"
)

func main() {
    c := api.DefaultClient("SOME-API-TOKEN")
    c.RestyClient().JSONUnmarshal = json.Marshal
    c.RestyClient().JSONUnmarshal = KokiUnmarshal
}

func KokiUnmarshal(data []byte, v interface{}) error {
    obj := make(map[string]interface{})

    err := json.Unmarshal(data, &obj)
    if err != nil {
        return err
  }

  err = jsonutil.UnmarshalMap(obj, &v)
  if err != nil {
    return err
  }

  extraneousPaths, err := jsonutil.ExtraneousFieldPaths(obj, v)
  if err != nil {
    return err
  }
  if len(extraneousPaths) > 0 {
    log.Printf("Extra JSON fields found: %v", len(extraneousPaths))
    for _, path := range extraneousPaths {
      log.Printf("%#v", path)
    }
  }

  return nil
}
```

## Alternative Golang Typesafe Client Options

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
