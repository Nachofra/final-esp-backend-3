package mysql

import "github.com/Nachofra/final-esp-backend-3/pkg/query_builder"

const (
	QueryGetAllAppointment = `SELECT a.id, a.patient_id, a.dentist_id, a.date, a.description, p.dni
	FROM clinic.appointment a INNER JOIN clinic.patient p on a.patient_id = p.id`

	QueryGetAppointmentByID = `SELECT id, patient_id, dentist_id, date, description
	FROM clinic.appointment WHERE id = ?`

	QueryInsertAppointment = `INSERT INTO clinic.appointment(patient_id,dentist_id,date,description)
	VALUES(?,?,?,?)`

	QueryUpdateAppointment = `UPDATE clinic.appointment SET patient_id = ?, dentist_id = ?, date = ?, description = ?
	WHERE id = ?`

	QueryDeleteAppointment = `DELETE FROM clinic.appointment WHERE id = ?`
)

// GenerateQuery handles query creation to filter dynamically based on params.
func GenerateQuery(filter map[string]string) string {
	// Hardcoded for now :(
	limit := 1000
	offset := 0

	var dqb query_builder.DynamicQueryBuilder
	query := dqb.And(
		dqb.NewExpression("a.patient_id", "=", filter["patient_id"]),
		dqb.NewExpression("a.dentist_id", "=", filter["dentist_id"]),
		dqb.NewExpression("p.dni", "=", filter["dni"]),
		dqb.NewExpression("a.date", ">=", filter["from_date"]),
		dqb.NewExpression("a.date", "<=", filter["to_date"]),
	).Limit(offset, limit).BindSql(QueryGetAllAppointment)
	return query
}
