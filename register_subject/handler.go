package register_subject

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"shrading/connection"
	"shrading/constant"
	"shrading/helper"
	"shrading/shard"
	"sync"
)

func RegistSubject(c *fiber.Ctx) error {
	body := new(RegisterSubjectBody)
	if err := c.BodyParser(body); err != nil {
		return c.JSON(helper.Response{
			Status:  false,
			Data:    nil,
			Message: "Missing params.",
			Error: helper.Error{
				ErrorCode:    constant.ErrorCode["ERROR_MISSING_PARAMS"],
				ErrorMessage: "Missing params.",
			},
		})
	}
	if errVC := helper.ValidateStruct(body); errVC != nil {
		return c.JSON(helper.Response{
			Status:  false,
			Data:    nil,
			Message: "Missing params.",
			Error: helper.Error{
				ErrorCode:    constant.ErrorCode["ERROR_MISSING_PARAMS"],
				ErrorMessage: "Missing params.",
			},
		})
	}

	var wg sync.WaitGroup
	var regist []*shard.RegisterSubject

	id := uint(helper.HashToInt(body.MaMon + body.NhomLop + "1464"))
	count, err := shard.Cluster.Shard(int64(id)).Model(&regist).Where("ma_sv = ? AND ma_mon_hoc = ?", body.MaSV, body.MaMon).Count()
	if err != nil {
		return c.JSON(helper.Response{
			Status:  true,
			Data:    nil,
			Message: "Fail to create",
			Error:   helper.Error{},
		})
	}
	if count > 0 {
		return c.JSON(helper.Response{
			Status:  false,
			Message: "Mon nay da dc dki",
			Data:    nil,
			Error:   helper.Error{},
		})
	}

	rs := &shard.RegisterSubject{
		ID:       id,
		MaSV:     body.MaSV,
		NhomLop:  body.NhomLop,
		MaMonHoc: body.MaMon,
	}
	tkb, err := shard.NewTKB().GetTKB(shard.Cluster, int64(id))
	if err != nil {
		return c.JSON(helper.Response{
			Status:  false,
			Data:    nil,
			Message: "get TKB fail",
			Error:   helper.Error{},
		})
	}
	if errVTS := validateTKBSlot(c, tkb); errVTS != nil {
		return errVTS
	}

	wg.Add(2)

	go func(wg *sync.WaitGroup) error {
		defer wg.Done()

		err = shard.NewRS().CreateRS(shard.Cluster, rs)
		if err != nil {
			return c.JSON(helper.Response{
				Status:  true,
				Data:    nil,
				Message: "Fail to create",
				Error:   helper.Error{},
			})
		}
		return nil
	}(&wg)

	go func(wg *sync.WaitGroup) error {
		defer wg.Done()
		tkb.SoChoConLai -= 1
		err1 := shard.NewTKB().UpdateTKB(shard.Cluster, tkb)
		if err1 != nil {
			return c.JSON(helper.Response{
				Status:  true,
				Data:    nil,
				Message: "Fail to update",
				Error:   helper.Error{},
			})
		}
		return nil
	}(&wg)

	wg.Wait()

	return c.JSON(helper.Response{
		Status:  true,
		Data:    nil,
		Message: "Unregist subject success",
		Error:   helper.Error{},
	})

}

func UnregistSubject(c *fiber.Ctx) error {

	body := new(RegisterSubjectBody)
	if err := c.BodyParser(body); err != nil {
		return c.JSON(helper.Response{
			Status:  false,
			Data:    nil,
			Message: "Missing params.",
			Error: helper.Error{
				ErrorCode:    constant.ErrorCode["ERROR_MISSING_PARAMS"],
				ErrorMessage: "Missing params.",
			},
		})
	}
	if errVC := helper.ValidateStruct(body); errVC != nil {
		return c.JSON(helper.Response{
			Status:  false,
			Data:    nil,
			Message: "Missing params.",
			Error: helper.Error{
				ErrorCode:    constant.ErrorCode["ERROR_MISSING_PARAMS"],
				ErrorMessage: "Missing params.",
			},
		})
	}

	var wg sync.WaitGroup

	rs := &shard.RegisterSubject{
		ID:       uint(helper.HashToInt(body.MaMon + body.NhomLop + "1464")),
		MaSV:     body.MaSV,
		NhomLop:  body.NhomLop,
		MaMonHoc: body.MaMon,
	}
	tkb, err := shard.NewTKB().GetTKB(shard.Cluster, int64(helper.HashToInt(body.MaMon+body.NhomLop+"1464")))
	if err != nil {
		return c.JSON(helper.Response{
			Status:  false,
			Data:    nil,
			Message: "get TKB fail",
			Error:   helper.Error{},
		})
	}
	if err := validateTKBSlot(c, tkb); err != nil {
		return err
	}

	wg.Add(2)

	go func(wg *sync.WaitGroup) error {
		defer wg.Done()
		err := shard.NewRS().DeleteRS(shard.Cluster, rs)
		if err != nil {
			return c.JSON(helper.Response{
				Status:  true,
				Data:    nil,
				Message: "Fail to delete",
				Error:   helper.Error{},
			})
		}
		return nil
	}(&wg)

	go func(wg *sync.WaitGroup) error {
		defer wg.Done()
		tkb.SoChoConLai += 1
		err1 := shard.NewTKB().UpdateTKB(shard.Cluster, tkb)
		if err1 != nil {
			return c.JSON(helper.Response{
				Status:  true,
				Data:    nil,
				Message: "Fail to update",
				Error:   helper.Error{},
			})
		}

		return nil
	}(&wg)

	wg.Wait()

	return c.JSON(helper.Response{
		Status:  true,
		Data:    nil,
		Message: "Unregist subject success",
		Error:   helper.Error{},
	})
}

func checkIntoDB(check int, numShard int) *gorm.DB {
	result := check % numShard
	switch result {
	case 0:
		return connection.PostgresConn(constant.Sharding1)
	case 1:
		return connection.PostgresConn(constant.Sharding2)
	case 2:
		return connection.PostgresConn(constant.Sharding3)
	}
	return nil
}
