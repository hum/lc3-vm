# Little Computer 3
This repository contains a simple implementation of the [Little Computer 3](https://en.wikipedia.org/wiki/Little_Computer_3). With its simple [instruction set](https://www.jmeiners.com/lc3-vm/supplies/lc3-isa.pdf) the VM is fairly straight-forward and the code speaks for itself.

There's a few examples of assembly projects in the `/asm/` folder.

### Hello World
```asm
.ORIG x3000                         ; Address of the first instruction
HELLO_STR .STRINGZ "Hello, World!"  ; Define "Hello, World!" string
LEA R0, HELLO_STR                   ; Load address of HELLO_STR to R0
PUTs                                ; Print R0 to STDOUT
HALT                                ; Halt program
.END
```

To actually run the code, it has to be built into an **.obj** file. Those are included in the `/obj/` folder.

# TODO
- [ ] Keyboard I/O

Currently the VM only handles a simple hello world program, because it doesn't implement a keyboard input handler.

- [ ] Logging

This section needs the biggest improvement. The only thing that should go into STDOUT is the actual print syscalls made by the programs. Everything else should be a debug log only available under a special flag.

## Run
Clone the repository
```bash
> git clone https://github.com/hum/lc3-vm
> cd lc3-vm
```

Build the binary
```bash
# or "make run" for the hello world example
> make build
```

Run the programs from `/obj/`
```bash
> ./bin/lc3_vm ./obj/2048.obj
```

## LC3 Tools
If you wish to convert your own `.asm` files into `.obj` files then you can do so by downloading the codebase from [here](https://highered.mheducation.com/sites/0072467509/student_view0/lc-3_simulator.html).
```bash
> cd lc3tools
> ./configure --installdir /path/to/loc/
> make
> make install
```
Now you will have `./lc3as` in your `/path/to/loc/` which you can use to generate the `.obj` files

## Resources
- [Justin Meiners' blog](https://www.jmeiners.com/lc3-vm)
- [Instructure Set Document](https://www.jmeiners.com/lc3-vm/supplies/lc3-isa.pdf)
- [This project doc by BYU](https://students.cs.byu.edu/~cs345ta/labs/P4-Virtual%20Memory.html)
