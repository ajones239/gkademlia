package gkademlia

type Peer struct {
    ID string `json:"id"`
    Host string `json:"host"`
    Port int `json:"port"`
    LastSeen int64 `json:"last_seen"`
}
    

