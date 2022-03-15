package main

import (
	"path"
	"testing"

	_ "github.com/Azure/go-autorest/autorest/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/require"
)

func TestVMAccessToSA(t *testing.T) {
	t.Parallel()

	const devPath = "../terragrunt/dev/"
	basementDir := path.Join(devPath, "basement")

	devEnv := &terraform.Options{
		TerraformBinary: "terragrunt",
		TerraformDir:    devPath,
	}

	basement := &terraform.Options{
		TerraformBinary: "terragrunt",
		TerraformDir:    basementDir,
	}

	defer terraform.TgDestroyAll(t, devEnv)

	terraform.TgApplyAll(t, devEnv)

	storageAccount := new(StorageAccountOutput)
	terraform.OutputStruct(t, basement, "storage", storageAccount)
	containerName := terraform.Output(t, basement, "storage_container_name")

	require.NoError(t, upload(storageAccount.Name, containerName, "helloworld.txt", []byte("Hello world"), nil), "upload data file")
	require.NoError(t, waitForBlobToBeRemoved(storageAccount.Name, containerName, "helloworld.txt"))
}

type StorageAccountOutput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
