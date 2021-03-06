package models

import "mime/multipart"

type Message struct {
	Code    uint32 `json:"code"`
	Message string `json:"message"`
}

type SingleFilmRollResponse struct {
	Success  bool      `json:"success"`
	Errors   []Message `json:"errors"`
	Messages []Message `json:"messages"`
	Result   struct {
		Roll_id     int    `json:"rollId"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Type_id     int    `json:"typeId"`
	} `json:"result"`
}

type FilmRollResponse struct {
	Success  bool      `json:"success"`
	Errors   []Message `json:"errors"`
	Messages []Message `json:"messages"`
	Result   []struct {
		Roll_id     int    `json:"rollId"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Type_id     int    `json:"typeId"`
	} `json:"result"`
}

type RollTypeResponse struct {
	Success  bool      `json:"success"`
	Errors   []Message `json:"errors"`
	Messages []Message `json:"messages"`
	Result   struct {
		Type_id   int    `json:"typeId"`
		StockName string `json:"stockName"`
		Format    string `json:"format"`
		M_id      int    `json:"mId"`
	} `json:"result"`
}

type MultipleRollTypeResponse struct {
	Success  bool      `json:"success"`
	Errors   []Message `json:"errors"`
	Messages []Message `json:"messages"`
	Result   []struct {
		Type_id   int    `json:"typeId"`
		StockName string `json:"stockName"`
		Format    string `json:"format"`
		M_id      int    `json:"mId"`
	} `json:"result"`
}

type FilmRollRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Type_Id     int    `json:"typeId"`
}

type FilmRollPhotosResponse struct {
	Success  bool      `json:"success"`
	Errors   []Message `json:"errors"`
	Messages []Message `json:"messages"`
	Result   []struct {
		PhotoId int    `json:"photoId"`
		Title   string `json:"title"`
		Uuid    string `json:"uuid"`
		RollId  int    `json:"rollId"`
	} `json:"result"`
}

type PhotoUpload struct {
	// Photo_id int                     `form:"photoid"`
	// UUID     string                  `form:"uuid"`
	Roll_id int                     `form:"rollId"`
	Files   []*multipart.FileHeader `form:"files"`
}

type PhotoDeletedResponse struct {
	Success  bool      `json:"success"`
	Errors   []Message `json:"errors"`
	Messages []Message `json:"messages"`
	Result   struct {
		PhotoId int    `json:"photoId"`
		Title   string `json:"title"`
		Uuid    string `json:"uuid"`
		RollId  int    `json:"rollId"`
	} `json:"result"`
}

type ManufacturerResponse struct {
	Success  bool      `json:"success"`
	Errors   []Message `json:"errors"`
	Messages []Message `json:"messages"`
	Result   struct {
		MId  int    `json:"mId"`
		Name string `json:"name"`
	} `json:"result"`
}
