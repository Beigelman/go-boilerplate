package handlers

import (
	"context"

	"github.com/truepay/go-boilerplate/modules/banking/application"
	pb "github.com/truepay/go-boilerplate/modules/banking/interface/grpc/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AccountHandler struct {
	CreateAccount application.CreateAccount
	pb.UnimplementedAccountServicesServer
}

func (self *AccountHandler) CreateAccountHandler(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	newAccount, err := self.CreateAccount(application.CreateAccountDTO{
		UserID: req.UserId,
	})

	if err != nil {
		return nil, err
	}

	return &pb.CreateAccountResponse{
		Id:        newAccount.ID,
		UserId:    newAccount.UserID,
		Balance:   int32(newAccount.Balance),
		Limit:     int32(newAccount.Limit),
		Status:    newAccount.Status,
		CreatedAt: timestamppb.New(newAccount.CreatedAt),
		UpdatedAt: timestamppb.New(newAccount.UpdatedAt),
	}, nil
}
