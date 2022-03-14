package main

import (
	"path"
	"testing"

	_ "github.com/Azure/go-autorest/autorest/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/require"
)

func TestVM(t *testing.T) {
	t.Parallel()

	const workingDir = "../terragrunt/dev/"
	basementDir := path.Join(workingDir, "basement")

	working := &terraform.Options{
		TerraformBinary: "terragrunt",
		TerraformDir:    workingDir,
	}

	base := &terraform.Options{
		TerraformBinary: "terragrunt",
		TerraformDir:    basementDir,
	}

	defer terraform.TgDestroyAll(t, working)

	terraform.TgApplyAll(t, working)

	storageAccount := new(StorageAccountOutput)
	terraform.OutputStruct(t, base, "storage", storageAccount)
	containerName := terraform.Output(t, base, "storage_container_name")

	require.NoError(t, upload(storageAccount.Name, containerName, "helloworld.txt", []byte("Hello world"), nil), "upload data file")
	require.NoError(t, waitForBlobToBeRemoved(storageAccount.Name, containerName, "helloworld.txt"))
}

type StorageAccountOutput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
