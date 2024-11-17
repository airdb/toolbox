package netutil_test

import (
	"strings"
	"testing"

	"github.com/airdb/toolbox/netutil"
)

func TestRun(t *testing.T) {
	t.Log("TestRun")

	fetcher := netutil.NewFetcher("whois.radb.net", 43)

	asnum := "AS4134"
	asnum = strings.ReplaceAll(asnum, "AS", "")
	ips, err := fetcher.Fetch(true, false, asnum)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("AS number %s ipv4 count %d", asnum, len(ips))
}
