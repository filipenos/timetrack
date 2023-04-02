// https://developers.track.toggl.com/docs/tracking
package toggl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Toggl struct {
	ApiToken string
	Client   *http.Client
	basePath string
}

func New() *Toggl {
	return &Toggl{Client: http.DefaultClient, basePath: "https://api.track.toggl.com/api/v9"}
}

func (t *Toggl) Path(path string) string {
	return fmt.Sprintf("%s%s", t.basePath, path)
}

func (t *Toggl) Send(method, path string, in, out any) error {
	var body io.Reader
	if in != nil {
		b, err := json.Marshal(in)
		if err != nil {
			return err
		}
		body = bytes.NewBuffer(b)
	}
	req, err := http.NewRequest(method, t.Path(path), body)
	if err != nil {
		return err
	}
	req.SetBasicAuth(t.ApiToken, "api_token")
	req.Header.Set("Content-type", "application/json")

	resp, err := t.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("server return :%s\n%s", resp.Status, string(b))
	}

	return json.NewDecoder(resp.Body).Decode(&out)
}

func (t *Toggl) Auth(apiToken string) error {
	if apiToken == "" {
		return fmt.Errorf("apiToken is required")
	}
	t.ApiToken = apiToken

	me, err := t.Me()
	if err != nil {
		return err
	}
	if me == nil || me.ID == 0 {
		return fmt.Errorf("invalid username/password")
	}
	return nil
}

func (t *Toggl) Me() (*Me, error) {
	var me Me
	if err := t.Send(http.MethodGet, "/me", nil, &me); err != nil {
		return nil, err
	}
	return &me, nil
}

func (t *Toggl) Workspaces() ([]Workspace, error) {
	var works []Workspace
	if err := t.Send(http.MethodGet, "/workspaces", nil, &works); err != nil {
		return nil, err
	}
	return works, nil
}

func (t *Toggl) Projects(workspaceID int) ([]Project, error) {
	var projects []Project
	if err := t.Send(http.MethodGet, fmt.Sprintf("/workspaces/%d/projects", workspaceID), nil, &projects); err != nil {
		return nil, err
	}
	return projects, nil
}

func (t *Toggl) Tags() ([]Tag, error) {
	var tags []Tag
	if err := t.Send(http.MethodGet, "/me/tags", nil, &tags); err != nil {
		return nil, err
	}
	return tags, nil
}

func (t *Toggl) CurrentTimeEntry() (*TimeEntry, error) {
	var cur TimeEntry
	if err := t.Send(http.MethodGet, "/me/time_entries/current", nil, &cur); err != nil {
		return nil, err
	}
	return &cur, nil
}

func (t *Toggl) NewTimeEntry(entry *NewTimeEntry) (*TimeEntry, error) {
	var resp TimeEntry
	if err := t.Send(http.MethodPost, fmt.Sprintf("/workspaces/%d/time_entries", entry.WorkspaceID), entry, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (t *Toggl) StopTimeEntry(workspaceID, timeEntryID int) error {
	return t.Send(http.MethodPatch, fmt.Sprintf("/workspaces/%d/time_entries/%d/stop", workspaceID, timeEntryID), nil, nil)
}
