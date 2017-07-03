package utils

type Page struct {
	PageNo     int
	PageSize   int
	TotalPage  int
	TotalCount int
	FirstPage  bool
	LastPage   bool
	List       interface{}
	Key        string
	Word       string
}

func PageUtil1(count int, pageNo int, pageSize int,key string, word string, list interface{}) Page {
	tp := count / pageSize
	if count%pageSize > 0 {
		tp = count/pageSize + 1
	}
	return Page{PageNo: pageNo, PageSize: pageSize, TotalPage: tp, TotalCount: count, FirstPage: pageNo == 1, LastPage: pageNo == tp, Key:key, Word:word, List: list}
}

func PageUtil(count int, pageNo int, pageSize int, list interface{}) Page {
	tp := count / pageSize
	if count%pageSize > 0 {
		tp = count/pageSize + 1
	}
	return Page{PageNo: pageNo, PageSize: pageSize, TotalPage: tp, TotalCount: count, FirstPage: pageNo == 1, LastPage: pageNo == tp, List: list}
}

/*func PageUser(p int, size int) utils.Page {
	o := orm.NewOrm()
	var user User
	var list []User
	qs := o.QueryTable(user)
	count, _ := qs.Limit(-1).Count()
	qs.RelatedSel().OrderBy("-InTime").Limit(size).Offset((p - 1) * size).All(&list)
	c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
	return utils.PageUtil(c, p, size, list)
}*/
