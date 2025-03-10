package types

const (
	// ModuleName defines the module name
	ModuleName = "kudos"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_kudos"

	// KudosKey defines the key to store the value
	KudosKey = "Kudos/value/"

	// KudosCountKey defines the key to store the count
	KudosCountKey = "Kudos/count/"
)

var (
	ParamsKey = []byte("p_kudos")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
