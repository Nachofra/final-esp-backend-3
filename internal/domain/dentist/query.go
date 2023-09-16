package dentist

var (
	QueryInsertDentist = `INSERT INTO clinic.dentist(first_name,last_name,registration_number)
	VALUES(?,?,?)`
	QueryGetAllDentist = `SELECT id, first_name, last_name, registration_number
	FROM clinic.dentist`
	QueryDeleteDentist  = `DELETE FROM clinic.dentist WHERE id = ?`
	QueryGetDentistByID = `SELECT id, first_name, last_name, registration_number
	FROM clinic.dentist WHERE id = ?`
	QueryGetDentistByRegistrationNumber = `SELECT id, first_name, last_name, registration_number
	FROM clinic.dentist WHERE registration_number = ?`
	QueryUpdateDentist = `UPDATE clinic.dentist SET first_name = ?, last_name = ?, registration_number = ?
	WHERE id = ?`
)
