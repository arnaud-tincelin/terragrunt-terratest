package main

import (
	"path"
	"testing"

	_ "github.com/Azure/go-autorest/autorest/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/require"
)

func TestVM(t *testing.T) {
	t.Parallel()

	const workingDir = "../terragrunt/dev/"
	basementDir := path.Join(workingDir, "basement")
	vmDir := path.Join(workingDir, "vm")

	test_structure.SaveTerraformOptions(t, workingDir, terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformBinary: "terragrunt",
		TerraformDir:    workingDir,
	}))

	test_structure.SaveTerraformOptions(t, basementDir, terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformBinary: "terragrunt",
		TerraformDir:    basementDir,
	}))

	test_structure.SaveTerraformOptions(t, vmDir, terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformBinary: "terragrunt",
		TerraformDir:    vmDir,
	}))

	defer test_structure.RunTestStage(t, "destroy", func() {
		terraformOptions := test_structure.LoadTerraformOptions(t, workingDir)
		terraform.TgDestroyAll(t, terraformOptions)
	})

	test_structure.RunTestStage(t, "apply", func() {
		terraformOptions := test_structure.LoadTerraformOptions(t, workingDir)
		terraform.TgApplyAll(t, terraformOptions)
	})

	test_structure.RunTestStage(t, "validate", func() {
		basementOptions := test_structure.LoadTerraformOptions(t, basementDir)

		storageAccount := new(StorageAccountOutput)
		terraform.OutputStruct(t, basementOptions, "storage", storageAccount)
		containerName := terraform.Output(t, basementOptions, "storage_container_name")

		require.NoError(t, upload(storageAccount.Name, containerName, "helloworld.txt", []byte("Hello world"), nil), "upload data file")
		require.NoError(t, waitForBlobToBeRemoved(storageAccount.Name, containerName, "helloworld.txt"))
	})
}

type StorageAccountOutput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
