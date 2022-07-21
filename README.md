# degorater package

This is a package to decorate of the text in the terminal.  
It expected improve visibility and user experience in interactive cli-tools.

## Installation
This package can be installed with the `go get` command:
```
go get github.com/ebiy0rom0/degorater
```

## How to Use
Simple and basic usage:
```
degorater.Print("message you want to decorate", nil)
```
For result:
```
+------------------------------+
| message you want to decorate |
+------------------------------+
```

Too long messages will be automatically wrapped.
```
degorater.Print("This is too long message. Want a automatically wrapped", nil)
```
For result:
```
+------------------------------+
| This is too long message. Wan|
| t a automatically wrapped    |
+------------------------------+
```

choices can output.
```
degorater.Print("Which one will you choose?", []string{"choice1", "choice2"})
```
For result:
```
+------------------------------+
| Which one will you choose?   |
|                              |
| 1. choice1                   |
| 2. choice2                   |
+------------------------------+
```

## Options
| Name | Type | Default | Description |
|------|------|---------|-------------|
| width | `int` | `terminal width` \| `30` | |
| sp | `string` | `_` |  |
| corner | `string` | `+` |  |
| top | `string` | `-` |  |
| side | `string` | `|` |  |
