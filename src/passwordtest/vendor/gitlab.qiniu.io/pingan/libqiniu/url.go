package libqiniu

import (
	"net/url"
	"strconv"
	"time"
)

type urlWithoutDeadline struct {
	baseURL *url.URL
	aksk    AccessKeySecretKey
	fop     string
}

func (u *urlWithoutDeadline) SetDeadline(deadline time.Time) *URL {
	return &URL{
		urlWithoutDeadline: u,
		deadline:           deadline.Unix(),
	}
}

func (u *urlWithoutDeadline) SetLifetime(lifetime time.Duration) *URL {
	return &URL{
		urlWithoutDeadline: u,
		deadline:           time.Now().Add(lifetime).Unix(),
	}
}

func (u *urlWithoutDeadline) HTTP() *urlWithoutDeadline {
	u.baseURL.Scheme = "http"
	return u
}

func (u *urlWithoutDeadline) HTTPs() *urlWithoutDeadline {
	u.baseURL.Scheme = "https"
	return u
}

func (u *urlWithoutDeadline) SetHTTPs(https bool) *urlWithoutDeadline {
	if https {
		return u.HTTPs()
	} else {
		return u.HTTP()
	}
}

func (u *urlWithoutDeadline) SetFop(cmd FopPipeline) *urlWithoutDeadline {
	u.fop = cmd.ToPipeline().String()
	return u
}

func (u *urlWithoutDeadline) PublicURL() string {
	publicURL := *u.baseURL
	if u.fop != "" {
		if publicURL.RawQuery != "" {
			publicURL.RawQuery += "&" + u.fop
		} else {
			publicURL.RawQuery = u.fop
		}
	}

	return publicURL.String()
}

type URL struct {
	*urlWithoutDeadline
	deadline int64
}

func (u *URL) HTTP() *URL {
	u.baseURL.Scheme = "http"
	return u
}

func (u *URL) HTTPs() *URL {
	u.baseURL.Scheme = "https"
	return u
}

func (u *URL) SetHTTPs(https bool) *URL {
	if https {
		return u.HTTPs()
	} else {
		return u.HTTP()
	}
}

func (u *URL) SetFop(cmd FopPipeline) *URL {
	u.fop = cmd.ToPipeline().String()
	return u
}

func (u *URL) PrivateURL() string {
	privateURL := *u.baseURL

	if u.fop != "" {
		if privateURL.RawQuery != "" {
			privateURL.RawQuery += "&" + u.fop
		} else {
			privateURL.RawQuery = u.fop
		}
	}

	deadline := strconv.FormatInt(u.deadline, 10)
	if privateURL.RawQuery != "" {
		privateURL.RawQuery += "&e=" + deadline
	} else {
		privateURL.RawQuery = "e=" + deadline
	}
	baseURL := privateURL.String()
	token := u.aksk.AccessKey + ":" + u.aksk.sign([]byte(baseURL))
	return baseURL + "&token=" + token
}
