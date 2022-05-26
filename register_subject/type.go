package register_subject

type RegisterSubjectBody struct {
	ID    int    `json:"id" xml:"id" form:"id"`
	MaSV  string `json:"ma_sv" xml:"ma_sv" form:"ma_sv"`
	IDMon int    `json:"id_mon" xml:"id_mon" form:"id_mon"`
	MaMon string `json:"ma_mon" xml:"ma_mon" form:"ma_mon"`
}
