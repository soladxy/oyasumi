package ssl

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/soladxy/oyasumi/biz/container"
	"time"
)

type _ssl struct {
	c *container.Container
}

func newSsl(c *container.Container) *_ssl {
	return &_ssl{
		c: c,
	}
}

func (w *_ssl) HostExpired(ctx context.Context, host, port string) (bool, string, error) {
	// 连接到目标主机
	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%s", host, port), &tls.Config{})
	if err != nil {
		hlog.CtxErrorf(ctx, "[GetLatestWebsiteInfo] tls dial err: %v", err)
		return false, "", err
	}
	defer conn.Close()

	// 获取服务端的证书链
	state := conn.ConnectionState()
	var expiredTime time.Time
	now := time.Now()
	for i, cert := range state.PeerCertificates {
		hlog.CtxInfof(ctx, "[GetLatestWebsiteInfo] Issuer: %s, Subject: %s, Expiry: %s", cert.Issuer, cert.Subject, cert.NotAfter.Format(time.DateTime))

		if i == 0 {
			expiredTime = cert.NotAfter
		}

		// 检查证书是否过期
		if now.Before(cert.NotBefore) {
			return true, fmt.Sprintf("证书未生效，距离生效还有%d，生效开始时间%s", cert.NotBefore.Sub(now)/(time.Hour*24), cert.NotBefore.Local().Format(time.DateTime)), nil
		}
		if now.After(cert.NotAfter) {
			return true, fmt.Sprintf("证书已过期，已过期%d天, 过期时间%s", now.Sub(cert.NotAfter)/(time.Hour*24)+1, cert.NotAfter.Local().Format(time.DateTime)), nil
		}
	}

	return false, fmt.Sprintf("证书生效中，剩余%d天, 过期时间%s", expiredTime.Sub(now)/(time.Hour*24), expiredTime.Local().Format(time.DateTime)), nil
}
