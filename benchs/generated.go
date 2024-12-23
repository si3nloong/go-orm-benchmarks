// Code generated by sqlgen, version v1.0.0-alpha.5; DO NOT EDIT.

package benchs

import (
	"database/sql/driver"

	"github.com/si3nloong/sqlgen/sequel"
	"github.com/si3nloong/sqlgen/sequel/types"
)

func (Model) TableName() string {
	return "`model`"
}
func (Model) HasPK()      {}
func (Model) IsAutoIncr() {}
func (v *Model) ScanAutoIncr(val int64) error {
	v.Id = int(val)
	return nil
}
func (v Model) PK() (string, int, any) {
	return "`id`", 0, (int64)(v.Id)
}
func (Model) Columns() []string {
	return []string{"`id`", "`name`", "`title`", "`fax`", "`web`", "`age`", "`right`", "`counter`"}
}
func (v Model) Values() []any {
	return []any{(int64)(v.Id), (string)(v.Name), (string)(v.Title), (string)(v.Fax), (string)(v.Web), (int64)(v.Age), (bool)(v.Right), (int64)(v.Counter)}
}
func (v *Model) Addrs() []any {
	return []any{types.Integer(&v.Id), types.String(&v.Name), types.String(&v.Title), types.String(&v.Fax), types.String(&v.Web), types.Integer(&v.Age), types.Bool(&v.Right), types.Integer(&v.Counter)}
}
func (Model) InsertPlaceholders(row int) string {
	return "(?,?,?,?,?,?,?)"
}
func (v Model) InsertOneStmt() (string, []any) {
	return "INSERT INTO `model` (`name`,`title`,`fax`,`web`,`age`,`right`,`counter`) VALUES (?,?,?,?,?,?,?);", []any{(string)(v.Name), (string)(v.Title), (string)(v.Fax), (string)(v.Web), (int64)(v.Age), (bool)(v.Right), (int64)(v.Counter)}
}
func (v Model) FindOneByPKStmt() (string, []any) {
	return "SELECT `id`,`name`,`title`,`fax`,`web`,`age`,`right`,`counter` FROM `model` WHERE `id` = ? LIMIT 1;", []any{(int64)(v.Id)}
}
func (v Model) UpdateOneByPKStmt() (string, []any) {
	return "UPDATE `model` SET `name` = ?,`title` = ?,`fax` = ?,`web` = ?,`age` = ?,`right` = ?,`counter` = ? WHERE `id` = ?;", []any{(string)(v.Name), (string)(v.Title), (string)(v.Fax), (string)(v.Web), (int64)(v.Age), (bool)(v.Right), (int64)(v.Counter), (int64)(v.Id)}
}
func (v Model) GetId() sequel.ColumnValuer[int] {
	return sequel.Column("`id`", v.Id, func(val int) driver.Value { return (int64)(val) })
}
func (v Model) GetName() sequel.ColumnValuer[string] {
	return sequel.Column("`name`", v.Name, func(val string) driver.Value { return (string)(val) })
}
func (v Model) GetTitle() sequel.ColumnValuer[string] {
	return sequel.Column("`title`", v.Title, func(val string) driver.Value { return (string)(val) })
}
func (v Model) GetFax() sequel.ColumnValuer[string] {
	return sequel.Column("`fax`", v.Fax, func(val string) driver.Value { return (string)(val) })
}
func (v Model) GetWeb() sequel.ColumnValuer[string] {
	return sequel.Column("`web`", v.Web, func(val string) driver.Value { return (string)(val) })
}
func (v Model) GetAge() sequel.ColumnValuer[int] {
	return sequel.Column("`age`", v.Age, func(val int) driver.Value { return (int64)(val) })
}
func (v Model) GetRight() sequel.ColumnValuer[bool] {
	return sequel.Column("`right`", v.Right, func(val bool) driver.Value { return (bool)(val) })
}
func (v Model) GetCounter() sequel.ColumnValuer[int64] {
	return sequel.Column("`counter`", v.Counter, func(val int64) driver.Value { return (int64)(val) })
}

func (Model2) Columns() []string {
	return []string{"`id`", "`name`", "`title`", "`fax`", "`web`", "`age`", "`right`", "`counter`"}
}
func (v Model2) Values() []any {
	return []any{(int64)(v.ID), (string)(v.Name), (string)(v.Title), (string)(v.Fax), (string)(v.Web), (int64)(v.Age), (bool)(v.Right), (int64)(v.Counter)}
}
func (v *Model2) Addrs() []any {
	return []any{types.Integer(&v.ID), types.String(&v.Name), types.String(&v.Title), types.String(&v.Fax), types.String(&v.Web), types.Integer(&v.Age), types.Bool(&v.Right), types.Integer(&v.Counter)}
}
func (Model2) InsertPlaceholders(row int) string {
	return "(?,?,?,?,?,?,?,?)"
}
func (v Model2) InsertOneStmt() (string, []any) {
	return "INSERT INTO " + v.TableName() + " (`id`,`name`,`title`,`fax`,`web`,`age`,`right`,`counter`) VALUES (?,?,?,?,?,?,?,?);", v.Values()
}
func (v Model2) GetID() sequel.ColumnValuer[int] {
	return sequel.Column("`id`", v.ID, func(val int) driver.Value { return (int64)(val) })
}
func (v Model2) GetName() sequel.ColumnValuer[string] {
	return sequel.Column("`name`", v.Name, func(val string) driver.Value { return (string)(val) })
}
func (v Model2) GetTitle() sequel.ColumnValuer[string] {
	return sequel.Column("`title`", v.Title, func(val string) driver.Value { return (string)(val) })
}
func (v Model2) GetFax() sequel.ColumnValuer[string] {
	return sequel.Column("`fax`", v.Fax, func(val string) driver.Value { return (string)(val) })
}
func (v Model2) GetWeb() sequel.ColumnValuer[string] {
	return sequel.Column("`web`", v.Web, func(val string) driver.Value { return (string)(val) })
}
func (v Model2) GetAge() sequel.ColumnValuer[int] {
	return sequel.Column("`age`", v.Age, func(val int) driver.Value { return (int64)(val) })
}
func (v Model2) GetRight() sequel.ColumnValuer[bool] {
	return sequel.Column("`right`", v.Right, func(val bool) driver.Value { return (bool)(val) })
}
func (v Model2) GetCounter() sequel.ColumnValuer[int64] {
	return sequel.Column("`counter`", v.Counter, func(val int64) driver.Value { return (int64)(val) })
}

func (Model3) Columns() []string {
	return []string{"`id`", "`name`", "`title`", "`fax`", "`web`", "`age`", "`right`", "`counter`"}
}
func (v Model3) Values() []any {
	return []any{(int64)(v.ID), (string)(v.Name), (string)(v.Title), (string)(v.Fax), (string)(v.Web), (int64)(v.Age), (bool)(v.Right), (int64)(v.Counter)}
}
func (v *Model3) Addrs() []any {
	return []any{types.Integer(&v.ID), types.String(&v.Name), types.String(&v.Title), types.String(&v.Fax), types.String(&v.Web), types.Integer(&v.Age), types.Bool(&v.Right), types.Integer(&v.Counter)}
}
func (Model3) InsertPlaceholders(row int) string {
	return "(?,?,?,?,?,?,?,?)"
}
func (v Model3) InsertOneStmt() (string, []any) {
	return "INSERT INTO " + v.TableName() + " (`id`,`name`,`title`,`fax`,`web`,`age`,`right`,`counter`) VALUES (?,?,?,?,?,?,?,?);", v.Values()
}
func (v Model3) GetID() sequel.ColumnValuer[int] {
	return sequel.Column("`id`", v.ID, func(val int) driver.Value { return (int64)(val) })
}
func (v Model3) GetName() sequel.ColumnValuer[string] {
	return sequel.Column("`name`", v.Name, func(val string) driver.Value { return (string)(val) })
}
func (v Model3) GetTitle() sequel.ColumnValuer[string] {
	return sequel.Column("`title`", v.Title, func(val string) driver.Value { return (string)(val) })
}
func (v Model3) GetFax() sequel.ColumnValuer[string] {
	return sequel.Column("`fax`", v.Fax, func(val string) driver.Value { return (string)(val) })
}
func (v Model3) GetWeb() sequel.ColumnValuer[string] {
	return sequel.Column("`web`", v.Web, func(val string) driver.Value { return (string)(val) })
}
func (v Model3) GetAge() sequel.ColumnValuer[int] {
	return sequel.Column("`age`", v.Age, func(val int) driver.Value { return (int64)(val) })
}
func (v Model3) GetRight() sequel.ColumnValuer[bool] {
	return sequel.Column("`right`", v.Right, func(val bool) driver.Value { return (bool)(val) })
}
func (v Model3) GetCounter() sequel.ColumnValuer[int64] {
	return sequel.Column("`counter`", v.Counter, func(val int64) driver.Value { return (int64)(val) })
}

func (Model4) TableName() string {
	return "`model_4`"
}
func (Model4) Columns() []string {
	return []string{"`id`", "`name`", "`title`", "`fax`", "`web`", "`age`", "`right`", "`counter`"}
}
func (v Model4) Values() []any {
	return []any{(int64)(v.ID), (string)(v.Name), (string)(v.Title), (string)(v.Fax), (string)(v.Web), (int64)(v.Age), (bool)(v.Right), (int64)(v.Counter)}
}
func (v *Model4) Addrs() []any {
	return []any{types.Integer(&v.ID), types.String(&v.Name), types.String(&v.Title), types.String(&v.Fax), types.String(&v.Web), types.Integer(&v.Age), types.Bool(&v.Right), types.Integer(&v.Counter)}
}
func (Model4) InsertPlaceholders(row int) string {
	return "(?,?,?,?,?,?,?,?)"
}
func (v Model4) InsertOneStmt() (string, []any) {
	return "INSERT INTO `model_4` (`id`,`name`,`title`,`fax`,`web`,`age`,`right`,`counter`) VALUES (?,?,?,?,?,?,?,?);", v.Values()
}
func (v Model4) GetID() sequel.ColumnValuer[int] {
	return sequel.Column("`id`", v.ID, func(val int) driver.Value { return (int64)(val) })
}
func (v Model4) GetName() sequel.ColumnValuer[string] {
	return sequel.Column("`name`", v.Name, func(val string) driver.Value { return (string)(val) })
}
func (v Model4) GetTitle() sequel.ColumnValuer[string] {
	return sequel.Column("`title`", v.Title, func(val string) driver.Value { return (string)(val) })
}
func (v Model4) GetFax() sequel.ColumnValuer[string] {
	return sequel.Column("`fax`", v.Fax, func(val string) driver.Value { return (string)(val) })
}
func (v Model4) GetWeb() sequel.ColumnValuer[string] {
	return sequel.Column("`web`", v.Web, func(val string) driver.Value { return (string)(val) })
}
func (v Model4) GetAge() sequel.ColumnValuer[int] {
	return sequel.Column("`age`", v.Age, func(val int) driver.Value { return (int64)(val) })
}
func (v Model4) GetRight() sequel.ColumnValuer[bool] {
	return sequel.Column("`right`", v.Right, func(val bool) driver.Value { return (bool)(val) })
}
func (v Model4) GetCounter() sequel.ColumnValuer[int64] {
	return sequel.Column("`counter`", v.Counter, func(val int64) driver.Value { return (int64)(val) })
}

func (Model5) TableName() string {
	return "`model_5`"
}
func (Model5) Columns() []string {
	return []string{"`id`", "`name`", "`title`", "`fax`", "`web`", "`age`", "`right`", "`counter`"}
}
func (v Model5) Values() []any {
	return []any{(int64)(v.Id), (string)(v.Name), (string)(v.Title), (string)(v.Fax), (string)(v.Web), (int64)(v.Age), (bool)(v.Right), (int64)(v.Counter)}
}
func (v *Model5) Addrs() []any {
	return []any{types.Integer(&v.Id), types.String(&v.Name), types.String(&v.Title), types.String(&v.Fax), types.String(&v.Web), types.Integer(&v.Age), types.Bool(&v.Right), types.Integer(&v.Counter)}
}
func (Model5) InsertPlaceholders(row int) string {
	return "(?,?,?,?,?,?,?,?)"
}
func (v Model5) InsertOneStmt() (string, []any) {
	return "INSERT INTO `model_5` (`id`,`name`,`title`,`fax`,`web`,`age`,`right`,`counter`) VALUES (?,?,?,?,?,?,?,?);", v.Values()
}
func (v Model5) GetId() sequel.ColumnValuer[int64] {
	return sequel.Column("`id`", v.Id, func(val int64) driver.Value { return (int64)(val) })
}
func (v Model5) GetName() sequel.ColumnValuer[string] {
	return sequel.Column("`name`", v.Name, func(val string) driver.Value { return (string)(val) })
}
func (v Model5) GetTitle() sequel.ColumnValuer[string] {
	return sequel.Column("`title`", v.Title, func(val string) driver.Value { return (string)(val) })
}
func (v Model5) GetFax() sequel.ColumnValuer[string] {
	return sequel.Column("`fax`", v.Fax, func(val string) driver.Value { return (string)(val) })
}
func (v Model5) GetWeb() sequel.ColumnValuer[string] {
	return sequel.Column("`web`", v.Web, func(val string) driver.Value { return (string)(val) })
}
func (v Model5) GetAge() sequel.ColumnValuer[int] {
	return sequel.Column("`age`", v.Age, func(val int) driver.Value { return (int64)(val) })
}
func (v Model5) GetRight() sequel.ColumnValuer[bool] {
	return sequel.Column("`right`", v.Right, func(val bool) driver.Value { return (bool)(val) })
}
func (v Model5) GetCounter() sequel.ColumnValuer[int64] {
	return sequel.Column("`counter`", v.Counter, func(val int64) driver.Value { return (int64)(val) })
}

func (Model7) TableName() string {
	return "`model_7`"
}
func (Model7) Columns() []string {
	return []string{"`id`", "`name`", "`title`", "`fax`", "`web`", "`age`", "`right`", "`counter`"}
}
func (v Model7) Values() []any {
	return []any{(int64)(v.ID), (string)(v.Name), (string)(v.Title), (string)(v.Fax), (string)(v.Web), (int64)(v.Age), (bool)(v.Right), (int64)(v.Counter)}
}
func (v *Model7) Addrs() []any {
	return []any{types.Integer(&v.ID), types.String(&v.Name), types.String(&v.Title), types.String(&v.Fax), types.String(&v.Web), types.Integer(&v.Age), types.Bool(&v.Right), types.Integer(&v.Counter)}
}
func (Model7) InsertPlaceholders(row int) string {
	return "(?,?,?,?,?,?,?,?)"
}
func (v Model7) InsertOneStmt() (string, []any) {
	return "INSERT INTO `model_7` (`id`,`name`,`title`,`fax`,`web`,`age`,`right`,`counter`) VALUES (?,?,?,?,?,?,?,?);", v.Values()
}
func (v Model7) GetID() sequel.ColumnValuer[int] {
	return sequel.Column("`id`", v.ID, func(val int) driver.Value { return (int64)(val) })
}
func (v Model7) GetName() sequel.ColumnValuer[string] {
	return sequel.Column("`name`", v.Name, func(val string) driver.Value { return (string)(val) })
}
func (v Model7) GetTitle() sequel.ColumnValuer[string] {
	return sequel.Column("`title`", v.Title, func(val string) driver.Value { return (string)(val) })
}
func (v Model7) GetFax() sequel.ColumnValuer[string] {
	return sequel.Column("`fax`", v.Fax, func(val string) driver.Value { return (string)(val) })
}
func (v Model7) GetWeb() sequel.ColumnValuer[string] {
	return sequel.Column("`web`", v.Web, func(val string) driver.Value { return (string)(val) })
}
func (v Model7) GetAge() sequel.ColumnValuer[int] {
	return sequel.Column("`age`", v.Age, func(val int) driver.Value { return (int64)(val) })
}
func (v Model7) GetRight() sequel.ColumnValuer[bool] {
	return sequel.Column("`right`", v.Right, func(val bool) driver.Value { return (bool)(val) })
}
func (v Model7) GetCounter() sequel.ColumnValuer[int64] {
	return sequel.Column("`counter`", v.Counter, func(val int64) driver.Value { return (int64)(val) })
}
