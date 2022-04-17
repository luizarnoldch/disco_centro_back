package domain

type Disco struct {
	Id                uint64 `db:"id_disco"`
	Nombre            string `db:"nombre_disco"`
	Album             string `db:"album_disco"`
	Artista           string `db:"artista_disco"`
	Fecha_Lanzamiento string `db:"lanzamiento_disco"`
	Estado            string `db:"estado_disco"`
	Stock             uint32 `db:"stock_disco"`
	Calidad           string `db:"calidad_disco"`
}

func (d Disco) ToResponse() DiscoResponse {
	return DiscoResponse{
		Id:                d.Id,
		Nombre:            d.Nombre,
		Album:             d.Album,
		Artista:           d.Artista,
		Fecha_Lanzamiento: d.Fecha_Lanzamiento,
		Estado:            d.Estado,
		Stock:             d.Stock,
		Calidad:           d.Calidad,
	}
}

func (d Disco) ToNewDiscoResponse() *DiscoResponse {
	return &DiscoResponse{
		Id:      d.Id,
		Nombre:  d.Nombre,
		Album:   d.Album,
		Artista: d.Artista,
		Estado:  d.Estado,
		Stock:   d.Stock,
		Calidad: d.Calidad,
	}
}

func NewDisco(id uint64, nombre string, album string, artista string, fecha_lanzamiento string, estado string, stock uint32, calidad string) Disco {
	return Disco{
		Id:                id,
		Nombre:            nombre,
		Album:             album,
		Artista:           artista,
		Fecha_Lanzamiento: fecha_lanzamiento,
		Estado:            estado,
		Stock:             stock,
		Calidad:           calidad,
	}
}
