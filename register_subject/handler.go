package register_subject

import (
	"github.com/gofiber/fiber/v2"
	"shrading/constant"
	"shrading/helper"
	"shrading/shard"
)

func RegisterSubject(c *fiber.Ctx) error {

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

	rs := &shard.RegisterSubject{
		ID:       5,
		MaSV:     body.MaSV,
		MaMonHoc: body.MaMon,
	}

	tkb, err := shard.GetTKB(shard.Cluster, 8)
	if err != nil {
		return c.JSON(helper.Response{
			Status:  false,
			Data:    nil,
			Message: err.Error(),
			Error:   helper.Error{},
		})
	}
	tkb.SoChoConLai = "86"
	//tkb := &shard.TKB{ID: 8, SoChoConLai: "86"}

	err = shard.CreateRS(shard.Cluster, rs)
	if err != nil {
		return c.JSON(helper.Response{
			Status:  false,
			Data:    nil,
			Message: "Create RS fail",
			Error:   helper.Error{},
		})
	}

	err1 := shard.UpdateTKB(shard.Cluster, tkb)
	if err1 != nil {
		return c.JSON(helper.Response{
			Status:  false,
			Data:    nil,
			Message: "Update TKB fail",
			Error:   helper.Error{},
		})
	}

	return nil
}
