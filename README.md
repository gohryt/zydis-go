## Motivation
There is an [existing library that wraps Zydis](https://github.com/zyantific/zydis-go), but it uses code generation extensively, which doesn't work as well as expected.

My approach is based on different principles:
- Code generation is only used for enums
- The code generation scripts are written in Python for simplicity
- I aim to provide a more Go-way API
