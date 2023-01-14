package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func initEnvs() {
	check := []string{"HEALTHCHECK_STATUS", "COUNTER_HIT_GOLANG", "UPLOADED_VT_FILE", "AWS_S3_BUCKET", "AWS_S3_REGION"}

	for _, v := range check {
		_, present := os.LookupEnv(v)
		if !present {
			msg := fmt.Sprintf("Environments %s not set", v)
			log.Fatal(msg)
		}
	}
}

func getIp(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")

	if forwarded != "" {
		return forwarded
	}

	return r.RemoteAddr
}

func hitCounter() int {
	flag := os.Getenv("COUNTER_HIT_GOLANG")
	count, err := strconv.Atoi(flag)

	if err != nil {
		panic(err)
	}

	quick_math := count + 1
	newValue := strconv.Itoa(quick_math)
	os.Setenv("COUNTER_HIT_GOLANG", newValue)

	return quick_math
}

func getHealth() bool {
	boolValue, err := strconv.ParseBool(os.Getenv("HEALTHCHECK_STATUS"))

	if err != nil {
		log.Println(err)
	}

	return boolValue
}

func switchHealth() string {
	os.Setenv("HEALTHCHECK_STATUS", strconv.FormatBool(!getHealth()))

	return "Switched"
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
