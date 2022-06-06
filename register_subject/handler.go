package register_subject

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"shrading/connection"
	"shrading/constant"
	"shrading/helper"
	"shrading/model"
	"shrading/shard"
	"strconv"
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
				ErrorMessage: "Missing params1.",
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
				ErrorMessage: "Missing params2.",
			},
		})
	}

	var wg sync.WaitGroup

	errorChanel := make(chan error, 3)
	id := uint(helper.HashToInt(body.MaMon + body.NhomLop + strconv.Itoa(body.IDMon)))

	db := shard.Cluster.Shard(int64(id))
	tx, err := db.Begin()

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{
						"match": map[string]interface{}{
							"MaSV": body.MaSV,
						},
					},
					{
						"match": map[string]interface{}{
							"MaMonHoc": body.MaMon,
						},
					},
				},
			},
		},
	}
	_, totalRecord, errorES := helper.QueryES("regist_subject", query)
	if errorES != nil {
		return c.JSON(helper.Response{
			Status:  false,
			Data:    nil,
			Message: "Fail to get in elasticsearch",
			Error:   helper.Error{},
		})
	}
	if totalRecord > 0 {
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
	if tkb.SoChoConLai < 1 {
		return c.JSON(helper.Response{
			Status:  false,
			Data:    nil,
			Message: "so cho dang ki mon hoc da het",
			Error:   helper.Error{},
		})
	}
	if tkb.SoChoConLai > uint(helper.StringToInt(tkb.SySo)) {
		return c.JSON(helper.Response{
			Status:  false,
			Data:    nil,
			Message: "huy cho khong thoa man",
			Error:   helper.Error{},
		})
	}

	wg.Add(3)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		//err = shard.NewRS().CreateRS(shard.Cluster, rs)
		errI := tx.Insert(rs)
		if errI != nil {
			errorChanel <- c.JSON(helper.Response{
				Status:  false,
				Data:    nil,
				Message: "Fail to create",
				Error:   helper.Error{},
			})
			return
		}

		errorChanel <- nil
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		_, err = helper.InsertToElastic(rs, "regist_subject", strconv.Itoa(int(rs.ID)), "_doc")
		if err != nil {
			errorChanel <- c.JSON(helper.Response{
				Status:  false,
				Data:    nil,
				Message: "Fail to create in elasticsearch",
				Error:   helper.Error{},
			})
			return
		}
		errorChanel <- nil
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		tkb.SoChoConLai -= 1
		errU := tx.Update(tkb)
		if errU != nil {
			errorChanel <- c.JSON(helper.Response{
				Status:  false,
				Data:    nil,
				Message: "Fail to update",
				Error:   helper.Error{},
			})
			return
		}
		errorChanel <- nil
	}(&wg)

	wg.Wait()

	for i := 0; i < len(errorChanel); i++ {
		if err := <-errorChanel; err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return c.JSON(helper.Response{
			Status:  true,
			Data:    nil,
			Message: "Commit transaction fail",
			Error:   helper.Error{},
		})
	}

	return c.JSON(helper.Response{
		Status:  true,
		Data:    nil,
		Message: "Regist subject success",
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
	errorChanel := make(chan error, 3)
	id := uint(helper.HashToInt(body.MaMon + body.NhomLop + strconv.Itoa(body.IDMon)))

	db := shard.Cluster.Shard(int64(id))
	tx, err := db.Begin()

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{
						"match": map[string]interface{}{
							"MaSV": body.MaSV,
						},
					},
					{
						"match": map[string]interface{}{
							"MaMonHoc": body.MaMon,
						},
					},
				},
			},
		},
	}
	_, totalRecord, errorES := helper.QueryES("regist_subject", query)
	if errorES != nil {
		return c.JSON(helper.Response{
			Status:  false,
			Data:    nil,
			Message: "Fail to get in elasticsearch",
			Error:   helper.Error{},
		})
	}
	if totalRecord == 0 {
		return c.JSON(helper.Response{
			Status:  false,
			Message: "Can't delete because this object doesn't exist",
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

	wg.Add(3)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		//err := shard.NewRS().DeleteRS(shard.Cluster, rs)
		err := tx.Delete(rs)
		if err != nil {
			errorChanel <- c.JSON(helper.Response{
				Status:  false,
				Data:    nil,
				Message: "Fail to delete",
				Error:   helper.Error{},
			})
			return
		}
		errorChanel <- nil
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		_, err = helper.DeleteByQueryES("regist_subject", query)
		if err != nil {
			errorChanel <- c.JSON(helper.Response{
				Status:  false,
				Data:    nil,
				Message: "Fail to delete in elasticsearch",
				Error:   helper.Error{},
			})
			return
		}
		errorChanel <- nil
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		tkb.SoChoConLai += 1
		err1 := tx.Update(tkb)
		if err1 != nil {
			errorChanel <- c.JSON(helper.Response{
				Status:  false,
				Data:    nil,
				Message: "Fail to update",
				Error:   helper.Error{},
			})
			return
		}

		errorChanel <- nil
	}(&wg)

	wg.Wait()

	for i := 0; i < len(errorChanel); i++ {
		if err := <-errorChanel; err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return c.JSON(helper.Response{
			Status:  true,
			Data:    nil,
			Message: "Commit transaction fail",
			Error:   helper.Error{},
		})
	}

	return c.JSON(helper.Response{
		Status:  true,
		Data:    nil,
		Message: "Unregist subject success",
		Error:   helper.Error{},
	})
}

func UploadDataToES(c *fiber.Ctx) error {
	err := model.SyncTKBToElasticSearch()
	if err != nil {
		return c.JSON(helper.Response{
			Status:  false,
			Data:    nil,
			Message: "fail to upload to elasticsearch",
			Error:   helper.Error{},
		})
	}

	return c.JSON(helper.Response{
		Status:  true,
		Data:    nil,
		Message: "upload to elasticsearch successfully",
		Error:   helper.Error{},
	})
}

func checkIntoDB(check uint, numShard uint) *gorm.DB {
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
