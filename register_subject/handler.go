package register_subject

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
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
	fmt.Println(helper.HashToInt(body.MaMon + body.NhomLop + "1464"))
	fmt.Println(tkb)

	wg.Add(2)

	go func(wg *sync.WaitGroup) error {
		defer wg.Done()
		err := shard.NewRS().CreateRS(shard.Cluster, rs)
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
