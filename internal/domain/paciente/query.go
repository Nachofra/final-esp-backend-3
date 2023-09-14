package paciente

var (
	QueryInsertPaciente = `INSERT INTO my_db.paciente(first_name,last_name,address,dni,discharge_date)
	VALUES(?,?,?,?,?)`
	QueryGetAllPaciente = `SELECT id, first_name, last_name, address, dni, discharge_date
	FROM my_db.paciente`
	QueryDeletePaciente  = `DELETE FROM my_db.paciente WHERE id = ?`
	QueryGetPacienteById = `SELECT id, first_name, last_name, address, dni, discharge_date
	FROM my_db.paciente WHERE id = ?`
	QueryGetPacienteByDni = `SELECT id, first_name, last_name, address, dni, discharge_date
	FROM my_db.paciente WHERE dni = ?`
	QueryUpdatePaciente = `UPDATE my_db.paciente SET first_name = ?, last_name = ?, address = ? , dni = ?, discharge_date = ?
	WHERE id = ?`
)
