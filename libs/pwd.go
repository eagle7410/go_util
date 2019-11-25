package lib

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	mathRand "math/rand"
	"time"

	// @docs https://ru.wikipedia.org/wiki/Argon2
	"golang.org/x/crypto/argon2"
)

const (
	PwdConverterSaltLen = 16
	//Memory parameter specifies the size of the memory in KiB. Memory=64*1024 sets the memory cost to ~64 MB.
	//Is 64*1024
	PwdConverterMemory      = 65536
	PwdConverterIterations  = 3
	PwdConverterParallelism = 2
	PwdConverterKeyLen      = 32

	PwdGeneCharsDigits            = "0123456789"
	PwdGeneCharsSpecials          = "~=+%^*/()[]{}/!@#$?|"
	PwdGeneCharsUp                = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	PwdGeneCharsDown              = "abcdefghijklmnopqrstuvwxyz"
	PwdGeneCharsMinDefault uint32 = 16
	PwdGeneCharsMin        uint32 = 4
)

type PasswordConverter struct {
	Hash []byte
	Salt []byte
}

func (i *PasswordConverter) getHashFromPassword(password *string) (hash []byte) {
	return argon2.IDKey(
		[]byte(*password), i.Salt,
		PwdConverterIterations, PwdConverterMemory, PwdConverterParallelism, PwdConverterKeyLen)
}

func (i *PasswordConverter) GeneSalt() (err error) {
	i.Salt, err = GenerateRandomBytes(PwdConverterSaltLen)
	return err
}

func (i *PasswordConverter) GeneHashFromPassword(password *string) {
	i.Hash = i.getHashFromPassword(password)
}

func (i *PasswordConverter) IsValidPass(password *string) (isValid bool) {
	return bytes.Equal(i.getHashFromPassword(password), i.Hash)
}

func (i *PasswordConverter) GetSaltHashBase64Encode() (salt, hash string) {
	return base64.RawStdEncoding.EncodeToString(i.Salt), base64.RawStdEncoding.EncodeToString(i.Hash)
}

func (i *PasswordConverter) SetSaltPassBase64Decode(saltBase64, hashBase64 *string) (err error) {

	if i.Salt, err = base64.RawStdEncoding.DecodeString(*saltBase64); err != nil {
		return fmt.Errorf("error decode salt: %v", err)
	}

	if i.Hash, err = base64.RawStdEncoding.DecodeString(*hashBase64); err != nil {
		return fmt.Errorf("error decode hash: %v", err)
	}

	return err
}

func (i *PasswordConverter) SetSaltPass(salt, hash *[]byte) {
	i.Salt = *salt
	i.Hash = *hash
}

func (i *PasswordConverter) GetSaltHash() (salt, hash []byte) {
	return i.Salt, i.Hash
}

type PasswordGeneOptions struct {
	Up, Down, Digits, Specials string
	Length                     uint32
}

func (i *PasswordGeneOptions) Validate() {
	if i.Length == 0 {
		i.Length = PwdGeneCharsMinDefault
	}

	if i.Length < PwdGeneCharsMin {
		panic("Min length must be 4 char. Because use 4 group: Up, Down, Digit, Special")
	}

	if len(i.Up) == 0 {
		i.Up = PwdGeneCharsUp
	}

	if len(i.Down) == 0 {
		i.Down = PwdGeneCharsDown
	}

	if len(i.Specials) == 0 {
		i.Specials = PwdGeneCharsSpecials
	}

	if len(i.Digits) == 0 {
		i.Digits = PwdGeneCharsDigits
	}
}

func (i *PasswordGeneOptions) getFullString() (str string) {
	return i.Digits + i.Up + i.Specials + i.Down
}

func (i *PasswordGeneOptions) GetRndDown() byte {
	return i.Down[mathRand.Intn(len(i.Down))]
}

func (i *PasswordGeneOptions) GetRndUp() byte {
	return i.Up[mathRand.Intn(len(i.Up))]
}

func (i *PasswordGeneOptions) GetRndSpecial() byte {
	return i.Specials[mathRand.Intn(len(i.Specials))]
}

func (i *PasswordGeneOptions) GetRndDigit() byte {
	return i.Digits[mathRand.Intn(len(i.Digits))]
}

func GenePassword(opt *PasswordGeneOptions) (pass []byte) {

	opt.Validate()

	mathRand.Seed(time.Now().UnixNano())

	full := opt.getFullString()
	fullLen := len(full)

	buf := make([]byte, opt.Length)

	// Get from each group
	buf[0] = opt.GetRndDigit()
	buf[1] = opt.GetRndDown()
	buf[2] = opt.GetRndSpecial()
	buf[3] = opt.GetRndUp()

	// Gene
	for i := uint32(4); i < opt.Length; i++ {
		buf[i] = full[mathRand.Intn(fullLen)]
	}

	// Recombine
	mathRand.Shuffle(len(buf), func(i, j int) { buf[i], buf[j] = buf[j], buf[i] })

	return buf
}

func GenerateRandomBytes(len uint32) (arrByte []byte, err error) {
	arrByte = make([]byte, len)
	_, err = rand.Read(arrByte)

	return arrByte, err
}
