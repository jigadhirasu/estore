package jwt

type JWTData struct {
	UUID     string   `json:",omitempty"` // 使用者編號
	Refresh  string   `json:",omitempty"` // 重刷用金鑰
	Business string   `json:",omitempty"` // 商場編號
	Store    string   `json:",omitempty"` // 商店編號
	Sales    string   `json:",omitempty"` // 銷售編號
	Group    string   `json:",omitempty"` // 會員群
	Roles    []string `json:",omitempty"` // 權限角色
}
