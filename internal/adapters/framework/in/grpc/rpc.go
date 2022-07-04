package grpc

import (
	"context"
	"github.com/Adetunjii/secrot/internal/adapters/framework/in/grpc/pb"
)

func (grpcAdapter *Adapter) Deposit(ctx context.Context, req *pb.DepositRequestParameters) (*pb.User, error) {
	var err error

	usr, err := grpcAdapter.api.Deposit(req.UserId, req.Amount)
	if err != nil {
		return &pb.User{}, err
	}

	user := &pb.User{
		Id:        usr.ID,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
		Balance:   usr.Balance,
	}

	return user, nil
}
