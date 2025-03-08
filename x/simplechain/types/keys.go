package types

const (
	// ModuleName defines the module name
	ModuleName = "simplechain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_simplechain"
)

var (
	ParamsKey = []byte("p_simplechain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
