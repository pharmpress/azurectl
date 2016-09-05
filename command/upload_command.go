package command

import (
	"fmt"
	"github.com/MSOpenTech/azure-sdk-for-go/storage"
	"github.com/codegangsta/cli"
	"os"
)

func NewUploadCommand() cli.Command {
	return cli.Command{
		Name:  "upload",
		Usage: "upload blob into a container",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "container", Usage: "Container Name"},
			cli.StringFlag{Name: "blob", Usage: "Blob name"},
		},
		Action: func(c *cli.Context) {
			uploadCommandFunc(c)
		},
	}
}

func uploadCommandFunc(c *cli.Context) error {
	accountName := c.GlobalString("account-name")
	accountKey := c.GlobalString("account-key")
	containerName := c.String("container")
	blobName := c.String("blob")
	fileToUpload := "/file_to_upload"

	if len(c.Args()) != 0 {
		fileToUpload = c.Args()[0]
	}

	client, err := storage.NewBasicClient(accountName, accountKey)

	if err != nil {
		fmt.Println(err)
		return err
	}

	blobService := client.GetBlobService()

	file, err := os.Open(fileToUpload)

	if err != nil {
		fmt.Println(err)
		return err
	}

	err = blobService.PutBlockBlob(containerName, blobName, file)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
