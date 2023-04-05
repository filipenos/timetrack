package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/filipenos/timetrack/internal/args"
	"github.com/filipenos/timetrack/internal/config"
	"github.com/filipenos/timetrack/internal/toggl"
)

var cfg *config.Config

func init() {
	cfg, _ = config.Read()
}

type Exec struct {
	t *toggl.Toggl
}

func main() {
	p, err := args.Parse(os.Args)
	if err != nil {
		log("%v", err)
		return
	}

	if err := p.Exec(&Exec{t: toggl.New()}); err != nil {
		log("%v", err)
		return
	}
}

func log(format string, a ...any) {
	fmt.Println(fmt.Sprintf(format, a...))
}

func (e *Exec) Auth() error {
	log("auth")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("ApiToken: ")
	apiToken, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	apiToken = strings.TrimSpace(apiToken)
	if err := e.t.Auth(apiToken); err != nil {
		return err
	}

	conf := config.Config{ApiToken: apiToken}

	works, err := e.t.Workspaces()
	if err != nil {
		return err
	}
	if len(works) > 0 {
		conf.WorkspaceID = fmt.Sprintf("%d", works[0].ID)
	}

	return config.Write(conf)
}

func (e *Exec) Start(args args.StartArgs) error {
	log("start")

	if err := e.t.Auth(cfg.ApiToken); err != nil {
		return err
	}

	var workspaceID = 0
	if cfg.WorkspaceID != "" {
		workspaceID, _ = strconv.Atoi(cfg.WorkspaceID)
	} else {
		works, err := e.t.Workspaces()
		if err != nil {
			return err
		}
		if len(works) == 0 {
			return fmt.Errorf("no workspaces found")
		}
		workspaceID = works[0].ID
	}
	if workspaceID == 0 {
		return fmt.Errorf("workspace id required")
	}

	entry := &toggl.NewTimeEntry{
		CreatedWith: "timetrack cli",
		Description: args.Message,
		Tags:        []string{},
		WorkspaceID: workspaceID,
		ProjectID:   0,
		Start:       time.Now(),
		Stop:        nil,
	}
	entry.Duration = int(entry.Start.Unix()) * -1

	if args.Project != "" {
		projects, err := e.t.Projects(entry.WorkspaceID)
		if err != nil {
			return err
		}
		for _, p := range projects {
			if p.Name == args.Project {
				entry.ProjectID = p.ID
				break
			}
		}
		if entry.ProjectID == 0 {
			return fmt.Errorf("project %s not found", args.Project)
		}
	}
	if args.Tag != "" {
		tags, err := e.t.Tags()
		if err != nil {
			return err
		}
		for _, t := range tags {
			if t.Name == args.Tag {
				entry.Tags = append(entry.Tags, args.Tag)
				break
			}
		}
		if len(entry.Tags) == 0 {
			return fmt.Errorf("tag %s not found", args.Tag)
		}
	}

	cur, err := e.t.CurrentTimeEntry()
	if err != nil {
		return err
	}
	if cur != nil && cur.ID > 0 {
		if cur.ProjectID == entry.ProjectID && idInArray(args.Tag, cur.Tags) && cur.Description == entry.Description {
			log("already running time track: %d", cur.ID)
			return nil
		}
		log("stop current, and start new")
		if err := e.t.StopTimeEntry(cur.WorkspaceID, cur.ID); err != nil {
			return err
		}
	}

	resp, err := e.t.NewTimeEntry(entry)
	if err != nil {
		return err
	}

	log("started track: %d", resp.ID)
	return nil
}

func (e *Exec) Stop() error {
	log("stop")

	if err := e.t.Auth(cfg.ApiToken); err != nil {
		return err
	}

	cur, err := e.t.CurrentTimeEntry()
	if err != nil {
		return err
	}
	if cur == nil || cur.ID == 0 {
		return fmt.Errorf("no running time track")
	}

	if err := e.t.StopTimeEntry(cur.WorkspaceID, cur.ID); err != nil {
		return err
	}

	log("stoped track: %d", cur.ID)
	return nil
}

func (e *Exec) Continue() error {
	log("continue")

	if err := e.t.Auth(cfg.ApiToken); err != nil {
		return err
	}

	cur, err := e.t.CurrentTimeEntry()
	if err != nil {
		return err
	}
	if cur != nil && cur.ID > 0 {
		log("already running time track: %d", cur.ID)
		return nil
	}

	list, err := e.t.List()
	if err != nil {
		return err
	}
	if len(list) == 0 {
		return fmt.Errorf("empty entrys")
	}
	last := list[0]
	log("continue: %d\n", last.ID)

	entry := &toggl.NewTimeEntry{
		CreatedWith: "timetrack cli",
		Description: last.Description,
		Tags:        last.Tags,
		WorkspaceID: last.WorkspaceID,
		ProjectID:   last.ProjectID,
		Start:       time.Now(),
		Stop:        nil,
	}
	entry.Duration = int(entry.Start.Unix()) * -1

	resp, err := e.t.NewTimeEntry(entry)
	if err != nil {
		return err
	}

	log("started track: %#v\n", resp.ID)

	return nil
}

func idInArray(id string, ids []string) bool {
	for _, i := range ids {
		if i == id {
			return true
		}
	}
	return false
}
