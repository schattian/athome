# athome

## Structure 


### Protocol buffers

- pb/: holds all protocol buffers related things.
- pb/<pkg>.proto: proto declarations.

- pb/go/: holds compiled pb code as subpkgs.
- pb/go/pb<pkg>/*.go: holds the compiled code from the protocol buffers to go.

- pb/js/: holds the compiled code from the protocol buffers to js.
- pb/js/pb_<pkg>.js: holds the compiled code to use the <pkg> pb.

