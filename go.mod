module github.com/mentos1386/ipfs-cloud

go 1.14

require (
	berty.tech/go-orbit-db v1.3.1 // indirect
	github.com/ProtonMail/gopenpgp/v2 v2.0.1-0.20200414132903-599adb6b2d6b
	github.com/gotk3/gotk3 v0.4.0
	github.com/ipfs/go-ipfs v0.5.0-rc2
	github.com/ipfs/go-ipfs-config v0.5.2
	github.com/ipfs/go-ipfs-files v0.0.8
	github.com/ipfs/interface-go-ipfs-core v0.2.6
	github.com/libp2p/go-libp2p-core v0.5.1
	github.com/libp2p/go-libp2p-peerstore v0.2.3
	github.com/multiformats/go-multiaddr v0.2.1
	github.com/spf13/cobra v1.0.0
	golang.org/x/crypto v0.0.0-20200406173513-056763e48d71
)

// We have to use this version as it introduces dependencies for ipfs
//  and later commits brake compatibility with ProtonMail/gopenpgp/v2
replace golang.org/x/crypto => github.com/ProtonMail/crypto v1.0.1-0.20200414132716-142ca7810b13
