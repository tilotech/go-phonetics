# go-phonetics

> A go phonetics algorithm library

Implemented algorithms:

* [Soundex](http://en.wikipedia.org/wiki/Soundex)
* [Metaphone](http://en.wikipedia.org/wiki/Metaphone)

## Usage

Get the latest tagged version:

```
go get github.com/tilotech/go-phonetics
```

Example:

```
import "github.com/tilotech/go-phonetics"

func main() {
  metaphoneEncoded := phonetics.EncodeMetaphone("Hello World!")
  soundexEncoded := phonetics.EncodeSoundex("Hello World!")
  fmt.Println(metaphoneEncoded, soundexEncoded)
}
```

More information can be found in the [godocs](https://pkg.go.dev/github.com/tilotech/go-phonetics).
