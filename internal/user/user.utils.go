package user

import (
	"github.com/google/uuid"
	proto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/auth/user/v1"
	stampProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/backend/stamp/v1"
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
		GroupId:     in.GroupID.String(),
		ReceiveGift: int32(in.ReceiveGift),
		Stamp: &stampProto.Stamp{
			Id:     in.Stamp.ID.String(),
			PointA: int32(in.Stamp.PointA),
			PointB: int32(in.Stamp.PointB),
			PointC: int32(in.Stamp.PointC),
			PointD: int32(in.Stamp.PointD),
			Stamp:  in.Stamp.Stamp,
		},
	}

	return protoUser
}

func UpdateRequestToModel(in *proto.UpdateUserRequest) (*model.User, error) {
	user := &model.User{
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
		ReceiveGift: int(in.ReceiveGift),
	}

	if in.GroupId != "" {
		groupId, err := uuid.Parse(in.GroupId)
		if err != nil {
			return nil, err
		}
		user.GroupID = &groupId
	}

	return user, nil
}

func NewStampModel(userId *uuid.UUID) *model.Stamp {
	return &model.Stamp{
		UserID: userId,
		PointA: 0,
		PointB: 0,
		PointC: 0,
		PointD: 0,
		Stamp:  "00000000000",
	}
}

func NewGroupModel(userId *uuid.UUID) *model.Group {
	return &model.Group{
		LeaderID:    userId,
		IsConfirmed: false,
	}
}
