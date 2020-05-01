package ipfs

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"log"

	icore "github.com/ipfs/interface-go-ipfs-core"
	"github.com/ipfs/interface-go-ipfs-core/options"
	"github.com/ipfs/interface-go-ipfs-core/path"
)

type Node struct {
	Links []Link
	Data  string
}

type Link struct {
	Name, Hash string
	Size       uint64
}

type Object struct {
	Hash  string `json:"Hash,omitempty"`
	Links []Link `json:"Links,omitempty"`
}

func Store(content string, node icore.CoreAPI) (path.Resolved, error) {
	ctx := context.Background()

	fileNode := &Node{
		Data: base64.StdEncoding.EncodeToString([]byte(content)),
	}

	fileNodeJson, err := json.Marshal(fileNode)
	if err != nil {
		return nil, err
	}

	log.Printf("%+v\n", fileNode)

	return node.Object().Put(ctx, bytes.NewReader(fileNodeJson),
		options.Object.InputEnc("json"),
		options.Object.DataType("base64"),
		options.Object.Pin(true))
}
