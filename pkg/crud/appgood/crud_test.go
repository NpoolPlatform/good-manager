package appgood

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/good-manager/pkg/db/ent"
	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	valuedef "github.com/NpoolPlatform/message/npool"
	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/appgood"
	commmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/commission"

	testinit "github.com/NpoolPlatform/good-manager/pkg/testinit"
	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var appGood = ent.AppGood{
	ID:                   uuid.New(),
	AppID:                uuid.New(),
	GoodID:               uuid.New(),
	Online:               true,
	Visible:              true,
	GoodName:             uuid.NewString(),
	Price:                decimal.RequireFromString("9999999999999999999.999999999999999999"),
	DisplayIndex:         100,
	PurchaseLimit:        100,
	CommissionPercent:    100,
	DailyRewardAmount:    decimal.RequireFromString("9999999999999999999.999999999999999999"),
	CommissionSettleType: commmgrpb.SettleType_NoCommission.String(),
}

var (
	id     = appGood.ID.String()
	appID  = appGood.AppID.String()
	goodID = appGood.GoodID.String()
	price  = appGood.Price.String()
	amount = appGood.DailyRewardAmount.String()
	req    = npool.AppGoodReq{
		ID:                &id,
		AppID:             &appID,
		GoodID:            &goodID,
		Online:            &appGood.Online,
		Visible:           &appGood.Visible,
		GoodName:          &appGood.GoodName,
		Price:             &price,
		DisplayIndex:      &appGood.DisplayIndex,
		PurchaseLimit:     &appGood.PurchaseLimit,
		CommissionPercent: &appGood.CommissionPercent,
		DailyRewardAmount: &amount,
	}
)

var info *ent.AppGood

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &req)
	if assert.Nil(t, err) {
		appGood.UpdatedAt = info.UpdatedAt
		appGood.CreatedAt = info.CreatedAt
		assert.Equal(t, info.String(), appGood.String())
	}
}

func createBulk(t *testing.T) {
	entities := []*ent.AppGood{
		{
			ID:                uuid.New(),
			AppID:             uuid.New(),
			GoodID:            uuid.New(),
			Online:            true,
			Visible:           true,
			GoodName:          uuid.NewString(),
			Price:             decimal.RequireFromString("9999999999999999999.999999999999999999"),
			DisplayIndex:      100,
			PurchaseLimit:     100,
			CommissionPercent: 100,
			DailyRewardAmount: decimal.RequireFromString("9999999999999999999.999999999999999999"),
		},
		{
			ID:                uuid.New(),
			AppID:             uuid.New(),
			GoodID:            uuid.New(),
			Online:            true,
			Visible:           true,
			GoodName:          uuid.NewString(),
			Price:             decimal.RequireFromString("9999999999999999999.999999999999999999"),
			DisplayIndex:      100,
			PurchaseLimit:     100,
			CommissionPercent: 100,
			DailyRewardAmount: decimal.RequireFromString("9999999999999999999.999999999999999999"),
		},
	}

	reqs := []*npool.AppGoodReq{}
	for _, _appGood := range entities {
		_id := _appGood.ID.String()
		_appID := _appGood.AppID.String()
		_goodID := _appGood.GoodID.String()
		_price := _appGood.Price.String()
		_amount := _appGood.DailyRewardAmount.String()
		reqs = append(reqs, &npool.AppGoodReq{
			ID:                &_id,
			AppID:             &_appID,
			GoodID:            &_goodID,
			Online:            &_appGood.Online,
			Visible:           &_appGood.Visible,
			GoodName:          &_appGood.GoodName,
			Price:             &_price,
			DisplayIndex:      &_appGood.DisplayIndex,
			PurchaseLimit:     &_appGood.PurchaseLimit,
			CommissionPercent: &_appGood.CommissionPercent,
			DailyRewardAmount: &_amount,
		})
	}
	infos, err := CreateBulk(context.Background(), reqs)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func update(t *testing.T) {
	var err error
	info, err = Update(context.Background(), &req)
	if assert.Nil(t, err) {
		appGood.UpdatedAt = info.UpdatedAt
		appGood.CreatedAt = info.CreatedAt
		assert.Equal(t, info.String(), appGood.String())
	}
}
func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), appGood.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), appGood.String())
	}
}

func rows(t *testing.T) {
	infos, total, err := Rows(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		}, 0, 0)
	if assert.Nil(t, err) {
		if assert.Equal(t, total, 1) {
			assert.Equal(t, infos[0].String(), appGood.String())
		}
	}
}

func rowOnly(t *testing.T) {
	var err error
	info, err = RowOnly(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), appGood.String())
	}
}

func count(t *testing.T) {
	count, err := Count(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, count, uint32(1))
	}
}

func exist(t *testing.T) {
	exist, err := Exist(context.Background(), appGood.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existConds(t *testing.T) {
	exist, err := ExistConds(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteA(t *testing.T) {
	info, err := Delete(context.Background(), appGood.ID.String())
	if assert.Nil(t, err) {
		appGood.DeletedAt = info.DeletedAt
		appGood.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info.String(), appGood.String())
	}
}

func TestDetail(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("create", create)
	t.Run("createBulk", createBulk)
	t.Run("update", update)
	t.Run("row", row)
	t.Run("rows", rows)
	t.Run("rowOnly", rowOnly)
	t.Run("exist", exist)
	t.Run("existConds", existConds)
	t.Run("count", count)
	t.Run("delete", deleteA)
}
