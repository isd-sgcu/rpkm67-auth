package user

import (
	"github.com/google/uuid"
	proto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/auth/user/v1"
	"github.com/isd-sgcu/rpkm67-model/constant"
	"github.com/isd-sgcu/rpkm67-model/model"
)

func ModelToProto(in *model.User) *proto.User {
	protoUser := &proto.User{
		Id:          in.ID.String(),
		Email:       in.Email,
		Nickname:    in.Nickname,
		Title:       in.Title,
		Firstname:   in.Firstname,
		Lastname:    in.Lastname,
		Year:        int32(in.Year),
		Faculty:     in.Faculty,
		Tel:         in.Tel,
		ParentTel:   in.ParentTel,
		Parent:      in.Parent,
		FoodAllergy: in.FoodAllergy,
		DrugAllergy: in.DrugAllergy,
		Illness:     in.Illness,
		Role:        in.Role.String(),
		PhotoKey:    in.PhotoKey,
		PhotoUrl:    in.PhotoUrl,
		Baan:        in.Baan,
		ReceiveGift: int32(in.ReceiveGift),
	}

	if in.GroupID != nil {
		protoUser.GroupId = in.GroupID.String()
	}

	return protoUser
}

func UpdateRequestToModel(in *proto.UpdateUserRequest) (*model.User, error) {
	groupId, err := uuid.Parse(in.GroupId)
	if err != nil {
		return nil, err
	}

	return &model.User{
		Email:       in.Email,
		Nickname:    in.Nickname,
		Title:       in.Title,
		Firstname:   in.Firstname,
		Lastname:    in.Lastname,
		Year:        int(in.Year),
		Faculty:     in.Faculty,
		Tel:         in.Tel,
		ParentTel:   in.ParentTel,
		Parent:      in.Parent,
		FoodAllergy: in.FoodAllergy,
		DrugAllergy: in.DrugAllergy,
		Illness:     in.Illness,
		Role:        constant.Role(in.Role),
		PhotoKey:    in.PhotoKey,
		PhotoUrl:    in.PhotoUrl,
		Baan:        in.Baan,
		GroupID:     &groupId,
		ReceiveGift: int(in.ReceiveGift),
	}, nil
}
