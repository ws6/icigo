package icigo

import (
	"context"

	"os"
	"testing"
)

func getConfig() map[string]string {
	ret := make(map[string]string)
	ret[AUTH_TOKEN] = os.Getenv("ICI_TEST_AUTH_TOKEN")
	ret[ILMN_DOMAIN] = os.Getenv("ICI_TEST_ILMN_DOMAIN")
	ret[ILMN_WORKGROUP] = os.Getenv("ICI_TEST_ILMN_WORKGROUP")
	ret[ILMN_PROJECT] = os.Getenv("ICI_TEST_ILMN_PROJECT")
	ret[BASE_URL] = os.Getenv("ICI_TEST_BASE_URL")

	return ret
}
func getNewClient() *Client {
	ret, err := NewClient(getConfig())
	if err != nil {
		panic(err)
	}
	return ret
}

func TestSearchAuditLog(t *testing.T) {
	client := getNewClient()
	queryParam := make(map[string]string)
	queryParam[`fromDate`] = `2021-10-25T14:12:06+0000`
	queryParam[`orderBy`] = `+createdDate`

	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	ret, err := client.SearchAuditLog(ctx, queryParam)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Logf(string(ret))

}
