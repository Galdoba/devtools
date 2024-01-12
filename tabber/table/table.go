package table

type table struct {
	Data [][]string
}

type Cell struct {
	PosX        int
	PosY        int
	Content     string
	FGcol       byte
	BGcol       byte
	Independent bool              //есть ли связи с другими ячейками
	ParentX     bool              //родитель слева
	ParentY     bool              //родитель сверху
	Action      map[string]string //при получении допустимого тригера (key) запускаем действие (val)
}
