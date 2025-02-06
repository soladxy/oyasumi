package ssl

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/soladxy/oyasumi/biz/container"
	"net"
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
	// 设置超时
	dialer := &net.Dialer{
		Timeout: 10 * time.Second,
	}
	// 建立TLS连接（忽略证书验证）
	conn, err := tls.DialWithDialer(dialer, "tcp", host+":"+port, &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		return true, "", fmt.Errorf("连接失败: %v", err)
	}
	defer conn.Close()

	// 获取证书链
	certs := conn.ConnectionState().PeerCertificates
	if len(certs) == 0 {
		return true, "", fmt.Errorf("未找到证书")
	}

	// 检查第一个证书（服务器证书）
	cert := certs[0]
	if time.Now().After(cert.NotAfter) {
		return true, fmt.Sprintf("证书已过期，过期时间: %s", cert.NotAfter), nil
	}
	return false,
		fmt.Sprintf("证书生效中，剩余%d天, 过期时间%s\n颁发给:%s\n颁发者: %s",
			int(cert.NotAfter.Sub(time.Now()).Hours()/24),
			cert.NotAfter.Local().Format(time.DateTime),
			cert.Subject.CommonName,
			cert.Issuer.CommonName,
		),
		nil
}
