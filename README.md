[![Docker Repository on Quay.io](https://quay.io/repository/pharmpress/azurectl/status "Docker Repository on Quay.io")](https://quay.io/repository/pharmpress/azurectl)

# azurectl
`azurectl` is a command line client for azure storage.

```
docker pull rpsrepo.ddns.net:9010/azurectl

#upload
docker run -v $(pwd):/tmp:ro  rpsrepo.ddns.net:9010/azurectl --account-name my_account --account-key XXXXX upload --container my_container --blob file_to_upload.txt /tmp/file_to_upload.txt

#ls
docker run rpsrepo.ddns.net:9010/azurectl --account-name my_account --account-key XXXXX ls --container my_container
```
