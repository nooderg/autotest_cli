package authentication

import (
	"autotest-cli/models"
	"crypto/rsa"
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

func init() {
	privateBytes, err := ioutil.ReadFile("./ssl/private.rsa")
	if err != nil {
		log.Fatal("No se pudo leer el archivo privado")
	}
	publicBytes, err := ioutil.ReadFile("./ssl/public.rsa.pub")
	if err != nil {
		log.Fatal("No se pudo leer el archivo publico")
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatal("no se pudo hacer el parse a privateKey")
	}

	publicKey, err = jwt.ParseRSAPublicKeyFromPEM([]byte(publicBytes))
	if err != nil {
		log.Fatal("no se pudo hacer el parse a publicKey")
	}

}

func GenerateJWT(user models.User) string {
	claims := models.Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(), //1 heure
			Issuer:    "Autotest",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	//convertir dans un string base 64
	result, err := token.SignedString(privateKey)
	if err != nil {
		log.Fatal("no se pudo firmar el token")
	}

	return result
}
func Login(name string, password string) string {
	var user models.User
	user.Name = name
	user.Password = password

	if user.Name == "test" && user.Password == "test" {
		user.Password = ""
		token := GenerateJWT(user)
		result := models.ResponseToken{token}
		//convertir le variable à json
		jsonResult, err := json.Marshal(result)

		if err != nil {
			return "Erreur generate json"
		}
		return string(jsonResult)

	}
	return string("Erreur connexion")
}
