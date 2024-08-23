package dto

type FileSearchDTO struct {
	PageCurrent int    `json:"pageCurrent"` // 当前所在页面
	PageTotal   int    `json:"pageTotal"`   //数据总页数
	PageSize    int    `json:"pageSize"`    //当前页面展示条数 默认是10
	RowsTotal   int    `json:"rowsTotal"`   //总数据量 - 条树
	SearchItem  string `json:"searchItem"`  //查询文件名称
}
