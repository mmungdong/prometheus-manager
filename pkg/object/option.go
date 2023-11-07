package object

// POWithList 设置对象列表
func POWithList[D, P any](list ...*P) POOption[D, P] {
	return func(po *PO[D, P]) {
		po.list = list
		if len(list) > 0 {
			po.one = list[0]
		}
	}
}

// POWithBuildFunc 设置对象构建函数
func POWithBuildFunc[D, P any](buildFunc POBuilder[D, P]) POOption[D, P] {
	return func(po *PO[D, P]) {
		po.buildFunc = buildFunc
	}
}

// DOWithList 设置对象列表
func DOWithList[D, P any](list ...*D) DOOption[D, P] {
	return func(do *DO[D, P]) {
		do.list = list
		if len(list) > 0 {
			do.one = list[0]
		}
	}
}

// DOWithBuildFunc 设置对象构建函数
func DOWithBuildFunc[D, P any](buildFunc DOBuilder[D, P]) DOOption[D, P] {
	return func(do *DO[D, P]) {
		do.buildFunc = buildFunc
	}
}

// NewDO 创建DO
func NewDO[D, P any](opts ...DOOption[D, P]) *DO[D, P] {
	do := &DO[D, P]{
		buildFunc: func(d *D) *P {
			panic("not implement")
		},
	}
	for _, opt := range opts {
		opt(do)
	}

	return do
}

// NewPO 创建PO
func NewPO[D, P any](opts ...POOption[D, P]) *PO[D, P] {
	po := &PO[D, P]{
		buildFunc: func(p *P) *D {
			panic("not implement")
		},
	}
	for _, opt := range opts {
		opt(po)
	}
	return po
}
