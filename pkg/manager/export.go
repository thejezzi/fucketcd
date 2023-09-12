package manager

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

type Node struct {
	Key           string  `json:"key"`
	Dir           bool    `json:"dir"`
	Value         string  `json:"value,omitempty"`
	Nodes         []*Node `json:"nodes,omitempty"`
	ModifiedIndex int     `json:"modifiedIndex"`
	CreatedIndex  int     `json:"createdIndex"`
}

type Response struct {
	Action string `json:"action"`
	Node   Node   `json:"node"`
}

func Export(cCtx *cli.Context) error {
	path := cCtx.String("path")
	if path == "" {
		path = "."
	}

	// Get all data from etcd server
	resp, err := http.Get("http://localhost:2379/v2/keys/?recursive=true")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		panic(err)
	}

	flatMap := buildMap(response.Node)

	flatJson, err := json.MarshalIndent(flatMap, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(path+"/etcd.json", flatJson, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Exported data to " + path + "/etcd.json")
	return nil
}

func buildMap(node Node) map[string]interface{} {
	m := make(map[string]interface{})
	for _, n := range node.Nodes {
		if n.Dir {
			lastKey := strings.Split(n.Key, "/")
			m[lastKey[len(lastKey)-1]] = buildMap(*n)
		} else {
			lastKey := strings.Split(n.Key, "/")
			m[lastKey[len(lastKey)-1]] = n.Value
		}
	}
	return m
}
