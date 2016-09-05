package command

import (
	"fmt"
	"github.com/MSOpenTech/azure-sdk-for-go/storage"
	"github.com/codegangsta/cli"
)

func NewLsCommand() cli.Command {
	return cli.Command{
		Name:  "ls",
		Usage: "retrieve container blobs",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "container", Usage: "Container Name"},
		},
		Action: func(c *cli.Context) {
			lsCommandFunc(c)
		},
	}
}

// lsCommandFunc executes the "ls" command.
func lsCommandFunc(c *cli.Context) error {
	accountName := c.GlobalString("account-name")
	accountKey := c.GlobalString("account-key")
	containerName := c.String("container")
	client, err := storage.NewBasicClient(accountName, accountKey)

	if err != nil {
		fmt.Println(err)
		return err
	}

	blobService := client.GetBlobService()

	blobs, err := blobService.ListBlobs(containerName, storage.ListBlobsParameters{})

	if err != nil {
		fmt.Println(err)
	} else {
		for _, file := range blobs.Blobs {
			fmt.Println(fmt.Sprintf("blob -> %+v", file))
		}
	}
	return err
}
