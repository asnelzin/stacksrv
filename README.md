# stacksrv

Stacksrv is an implementation of a stack which works through TCP and meets the following requirements:

1. stack should have `POP` and `PUSH` operations
2. POP should block until stack has data
3. stack should work through TCP with protocol:
    * first bit of an incoming message determines the command (`0` for `POP` and `1` for `PUSH`)
    * next `7` bits determines size of a payload (for `PUSH` only)
    * following bytes is a payload.
