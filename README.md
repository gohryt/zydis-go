## Motivation
There is an [existing library that wraps Zydis](https://github.com/zyantific/zydis-go), but it uses code generation extensively, which doesn't work as well as expected.

This approach is based on different principles:
- Code generation is only used for enums
- The code generation scripts are written in Python for simplicity
- Aim to provide a more Go-way API
