package ipfs

import (
	"context"

	files "github.com/ipfs/go-ipfs-files"
	icore "github.com/ipfs/interface-go-ipfs-core"
	"github.com/ipfs/interface-go-ipfs-core/path"
)

func Store(ctx context.Context, content string, node icore.CoreAPI) (path.Resolved, error) {
	return node.Unixfs().Add(ctx, files.NewBytesFile([]byte(content)))
}
