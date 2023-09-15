package odontologo

var (
	QueryInsertOdontologo = `INSERT INTO my_db.odontologo(first_name,last_name,registration_number)
	VALUES(?,?,?)`
	QueryGetAllOdontologo = `SELECT id, first_name, last_name, registration_number
	FROM my_db.odontologo`
	QueryDeleteOdontologo  = `DELETE FROM my_db.odontologo WHERE id = ?`
	QueryGetOdontologoById = `SELECT id, first_name, last_name, registration_number
	FROM my_db.odontologo WHERE id = ?`
	QueryGetOdontologoByRegistrationNumber = `SELECT id, first_name, last_name, registration_number
	FROM my_db.odontologo WHERE registration_number = ?`
	QueryUpdateOdontologo = `UPDATE my_db.odontologo SET first_name = ?, last_name = ?, registration_number = ?
	WHERE id = ?`
)
