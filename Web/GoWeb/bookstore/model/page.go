package model

// Page 需要当前页的数据,当前页数,每页记录数,总页数,总记录数
type Page struct {
	Books       []*Book //每页查询出来的图书存放的切片
	PageNo      int     //当前页
	PageSize    int     //每页显示的条数
	TotalPageNo int     //总页数
	TotalRecord int     //总记录数
	Min         string //最小价格
	Max         string //最大价格
}

// IsHasPrev 是否有上一页
func (p *Page) IsHasPrev() bool {
	return p.PageNo > 1
}

// IsHasNext 是否有下一页
func (p *Page) IsHasNext() bool {
	return p.PageNo < p.TotalPageNo
}

// GetPrevPageNo 获取前一页
func (p *Page) GetPrevPageNo() int {
	if p.IsHasPrev() {
		return p.PageNo - 1
	}
	return 1
}

// GetNextPageNo 获取下一页
func (p *Page) GetNextPageNo() int {
	if p.IsHasNext() {
		return p.PageNo + 1
	}
	return p.PageSize
}
