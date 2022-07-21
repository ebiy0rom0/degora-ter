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
Result is:
```
+------------------------------+
| message you want to decorate |
+------------------------------+
```

Too long messages will be automatically wrapped.  
Note: The word breaks when wrapped.[#]()
```
degorater.Print("This is too long message. Want a automatically wrapped.", nil)
```
Result is:
```
+------------------------------+
| This is too long message. Wan|
| t a automatically wrapped.   |
+------------------------------+
```

It can be lined up choices.
```
degorater.Print("Which one will you choose?", []string{"choice1", "choice2"})
```
Result is:
```
+------------------------------+
| Which one will you choose?   |
|                              |
| 1. choice1                   |
| 2. choice2                   |
+------------------------------+
```

And if there are more than four choices, they may be lined up side by side.  
This depends on the string length of each column.
```
degorater.Print("Which one will you choose?", []string{"choice1", "choice2", "choice3", "choice4"})
```
Result is:
```
+------------------------------+
| Which one will you choose?   |
|                              |
| 1. choice1     2. choice2    |
| 3. choice3     4. choice4    |
+------------------------------+
```

## Features
| Name | Type | Default | Description |
|------|------|---------|-------------|
| width | `int` | `terminal width` \| `30` | wip |
| sp | `string` | `_` | wip |
| corner | `string` | `+` | wip |
| side | `string` | `\|` | wip |
| top | `string` | `-` | wip |

## License
MIT License