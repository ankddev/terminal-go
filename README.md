# terminal-go

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A Go library for working with ANSI/VT terminal sequences.

## Installation

```bash
go get github.com/ankddev/terminal-go
```

## Usage

```go
package main

import (
    "fmt"
    term "github.com/ankddev/terminal-go"
)

func main() {
    // Move cursor to position (5,10)
    fmt.Print(term.CursorPosition(5, 10))
    
    // Set text color to red
    fmt.Print(term.SetTextColor(1))
    fmt.Println("This text is red")
    
    // Reset all attributes
    fmt.Print(term.ResetAllAttributes)
    
    // Clear screen
    fmt.Print(term.EraseInDisplay)
    
    // Save cursor position
    fmt.Print(term.SaveCursorPosition)
    
    // Move cursor and restore position
    fmt.Print(term.CursorForward(10))
    fmt.Println("This text is red")
}
```

## Features

- Cursor movement and positioning
- Text colors and attributes
- Screen clearing and line manipulation
- Scrolling and margins
- Window manipulation
- Character sets
- And more...

## Documentation

For detailed documentation of all available functions and constants, please see:
- [GoDoc](https://pkg.go.dev/github.com/ankddev/terminal-go)

## Building from Source

```bash
git clone https://github.com/ankddev/terminal-go.git
cd terminal-go
go build
```

## Testing

To run tests:
```bash
go test -v
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

**ANKDDEV**

## See Also

- [codewars-api-rs](https://github.com/ankddev/codewars-api-rs) - Rust library for Codewars API
- [conemu-progressbar-go](https://github.com/ankddev/conemu-progressbar-go) - Progress bar for ConEmu for Go
- [envfetch](https://github.com/ankddev/envfetch) - Lightweight crossplatform CLI tool for working with environment variables
- [zapret-discord-youtube](https://github.com/ankddev/zapret-discord-youtube) - Zapret build for Windows for fixing Discord and YouTube in Russia or othher services
## Support

If you have any questions or suggestions, please open an issue on GitHub.
