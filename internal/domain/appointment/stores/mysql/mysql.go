package mysql

const (
	QueryInsertAppointment = `INSERT INTO clinic.appointment(patient_id,dentist_id,date,description)
	VALUES(?,?,?,?)`
	QueryGetAllAppointment = `SELECT id, patient_id, dentist_id, date, description
	FROM clinic.appointment`
	QueryDeleteAppointment  = `DELETE FROM clinic.appointment WHERE id = ?`
	QueryGetAppointmentByID = `SELECT id, patient_id, dentist_id, date, description
	FROM clinic.appointment WHERE id = ?`
	QueryGetAppointmentByPatientID = `SELECT id, patient_id, dentist_id, date, description
	FROM clinic.appointment WHERE patient_id = ?`
	QueryUpdateAppointment = `UPDATE clinic.appointment SET patient_id = ?, dentist_id = ?, date = ?, description = ?
	WHERE id = ?`
)
