package dto

type ResponseParam struct{
	StatusCode int 
	Message string 
	Paginate *Paginate 
	Data any
}

type FilterParam struct{
	Page int
	Limit int
	Offset int
	Search string
}