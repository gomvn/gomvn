package entity

import (
	"strings"
	"time"
)

func NewArtifact(path string, mod time.Time) *Artifact {
	parts := strings.Split(path, "/")
	last := len(parts) - 1
	return &Artifact{
		Group: strings.Join(parts[0:last-2], "."),
		Artifact: parts[last-2],
		Version: parts[last-1],
		Modified: mod,
	}
}

type Artifact struct {
	Group    string
	Artifact string
	Version  string
	Modified time.Time
}

func (a *Artifact) GetPath() string {
	return strings.Replace(a.Group, ".", "/", -1) + "/" + a.Artifact + "/" + a.Version
}
