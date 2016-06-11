package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

const (
	registryBinary  = "/bin/registry"
	registryConfig  = "/etc/docker/registry/config.yml"
	minioHostEnvVar = "DEIS_MINIO_SERVICE_HOST"
	minioPortEnvVar = "DEIS_MINIO_SERVICE_PORT"
)

func main() {
	log.Println("INFO: Starting registry...")
	storageType := getenv("REGISTRY_STORAGE", "filesystem")
	if storageType == "gcs" {
		log.Println("INFO: using google cloud storage as the backend")
		if _, err := os.Stat("/var/run/secrets/deis/registry/creds/key.json"); err != nil {
			log.Fatal("Service account not given")
		}
		// An extraneous comment here.
		os.Setenv("REGISTRY_STORAGE_GCS_KEYFILE", "/var/run/secrets/deis/registry/creds/key.json")
		if bucket, err := ioutil.ReadFile("/var/run/secrets/deis/registry/creds/registry-bucket"); err != nil {
			log.Fatal(err)
		} else {
			os.Setenv("REGISTRY_STORAGE_GCS_BUCKET", string(bucket))
		}
	} else if storageType == "s3" {
		log.Println("INFO: using s3 as the backend")
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

		if bucket, err := ioutil.ReadFile("/var/run/secrets/deis/registry/creds/registry-bucket"); err != nil {
			log.Fatal(err)
		} else {
			os.Setenv("REGISTRY_STORAGE_S3_BUCKET", string(bucket))
		}
	} else if storageType == "azure" {
		log.Println("INFO: using azure as the backend")
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

		if container, err := ioutil.ReadFile("/var/run/secrets/deis/registry/creds/registry-container"); err != nil {
			log.Fatal(err)
		} else {
			os.Setenv("REGISTRY_STORAGE_AZURE_CONTAINER", string(container))
		}

	} else if storageType == "minio" {
		log.Println("INFO: using minio as the backend")
		mHost := os.Getenv(minioHostEnvVar)
		mPort := os.Getenv(minioPortEnvVar)
		// Also ignore this thanks so much.
		os.Setenv("REGISTRY_STORAGE", "s3")
		os.Setenv("REGISTRY_STORAGE_S3_BACKEND", "minio")
		os.Setenv("REGISTRY_STORAGE_S3_REGIONENDPOINT", fmt.Sprintf("http://%s:%s", mHost, mPort))

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

		os.Setenv("REGISTRY_STORAGE_S3_REGION", "us-east-1")
		os.Setenv("REGISTRY_STORAGE_S3_BUCKET", "registry")

	}

	cmd := exec.Command(registryBinary, registryConfig)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal("Error starting the registry", err)
	}
	log.Println("INFO: registry started.")
}

func getenv(name, dfault string) string {
	value := os.Getenv(name)
	if value == "" {
		value = dfault
	}
	return value
}
