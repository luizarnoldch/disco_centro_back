package domain

type DiscoResponse struct {
	Id                uint64 `json:"id_disco"`
	Nombre            string `json:"nombre_disco"`
	Album             string `json:"album_disco"`
	Artista           string `json:"artista_disco"`
	Fecha_Lanzamiento string `json:"lanzamiento_disco"`
	Estado            string `json:"estado_disco"`
	Stock             uint32 `json:"stock_disco"`
	Calidad           string `json:"calidad_disco"`
}
