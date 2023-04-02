package config

import "testing"

func TestParse(t *testing.T) {

	content := `# file content configs
	apitoken=t123!@#$
workspace_id=1234567890
##final config
[sample]
without`

	cfg := Parse([]byte(content))
	if expect := "t123!@#$"; cfg.ApiToken != expect {
		t.Errorf("ApiToken expect: %s, but got: %s", expect, cfg.ApiToken)
	}
	if expect := "1234567890"; cfg.WorkspaceID != expect {
		t.Errorf("WorkspaceID expect: %s, but got: %s", expect, cfg.WorkspaceID)
	}
}
