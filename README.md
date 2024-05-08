# emoji

A basic wrapper around the [Unicode.org emoji data files](http://unicode.org/Public/emoji/).

## Import

```go
import "libdb.so/go-emoji"
```

## Documentation

The only function that matters:

```go
// IsEmoji returns true if the specified rune has the (single-character)
// Emoji property in the latest Emoji version, false otherwise
func IsEmoji(r rune) bool
```

**Go Reference**: [pkg.go.dev/libdb.so/go-emoji](https://pkg.go.dev/libdb.so/go-emoji)
