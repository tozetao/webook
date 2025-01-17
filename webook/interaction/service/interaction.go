package service

import (
	"context"
	"learn_go/webook/interaction/domain"
	repository "learn_go/webook/interaction/repository"
)

//go:generate mockgen -source=./interaction.go -package=svcmocks -destination=./mocks/interaction.mock.go InteractionService
type InteractionService interface {
	View(ctx context.Context, biz string, bizID int64) error

	Like(ctx context.Context, uid int64, biz string, bizID int64) error
	CancelLike(ctx context.Context, uid int64, biz string, bizID int64) error
	Favorite(ctx context.Context, uid int64, favoriteID int64, biz string, bizID int64) error

	// Get 查询bizID的交互数据，以及用户id（uid）对应的交互数据
	Get(ctx context.Context, uid int64, biz string, bizID int64) (domain.Interaction, error)

	// Liked 用户是否点赞
	Liked(ctx context.Context, uid int64, biz string, bizID int64) (bool, error)
	// Collected 用户是否收藏
	Collected(ctx context.Context, uid int64, biz string, bizID int64) (bool, error)

	GetByIDs(ctx context.Context, biz string, bizIDs []int64) (map[int64]domain.Interaction, error)
}

func (svc *interactionService) Liked(ctx context.Context, uid int64, biz string, bizID int64) (bool, error) {
	_, err := svc.repo.GetUserLikeInfo(ctx, uid, biz, bizID)
	switch err {
	case nil:
		return true, nil
	case repository.ErrNotFound:
		// 吞掉找不到的错误
		return false, nil
	default:
		return false, err
	}
}

func (svc *interactionService) Collected(ctx context.Context, uid int64, biz string, bizID int64) (bool, error) {
	// TODO: implement me
	return false, nil
}

func (svc *interactionService) Favorite(ctx context.Context, uid int64, favoriteID int64, biz string, bizID int64) error {
	return svc.repo.AddFavoriteItem(ctx, uid, favoriteID, biz, bizID)
}

type interactionService struct {
	repo repository.InteractionRepository
}

func (svc *interactionService) GetByIDs(ctx context.Context, biz string, bizIDs []int64) (map[int64]domain.Interaction, error) {
	inters, err := svc.repo.GetByIDs(ctx, biz, bizIDs)
	if err != nil {
		return nil, err
	}
	res := make(map[int64]domain.Interaction, len(inters))
	for _, inter := range inters {
		res[inter.ID] = inter
	}
	return res, nil
}

func (svc *interactionService) Get(ctx context.Context, uid int64, biz string, bizID int64) (domain.Interaction, error) {
	return svc.repo.Get(ctx, uid, biz, bizID)
}

func (svc *interactionService) Like(ctx context.Context, uid int64, biz string, bizID int64) error {
	return svc.repo.IncrLike(ctx, uid, biz, bizID)
}

func (svc *interactionService) CancelLike(ctx context.Context, uid int64, biz string, bizID int64) error {
	return svc.repo.DecrLike(ctx, uid, biz, bizID)
}

// View 查看文章：增加文章点击量
func (svc *interactionService) View(ctx context.Context, biz string, bizID int64) error {
	return svc.repo.IncrReadCnt(ctx, biz, bizID)
}

func NewInteractionService(repo repository.InteractionRepository) InteractionService {
	return &interactionService{
		repo: repo,
	}
}
