package rsa

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"github.com/nine-monsters/go-hutool/hutool-crypto/asymmetric"
	"strings"
)

func NewRsa(publicKey, privateKey string) *asymmetric.Rsa {
	rsaObj := &asymmetric.Rsa{
		privateKey: privateKey,
		publicKey:  publicKey,
	}
	rsaObj.init() //初始化，如果存在公钥私钥，将其解析
	return rsaObj
}

func (r *asymmetric.Rsa) init() {
	if r.privateKey != "" {
		//将私钥解码
		block, _ := pem.Decode([]byte(r.privateKey))
		//pkcs1   //判断是否包含 BEGIN RSA 字符串,这个是由下面生成的时候定义的
		if strings.Index(r.privateKey, "BEGIN RSA") > 0 {
			//解析私钥
			r.rsaPrivateKey, _ = x509.ParsePKCS1PrivateKey(block.Bytes)
		} else { //pkcs8
			//解析私钥
			privateKey, _ := x509.ParsePKCS8PrivateKey(block.Bytes)
			//转换格式  类型断言
			r.rsaPrivateKey = privateKey.(*rsa.PrivateKey)
		}
	}

	if r.publicKey != "" {
		//将公钥解码 解析 转换格式
		block, _ := pem.Decode([]byte(r.publicKey))
		publicKey, _ := x509.ParsePKIXPublicKey(block.Bytes)
		r.rsaPublicKey = publicKey.(*rsa.PublicKey)
	}
}

// Encrypt 加密
func (r *asymmetric.Rsa) Encrypt(data []byte) ([]byte, error) {
	// blockLength = 密钥长度 = 一次能加密的明文长度
	// "/8" 将bit转为bytes
	// "-11" 为 PKCS#1 建议的 padding 占用了 11 个字节
	blockLength := r.rsaPublicKey.N.BitLen()/8 - 11
	//如果明文长度不大于密钥长度，可以直接加密
	if len(data) <= blockLength {
		//对明文进行加密
		return rsa.EncryptPKCS1v15(rand.Reader, r.rsaPublicKey, []byte(data))
	}
	//否则分段加密
	//创建一个新的缓冲区
	buffer := bytes.NewBufferString("")
	pages := len(data) / blockLength //切分为多少块
	//循环加密
	for i := 0; i <= pages; i++ {
		start := i * blockLength
		end := (i + 1) * blockLength
		if i == pages { //最后一页的判断
			if start == len(data) {
				continue
			}
			end = len(data)
		}
		//分段加密
		chunk, err := rsa.EncryptPKCS1v15(rand.Reader, r.rsaPublicKey, data[start:end])
		if err != nil {
			return nil, err
		}
		//写入缓冲区
		buffer.Write(chunk)
	}
	//读取缓冲区内容并返回，即返回加密结果
	return buffer.Bytes(), nil
}

// Decrypt 解密
func (r *asymmetric.Rsa) Decrypt(data []byte) ([]byte, error) {
	//加密后的密文长度=密钥长度。如果密文长度大于密钥长度，说明密文非一次加密形成
	//1、获取密钥长度
	blockLength := r.rsaPublicKey.N.BitLen() / 8
	if len(data) <= blockLength { //一次形成的密文直接解密
		return rsa.DecryptPKCS1v15(rand.Reader, r.rsaPrivateKey, data)
	}

	buffer := bytes.NewBufferString("")
	pages := len(data) / blockLength
	for i := 0; i <= pages; i++ { //循环解密
		start := i * blockLength
		end := (i + 1) * blockLength
		if i == pages {
			if start == len(data) {
				continue
			}
			end = len(data)
		}
		chunk, err := rsa.DecryptPKCS1v15(rand.Reader, r.rsaPrivateKey, data[start:end])
		if err != nil {
			return nil, err
		}
		buffer.Write(chunk)
	}
	return buffer.Bytes(), nil
}

// Sign 签名
func (r *asymmetric.Rsa) Sign(data []byte, sHash crypto.Hash) ([]byte, error) {
	hash := sHash.New()
	hash.Write(data)
	sign, err := rsa.SignPKCS1v15(rand.Reader, r.rsaPrivateKey, sHash, hash.Sum(nil))
	if err != nil {
		return nil, err
	}
	return sign, nil
}

// Verify 验签
func (r *asymmetric.Rsa) Verify(data []byte, sign []byte, sHash crypto.Hash) bool {
	h := sHash.New()
	h.Write(data)
	return rsa.VerifyPKCS1v15(r.rsaPublicKey, sHash, h.Sum(nil), sign) == nil
}

// CreateKeys 生成pkcs1 格式的公钥私钥
func (r *asymmetric.Rsa) CreateKeys(keyLength int) (privateKey, publicKey string) {
	//根据 随机源 与 指定位数，生成密钥对。rand.Reader = 密码强大的伪随机生成器的全球共享实例
	rsaPrivateKey, err := rsa.GenerateKey(rand.Reader, keyLength)
	if err != nil {
		return
	}
	//编码私钥
	privateKey = string(pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY", //自定义类型
		Bytes: x509.MarshalPKCS1PrivateKey(rsaPrivateKey),
	}))
	//编码公钥
	objPkix, err := x509.MarshalPKIXPublicKey(&rsaPrivateKey.PublicKey)
	if err != nil {
		return
	}
	publicKey = string(pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: objPkix,
	}))
	return
}

// CreatePkcs8Keys 生成pkcs8 格式公钥私钥
func (r *asymmetric.Rsa) CreatePkcs8Keys(keyLength int) (privateKey, publicKey string) {
	rsaPrivateKey, err := rsa.GenerateKey(rand.Reader, keyLength)
	if err != nil {
		return
	}
	//两种方式
	//一：1、生成pkcs1格式的密钥 2、将其转化为pkcs8格式的密钥（使用自定义方法）
	//  objPkcs1 := x509.MarshalPKCS1PrivateKey(rsaPrivateKey)
	//  objPkcs8 := r.Pkcs1ToPkcs8(objPkcs1)
	//二：直接使用 x509 包 MarshalPKCS8PrivateKey 生成pkcs8密钥
	objPkcs8, _ := x509.MarshalPKCS8PrivateKey(rsaPrivateKey)
	//fmt.Println("对比两种结果",strings.Compare(string(objPkcs8),string(rr)))

	privateKey = string(pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: objPkcs8,
	}))

	objPkix, err := x509.MarshalPKIXPublicKey(&rsaPrivateKey.PublicKey)
	if err != nil {
		return
	}

	publicKey = string(pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: objPkix,
	}))
	return
}

// Pkcs1ToPkcs8 将pkcs1 转到 pkcs8 自定义
func (r *asymmetric.Rsa) Pkcs1ToPkcs8(key []byte) []byte {
	info := struct {
		Version             int
		PrivateKeyAlgorithm []asn1.ObjectIdentifier
		PrivateKey          []byte
	}{}
	info.Version = 0
	info.PrivateKeyAlgorithm = make([]asn1.ObjectIdentifier, 1)
	info.PrivateKeyAlgorithm[0] = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 1}
	info.PrivateKey = key
	k, _ := asn1.Marshal(info)
	return k
}
