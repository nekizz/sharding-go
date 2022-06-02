package helper

type ESResultRegistSubject struct {
	Hits ESHitsRegistSubject `json:"hits"`
}

type ESHitsRegistSubject struct {
	Total struct {
		Value int `json:"value"`
	} `json:"total"`
	Hits []struct {
		Source ESSourceRegistSubject `json:"_source"`
	} `json:"hits"`
}

type ESSourceRegistSubject struct {
	ID       int    `json:"ID"`
	MaSV     string `json:"MaSV"`
	MaMonHoc string `json:"MaMonHoc"`
	NhomLop  string `json:"NhomLop"`
}
