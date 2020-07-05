package payment

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"strconv"

	"github.com/athomecomar/athome/backend/checkout/checkoutconf"
	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/storeql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type Card struct {
	Id     uint64 `json:"id,omitempty"`
	UserId uint64 `json:"user_id,omitempty"`

	LastFourDigits uint64 `json:"last_four_digits,omitempty"`
	NumberHash     string `json:"number_hash,omitempty"`
	CVVHash        string `json:"cvv_hash,omitempty"`
	ExpiryMonth    uint64 `json:"expiry_month,omitempty"`
	ExpiryYear     uint64 `json:"expiry_year,omitempty"`

	HolderDNI  uint64 `json:"holder_dni,omitempty"`
	HolderName string `json:"holder_name,omitempty"`
}

func NewCard(ctx context.Context, userId uint64, number uint64, CVV uint64, expiryMonth, expiryYear, dni uint64, name string) (*Card, error) {
	card := &Card{UserId: userId, ExpiryMonth: expiryMonth, ExpiryYear: expiryYear, HolderDNI: dni, HolderName: name}
	var err error
	card.LastFourDigits = number % 1000
	card.CVVHash, err = hash(CVV)
	if err != nil {
		return nil, errors.Wrap(err, "hash cvv")
	}
	card.NumberHash, err = encryptNumber(number)
	if err != nil {
		return nil, errors.Wrap(err, "encrypt number")
	}
	return card, nil
}

func (c *Card) ToPb() *pbcheckout.Card {
	return &pbcheckout.Card{
		LastFourDigits: c.LastFourDigits,
		ExpiryMonth:    c.ExpiryMonth,
		ExpiryYear:     c.ExpiryYear,
		Holder: &pbcheckout.CardHolder{
			Dni:  c.HolderDNI,
			Name: c.HolderName,
		},
	}
}

func NewCardFromPb(ctx context.Context, in *pbcheckout.CardInput, userId uint64) (*Card, error) {
	card := &Card{UserId: userId, ExpiryMonth: in.GetExpiryMonth(), ExpiryYear: in.GetExpiryYear(),
		HolderDNI: in.GetHolder().GetDni(), HolderName: in.GetHolder().GetName(),
	}
	var err error
	card.LastFourDigits = in.GetNumber() % 1000
	card.CVVHash, err = hash(in.GetCvv())
	if err != nil {
		return nil, errors.Wrap(err, "hash cvv")
	}
	card.NumberHash, err = hash(in.GetNumber())
	if err != nil {
		return nil, errors.Wrap(err, "hash number")
	}
	return card, nil
}

func decryptNumber(s []byte) (uint64, error) {
	c, err := aes.NewCipher(checkoutconf.GetNUMBER_SECRET())
	if err != nil {
		return 0, errors.Wrap(err, "aes.NewCipher")
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return 0, errors.Wrap(err, "cipher.NewGCM")
	}
	nonceSize := gcm.NonceSize()
	if len(s) < nonceSize {
		return 0, errors.New("invalid len of given text")
	}
	nonce, s := s[:nonceSize], s[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, s, nil)
	if err != nil {
		return 0, errors.Wrap(err, "gcm.Open")
	}
	num, err := strconv.Atoi(string(plaintext))
	if err != nil {
		return 0, errors.Wrap(err, "strconv.Atoi")
	}
	return uint64(num), nil
}

func encryptNumber(num uint64) (string, error) {
	c, err := aes.NewCipher(checkoutconf.GetNUMBER_SECRET())
	if err != nil {
		return "", errors.Wrap(err, "aes.NewCipher")
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", errors.Wrap(err, "cipher.NewGCM")
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", errors.Wrap(err, "io.ReadFull")
	}
	return string(gcm.Seal(nonce, nonce, []byte(strconv.Itoa(int(num))), nil)), nil
}

func hash(num uint64) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(strconv.Itoa(int(num))), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.Wrap(err, "bcrypt.GenerateFromPassword")
	}
	return string(h), nil
}

func FindUserCards(ctx context.Context, db *sqlx.DB, userId uint64) ([]*Card, error) {
	rows, err := storeql.WhereMany(ctx, db, &Card{}, `user_id=$1`, userId)
	if err != nil {
		return nil, errors.Wrap(err, "WhereMany")
	}
	defer rows.Close()
	var cards []*Card
	for rows.Next() {
		card := &Card{}
		err = rows.StructScan(card)
		if err != nil {
			return nil, errors.Wrap(err, "StructScan")
		}
		cards = append(cards, card)
	}

	return cards, nil
}

func FindCard(ctx context.Context, db *sqlx.DB, oId uint64, userId uint64) (*Card, error) {
	order := &Card{}
	row := storeql.Where(ctx, db, order, `id=$1 AND user_id=$2`, oId, userId)
	err := row.StructScan(order)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return order, nil
}
