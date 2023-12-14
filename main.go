package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/blynn-dn/webhook/config"
	"github.com/blynn-dn/webhook/model"
)

var (
	sc = &config.SafeConfig{
		C: &config.Config{},
	}
)

// VerifySignature - Verifies a SHA256 signature digest. This is primary to support signatures
// from nautobot webhooks.
// secret: key used to sign
// signature: signature retrieved from the request header
// payload: the request body
func VerifySignature(secret string, signature string, payload []byte) bool {
	hash := hmac.New(sha512.New, []byte(secret))
	if _, err := hash.Write(payload); err != nil {
		log.Printf("Cannot compute the HMAC for request: %s\n", err)
		return false
	}
	expectedHash := hex.EncodeToString(hash.Sum(nil))
	log.Println("EXPECTED HASH:", expectedHash)
	return signature == expectedHash
}

func run() int {
	userHome, _ := os.UserHomeDir()
	configFile := ""

	// listen for SIGTERM signal
	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)

	// listen for SIGHUP signal
	hup := make(chan os.Signal, 1)
	signal.Notify(hup, syscall.SIGHUP)

	// use a channel for closing the http service on an unexpected interrupt
	service := make(chan struct{})

	// load the config file
	configPaths := []string{".webhook.yaml", fmt.Sprintf("%s/.webhook.yaml", userHome), "/etc/.webhook.yaml"}
	for _, fn := range configPaths {
		if err := sc.LoadConfig(fn); err == nil {
			configFile = fn

			break
		}
	}
	log.Printf("using config: %s", configFile)

	// reload config on SIGHUP
	go func() {
		<-hup
		if err := sc.LoadConfig(configFile); err != nil {
			log.Fatalf("Failed to read config: %s", err)
		}
		log.Println("config reloaded")
	}()

	// webhook handler
	http.HandleFunc(
		fmt.Sprintf("%sv1/webhook", sc.C.BasePath),
		func(w http.ResponseWriter, r *http.Request) {
			// read-in the request body which will contain a JSON payload
			data, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Printf("Failed to read request: %s", err)
				return
			}

			// get signature from request header
			sig := r.Header.Get("X-Hook-Signature")
			log.Println("X-Hook-Signature: ", sig)

			// verify the signature
			if !VerifySignature(sc.C.Secret, sig, data) {
				log.Println("Unable to verify signature")
				return
			}

			// print the payload
			// log.Println("Data: ", string(data))

			var payload model.WebhookEvent
			err = json.Unmarshal(data, &payload)
			if err != nil {
				log.Println("error: ", err)
				return
			}

			// just print the device display field vs the entire structure
			log.Printf("Display: %s", payload.Data.Display)
			//log.Println("payload: \n", payload)
		})

	// create the http server
	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", sc.C.Port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
	log.Println("Listening on ", sc.C.Port)

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println("Error starting HTTP server: ", err)

			// close channel -- this will signal the parent to exit
			close(service)
		}
	}()

	// handle SIGTERM or service interruption
	for {
		select {
		case <-term:
			log.Println("Received SIGTERM, exiting gracefully...")
			_ = srv.Close()
			return 0
		case <-service:
			// http server unexpectedly interrupted/exited
			return 1
		}
	}
}

func main() {
	// call run(); exit with the appropriate exit error code
	os.Exit(run())
}
