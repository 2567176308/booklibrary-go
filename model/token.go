package model

import (
	// "crypto/hmac"
	"crypto/rand"
	// "crypto/sha256"
	"encoding/base64"
	// "encoding/json"
	// "errors"
	// "strings"
	"time"

	"gorm.io/gorm"
)

type Token struct {
	gorm.Model 
	UserId uint	`json:"user_id"`
	Value string `json:"token_Value"`
	ExpiresAt time.Time
}

func (table Token) TableName() string {
	return "token"
}

//创建token前函数（自动执行）
func (t *Token) BeforeCreate(tx *gorm.DB) error {
    t.ExpiresAt = time.Now().Add(time.Hour) // 设置token过期时间为1小时
    return nil
}


// 生成 token 的函数
// func GenerateToken(secretKey string, data *Token) (string, error) {
//     // 将 token 数据转换为 JSON 格式
//     jsonData, err := json.Marshal(data)
//     if err != nil {
//         return "", err
//     }

//     // 使用 HMAC-SHA256 算法对 token 数据进行签名
//     mac := hmac.New(sha256.New, []byte(secretKey))
//     mac.Write(jsonData)
//     signature := mac.Sum(nil)

//     // 将签名和 token 数据拼接成一个字符串
//     token := base64.StdEncoding.EncodeToString(signature) + "." + base64.StdEncoding.EncodeToString(jsonData)

//     return token, nil
// }

// // 解析 token 的函数
// func ParseToken(secretKey string, token string) (*Token, error) {
//     // 将签名和 token 数据拆分开来
//     parts := strings.Split(token, ".")
//     if len(parts) != 2 {
//         return nil, errors.New("invalid token format")
//     }

//     // 解码签名和 token 数据
//     signature, err := base64.StdEncoding.DecodeString(parts[0])
//     if err != nil {
//         return nil, err
//     }

//     jsonData, err := base64.StdEncoding.DecodeString(parts[1])
//     if err != nil {
//         return nil, err
//     }

//     // 验证签名，确保 token 数据没有被篡改
//     mac := hmac.New(sha256.New, []byte(secretKey))
//     mac.Write(jsonData)
//     expectedSignature := mac.Sum(nil)
//     if !hmac.Equal(signature, expectedSignature) {
//         return nil, errors.New("invalid token signature")
//     }

//     // 将 JSON 数据解析为 TokenData 结构体
//     var data Token
//     err = json.Unmarshal(jsonData, &data)
//     if err != nil {
//         return nil, err
//     }

//     // 验证 token 是否已经过期
//     if time.Now().After(data.ExpiresAt) {
//         return nil, errors.New("token has expired")
//     }

//     return &data, nil
// }


//生成指定长度的随机字符串
func GenerateRandomString(length int) string {
    bytes := make([]byte, length)
    if _, err := rand.Read(bytes); err != nil {
        return ""
    }
    return base64.URLEncoding.EncodeToString(bytes)[:length]
}
