package upper

import (
	"encoding/json"
	"fmt"
	"github.com/kruftik/nebius-instance-upper/internal/config"
	"github.com/pkg/errors"

	ycsdk "github.com/yandex-cloud/go-sdk"
	"github.com/yandex-cloud/go-sdk/iamkey"

	nccm "github.com/kruftik/nebius-instance-upper/pkg/nebius-cloud-utils"
)

// NewCloudConfig creates a new instance of CloudConfig object
func NewCloudConfig(cfg config.AppConfig) (ycsdk.Credentials, string, error) {
	var (
		iamKey iamkey.Key

		err error
	)

	if err = json.Unmarshal(cfg.IAMServiceAccountKey, &iamKey); err != nil {
		return nil, "", fmt.Errorf("malformed service account json: %w", err)
	}

	credentials, err := ycsdk.ServiceAccountKey(&iamKey)
	if err != nil {
		return nil, "", fmt.Errorf("invalid auth credentials: %w", err)
	}

	if cfg.FolderID == "" {
		metadata := nccm.NewMetadataService()

		cfg.FolderID, err = metadata.GetFolderID()
		if err != nil {
			return nil, "", errors.Wrap(err, "cannot get FolderID from instance metadata")
		}
	}

	return credentials, cfg.FolderID, nil
}
