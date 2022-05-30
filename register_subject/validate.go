package register_subject

import (
	"github.com/gofiber/fiber/v2"
	"shrading/constant"
	"shrading/helper"
	"shrading/shard"
)

func validateTKBSlot(c *fiber.Ctx, tkb *shard.TKB) error {
	if tkb.SoChoConLai < 1 {
		return c.JSON(helper.Response{
			Status:  false,
			Data:    nil,
			Message: "so cho dang ki mon hoc da het",
			Error:   helper.Error{},
		})
	}

	if tkb.SoChoConLai >= uint(helper.StringToInt(tkb.SySo)) {
		return c.JSON(helper.Response{
			Status:  false,
			Data:    nil,
			Message: "huy cho khong thoa man",
			Error:   helper.Error{},
		})
	}

	return nil
}

func validateStruct(c *fiber.Ctx, body *RegisterSubjectBody) error {
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

	return nil
}

func getFail(c *fiber.Ctx) error {
	return c.JSON(helper.Response{
		Status:  false,
		Data:    nil,
		Message: "get TKB fail",
		Error:   helper.Error{},
	})
}

func registSuccess(c *fiber.Ctx) error {
	return c.JSON(helper.Response{
		Status:  true,
		Data:    nil,
		Message: "Regist subject success",
		Error:   helper.Error{},
	})
}

func unregistSuccess(c *fiber.Ctx) error {
	return c.JSON(helper.Response{
		Status:  true,
		Data:    nil,
		Message: "Unregist subject success",
		Error:   helper.Error{},
	})
}

func updateFail(c *fiber.Ctx) error {
	return c.JSON(helper.Response{
		Status:  true,
		Data:    nil,
		Message: "Fail to update",
		Error:   helper.Error{},
	})
}

func createFail(c *fiber.Ctx) error {
	return c.JSON(helper.Response{
		Status:  true,
		Data:    nil,
		Message: "Fail to create",
		Error:   helper.Error{},
	})
}

func deleteFail(c *fiber.Ctx) error {
	return c.JSON(helper.Response{
		Status:  true,
		Data:    nil,
		Message: "Fail to delete",
		Error:   helper.Error{},
	})
}

func checkRegistSubject(count int, c *fiber.Ctx) error {
	if count >= 1 {
		return c.JSON(helper.Response{
			Status:  false,
			Message: "Mon nay da dc dki",
			Data:    nil,
			Error:   helper.Error{},
		})
	}

	return nil
}
