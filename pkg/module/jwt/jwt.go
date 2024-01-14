package jwt

// func NewClaim(data JWTData) *Claim {
// 	cl := &Claim{
// 		Iat:  time.Now().Unix(),
// 		Exp:  time.Now().AddDate(0, 0, 10).Unix(),
// 		Iss:  "rex",
// 		Sub:  uuid.Gen().Hex(),
// 		Jti:  uuid.New(),
// 		Data: types.JSON(data),
// 	}

// 	dbr := redis.Connent()
// 	// jwtKey := fmt.Sprintf("jwt:%s:%s", data.UUID, cl.Sub)
// 	// if err := dbr.Set(dbr.Context(), jwtKey, 1, time.Hour).Err(); err != nil {
// 	// 	panic(err)
// 	// }
// 	freshKey := fmt.Sprintf("jwt:refresh:%s", data.Refresh)
// 	if err := dbr.Set(dbr.Context(), freshKey, data.UUID, time.Hour*24).Err(); err != nil {
// 		panic(err)
// 	}
// 	return cl
// }

// type Claim struct {
// 	Iat  int64       `json:"iat,omitempty"` // 頒發時間
// 	Nbf  int64       `json:"nbf,omitempty"` // 生效時間
// 	Exp  int64       `json:"exp,omitempty"` // 到期時間
// 	Iss  string      `json:"iss,omitempty"` // 頒發者
// 	Sub  string      `json:"sub,omitempty"` // 主題
// 	Jti  string      `json:"jti,omitempty"` // token編號
// 	Data types.Bytes `json:"data,omitempty"`
// }

// // default time.Now().Unix()
// func (cl *Claim) WithIat(iat int64) *Claim {
// 	cl.Iat = iat
// 	return cl
// }

// // default 0
// func (cl *Claim) WithNbf(nbf int64) *Claim {
// 	cl.Nbf = nbf
// 	return cl
// }

// // default 3h
// func (cl *Claim) WithExp(exp int64) *Claim {
// 	cl.Exp = exp
// 	return cl
// }

// // default ding4pro
// func (cl *Claim) WithIss(iss string) *Claim {
// 	cl.Iss = iss
// 	return cl
// }

// // default uuid.New()
// func (cl *Claim) WithJti(jti string) *Claim {
// 	cl.Jti = jti
// 	return cl
// }

// func (cl Claim) Valid() *error4.Error {
// 	if cl.Exp > 0 && cl.Exp < time.Now().Unix() {
// 		return error4.New("A000", "過期了重新要一個吧")
// 	}
// 	if time.Now().Unix() < cl.Nbf {
// 		return error4.New("A001", "還沒到生效時間再等等吧")
// 	}
// 	if cl.Iss != "rex" {
// 		return error4.New("A002", "這不是我發的,別亂喔")
// 	}

// 	// key := fmt.Sprintf("jwt:%s", cl.Sub)
// 	// dbr := redis.Connent()
// 	// if exists := dbr.Exists(dbr.Context(), key).Val(); exists < 1 {
// 	// 	return error4.New("S602", "過期了重新要一個吧")
// 	// }
// 	// dbr.Expire(dbr.Context(), key, time.Hour*12)
// 	// // data裡面的內容個應用程式自己檢查
// 	return nil
// }
