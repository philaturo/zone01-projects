# Textools - Go Text Transformation Pipeline

TextTools is a Go-based text processing utility that tokenizes input text and applies a transformation pipeline including:

-Case transformations
-Hexadecimal and binary number conversions
-Punctuation and normalization
-Quote normalization
-Article correction ('a' -> 'an')

The project is designed to be **modular**, **testdriven**, and **easy to extend**

---

##  Features

### 1: Case Transformations
Supported commands:
- `(up)` – uppercase
- `(low)` – lowercase
- `(cap)` – capitalize
- `(cmd, N)` – apply to the previous **N words**

Example: hello world (cap, 2) → Hello World

---

### 2: Hex & Binary Conversion

- `(hex)` converts hexadecimal numbers
- `(bin)` converts binary numbers

Example: 42 (hex) → 66 , 10 (bin) → 2

---

### 3: Punctuation Normalization

- Removes extra spaces before punctuation
- Collapses repeated punctuation

Example: Hello , world ! → Hello, world!, Wait . . . → Wait... 

---

### 4: Quote Normalization

- Removes unnecessary spacing inside quotes

Example: ' I am the best ' → 'I am the best'

---

### 5: Article Correction

Automatically fixes incorrect articles:

Example: a apple → an , a amazing rock → an amazing rock

---

## Testing

All transformations are **fully tested** using Go’s testing framework.

Run all tests:
```bash
go test ./...

Sample output:

ok   go-reloaded/texttools/tokenizer
ok   go-reloaded/texttools/transform

Usage

Run the program with an input and an output file: 
go run ./texttools sample.txt results.txt

Requirements

-Go 1.20+

