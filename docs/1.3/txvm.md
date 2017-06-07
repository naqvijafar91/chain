# Stacks

0. Data stack
1. Alt stack
2. Inputs stack
3. Values stack
4. Outputs stack
5. Conditions stack
6. Nonces stack
7. Anchors stack
8. Transaction header stack
9. Version 1 Inputs Stack
10. Version 1 Values Stack
11. Muxes Stacks
12. Version 1 Outputs Stack
13. Version 1 Conditions Stack
14. Version 1 Nonces Stack
15. Version 1 Anchors Stack

# Types

0. Int64
1. String
2. Tuple

"Boolean" is not a separate type, but rather the two Int64 values `0` (for false) and `1` (for true). Operations that expect booleans fail if the value on top of the stack is not `0` or `1`.

# Operations

## Control flow operations

# Fail

Halts VM execution, returning false.

# PC

Pushes the current program counter (after incrementing for this instruction) to the data stack.

# JumpIf

Pops an integer `destination`, then a boolean `cond` from the data stack. If `cond` is false, do nothing. If `cond` is true, set program counter to `destination`. Fail if `destination` is negative, if `destination` is greater than or equal to the length of the current program.

## Stack operations 

# Roll

Pops an integer `stackid` from the data stack, representing a [stack identifier](#stacks), and pops another integer `n` from the data stack. On the stack identified by `stackid`, moves the `n`th item from the top from its current position to the top of the stack.

Fails if `stackid` does not correspond to a valid stack, or if the stack has fewer than `n + 1` items.

# Bury

Pops an integer `stackid` from the data stack, representing a [stack identifier](#stacks), and pops a number `n` from the data stack. On the stack identified by `stackid`, moves the top item and inserts it at the `n`th-from-top position.

# Depth

Pops an integer `stackid` from the data stack, representing a [stack identifier](#stacks). Counts the number of items on the stack identified by `stackid`, and pushes it to the data stack.

## Data stack operations

### Equal

Pops two items `val1` and `val2` from the data stack. If they have different types, or if either is a tuple, fails execution. If they have the same type: if they are equal, pushes `true` to the stack; otherwise, pushes `false` to the stack.

### Type

Looks at the top item on the data stack. Pushes a number to the stack corresponding to that item's [type](#type).

### Encode

Pops a string or integer `val` from the data stack. Pushes a string to the data stack that, when executed on the VM, would push `val` to the data stack. Fails if `val` is a tuple.

### Len

Pops a string or tuple `val` from the data stack. If `val` is a tuple, pushes the number of fields in that tuple to the data stack. If `val` is a string, pushes the length of that string to the data stack. Fails if `val` is a number.

# Drop

Drops an item from the data stack.

# Dup
	
Pops the top item from the data stack, and pushes two copies of that item to the data stack.

# ToAlt

Pops an item from the data stack and pushes it to the alt stack.

# FromAlt

Pops an item from the alt stack and pushes it to the data stack.

## Tuple operations 

### Tuple

Pops an integer `len` from the data stack. Pops `len` items from the data stack and creates a tuple of length `len` on the data stack.

### Untuple

Pops a tuple `tuple` from the data stack. Pushes each of the fields in `tuple` to the stack in reverse order (so that the 0th item in the tuple ends up on top of the stack).

### Field

Pops an integer `stackid` from the data stack, representing a [stack identifier](#stack-identifier), and pops another integer `i` from the top of the data stack. Looks at the tuple on top of the stack identified by `stackid`, and pushes the item in its `i`th field to the top of the data stack.

Fails if the stack identified by `stackid` is empty or does not have a tuple of at least length `i + 1` on top of it, or if `i` is negative.

## Boolean operations

### Not

Pop a boolean `p` from the stack. If `p` is `true`, push `false`. If `p` is `false`, push `true`.

### And

Pops two booleans `p` and `q` from the stack. If both `p` and `q` are true, push `true`. Otherwise, push `false`.

### Or

Pops two booleans `p` and `q` from the stack. If both `p` and `q` are false, push `false`. Otherwise, push `true`.

## Math operations

### Add

Pops two integers `a` and `b` from the stack, adds them, and pushes their sum `a + b` to the stack.

### Sub

Pops two integers `a` and `b` from the stack, subtracts the second frmo the first, and pushes their difference `a - b` to the stack.

### Mul

Pops two integers `a` and `b` from the stack, multiplies them, and pushes their product `a * b` to the stack.

### Div

Pops two integers `a` and `b` from the stack, divides them truncated toward 0, and pushes their quotient `a / b` to the stack.

### Mod

Pops two integers `a` and `b` from the stack, computes their remainder `a % b`, and pushes it to the stack.

TODO: clarify behavior. Should act like Go.

### LeftShift

Pops two integers `a` and `b` from the stack, shifts `a` to the left by 

### RightShift

### GreaterThan

Pops two numbers `a` and `b` from the stack. If `a` is greater than `b`, pushes `true` to the stack. Otherwise, pushes `false`.

## String operations

### Cat

Pops two strings, `a`, then `b`, from the stack, concatenates them, and pushes the result, `a ++ b` to the stack.

### Slice

Pops two integers, `start`, then `end`, from the stack. Pops a string `str` from the stack. Pushes the string `str[start:end]` (with the first character being the one at index `start`, and the second character being the one at index `end`). Fails if `end` is less than `start`, if `start` is less than 0, or if `end` is greater thanss the length of `str`.

## Bitwise operations

### BitNot

Pops a string `a` from the data stack, inverts its bits, and pushes the result `~a`.

### BitAnd

Pops two strings `a` and `b` from the data stack. Fails if they do not have the same length. Performs a "bitwise and" operation on `a` and `b` and pushes the result `a & b`.

### BitOr

Pops two strings `a` and `b` from the data stack. Fails if they do not have the same length. Performs a "bitwise or" operation on `a` and `b` and pushes the result `a | b`.

### BitXor

Pops two strings `a` and `b` from the data stack. Fails if they do not have the same length. Performs a "bitwise xor" operation on `a` and `b` and pushes the result `a ^ b`.

## Crypto operations

### SHA256

Pops a string `a` from the data stack. Performs a Sha2-256 hash on it, and pushes the result `sha256(a)` to the data stack.

### SHA3

Pops a string `a` from the data stack. Performs a Sha3-256 hash on it, and pushes the result `sha3-256(a)` to the data stack.

### CheckSig

Pops a string `pubKey`, a string `msg`, and a string `sig` from the data stack. Performs an Ed25519 signature check with `pubKey` as the public key, `msg` as the message, and `sig` as the signature. Pushes `true` to the data stack if the signature check succeeded, and `false` otherwise.

TODO: Should we switch order of `pubKey` and `msg`.

### PointAdd

### PointSub

### PointMul

## Entry operations

TODO: add descriptions.

	Cond         = 46 // prog => cond
	Unlock       = 47 // inputid + data => value + cond
	UnlockOutput = 48 // outputid + data => value + cond
	Merge        = 49 // value value => value
	Split        = 50 // value + amount => value value
	ProveRange   = 51 // TODO(kr): review for CA
	ProveValue   = 52 // TODO(kr): review for CA
	ProveAsset   = 53 // TODO(kr): review for CA
	Blind        = 54 // TODO(kr): review for CA
	Lock         = 55 // value + prog => outputid
	Satisfy      = 56 // cond => {}
	Anchor       = 57 // nonce + data => anchor + cond
	Issue        = 58 // anchor + data => value + cond
	IssueCA      = 59 // TODO(kr): review for CA
	Retire       = 60 // valud + refdata => {}

## Legacy operations

	VM1CheckPredicate = 63 // list vm1prog => bool
	VM1Unlock         = 64 // vm1inputid + data => vm1value + cond
	VM1Nonce          = 65 // vm1nonce => vm1anchor + cond
	VM1Issue          = 66 // vm1anchor => vm1value + cond
	VM1Mux            = 67 // entire vm1value stack => vm1mux
	VM1Withdraw       = 68 // vm1mux + amount asset => vm1mux + value

## Extension opcodes

	Nop0    = 69
	Nop1    = 70
	Nop2    = 71
	Nop3    = 72
	Nop4    = 73
	Nop5    = 74
	Nop6    = 75
	Nop7    = 76
	Nop8    = 77
	Private = 78

## Encoding opcodes

### Negate

### Small integers

### Pushdata
