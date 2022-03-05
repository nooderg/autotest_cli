package authentication

import (
	"autotest-cli/models"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
		log.Fatal("unable to read the file key private")
	}
	publicBytes, err := ioutil.ReadFile("./ssl/public.rsa.pub")
	if err != nil {
		log.Fatal("unable to read the file key public")
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatal("Error executing command parse in file key private")
	}

	publicKey, err = jwt.ParseRSAPublicKeyFromPEM([]byte(publicBytes))
	if err != nil {
		log.Fatal("Error executing command parse in file key public")
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
		log.Fatal("could not sign token")
	}

	return result
}
func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Fprintln(w, "Error reading user", err)
		return
	}

	if user.Name == "test" && user.Password == "test" {
		user.Password = ""
		token := GenerateJWT(user)

		result := models.ResponseToken{token}
		//convert the variable to json
		jsonResult, err := json.Marshal(result)

		if err != nil {
			fmt.Fprintln(w, "Erreur generate json", err)
			return
		}
		//We send the token

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-type", "application/json")
		w.Write(jsonResult)
	} else {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "User o password invalid", err)
	}

}
