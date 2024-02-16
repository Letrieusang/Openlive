package utils

import (
	"api-openlive/common"
	"api-openlive/config"
	usermodel "api-openlive/module/user/model"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	cryptocommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/twinj/uuid"
	"time"
)

// TokenValid method
func TokenValid(bearerToken string) (*jwt.Token, error) {
	token, err := verifyToken(bearerToken)
	if err != nil {
		if token != nil {
			return token, err
		}
		return nil, err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return nil, fmt.Errorf("Unauthorized")
	}
	return token, nil
}

//verifyToken verify token
func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.GetConfig().GetString("server.private_key")), nil
	})
	if err != nil {
		return token, fmt.Errorf("Unauthorized")
	}
	return token, nil
}

func CreateToken(userId int64, userType string) (*usermodel.Credential, error) {
	var res usermodel.Credential

	token, err := getToken(userId, userType, common.EXPIRED_TOKEN)
	if err != nil {
		return nil, err
	}
	res.Token = token

	refreshToken, err := getToken(userId, userType, common.EXPIRED_REFRESH_TOKEN)
	if err != nil {
		return nil, err
	}
	res.RefreshToken = refreshToken
	return &res, nil
}

func getToken(userId int64, userType string, t int64) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = userId
	atClaims["user_type"] = userType
	uuid := uuid.NewV4().String()
	atClaims["uuid"] = uuid
	tToken := time.Now().Add(time.Duration(t) * time.Minute)
	atClaims["exp"] = tToken.Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(config.GetConfig().GetString("server.private_key")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func VerifySig(req usermodel.UserLoginRequest, nonce string) bool {
	fromAddr := cryptocommon.HexToAddress(req.WalletAddress)

	sig := hexutil.MustDecode(req.HashData)
	if len(sig) < 65 {
		return false
	}
	if sig[64] != 27 && sig[64] != 28 {
		return false
	}
	sig[64] -= 27
	msg := []byte(nonce)
	pubKey, err := crypto.SigToPub(signHash(msg), sig)
	if err != nil {
		return false
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	return fromAddr == recoveredAddr
}

func signHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}
