package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

const (
	registryBinary = "/bin/registry"
	registryConfig = "/etc/docker/registry/config.yml"
)

func main() {
	log.Println("INFO: Starting registry...")
	storageType := getenv("REGISTRY_STORAGE", "filesystem")
	if storageType == "gcs" {
		if _, err := os.Stat("/var/run/secrets/deis/registry/creds/key.json"); err != nil {
			log.Fatal("Service account not given")
		}
		os.Setenv("REGISTRY_STORAGE_GCS_KEYFILE", "/var/run/secrets/deis/registry/creds/key.json")
		if bucket, err := ioutil.ReadFile("/var/run/secrets/deis/registry/creds/bucket"); err != nil {
			log.Fatal(err)
		} else {
			os.Setenv("REGISTRY_STORAGE_GCS_BUCKET", string(bucket))
		}
	} else if storageType == "generic" {
		if accesskey, err := ioutil.ReadFile("/var/run/secrets/deis/registry/creds/accesskey"); err != nil {
			log.Fatal(err)
		} else {
			os.Setenv("REGISTRY_STORAGE_S3_ACCESSKEY", string(accesskey))
		}

		if secretkey, err := ioutil.ReadFile("/var/run/secrets/deis/registry/creds/secretkey"); err != nil {
			log.Fatal(err)
		} else {
			os.Setenv("REGISTRY_STORAGE_S3_SECRETKEY", string(secretkey))
		}

		if region, err := ioutil.ReadFile("/var/run/secrets/deis/registry/creds/region"); err != nil {
			log.Fatal(err)
		} else {
			os.Setenv("REGISTRY_STORAGE_S3_REGION", string(region))
		}

		if bucket, err := ioutil.ReadFile("/var/run/secrets/deis/registry/creds/bucket"); err != nil {
			log.Fatal(err)
		} else {
			os.Setenv("REGISTRY_STORAGE_S3_BUCKET", string(bucket))
		}
	} else if storageType == "azure" {
		if accountname, err := ioutil.ReadFile("/var/run/secrets/deis/registry/creds/accountname"); err != nil {
			log.Fatal(err)
		} else {
			os.Setenv("REGISTRY_STORAGE_AZURE_ACCOUNTNAME", string(accountname))
		}

		if accountkey, err := ioutil.ReadFile("/var/run/secrets/deis/registry/creds/accountkey"); err != nil {
			log.Fatal(err)
		} else {
			os.Setenv("REGISTRY_STORAGE_AZURE_ACCOUNTKEY", string(accountkey))
		}

		if container, err := ioutil.ReadFile("/var/run/secrets/deis/registry/creds/container"); err != nil {
			log.Fatal(err)
		} else {
			os.Setenv("REGISTRY_STORAGE_AZURE_CONTAINER", string(container))
		}

	}

	cmd := exec.Command(registryBinary, registryConfig)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		log.Fatal("Error starting the registry", err)
	}
	log.Println("INFO: registry started.")
	for {
	}
}

func getenv(name, dfault string) string {
	value := os.Getenv(name)
	if value == "" {
		value = dfault
	}
	return value
}
