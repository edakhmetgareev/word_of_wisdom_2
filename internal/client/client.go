package client

type Client struct {
	tcpConn Conn
}

func NewClient(conn Conn) *Client {
	return &Client{tcpConn: conn}
}
