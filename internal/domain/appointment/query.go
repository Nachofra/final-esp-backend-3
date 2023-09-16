package appointment

var (
	QueryInsertTurno = `INSERT INTO my_db.turno(paciente_id,odontologo_id,date,description)
	VALUES(?,?,?,?)`
	QueryGetAllTurno = `SELECT id, paciente_id, odontologo_id, date, description
	FROM my_db.turno`
	QueryDeleteTurno  = `DELETE FROM my_db.turno WHERE id = ?`
	QueryGetTurnoById = `SELECT id, paciente_id, odontologo_id, date, description
	FROM my_db.turno WHERE id = ?`
	QueryGetTurnoByPacienteId = `SELECT id, paciente_id, odontologo_id, date, description
	FROM my_db.turno WHERE paciente_id = ?`
	QueryUpdateTurno = `UPDATE my_db.turno SET paciente_id = ?, odontologo_id = ?, date = ?, description = ?
	WHERE id = ?`
)
