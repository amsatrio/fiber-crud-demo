package request

type Filter struct {
	Id        string          `json:"id" example:"fullName"`
	Value     interface{}     `json:"value" example:"Adi"`
	MatchMode FilterMatchMode `json:"matchMode" example:"CONTAINS"`
	DataType  FilterDataType  `json:"dataType" example:"TEXT"`
	Mode      FilterMode      `json:"mode" example:"AND"`
}

type FilterDataType string

const (
	TEXT    FilterDataType = "TEXT"
	NUMBER  FilterDataType = "NUMBER"
	DATE    FilterDataType = "DATE"
	BOOLEAN FilterDataType = "BOOLEAN"
)

func (f FilterDataType) String() string {
	return string(f)
}

type FilterMatchMode string

const (
	CONTAINS     FilterMatchMode = "CONTAINS"
	BETWEEN      FilterMatchMode = "BETWEEN"
	EQUALS       FilterMatchMode = "EQUALS"
	NOT          FilterMatchMode = "NOT"
	LESS_THAN    FilterMatchMode = "LESS_THAN"
	GREATER_THAN FilterMatchMode = "GREATER_THAN"
)

func (f FilterMatchMode) String() string {
	return string(f)
}

type FilterMode string

const (
	AND FilterMode = "AND"
	OR  FilterMode = "OR"
)

func (f FilterMode) String() string {
	return string(f)
}

type Sort struct {
	Id   string `json:"id" example:"fullName"`
	Desc bool   `json:"desc" example:"true"`
}
