package jwt

// func IsAllow(tags types.Tags, service, method string) *error4.Error {

// 	serialized := tags.String("authorization")
// 	if strings.Contains(serialized, "Bearer") {
// 		serialized = serialized[len("Bearer"):]
// 	}

// 	cl, err4 := JWSVerify(serialized)
// 	if err4 != nil {
// 		return err4
// 	}

// 	data := JWTData{}
// 	types.STRUCT(cl.Data, &data)
// 	if data.UUID == "" {
// 		return error4.New("A010", "UUID is invalid")
// 	}

// 	tags["UUID"] = data.UUID
// 	tags["Refresh"] = data.Refresh

// 	data.Roles = append(data.Roles, "guest")
// 	// permission check

// 	return nil
// }
