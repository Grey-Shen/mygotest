package libqiniu

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"gitlab.qiniu.io/pingan/libhttpclient/form_postman"
	"gitlab.qiniu.io/pingan/libqiniu/op"
	ctx "golang.org/x/net/context"
)

type PfopClient struct {
	*Zone
	aksk    AccessKeySecretKey
	postman form_postman.Postman
}

func NewPfopClient(postman form_postman.Postman, aksk AccessKeySecretKey, zone *Zone) PfopClient {
	return PfopClient{Zone: zone, aksk: aksk, postman: postman}
}

type (
	PfopParams struct {
		Fops      MultiFopCommands
		NotifyURL string
		Pipeline  string
		Force     bool
	}

	idStruct struct {
		Id string `json:"persistentID"`
	}
)

func (client PfopClient) Pfop(context ctx.Context, entry op.Entry, params PfopParams) (PersistentID, error) {
	var (
		id   idStruct
		auth string
		err  error
	)

	values := make(url.Values)
	values.Set("bucket", entry.Bucket)
	values.Set("key", entry.Key)
	values.Set("fops", params.Fops.ToMultiCommands().String())
	if params.NotifyURL != "" {
		values.Set("notifyURL", params.NotifyURL)
	}
	if params.Force {
		values.Set("force", "1")
	}
	if params.Pipeline != "" {
		values.Set("pipeline", params.Pipeline)
	}

	if auth, err = client.aksk.generateAuthorization("/pfop", []byte(values.Encode())); err != nil {
		return PersistentID{}, err
	} else if err = client.postman.PostForm(context, client.PfopHost+"/pfop", auth, values, &id); err != nil {
		return PersistentID{}, err
	} else {
		return PersistentID{id: id.Id, client: client}, nil
	}
}

func (client PfopClient) Prefop(context ctx.Context, id string) (*PfopStatus, error) {
	persistentID := PersistentID{id: id, client: client}
	return persistentID.QueryStatus(context)
}

type PersistentID struct {
	id     string
	client PfopClient
}

func (persistentID PersistentID) String() string {
	return string(persistentID.id)
}

func (persistentID PersistentID) QueryStatus(context ctx.Context) (*PfopStatus, error) {
	var (
		status   PfopStatus
		response *http.Response
		body     []byte
	)
	if request, err := http.NewRequest("GET", persistentID.client.StatusHost+"/status/get/prefop?id="+persistentID.id, http.NoBody); err != nil {
		return nil, err
	} else {
		if context != nil {
			request = request.WithContext(context)
		}
		if response, err = persistentID.client.postman.SendRequest(request, &status); err != nil {
			return nil, err
		} else if body, err = ioutil.ReadAll(response.Body); err != nil {
			return nil, err
		} else {
			status.originalString = string(body)
			return &status, nil
		}
	}
}
