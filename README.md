# ABSTRACT FORTH MACHINE 

Simon Kirkby
tigger@interthingy.com
20170604

This is a first cut of a abstract forth machine. An attempt to make a forth machine that is not bound to a bit width or memory model.
At the moment it is just a list processor. 


completer.go - auto completion functions for readline

define.go - colon and semi colon defns 

dictionary.go - the searchable dictonary struct

extra.go - extra words for visiablilty

fth.go - the main functions for the Abstract Forth machine

init.go - base setup words

interface.go - external interface functions

machine.go - the forth and options struct

rstack.go - return stack (need to work towards normal stack impl)

stack.go - the data stack

word.go - the word struct and interface (needs more work)
