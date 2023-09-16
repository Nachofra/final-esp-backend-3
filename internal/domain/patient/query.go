package patient

var (
	QueryInsertPatient = `INSERT INTO clinic.patient(first_name,last_name,address,dni,discharge_date)
	VALUES(?,?,?,?,?)`
	QueryGetAllPatient = `SELECT id, first_name, last_name, address, dni, discharge_date
	FROM clinic.patient`
	QueryDeletePatient  = `DELETE FROM clinic.patient WHERE id = ?`
	QueryGetPatientByID = `SELECT id, first_name, last_name, address, dni, discharge_date
	FROM clinic.patient WHERE id = ?`
	QueryGetPatientByDNI = `SELECT id, first_name, last_name, address, dni, discharge_date
	FROM clinic.patient WHERE dni = ?`
	QueryUpdatePatient = `UPDATE clinic.patient SET first_name = ?, last_name = ?, address = ? , dni = ?, discharge_date = ?
	WHERE id = ?`
)
