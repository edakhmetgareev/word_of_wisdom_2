package server

type Server struct {
	tcpConn Conn
}

func NewServer(conn Conn) *Server {
	return &Server{tcpConn: conn}
}
