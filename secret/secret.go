package secret

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

const secretAddrFmt = "projects/%s/secrets/%s/versions/%s"

type Config struct {
	Version   string
	ProjectID string
	Name      string
	IsFile    bool
	File      string
}

// NewSecret create Secret Data from GCP Secret Manager
//
// You can set from json file if you set 'IsFile' as true.
// In that case, you must set 'File'.
func NewSecret(conf Config, v interface{}) error {
	if conf.IsFile {
		return fromFile(conf, v)

	}
	return fromSecretManager(conf, v)
}

func fromSecretManager(conf Config, v interface{}) error {
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to create secretmanager client: %v", err)
	}

	name := fmt.Sprintf(secretAddrFmt, conf.ProjectID, conf.Name, conf.Version)
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: name,
	}

	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to get secret: %v", err)
	}

	return json.Unmarshal(result.Payload.Data, &v)

}

func fromFile(conf Config, v interface{}) error {
	f, err := os.Open(conf.File)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewDecoder(f).Decode(&v)
}
