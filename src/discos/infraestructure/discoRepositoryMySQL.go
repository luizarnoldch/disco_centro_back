package infraestructure

import (
	//"database/sql"
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/luizarnoldch/disco_centro_back/src/discos/domain"
	"github.com/luizarnoldch/disco_centro_lib/errs"
	"github.com/luizarnoldch/disco_centro_lib/logger"
)

type DiscoRepositoryMySQL struct {
	client *sqlx.DB
}

func (db DiscoRepositoryMySQL) RepoGetAllDiscos() ([]domain.Disco, *errs.AppError) {
	var err error
	discos := make([]domain.Disco, 0)

	getAllDiscos := "SELECT * FROM discos"
	//findAllDiscos := "CALL GetAllDiscos()"
	err = db.client.Select(&discos, getAllDiscos)

	if err != nil {
		logger.Error("Error al consultar la tabla discos: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database while querying")
	}
	return discos, nil
}

func (db DiscoRepositoryMySQL) RepoGetDisco(id uint64) (*domain.Disco, *errs.AppError) {
	getDisco := "SELECT * FROM discos WHERE id_disco = ?"
	//customerSql := "Call GetDiscoById(?)"
	var d domain.Disco
	err := db.client.Get(&d, getDisco, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Disco not found")
		} else {
			logger.Error("Error while querying discos " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected error from database while querying")
		}
	}
	return &d, nil
}

func (db DiscoRepositoryMySQL) RepoPostDisco(disco domain.Disco) (*domain.Disco, *errs.AppError) {
	postDisco := "INSERT INTO discos (nombre_disco, album_disco, artista_disco, lanzamiento_disco, estado_disco, stock_disco, calidad_disco) values (?, ?, ?, ?, ?, ?, ?)"
	//postDisco := "CALL PostDisco(?)"
	result, err := db.client.Exec(postDisco, disco.Nombre, disco.Album, disco.Artista, disco.Fecha_Lanzamiento, disco.Estado, disco.Stock, disco.Calidad)
	//result, err := db.client.Exec(postDisco, disco)
	if err != nil {
		logger.Error("Error while saving disco: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database while saving")
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new disco: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database while inserting")
	}
	disco.Id = uint64(id)
	return &disco, nil
}

func (db DiscoRepositoryMySQL) RepoUpdateDisco(disco domain.Disco) (*domain.Disco, *errs.AppError) {
	postDisco := `UPDATE discos SET nombre_disco = ?,	album_disco = ?, artista_disco = ?, lanzamiento_disco = ?, estado_disco = ?, stock_disco = ?, calidad_disco = ? WHERE id_disco = ?`
	//postDisco := "CALL UpdateDisco(?)"
	result, err := db.client.Exec(postDisco, disco.Nombre, disco.Album, disco.Artista, disco.Fecha_Lanzamiento, disco.Estado, disco.Stock, disco.Calidad, disco.Id)
	//result, err := db.client.Exec(postDisco, disco)
	if err != nil {
		logger.Error("Error while updating disco: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new disco: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database while updating")
	}
	if id != 0 {
		disco.Id = uint64(id)
	}
	return &disco, nil
}

func (db DiscoRepositoryMySQL) RepoDeleteDisco(id uint64) (*json.RawMessage, *errs.AppError) {
	postDisco, err := db.client.Prepare(`
	DELETE FROM
		discos		
	WHERE
		id_disco = ?
	`)
	if err != nil {
		logger.Error("Error while deleting disco: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database while deleting")
	}
	//postDisco := "CALL UpdateDisco(?)"
	postDisco.Exec(id)
	/*
		var jsonString = []byte(fmt.Sprintf(`{
			"id" : %d,
			"msg": "disco %d eliminado"
		}`, id, id))
	*/
	jsonString := json.RawMessage(fmt.Sprintf(`{
		"id" : %d,
		"msg": "disco %d eliminado"
	}`, id, id))
	return &jsonString, nil
}

func NewDiscoRepositoryMySQL(DB *sqlx.DB) DiscoRepositoryMySQL {
	return DiscoRepositoryMySQL{DB}
}
