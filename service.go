package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/savageking-io/ogbsteam/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"io"
	"net"
	"net/http"
)

type Service struct {
	proto.UnimplementedSteamServiceServer
	config *Config
}

func NewService(config *Config) *Service {
	return &Service{
		config: config,
	}
}

func (s *Service) Init() error {
	if s.config == nil {
		return fmt.Errorf("config is nil")
	}

	if s.config.Steam.AppId == 0 {
		return fmt.Errorf("steam app id is not set")
	}

	if s.config.Steam.PublisherId == "" {
		return fmt.Errorf("steam publisher id is empty")
	}

	if s.config.Steam.Key == "" {
		return fmt.Errorf("steam key is empty")
	}

	return nil
}

func (s *Service) Start() error {
	log.Infof("Starting service on %s:%d", s.config.Rpc.Hostname, s.config.Rpc.Port)
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.config.Rpc.Hostname, s.config.Rpc.Port))
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	proto.RegisterSteamServiceServer(grpcServer, s)
	if err := grpcServer.Serve(lis); err != nil {
		return err
	}
	return nil
}

func (s *Service) AuthUserTicket(ctx context.Context, req *proto.UserTicket) (*proto.AuthUserTicketResponse, error) {
	ticket := string(req.Ticket)

	data := &AuthTicketRequestPayload{
		Key:      s.config.Steam.Key,
		AppId:    s.config.Steam.AppId,
		Ticket:   ticket,
		Identity: s.config.Steam.PublisherId,
	}

	payload := fmt.Sprintf("key=%s&appid=%d&ticket=%s&identity=%s", data.Key, data.AppId, data.Ticket, data.Identity)
	url := fmt.Sprintf("%s%s?%s", APIBackend, "/ISteamUserAuth/AuthenticateUserTicket/v1/", payload)

	request, err := http.NewRequest("GET", url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		log.Errorf("Failed to create request: %s", err.Error())
		return nil, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Errorf("Failed to send request: %s", err.Error())
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Errorf("Failed to read response: %s", err.Error())
		return nil, err
	}

	result := &AuthTicketResponsePayload{}
	err = json.Unmarshal(body, result)
	if err != nil {
		log.Errorf("Failed to unmarshal response: %s", err.Error())
		return nil, err
	}

	return &proto.AuthUserTicketResponse{
		Result:            result.Result,
		SteamId:           result.SteamId,
		OwnerSteamId:      result.OwnerSteamId,
		IsVacBanned:       result.VacBanned,
		IsPublisherBanned: result.PublisherBanned,
	}, nil
}
