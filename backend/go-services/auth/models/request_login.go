// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    requestLogin, err := UnmarshalRequestLogin(bytes)
//    bytes, err = requestLogin.Marshal()

package models

import "encoding/json"

func UnmarshalRequestLogin(data []byte) (RequestLogin, error) {
	var r RequestLogin
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RequestLogin) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type RequestLogin struct {
	// 设备ID
	DeviceID string `json:"deviceId"`
	// 用户名/用户ID
	ID string `json:"id"`
	// 随机数
	Nonce string `json:"nonce"`
	// 密码的md5摘要
	Password string `json:"password"`
	// 时间戳
	Timestamp string `json:"timestamp"`
}
