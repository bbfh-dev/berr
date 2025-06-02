# [B]etter [err]or handling in Go

Simple library that focuses on improving the errors in two ways:

- [📜 Error trace](#-error-trace)
- [🖨️ Error context](#%EF%B8%8F-error-context)

> [!NOTE]
> Both `berr.New()` and `berr.WithContext()` return `error` interface, meaning that no function signature needs to change to use this library.

## 📜 Error trace

Allow errors to be traced using:

```go
// Will result in `nil` if `err == nil`, otherwise in `prefix this error: <err>`
berr.New("prefix this error", err)
```

You can print a pretty message using `berr.Expand(error)` or `berr.Fexpand(io.Writer, error)`.
Example:

```
[Error]
another example: Hello World!: Yet another error!

[Trace]
1. "another example"
2. "Hello World!"
3. "Yet another error!"
```

## 🖨️ Error context

You can add context to an error, which is a list of variables to be included with the error:

```go
berr.WithContext(
    "Hello World!",
    err,
    "c", "Something Something",
    "d", map[string]bool{"x": true, "y": false, "z": true},
),
```

You can print a pretty message using `berr.Expand(error)` or `berr.Fexpand(io.Writer, error)`.
Example:

```
[Error]
another example: Hello World!: Yet another error!

[Trace]
1. "another example"
└── a: 123
└── b: 456
2. "Hello World!"
└── c: "Something Something"
└── d: map[string]bool{"x":true, "y":false, "z":true}
3. "Yet another error!"
```
