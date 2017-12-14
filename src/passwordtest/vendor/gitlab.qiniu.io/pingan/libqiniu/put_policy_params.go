package libqiniu

type PutPolicyCallbackParams struct {
	urls        []string
	host        string
	contentType string
	body        string
}

func NewPutPolicyCallbackParams(urls ...string) PutPolicyCallbackParams {
	return PutPolicyCallbackParams{
		urls: urls,
	}
}

func (self PutPolicyCallbackParams) SetHost(host string) PutPolicyCallbackParams {
	self.host = host
	return self
}

func (self PutPolicyCallbackParams) SetBody(body, contentType string) PutPolicyCallbackParams {
	self.contentType = contentType
	self.body = body
	return self
}

type PutPolicyPfopParams struct {
	cmd       string
	notifyURL string
	pipeline  string
}

func NewPutPolicyPfopParams(fops MultiFopCommands) PutPolicyPfopParams {
	return PutPolicyPfopParams{
		cmd: fops.ToMultiCommands().String(),
	}
}

func (self PutPolicyPfopParams) SetNotifyURL(notifyURL string) PutPolicyPfopParams {
	self.notifyURL = notifyURL
	return self
}

func (self PutPolicyPfopParams) SetPipeline(pipeline string) PutPolicyPfopParams {
	self.pipeline = pipeline
	return self
}
