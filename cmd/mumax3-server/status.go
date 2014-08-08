package main

// Serves human-readable status information over http.

import (
	"html/template"
	"net/http"
	"time"
)

var (
	templ = template.Must(template.New("status").Parse(templText))
)

func (n *Node) HandleStatus(w http.ResponseWriter, r *http.Request) {
	err := templ.Execute(w, node.Status())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type NodeStatus struct {
	NodeInfo
	MumaxVersion string // which mumax version this node runs, if any
	GPUs         []GPU  // number of available GPUs
	Queue        []string
	Peers        []string
	Uptime       time.Duration
}

func (n *Node) Status() NodeStatus {
	n.lock()
	defer n.unlock()
	peers := make([]string, 0, len(n.peers))
	for k, _ := range n.peers {
		peers = append(peers, k)
	}
	return NodeStatus{
		NodeInfo: n.inf,
		Queue:    n.jobs.ListFiles(),
		Uptime:   time.Since(n.upSince),
		Peers:    peers,
	}
}

const templText = `
<html>

<head>
	<style>
		body{font-family:monospace}
	</style>
</head>

<body>

<h1>{{.Addr}}</h1>

Uptime: {{.Uptime}} <br/>

{{with .MumaxVersion}}
	Running {{.}}
{{end}}

{{with .GPUs}}
	<h2>GPUs</h2>
	{{range .}}
		{{.Info}}</br>
	{{end}}
{{end}}

<h2>Queue</h2>
{{range .Queue}}
	{{.}}</br>
{{end}}

<h2>Peers</h2>
{{range .Peers}}
	{{.}}</br>
{{end}}

</body>
</html>
`
