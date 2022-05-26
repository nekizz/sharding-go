package register_subject

import (
	"github.com/gofiber/fiber/v2"
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
		ID:       uint(body.ID),
		MaSV:     body.MaSV,
		IDMon:    uint(body.IDMon),
		MaMonHoc: body.MaMon,
	}
	tkb, err := shard.GetTKB(shard.Cluster, int64(rs.IDMon))
	if err != nil {
		return getFail(c)
	}
	if err := validateTKBSlot(c, tkb); err != nil {
		return err
	}

	wg.Add(2)

	go func(wg *sync.WaitGroup) error {
		defer wg.Done()
		err := shard.CreateRS(shard.Cluster, rs)
		if err != nil {
			return createFail(c)
		}
		return nil
	}(&wg)

	go func(wg *sync.WaitGroup) error {
		defer wg.Done()
		tkb.SoChoConLai -= 1
		err1 := shard.UpdateTKB(shard.Cluster, tkb)
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
		ID:       uint(body.ID),
		IDMon:    uint(body.IDMon),
		MaSV:     body.MaSV,
		MaMonHoc: body.MaMon,
	}
	tkb, err := shard.GetTKB(shard.Cluster, int64(rs.IDMon))
	if err != nil {
		return getFail(c)
	}
	if err := validateTKBSlot(c, tkb); err != nil {
		return err
	}

	wg.Add(2)

	go func(wg *sync.WaitGroup) error {
		defer wg.Done()
		err := shard.DeleteRS(shard.Cluster, rs)
		if err != nil {
			return deleteFail(c)
		}
		return nil
	}(&wg)

	go func(wg *sync.WaitGroup) error {
		defer wg.Done()
		tkb.SoChoConLai += 1
		err1 := shard.UpdateTKB(shard.Cluster, tkb)
		if err1 != nil {
			return updateFail(c)
		}

		return nil
	}(&wg)

	wg.Wait()

	return unregistSuccess(c)
}

func deHashToID(s string) int {
	return 1
}
