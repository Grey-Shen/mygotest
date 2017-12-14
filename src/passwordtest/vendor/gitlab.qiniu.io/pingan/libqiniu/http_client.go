package libqiniu

import (
	"net/http"

	"gitlab.qiniu.io/pingan/libhttpclient/form_postman"
)

type Client struct {
	Downloader
	client *http.Client
}

func NewClient(client *http.Client) Client {
	return Client{
		Downloader: NewDownloader(client),
		client:     client,
	}
}

func (client Client) Authorize(aksk AccessKeySecretKey) AuthorizedClient {
	return AuthorizedClient{
		postman:            form_postman.NewPostman(client.client),
		AccessKeySecretKey: aksk,
		Client:             client,
	}
}

type AuthorizedClient struct {
	Client
	AccessKeySecretKey
	postman form_postman.Postman
}

func (client AuthorizedClient) Zone(zone *Zone) ZonedClient {
	return ZonedClient{
		postman:          client.postman,
		Zone:             zone,
		AuthorizedClient: client,
		FormUploader:     NewFormUploader(client.postman, zone),
		PfopClient:       NewPfopClient(client.postman, client.AccessKeySecretKey, zone),
	}
}

type ZonedClient struct {
	*Zone
	AuthorizedClient
	FormUploader
	PfopClient
	postman form_postman.Postman
}

func (client ZonedClient) Bucket(name string) *Bucket {
	return &Bucket{
		name:               name,
		postman:            client.postman,
		pfopClient:         client.PfopClient,
		Zone:               client.Zone,
		AccessKeySecretKey: client.AccessKeySecretKey,
	}
}
