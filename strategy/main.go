package main

import "fmt"

type PasswordProtector struct {
	username      string
	passwordName  string
	hashAlgorithm HashAlgorithm
}

func (p *PasswordProtector) SetHashAlgorithm(hash HashAlgorithm) {
	p.hashAlgorithm = hash
}

func (p *PasswordProtector) Hash() {
	p.hashAlgorithm.Hash(p)

}

type HashAlgorithm interface {
	Hash(p *PasswordProtector)
}

type SHA struct{}

func (SHA) Hash(p *PasswordProtector) {
	fmt.Printf("Hashing using SHA for %s.\n", p.passwordName)
}

type MD5 struct{}

func (MD5) Hash(p *PasswordProtector) {
	fmt.Printf("Hashing using MD5 for %s.\n", p.passwordName)
}

func NewPasswordProtector(username string, password string, hash HashAlgorithm) *PasswordProtector {
	return &PasswordProtector{username, password, hash}
}

func main() {

	sha := &SHA{}

	md5 := &MD5{}

	pwdProtector := NewPasswordProtector("zero", "me@ultirequiem.com", sha)

	pwdProtector.Hash()

	pwdProtector.SetHashAlgorithm(md5)

	pwdProtector.Hash()

}
