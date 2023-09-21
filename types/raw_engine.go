package types

type RawEngineReq struct {
	ID     string
	Op     string
	Params []byte
}

type RawEngineResp struct {
	ID   string
	Data []byte
}
