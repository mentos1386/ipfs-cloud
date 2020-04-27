package state

import (
	"sync"

	gopenpgp "github.com/ProtonMail/gopenpgp/v2/crypto"
)

type State struct {
	OpenPGPPrivateKeyPath string
	OpenPGPDecryptedKey   *gopenpgp.Key
}

var instance *State
var once sync.Once

func GetState() *State {
	once.Do(func() {
		instance = &State{}
	})
	return instance
}
