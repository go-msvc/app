package app

type DisplayItem interface {
	Parent() DisplayItem
}

type DisplayText interface {
	DisplayItem
	Text() TextFunc
}

type DisplayButton interface {
	DisplayText
	Action() interface{} //?function
}

type DisplayInput interface {
	DisplayText
	FieldName() string
}

type DisplayGroup interface {
	DisplayItem
	Items() []DisplayItem
}

type DisplaySection interface {
	DisplayGroup
	Heading() TextFunc
}

//=====[ constructors ]========================================================

func Item() DisplayItem {
	return displayItem{
		parent: nil,
	}
}

type displayItem struct {
	parent *displayItem
}

func (di displayItem) Parent() DisplayItem {
	return di.parent
}

func Text(tf TextFunc) DisplayText {
	return displayText{
		DisplayItem: Item(),
		tf:          tf,
	}
}

type displayText struct {
	DisplayItem
	tf TextFunc
}

func (dt displayText) Text() TextFunc {
	return dt.tf
}

func Button(tf TextFunc, action interface{}) DisplayButton {
	return displayButton{
		DisplayText: Text(tf),
		action:      action,
	}
}

type displayButton struct {
	DisplayText
	action interface{} //todo...
}

func (db displayButton) Action() interface{} {
	return db.action
}

func Input(tf TextFunc, fieldName string) DisplayInput {
	return displayInput{
		DisplayText: Text(tf),
		fieldName:   fieldName,
	}
}

type displayInput struct {
	DisplayText
	fieldName string
}

func (db displayInput) FieldName() string {
	return db.fieldName
}

func Group(items ...DisplayItem) DisplayGroup {
	return &displayGroup{
		DisplayItem: Item(),
		vert:        true,
		items:       items,
	}
}

type displayGroup struct {
	DisplayItem
	vert  bool
	items []DisplayItem
}

func (dg displayGroup) Items() []DisplayItem {
	return dg.items
}

func Section(heading TextFunc, items ...DisplayItem) DisplaySection {
	return displaySection{
		DisplayGroup: Group(items...),
		heading:      heading,
	}
}

type displaySection struct {
	DisplayGroup
	level   int
	heading TextFunc
}

func (ds displaySection) Heading() TextFunc {
	return ds.heading
}

func (ds displaySection) Level() int {
	return ds.level
}
