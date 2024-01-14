package jwt

// func jwsRSApair() {
// 	dbr := redis.Connent()
// 	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
// 	if err != nil {
// 		panic(err)
// 	}

// 	privateBlock := &pem.Block{
// 		Type:  "JWS PRIVATE KEY",
// 		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
// 	}
// 	privateBuffer := &bytes.Buffer{}
// 	pem.Encode(privateBuffer, privateBlock)

// 	publicBlock := &pem.Block{
// 		Type:  "JWS PUBLIC KEY",
// 		Bytes: x509.MarshalPKCS1PublicKey(&privateKey.PublicKey),
// 	}
// 	publicBuffer := &bytes.Buffer{}
// 	pem.Encode(publicBuffer, publicBlock)

// 	dbr.HSet(dbr.Context(), "jws:pair", "private", privateBuffer.Bytes(), "public", publicBuffer.Bytes())
// }

// var jwsPrivateKey *rsa.PrivateKey

// func LoadJWSPrivateKey() *rsa.PrivateKey {
// 	dbr := redis.Connent()
// 	if count := dbr.Exists(dbr.Context(), "jws:pair").Val(); count < 1 {
// 		jwsRSApair()
// 	}
// 	b, _ := dbr.HGet(dbr.Context(), "jws:pair", "private").Bytes()
// 	block, _ := pem.Decode(b)
// 	jwsPrivateKey, _ = x509.ParsePKCS1PrivateKey(block.Bytes)
// 	return jwsPrivateKey
// }

// var jwsPublicKey *rsa.PublicKey

// func LoadJWSPublicKey() *rsa.PublicKey {
// 	dbr := redis.Connent()
// 	b, _ := dbr.HGet(dbr.Context(), "jws:pair", "public").Bytes()
// 	block, _ := pem.Decode(b)
// 	jwsPublicKey, _ = x509.ParsePKCS1PublicKey(block.Bytes)
// 	return jwsPublicKey
// }
// func JWSSign(claim *Claim) string {
// 	if jwsPrivateKey == nil {
// 		panic("load jws private key")
// 	}

// 	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.PS512, Key: jwsPrivateKey}, nil)
// 	if err != nil {
// 		panic(err)
// 	}

// 	object, err := signer.Sign(types.JSON(claim))
// 	if err != nil {
// 		panic(err)
// 	}

// 	bearer, _ := object.CompactSerialize()
// 	return bearer
// }

// func JWSVerify(bearer string) (*Claim, *error4.Error) {
// 	key := jwsPublicKey
// 	if jwsPrivateKey != nil {
// 		key = &jwsPrivateKey.PublicKey
// 	}

// 	object, err4 := jose.ParseSigned(bearer)
// 	if err4 != nil {
// 		return nil, error4.New("S505", err4)
// 	}
// 	output, err4 := object.Verify(key)
// 	if err4 != nil {
// 		return nil, error4.New("S505", err4)
// 	}
// 	cl := &Claim{}
// 	types.STRUCT(output, cl)
// 	return cl, cl.Valid()
// }

// func GetClaim(bearer string) *Claim {
// 	object, err := jose.ParseSigned(bearer)
// 	if err != nil {
// 		return nil
// 	}
// 	output, _ := object.Verify("")
// 	cl := &Claim{}
// 	types.STRUCT(output, cl)
// 	return cl
// }
