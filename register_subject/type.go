package register_subject

type RegisterSubjectBody struct {
	MaSV    string `json:"ma_sv" xml:"ma_sv" form:"ma_sv"`
	IDMon   int    `json:"id_mon" xml:"id_mon" form:"id_mon"`
	MaMon   string `json:"ma_mon" xml:"ma_mon" form:"ma_mon"`
	NhomLop string `json:"nhom_lop" xml:"nhom_lop" form:"nhom_lop"`
}

type Hello struct {
	MaSV    string `json:"ma_sv" xml:"ma_sv" form:"ma_sv"`
	IDMon   int    `json:"id_mon" xml:"id_mon" form:"id_mon"`
	MaMon   string `json:"ma_mon" xml:"ma_mon" form:"ma_mon"`
	NhomLop string `json:"nhom_lop" xml:"nhom_lop" form:"nhom_lop"`
}
