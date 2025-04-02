package model

type PaginationRequest struct {
	PageID   int32 `form:"page_id"`
	PageSize int32 `form:"page_size"`
}

type GetIDRequest struct {
	ID string `from:"id"`
}
