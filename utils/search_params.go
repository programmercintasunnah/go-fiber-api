package utils

type SearchParams struct {
	Page        int                    `json:"page"`
	PageSize    int                    `json:"pageSize"`
	OrderBy     OrderBy                `json:"orderBy"`
	SearchValue string                 `json:"search"`
	Filters     map[string]interface{} `json:"filters"`
	Populate    []string               `json:"populate"`
}

type OrderBy struct {
	Field string `json:"field"`
	Type  string `json:"type"`
}
