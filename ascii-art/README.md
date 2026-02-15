# Ascii-art

## Project Overview

**Ascii-art** is a Go program that receives a string as input and returns it in a **graphic representation using ASCII characters**. It allows you to display characters by rendering text with ASCII art.  

The program supports letters, numbers, spaces, special characters, and newline characters (`\n`). ASCII art is rendered according to predefined banner templates: shadow.txt, standard.txt and thinketoy.txt.

---

## Objectives

- Learn to manipulate strings and text in Go.
- Work with the **Go file system (fs) API** to load banner files.
- Implement modular code with unit testing.
- Handle multi-line input and special characters in ASCII output.

---

## Instructions

1. **Project Language:** Go  
2. **Coding Practices:** Follow Go best practices.  
3. **Unit Testing:** Test files (`*_test.go`) are included.  
4. **Banner Files:** Provided in the `banners` directory and should **not be modified**:
   - `shadow.txt`
   - `standard.txt`
   - `thinkertoy.txt`  

---

## Banner Format

- Each character has a **height of 8 lines**.  
- Characters are separated by a newline (`\n`).  

Example for `' '`, `'!'`, and `'"'` (dots `.` represent spaces):

```
......
......
......
......
......
......
......
......

...
|.|.
|.|.
|.|.
||.
(_).
....
....

....
(.|.).
.V.V..
......
......
......
......
......

---

```

## Usage

Run the program from the project root:

```bash
# Empty input
go run . "" | cat -e
```

## Single newline
```bash
go run . "\n" | cat -e
```

## Single word
```bash
go run . "Hello\n" | cat -e
```

## Two words
```bash
go run . "Hello There" | cat -e
```

## Mixed letters and numbers
```bash
go run . "1Hello 2There" | cat -e
```

## Example Output
```bash

student$ go run . "Hello\n" | cat -e
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $

  ```                              
 Multiple lines with \n are supported, and an empty line is printed between blocks of ASCII art.

## Testing

Unit tests are provided for each module. Run all tests with:

```bash
go test -v
```

Tests include:

* Rendering empty input

* Rendering input with only newline(s)

* Rendering single and multi-line text

## Correct handling of trailing and double newlines

```bash
File Structure
ascii-art/
├── ascii/                # Main source code
│   ├── fs.go
│   ├── fs_test.go
│   ├── parser.go
│   ├── parser_test.go
│   ├── renderer.go
│   └── renderer_test.go
├── banners/              # Banner templates
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
├── go.mod
├── go.sum
└── main.go               # Entry point
```

## Allowed Packages

Only standard Go packages are allowed.