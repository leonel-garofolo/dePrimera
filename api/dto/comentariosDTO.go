package models

type Comentarios struct {
	IDComentario int64   `json:"id_comentario"`
	Mail         string  `json:"mail"`
	Puntaje      float64 `json:"puntaje"`
	Comentario   string  `json:"comentario"`
}

// TableName sets the insert table name for this struct type
func (z *Comentarios) TableName() string {
	return "comentarios"
}
