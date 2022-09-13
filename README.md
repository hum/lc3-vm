# Little Computer 3
This repository contains a simple implementation of the [Little Computer 3](https://en.wikipedia.org/wiki/Little_Computer_3). With its simple [instruction set](https://www.jmeiners.com/lc3-vm/supplies/lc3-isa.pdf) the VM is fairly straight-forward and the code speaks for itself.

There's a few examples of assembly projects in the `/asm/` folder

### Hello World
```asm
.ORIG x3000                         ; Address of the first instruction
HELLO_STR .STRINGZ "Hello, World!"  ; Define "Hello, World!" string
LEA R0, HELLO_STR                   ; Load address of HELLO_STR to R0
PUTs                                ; Print R0 to STDOUT
HALT                                ; Halt program
.END
```
