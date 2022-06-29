package register_subject

type RegisterSubjectBody struct {
	MaSV    string `json:"ma_sv" validate:"required"`
	IDMon   int    `json:"id_mon" validate:"required"`
	MaMon   string `json:"ma_mon" validate:"required"`
	NhomLop string `json:"nhom_lop" validate:"required"`
}
