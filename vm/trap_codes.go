package vm

// TRAP CODES
const (
	TRAP_GETC  = 0x20 // Get char from keyboard
	TRAP_OUT   = 0x21 // Output a char
	TRAP_PUTS  = 0x22 // Output a string
	TRAP_IN    = 0x23 // Get char from keyboard, echo to terminal
	TRAP_PUTSP = 0x24 // Output a byte string
	TRAP_HALT  = 0x25 // Halt
)
