package object

type (
	DOBuilder[D, P any] func(*D) *P
	POBuilder[D, P any] func(*P) *D

	DO[D, P any] struct {
		one       *D
		list      []*D
		buildFunc DOBuilder[D, P]
	}

	PO[D, P any] struct {
		one       *P
		list      []*P
		buildFunc POBuilder[D, P]
	}

	DOOption[D, P any] func(*DO[D, P])
	POOption[D, P any] func(*PO[D, P])
)

// PO 转换为PO
func (l *DO[D, P]) PO() *PO[D, P] {
	return NewPO[D, P](
		POWithList[D, P](func(list []*D) []*P {
			res := make([]*P, 0, len(list))
			for _, item := range list {
				res = append(res, l.buildFunc(item))
			}
			return res
		}(l.list)...),
	)
}

// DO 转换为DO
func (l *PO[D, P]) DO() *DO[D, P] {
	return NewDO[D, P](
		DOWithList[D, P](func(list []*P) []*D {
			res := make([]*D, 0, len(list))
			for _, item := range list {
				res = append(res, l.buildFunc(item))
			}
			return res
		}(l.list)...),
	)
}

// One 获取单个对象
func (l *DO[D, P]) One() *D {
	return l.one
}

// List 获取对象列表
func (l *DO[D, P]) List() []*D {
	return l.list
}

// One 获取单个对象
func (l *PO[D, P]) One() *P {
	return l.one
}

// List 获取对象列表
func (l *PO[D, P]) List() []*P {
	return l.list
}
