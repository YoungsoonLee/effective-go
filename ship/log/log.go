package main

import "go.uber.org/zap"

// Server represents the server.
type Server struct {
	logger *zap.Logger
}

// Login authenticates user from token and returns user id.
func (s *Server) Login(token string) (int, error) {
	uid, err := auth.Authenticate(token)
	if err != nil {
		return 0, err
	}

	// s.logger.Info("user logged in", zap.Any("user", users.Load(uid)))
	if info := s.logger.Check(zap.InfoLevel, "user logged in"); info != nil {
		info.Write(zap.Int("user", users.Load(uid)))
	}

	return uid, nil
}

type auth struct{}

func (a auth) Authenticate(token string) (int, error) {
	return 0, nil
}

type users struct{}

func (u users) Load(uid int) int {
	return 0
}

func main() {

}
