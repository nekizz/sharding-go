package register_subject

import (
	"fmt"
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
	if err := validateStruct(c, body); err != nil {
		return err
	}

	var wg sync.WaitGroup
	var regist []*shard.RegisterSubject

	id := uint(helper.HashToInt(body.MaMon + body.NhomLop + "1464"))
	rs := &shard.RegisterSubject{
		ID:       id,
		MaSV:     body.MaSV,
		NhomLop:  body.NhomLop,
		MaMonHoc: body.MaMon,
	}
	tkb, err := shard.NewTKB().GetTKB(shard.Cluster, int64(id))
	if err != nil {
		return getFail(c)
	}

	if err := validateTKBSlot(c, tkb); err != nil {
		return err
	}

	count, err := shard.Cluster.Shard(int64(id)).Model(&regist).Where("ma_sv = ? AND ma_mon_hoc = ?", body.MaSV, body.MaMon).Count()
	fmt.Println(count)
	if err != nil {
		fmt.Println(err)
	}
	if err := checkRegistSubject(count, c); err != nil {
		return err
	} else {

		wg.Add(2)

		go func(wg *sync.WaitGroup) error {
			defer wg.Done()

			err = shard.NewRS().CreateRS(shard.Cluster, rs)
			if err != nil {
				return createFail(c)
			}
			return nil
		}(&wg)

		go func(wg *sync.WaitGroup) error {
			defer wg.Done()
			tkb.SoChoConLai -= 1
			err1 := shard.NewTKB().UpdateTKB(shard.Cluster, tkb)
			if err1 != nil {
				return updateFail(c)
			}
			return nil
		}(&wg)

		wg.Wait()

		return registSuccess(c)
	}
}

func UnregistSubject(c *fiber.Ctx) error {

	body := new(RegisterSubjectBody)
	if err := validateStruct(c, body); err != nil {
		return err
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
		return getFail(c)
	}
	if err := validateTKBSlot(c, tkb); err != nil {
		return err
	}

	wg.Add(2)

	go func(wg *sync.WaitGroup) error {
		defer wg.Done()
		err := shard.NewRS().DeleteRS(shard.Cluster, rs)
		if err != nil {
			return deleteFail(c)
		}
		return nil
	}(&wg)

	go func(wg *sync.WaitGroup) error {
		defer wg.Done()
		tkb.SoChoConLai += 1
		err1 := shard.NewTKB().UpdateTKB(shard.Cluster, tkb)
		if err1 != nil {
			return updateFail(c)
		}

		return nil
	}(&wg)

	wg.Wait()

	return unregistSuccess(c)
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
