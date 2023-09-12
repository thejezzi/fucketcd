package manager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/thejezzi/fucketcd/pkg/util"
	"github.com/urfave/cli/v2"
)

type BranchList = map[string]interface{}

func Import(cCtx *cli.Context) error {
	path := cCtx.Args().First()

	isValid := util.ValidatePath(path)
	if !isValid {
		return util.ErrInvalidPath()
	}

	isValidImportFile := util.ValidateImportFile(path)
	if !isValidImportFile {
		return util.ErrInvalidImportFile()
	}

	data, _ := os.ReadFile(path)

	branchList, err := IntoBranches(&data)
	if err != nil {
		return err
	}

	fmt.Println("---")
	for k, v := range *branchList {
		fmt.Println(k, ":", v)
	}
	fmt.Println("---\n")

	if !cCtx.Bool("yes") {
		confirmation, err := util.GetUserConfirmation("I will push these values under the corresponding paths to etcd. Continue?")
		if err != nil {
			return err
		}

		if !confirmation {
			fmt.Println("Aborting ...")
			return nil
		}
	}

	for k, v := range *branchList {
		err := PushToETCD(k, v)
		if err != nil {
			return err
		}
	}

	return nil
}

func IntoBranches(content *[]byte) (*BranchList, error) {
	data, err := jsonToMap(content)
	if err != nil {
		return nil, err
	}

	flattenedMap := make(map[string]interface{})
	flattenMap(data, "", flattenedMap)

	return &flattenedMap, nil
}

func jsonToMap(jsonBlob *[]byte) (map[string]interface{}, error) {
	var data map[string]interface{}
	if err := json.Unmarshal(*jsonBlob, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func flattenMap(input map[string]interface{}, currentPath string, output map[string]interface{}) {
	for key, value := range input {
		newPath := currentPath + "/" + key
		switch v := value.(type) {
		case map[string]interface{}:
			flattenMap(v, newPath, output)
		default:
			output[newPath] = v
		}
	}
}

func PushToETCD(path string, value interface{}) error {
	stringValue, ok := value.(string)
	if !ok {
		return fmt.Errorf("value is not a string: %v", value)
	}

	pathWithOutPrefixDelimiter := strings.TrimPrefix(path, "/")

	fmt.Println("Pushing", path, ":", value)

	url := fmt.Sprintf("http://localhost:2379/v2/keys/%s", pathWithOutPrefixDelimiter)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer([]byte("value="+stringValue)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send the request via a client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return err
	}
	defer resp.Body.Close()

	return nil
}
