# athome

[![Build Status](https://travis-ci.com/athomecomar/athome.svg?token=e2HswtC3ixASrSJvAGtB&branch=master)](https://travis-ci.com/athomecomar/athome)

## Structure 


### Protocol buffers

- pb/: holds all protocol buffers related things.
- pb/<pkg>.proto: proto declarations.

- ./backend/<svc>/pb<svc>/: holds compiled pb code.

- pb/js/: holds the compiled code from the protocol buffers to js.
- pb/js/pb_<pkg>.js: holds the compiled code to use the <pkg> pb.

