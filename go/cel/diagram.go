// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package cel

import (
	"fmt"
	"io"
	"text/template"
)

type node struct {
	Label string

	// Ids of nodes for which there are outgoing edges from this node.
	Edges []string

	visited bool
}

type graph struct {
	Nodes map[string]*node
}

var nextId = 0

const kDotGraph = `strict digraph {
  node [shape=rectangle]
{{range $id, $node := .Nodes}}
  {{$id}} [label="{{$node.Label}}"]
  {{- range .Edges}}
  {{$id}} -> {{.}}{{end}}
{{end}}
}
`

type nodeset map[string]bool

func (n nodeset) mergeFrom(o nodeset) {
	for a, _ := range o {
		n[a] = true
	}
}

func (n nodeset) add(s string) {
	n[s] = true
}

func (n nodeset) remove(s string) {
	delete(n, s)
}

func getIdForNode(d *DependencyNode, m map[*DependencyNode]string) string {
	if id, ok := m[d]; ok {
		return id
	}

	id := fmt.Sprintf("L%d", nextId)
	nextId += 1

	m[d] = id
	return id
}

// simplifyGraphAt removes any edges from node |n| that another edge is able to
// reach. Assumes that |n| is a member of |g| and that |g| is a DAG.
//
// E.g.: Imagine we are dealing with the following subgraph, where edges are
// directed downward:
//
//                 a
//                /|
//               / |
//              b  |
//               \ |
//                \|
//                 c
//
// Both |a| and |b| have edges to |c|. Calling simplifyGraphAt() on |a| would
// remove the edge from |a| to |c| since |c| is reachable via either |b| or
// |a|. Thus yielding the following:
//
//                 a
//                /
//               /
//              b
//               \
//                \
//                 c
//
// Assuming that we are dealing with an acyclic dependency graph, this results
// in a minimal graph that's easier to read than one where all transitive
// dependencies are visible.
func simplifyGraphAt(g *graph, id string, n *node) (ns nodeset) {
	if n.visited {
		return
	}

	ns = make(nodeset)
	for _, e := range n.Edges {
		reachables := simplifyGraphAt(g, e, g.Nodes[e])
		reachables.remove(e)

		// |reachables| - 'e' is now a set of edges that can be removed from
		// |n|. 'e' itself may end up being removed if another edge leads to
		// 'e', but such an edge cannot be a member of |reachables|. Otherwise
		// |g| is not a DAG.
		ns.mergeFrom(reachables)
	}

	new_e := []string{}
	for _, e := range n.Edges {
		if _, ok := ns[e]; !ok {
			new_e = append(new_e, e)
		}
	}
	n.Edges = new_e

	for _, e := range new_e {
		ns.add(e)
	}
	ns.add(id)
	return
}

// DumpAssetDepsInDotFormat writes the dependencies in |A| to |w| in DOT
// format. See https://en.wikipedia.org/wiki/DOT_(graph_description_language)
// for details on the DOT format.
func DumpAssetDepsInDotFormat(A *Assets, w io.Writer) error {
	// has_incoming[node_id] is true if node_id has incoming edges.
	has_incoming := make(map[string]bool)

	g := graph{Nodes: make(map[string]*node)}

	id_map := make(map[*DependencyNode]string)

	for _, d := range A.All {
		id := getIdForNode(d, id_map)
		n := &node{Label: fmt.Sprintf("%s : %s", d.Asset.Namespace(), d.Asset.Id())}
		for _, t := range d.Dependents {
			edge_id := getIdForNode(t, id_map)
			n.Edges = append(n.Edges, edge_id)
			has_incoming[edge_id] = true
		}
		g.Nodes[id] = n
	}

	// Look for root nodes and start simplifying the graph by removing
	// transitively reachable edges.
	for id, n := range g.Nodes {
		if _, ok := has_incoming[id]; !ok {
			simplifyGraphAt(&g, id, n)
		}
	}

	t, err := template.New("g").Parse(kDotGraph)
	if err != nil {
		return err
	}

	return t.Execute(w, g)
}
