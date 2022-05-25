package register_subject

type RegisterSubjectBody struct {
	MaSV  string `json:"ma_sv" xml:"ma_sv" form:"ma_sv"`
	MaMon string `json:"ma_mon" xml:"ma_mon" form:"ma_mon"`
}
