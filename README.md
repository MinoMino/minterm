# minterm

Just a small library to help manipulate the terminal in Go.
I needed something like Python's `shutil.get_terminal_size()`,
so I wrote this. As of writing all it can do is get the terminal
size on Windows, but it'll do more soon.

# Usage

Import it:
```go
import "github.com/MinoMino/minterm"
```

Terminal size:
```go
columns, rows, err := minterm.TerminalSize()
```

# License

MIT. See `LICENSE` for details.
