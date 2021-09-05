// +build postgres

package constant

const (
	GetDbsSql     = "SELECT datname as database FROM pg_database WHERE datistemplate = false;"
	GetTablesSql  = "select table_name as table_name from information_schema.tables where table_catalog = ? and table_schema = ?"
	GetColumnsSql = `
SELECT columns.COLUMN_NAME as column_name,
       columns.DATA_TYPE   as data_type,
       CASE
           columns.DATA_TYPE
           WHEN 'text' THEN
               concat_ws('', '', columns.CHARACTER_MAXIMUM_LENGTH)
           WHEN 'varchar' THEN
               concat_ws('', '', columns.CHARACTER_MAXIMUM_LENGTH)
           WHEN 'smallint' THEN
               concat_ws(',', columns.NUMERIC_PRECISION, columns.NUMERIC_SCALE)
           WHEN 'decimal' THEN
               concat_ws(',', columns.NUMERIC_PRECISION, columns.NUMERIC_SCALE)
           WHEN 'integer' THEN
               concat_ws('', '', columns.NUMERIC_PRECISION)
           WHEN 'bigint' THEN
               concat_ws('', '', columns.NUMERIC_PRECISION)
           ELSE ''
           END             AS data_type_long
FROM INFORMATION_SCHEMA.COLUMNS columns
WHERE table_catalog = ?
  and table_schema = ?
  and table_name = ?
`
)
