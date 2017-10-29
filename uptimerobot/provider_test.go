package uptimerobot

import (
	"testing"
	"github.com/hashicorp/terraform/helper/schema"
)

func init() {
}


func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
}

func testAccPreCheck(t *testing.T) {
}
