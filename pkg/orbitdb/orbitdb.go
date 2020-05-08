package orbitdb

import (
	"context"

	orbitdb "berty.tech/go-orbit-db"
	"berty.tech/go-orbit-db/baseorbitdb"
	"berty.tech/go-orbit-db/iface"
	icore "github.com/ipfs/interface-go-ipfs-core"
)

func NewOrbitDB(ctx context.Context, ipfs icore.CoreAPI) (iface.OrbitDB, error) {
	return orbitdb.NewOrbitDB(ctx, ipfs, &baseorbitdb.NewOrbitDBOptions{})
}

func NewDocs(ctx context.Context, orbitdb iface.OrbitDB) {
	//return orbitdb.Doc
}
