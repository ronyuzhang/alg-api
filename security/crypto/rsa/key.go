package RSA

import (
	rsa"crypto/rsa"
	"crypto/rand"
	ckey"security/crypto/key"
	"errors"
	"encoding/json"
//	"fmt"
)


type RSAPrivateKey interface {
	ckey.PrivateKey
	GetRSAPrivateKey() (*rsa.PrivateKey)
	SetRSAPrivateKey(*rsa.PrivateKey)
}


func (rpk *rsaPrivKey)SetRSAPrivateKey(privkey *rsa.PrivateKey){
	rpk.priv = privkey
}

type rsaPrivKey struct {
	priv *rsa.PrivateKey
}


func (rpk *rsaPrivKey)GetRSAPrivateKey() (*rsa.PrivateKey){
	return rpk.priv
}


func (rpk *rsaPrivKey)GetType() ckey.KeyType{
	return ckey.KEYTYPE_RSA_PRIVATE
}

//type rsaMarshalKey struct {
//
//}
func (rpk *rsaPrivKey)ExportKey()([]byte,error){
	return json.Marshal(rpk.priv)
}

func (rpk *rsaPrivKey)ImportKey(key []byte) error{
	rsapriv:=rsa.PrivateKey{}
	err := json.Unmarshal(key,&rsapriv)
	rpk.priv = &rsapriv
	return err
}

type RSAPublicKey interface {
	ckey.PublicKey
	GetRSAPublicKey() (*rsa.PublicKey)
	SetRSAPublicKey(pubkey rsa.PublicKey)
}

type rsaPubKey struct {
	pub *rsa.PublicKey
}

func (rpk *rsaPubKey)SetRSAPublicKey(pubkey rsa.PublicKey){
	rpk.pub = &pubkey
}

func (rpk *rsaPubKey)GetRSAPublicKey() (*rsa.PublicKey){
	return rpk.pub
}


func (rpk *rsaPubKey)GetType() ckey.KeyType{
	return ckey.KEYTYPE_RSA_PUBLIC
}

func (rpk *rsaPubKey)ExportKey()([]byte,error){
	return json.Marshal(rpk.pub)
}

func (rpk *rsaPubKey)ImportKey(key []byte) error{
	rsapub:= rsa.PublicKey{}
	err:=json.Unmarshal(key,&rsapub)
	rpk.pub = &rsapub
	return err
}


type keypairKey struct {
	rsaPriv rsaPrivKey
	rsaPub rsaPubKey
}

func (kpk *keypairKey)GetPublic() (ckey.Key){
	return &(kpk.rsaPub)
}

func (kpk *keypairKey)GetPrivate() (ckey.Key){
	return &(kpk.rsaPriv)
}


func RSAKeyPair(bits int) (ckey.KeyPair,error){
	//为了让公私密钥无关,需要考虑把rsa不该有的指针设置为空
	kp :=keypairKey{}
	rkp,err := rsa.GenerateKey(rand.Reader,bits)
	if(err != nil){
		return nil,err
	}

	kp.rsaPriv.priv = rkp
	pk:=rsa.PublicKey(rkp.PublicKey)
	kp.rsaPub.pub = &pk
	return &kp,err
}

//该函数目前用不上 后续看情况
func BuildKey(keyType ckey.KeyType)(ckey.Key,error){
	switch keyType {
	case ckey.KEYTYPE_RSA_PUBLIC:
		rsaPub := rsaPubKey{}
		return &rsaPub,nil
	case ckey.KEYTYPE_RSA_PRIVATE:
		rsaPriv := rsaPrivKey{}
		return &rsaPriv,nil
	default:
		return nil,errors.New("Error: invalid rsa key type in BuildKey")
	}
}

