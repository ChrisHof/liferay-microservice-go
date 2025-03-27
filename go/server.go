package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jws"
	"github.com/lestrrat-go/jwx/v3/jwt"
)

type OAuth2Application struct {
	Client_id string `json:"client_id"`
}

func IsValidRequest(w http.ResponseWriter, r *http.Request) bool {
	authorization := r.Header.Get("Authorization")
	if !strings.HasPrefix(authorization, "Bearer ") || !isValidJWT(authorization[len("Bearer "):]) {
		w.WriteHeader(http.StatusUnauthorized)
		return false
	} else if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return false
	} else if r.Body == http.NoBody {
		w.WriteHeader(http.StatusBadRequest)
		return false
	} else if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return false
	}
	return true
}

func isValidJWT(token string) bool {
	oAuth2Baseurl := os.Getenv("LIFERAY_BASE_URL") + "/o/oauth2"
	jwkSet, err := jwk.Fetch(context.Background(), oAuth2Baseurl+"/jwks")
	if err != nil {
		Logger.Printf("Error fetching JWKS: %s\n", err)
		return false
	}
	tok, err := jwt.Parse(
		[]byte(token),
		jwt.WithKeySet(jwkSet, jws.WithRequireKid(false)),
	)
	if err != nil {
		Logger.Printf("Error parsing token: %s\n", err)
		return false
	}
	if !tok.Has("client_id") {
		Logger.Printf("JWT missing client_id")
		return false
	}
	var client_id string
	err = tok.Get("client_id", &client_id)
	if err != nil {
		Logger.Printf("Error getting JWT client_id: %s\n", err)
		return false
	}
	resp, err := http.Get(oAuth2Baseurl + "/application?externalReferenceCode=" + os.Getenv("OAUTH2_APPLICATION_REFERENCE_CODE"))
	if err != nil {
		Logger.Printf("Error getting OAuth2 Application: %s\n", err)
		return false
	}
	var oAuth2Application OAuth2Application
	json.NewDecoder(resp.Body).Decode(&oAuth2Application)
	if oAuth2Application.Client_id != client_id {
		Logger.Println("JWT client_id does not match OAuth2 application")
		return false
	}
	return true
}

func StartHttpServer(port string) {
	Logger.Println("Starting HTTP Server on Port " + port + "...")
	httpServer := &http.Server{
		Addr:         ":" + port,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	httpError := httpServer.ListenAndServe()
	if httpError != nil {
		Logger.Println(httpError)
	}
}
