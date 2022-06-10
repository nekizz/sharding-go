package register_subject

import (
	"errors"
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
				ErrorMessage: "Missing params.",
			},
		})
	}

	var wg sync.WaitGroup
	var tkbLock shard.TKB

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
			Message: "This subject have already registed",
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
			Message: "This subject doesn't exist",
			Error:   helper.Error{},
		})
	}
	if tkb.SoChoConLai < 1 {
		return c.JSON(helper.Response{
			Status:  false,
			Data:    nil,
			Message: "This subject is out of slot",
			Error:   helper.Error{},
		})
	}
	if tkb.SoChoConLai > uint(helper.StringToInt(tkb.SySo)) {
		return c.JSON(helper.Response{
			Status:  false,
			Data:    nil,
			Message: "Invalid regist action with this subject",
			Error:   helper.Error{},
		})
	}

	errA := db.Model(&tkbLock).Where("id = ?", id).For("UPDATE").Column("so_cho_con_lai").Select()
	if errA != nil {
		return c.JSON(helper.Response{
			Status:  false,
			Data:    nil,
			Message: errA.Error(),
			Error:   helper.Error{},
		})
	}

	wg.Add(3)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		//err = shard.NewRS().CreateRS(shard.Cluster, rs)
		errI := tx.Insert(rs)
		if errI != nil {
			errorChanel <- errors.New("Fail to insert to database")
			return
		}

		errorChanel <- nil
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		_, err = helper.InsertToElastic(rs, "regist_subject", strconv.Itoa(int(rs.ID)), "_doc")
		if err != nil {
			errorChanel <- errors.New("Fail to insert to elasticsearch")
			return
		}
		errorChanel <- nil
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		tkb.SoChoConLai -= 1
		errU := tx.Update(tkb)
		if errU != nil {
			errorChanel <- errors.New("Fail to update to tkb")
			return
		}
		errorChanel <- nil
	}(&wg)

	wg.Wait()

	for i := 0; i < len(errorChanel); i++ {
		if err := <-errorChanel; err != nil {
			if errR := tx.Rollback(); errR != nil {
				return c.JSON(helper.Response{
					Status:  false,
					Data:    nil,
					Message: "Rollback fail",
					Error:   helper.Error{},
				})
			}
			_, errES := helper.DeleteByQueryES("regist_subject", query)
			if errES != nil {
				return c.JSON(helper.Response{
					Status:  false,
					Data:    nil,
					Message: "Rollback ES fail",
					Error:   helper.Error{},
				})
			}
			return c.JSON(helper.Response{
				Status:  false,
				Data:    nil,
				Message: err.Error(),
				Error:   helper.Error{},
			})
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
	var tkbLock shard.TKB

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
			Message: "Fail to get TKB",
			Error:   helper.Error{},
		})
	}
	if tkb.SoChoConLai < 1 {
		return c.JSON(helper.Response{
			Status:  false,
			Data:    nil,
			Message: "This subject is out of slot",
			Error:   helper.Error{},
		})
	}
	if tkb.SoChoConLai >= uint(helper.StringToInt(tkb.SySo)) {
		return c.JSON(helper.Response{
			Status:  false,
			Data:    nil,
			Message: "Invalid unregist action of this subject",
			Error:   helper.Error{},
		})
	}

	errA := db.Model(&tkbLock).Where("id = ?", id).For("UPDATE").Column("so_cho_con_lai").Select()
	if errA != nil {
		return c.JSON(helper.Response{
			Status:  false,
			Data:    nil,
			Message: "Fail to lock record slot",
			Error:   helper.Error{},
		})
	}

	wg.Add(3)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		//err := shard.NewRS().DeleteRS(shard.Cluster, rs)
		err := tx.Delete(rs)
		if err != nil {
			errorChanel <- errors.New("Fail to delete from database")
			return
		}
		errorChanel <- nil
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		_, err = helper.DeleteByQueryES("regist_subject", query)
		if err != nil {
			errorChanel <- errors.New("Fail to delete from elasticsearch")
			return
		}
		errorChanel <- nil
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		tkb.SoChoConLai += 1
		err1 := tx.Update(tkb)
		if err1 != nil {
			errorChanel <- errors.New("Fail to update to tkb")
			return
		}

		errorChanel <- nil
	}(&wg)

	wg.Wait()

	for i := 0; i < len(errorChanel); i++ {
		if err := <-errorChanel; err != nil {
			_ = tx.Rollback()
			_, _, errES := helper.QueryES("regist_subject", query)
			if errES != nil {
				return c.JSON(helper.Response{
					Status:  false,
					Data:    nil,
					Message: "Rollback ES fail",
					Error:   helper.Error{},
				})
			}

			return c.JSON(helper.Response{
				Status:  false,
				Data:    nil,
				Message: err.Error(),
				Error:   helper.Error{},
			})
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
