package middleware

import (
	"backend/pkg/schema"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	opmiddleware "github.com/oapi-codegen/echo-middleware"
)

type JWSValidator interface {
	ValidateJWS(jws string) (*jwt.Token, error)
}

type JWSValidatorImpl struct {
	PrivateKey *ecdsa.PrivateKey
}

// なにこれ
const JWTClaimsContextKey = "jwt_claims"

// 型アサーション用のダミー変数
var _ JWSValidator = JWSValidatorImpl{}

// AuthenticationFunc を返す
func NewAuthenticator(v JWSValidator) openapi3filter.AuthenticationFunc {
	return func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
		return Authenticate(v, ctx, input)
	}
}

// middlewareを返す
// TODO JWSValidatorをここで初期化する
func Validator() echo.MiddlewareFunc {
	swagger, err := schema.GetSwagger()
	if err != nil {
		panic(fmt.Errorf("loading spec : %w", err))
	}
	v := JWSValidatorImpl{}

	validator := opmiddleware.OapiRequestValidatorWithOptions(swagger,
		&opmiddleware.Options{
			Options: openapi3filter.Options{
				AuthenticationFunc: NewAuthenticator(v),
			},
		})

	return echo.MiddlewareFunc(validator)
}

var (
	ErrNoAuthHeader      = errors.New("authorization header is missing")
	ErrInvalidAuthHeader = errors.New("authorization header is malformed")
	ErrClaimsInvalid     = errors.New("provided claims do not match expected scopes")
)

// リクエストヘッダから JWS を取得
func GetJWSFromRequest(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeader
	}
	const prefix = "Bearer "
	if !strings.HasPrefix(authHeader, prefix) {
		return "", ErrInvalidAuthHeader
	}
	return strings.TrimPrefix(authHeader, prefix), nil
}

// JWSを検証してJWTを返す
func (v JWSValidatorImpl) ValidateJWS(jws string) (*jwt.Token, error) {
	token, err := jwt.Parse(jws, func(token *jwt.Token) (any, error) {
		// 署名アルゴリズムの検証
		// パースされたトークンを受け取り署名の検証に必要な暗号化キーを返す
		// 暗号あるごリズムがHMAC、ES256どれがよいか調べる
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// 署名キーを返す（例: HMACの場合は共有秘密鍵）
		return []byte("your-256-bit-secret"), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodES256.Alg()}))

	if err != nil {
		return nil, fmt.Errorf("failed to parse/validate JWS: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}

// func (f JWSValidatorImpl) SignToken(t jwt.Token) ([]byte, error) {

// }

// func (f JWSValidatorImpl) CreateJWSWithClaims(claims jwt.Claims) (string, error) {
// }

// 認証処理
func Authenticate(v JWSValidator, ctx context.Context, input *openapi3filter.AuthenticationInput) error {
	// debug用に無条件で通す
	return nil
	if input.SecuritySchemeName != "BearerAuth" {
		return fmt.Errorf("security scheme %s != BearerAuth", input.SecuritySchemeName)
	}

	jws, err := GetJWSFromRequest(input.RequestValidationInput.Request)
	if err != nil {
		return err
	}

	// 検証
	token, err := v.ValidateJWS(jws)
	if err != nil {
		return err
	}

	err = CheckTokenClaims(input.Scopes, *token)
	if err != nil {
		return err
	}

	eCtx := opmiddleware.GetEchoContext(ctx)
	eCtx.Set(JWTClaimsContextKey, token)

	return nil
}

// クレームが全て含まれているか確認
func CheckTokenClaims(expectedClaims []string, t jwt.Token) error {
	claims := t.Claims
	var _ jwt.Claims
	claimMap, ok := claims.(jwt.MapClaims)
	if !ok {
		return fmt.Errorf("invalid claims type")
	}

	for _, e := range expectedClaims {
		if claimMap[e] != true {
			log.Printf("missing claim: %s", e)
			return ErrClaimsInvalid
		}
	}

	return nil

}
