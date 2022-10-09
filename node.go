package gkademlia

import (
    "gkademlia/pkg/kademlia"
)

type Node struct {
    Peer
    Buckets [kademlia.M][kademlia.K]Peer
    PeerCount uint64
    Data KData
}
