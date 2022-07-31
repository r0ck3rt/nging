// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"github.com/webx-top/com"
	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/db/lib/factory/pagination"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/param"
)

type Slice_NgingUserRolePermission []*NgingUserRolePermission

func (s Slice_NgingUserRolePermission) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingUserRolePermission) RangeRaw(fn func(m *NgingUserRolePermission) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingUserRolePermission) GroupBy(keyField string) map[string][]*NgingUserRolePermission {
	r := map[string][]*NgingUserRolePermission{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*NgingUserRolePermission{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (s Slice_NgingUserRolePermission) KeyBy(keyField string) map[string]*NgingUserRolePermission {
	r := map[string]*NgingUserRolePermission{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (s Slice_NgingUserRolePermission) AsKV(keyField string, valueField string) param.Store {
	r := param.Store{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (s Slice_NgingUserRolePermission) Transform(transfers map[string]param.Transfer) []param.Store {
	r := make([]param.Store, len(s))
	for idx, row := range s {
		r[idx] = row.AsMap().Transform(transfers)
	}
	return r
}

func (s Slice_NgingUserRolePermission) FromList(data interface{}) Slice_NgingUserRolePermission {
	values, ok := data.([]*NgingUserRolePermission)
	if !ok {
		for _, value := range data.([]interface{}) {
			row := &NgingUserRolePermission{}
			row.FromRow(value.(map[string]interface{}))
			s = append(s, row)
		}
		return s
	}
	s = append(s, values...)

	return s
}

func NewNgingUserRolePermission(ctx echo.Context) *NgingUserRolePermission {
	m := &NgingUserRolePermission{}
	m.SetContext(ctx)
	return m
}

// NgingUserRolePermission
type NgingUserRolePermission struct {
	base    factory.Base
	objects []*NgingUserRolePermission

	RoleId     uint   `db:"role_id,pk" bson:"role_id" comment:"角色ID" json:"role_id" xml:"role_id"`
	Type       string `db:"type,pk" bson:"type" comment:"权限类型" json:"type" xml:"type"`
	Permission string `db:"permission" bson:"permission" comment:"权限值" json:"permission" xml:"permission"`
}

// - base function

func (a *NgingUserRolePermission) Trans() *factory.Transaction {
	return a.base.Trans()
}

func (a *NgingUserRolePermission) Use(trans *factory.Transaction) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *NgingUserRolePermission) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *NgingUserRolePermission) EventON(on ...bool) factory.Model {
	a.base.EventON(on...)
	return a
}

func (a *NgingUserRolePermission) EventOFF(off ...bool) factory.Model {
	a.base.EventOFF(off...)
	return a
}

func (a *NgingUserRolePermission) Context() echo.Context {
	return a.base.Context()
}

func (a *NgingUserRolePermission) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *NgingUserRolePermission) SetNamer(namer func(factory.Model) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *NgingUserRolePermission) Namer() func(factory.Model) string {
	return a.base.Namer()
}

func (a *NgingUserRolePermission) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *NgingUserRolePermission) Param(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam().SetMiddleware(mw).SetArgs(args...)
	}
	return a.base.Param().SetMiddleware(mw).SetArgs(args...)
}

// - current function

func (a *NgingUserRolePermission) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName, connID[0]).Use(a.base.Trans())
	}
	return factory.NewModel(structName, a.base.ConnID()).Use(a.base.Trans())
}

func (a *NgingUserRolePermission) Objects() []*NgingUserRolePermission {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *NgingUserRolePermission) XObjects() Slice_NgingUserRolePermission {
	return Slice_NgingUserRolePermission(a.Objects())
}

func (a *NgingUserRolePermission) NewObjects() factory.Ranger {
	return &Slice_NgingUserRolePermission{}
}

func (a *NgingUserRolePermission) InitObjects() *[]*NgingUserRolePermission {
	a.objects = []*NgingUserRolePermission{}
	return &a.objects
}

func (a *NgingUserRolePermission) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *NgingUserRolePermission) Short_() string {
	return "nging_user_role_permission"
}

func (a *NgingUserRolePermission) Struct_() string {
	return "NgingUserRolePermission"
}

func (a *NgingUserRolePermission) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *NgingUserRolePermission) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.Use(source.Trans())
	a.SetNamer(source.Namer())
	return a
}

func (a *NgingUserRolePermission) Get(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	base := a.base
	if !a.base.Eventable() {
		err = a.Param(mw, args...).SetRecv(a).One()
		a.base = base
		return
	}
	queryParam := a.Param(mw, args...).SetRecv(a)
	if err = DBI.FireReading(a, queryParam); err != nil {
		return
	}
	err = queryParam.One()
	a.base = base
	if err == nil {
		err = DBI.FireReaded(a, queryParam)
	}
	return
}

func (a *NgingUserRolePermission) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetPage(page).SetSize(size).SetRecv(recv).List()
	}
	queryParam := a.Param(mw, args...).SetPage(page).SetSize(size).SetRecv(recv)
	if err := DBI.FireReading(a, queryParam); err != nil {
		return nil, err
	}
	cnt, err := queryParam.List()
	if err == nil {
		switch v := recv.(type) {
		case *[]*NgingUserRolePermission:
			err = DBI.FireReaded(a, queryParam, Slice_NgingUserRolePermission(*v))
		case []*NgingUserRolePermission:
			err = DBI.FireReaded(a, queryParam, Slice_NgingUserRolePermission(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingUserRolePermission) GroupBy(keyField string, inputRows ...[]*NgingUserRolePermission) map[string][]*NgingUserRolePermission {
	var rows Slice_NgingUserRolePermission
	if len(inputRows) > 0 {
		rows = Slice_NgingUserRolePermission(inputRows[0])
	} else {
		rows = Slice_NgingUserRolePermission(a.Objects())
	}
	return rows.GroupBy(keyField)
}

func (a *NgingUserRolePermission) KeyBy(keyField string, inputRows ...[]*NgingUserRolePermission) map[string]*NgingUserRolePermission {
	var rows Slice_NgingUserRolePermission
	if len(inputRows) > 0 {
		rows = Slice_NgingUserRolePermission(inputRows[0])
	} else {
		rows = Slice_NgingUserRolePermission(a.Objects())
	}
	return rows.KeyBy(keyField)
}

func (a *NgingUserRolePermission) AsKV(keyField string, valueField string, inputRows ...[]*NgingUserRolePermission) param.Store {
	var rows Slice_NgingUserRolePermission
	if len(inputRows) > 0 {
		rows = Slice_NgingUserRolePermission(inputRows[0])
	} else {
		rows = Slice_NgingUserRolePermission(a.Objects())
	}
	return rows.AsKV(keyField, valueField)
}

func (a *NgingUserRolePermission) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetOffset(offset).SetSize(size).SetRecv(recv).List()
	}
	queryParam := a.Param(mw, args...).SetOffset(offset).SetSize(size).SetRecv(recv)
	if err := DBI.FireReading(a, queryParam); err != nil {
		return nil, err
	}
	cnt, err := queryParam.List()
	if err == nil {
		switch v := recv.(type) {
		case *[]*NgingUserRolePermission:
			err = DBI.FireReaded(a, queryParam, Slice_NgingUserRolePermission(*v))
		case []*NgingUserRolePermission:
			err = DBI.FireReaded(a, queryParam, Slice_NgingUserRolePermission(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingUserRolePermission) Insert() (pk interface{}, err error) {

	if a.base.Eventable() {
		err = DBI.Fire("creating", a, nil)
		if err != nil {
			return
		}
	}
	pk, err = a.Param(nil).SetSend(a).Insert()

	if err == nil && a.base.Eventable() {
		err = DBI.Fire("created", a, nil)
	}
	return
}

func (a *NgingUserRolePermission) Update(mw func(db.Result) db.Result, args ...interface{}) (err error) {

	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(a).Update()
	}
	if err = DBI.Fire("updating", a, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(a).Update(); err != nil {
		return
	}
	return DBI.Fire("updated", a, mw, args...)
}

func (a *NgingUserRolePermission) Updatex(mw func(db.Result) db.Result, args ...interface{}) (affected int64, err error) {

	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(a).Updatex()
	}
	if err = DBI.Fire("updating", a, mw, args...); err != nil {
		return
	}
	if affected, err = a.Param(mw, args...).SetSend(a).Updatex(); err != nil {
		return
	}
	err = DBI.Fire("updated", a, mw, args...)
	return
}

func (a *NgingUserRolePermission) UpdateByFields(mw func(db.Result) db.Result, fields []string, args ...interface{}) (err error) {

	if !a.base.Eventable() {
		return a.Param(mw, args...).UpdateByStruct(a, fields...)
	}
	editColumns := make([]string, len(fields))
	for index, field := range fields {
		editColumns[index] = com.SnakeCase(field)
	}
	if err = DBI.FireUpdate("updating", a, editColumns, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).UpdateByStruct(a, fields...); err != nil {
		return
	}
	err = DBI.FireUpdate("updated", a, editColumns, mw, args...)
	return
}

func (a *NgingUserRolePermission) UpdatexByFields(mw func(db.Result) db.Result, fields []string, args ...interface{}) (affected int64, err error) {

	if !a.base.Eventable() {
		return a.Param(mw, args...).UpdatexByStruct(a, fields...)
	}
	editColumns := make([]string, len(fields))
	for index, field := range fields {
		editColumns[index] = com.SnakeCase(field)
	}
	if err = DBI.FireUpdate("updating", a, editColumns, mw, args...); err != nil {
		return
	}
	if affected, err = a.Param(mw, args...).UpdatexByStruct(a, fields...); err != nil {
		return
	}
	err = DBI.FireUpdate("updated", a, editColumns, mw, args...)
	return
}

func (a *NgingUserRolePermission) UpdateField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.UpdateFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *NgingUserRolePermission) UpdateFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {

	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(kvset).Update()
	}
	m := *a
	m.FromRow(kvset)
	var editColumns []string
	for column := range kvset {
		editColumns = append(editColumns, column)
	}
	if err = DBI.FireUpdate("updating", &m, editColumns, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(kvset).Update(); err != nil {
		return
	}
	return DBI.FireUpdate("updated", &m, editColumns, mw, args...)
}

func (a *NgingUserRolePermission) UpdateValues(mw func(db.Result) db.Result, keysValues *db.KeysValues, args ...interface{}) (err error) {
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(keysValues).Update()
	}
	m := *a
	m.FromRow(keysValues.Map())
	if err = DBI.FireUpdate("updating", &m, keysValues.Keys(), mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(keysValues).Update(); err != nil {
		return
	}
	return DBI.FireUpdate("updated", &m, keysValues.Keys(), mw, args...)
}

func (a *NgingUserRolePermission) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param(mw, args...).SetSend(a).Upsert(func() error {
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("updating", a, mw, args...)
	}, func() error {
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("creating", a, nil)
	})

	if err == nil && a.base.Eventable() {
		if pk == nil {
			err = DBI.Fire("updated", a, mw, args...)
		} else {
			err = DBI.Fire("created", a, nil)
		}
	}
	return
}

func (a *NgingUserRolePermission) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

	if !a.base.Eventable() {
		return a.Param(mw, args...).Delete()
	}
	if err = DBI.Fire("deleting", a, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).Delete(); err != nil {
		return
	}
	return DBI.Fire("deleted", a, mw, args...)
}

func (a *NgingUserRolePermission) Deletex(mw func(db.Result) db.Result, args ...interface{}) (affected int64, err error) {

	if !a.base.Eventable() {
		return a.Param(mw, args...).Deletex()
	}
	if err = DBI.Fire("deleting", a, mw, args...); err != nil {
		return
	}
	if affected, err = a.Param(mw, args...).Deletex(); err != nil {
		return
	}
	err = DBI.Fire("deleted", a, mw, args...)
	return
}

func (a *NgingUserRolePermission) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param(mw, args...).Count()
}

func (a *NgingUserRolePermission) Exists(mw func(db.Result) db.Result, args ...interface{}) (bool, error) {
	return a.Param(mw, args...).Exists()
}

func (a *NgingUserRolePermission) Reset() *NgingUserRolePermission {
	a.RoleId = 0
	a.Type = ``
	a.Permission = ``
	return a
}

func (a *NgingUserRolePermission) AsMap(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["RoleId"] = a.RoleId
		r["Type"] = a.Type
		r["Permission"] = a.Permission
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "RoleId":
			r["RoleId"] = a.RoleId
		case "Type":
			r["Type"] = a.Type
		case "Permission":
			r["Permission"] = a.Permission
		}
	}
	return r
}

func (a *NgingUserRolePermission) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "role_id":
			a.RoleId = param.AsUint(value)
		case "type":
			a.Type = param.AsString(value)
		case "permission":
			a.Permission = param.AsString(value)
		}
	}
}

func (a *NgingUserRolePermission) Set(key interface{}, value ...interface{}) {
	switch k := key.(type) {
	case map[string]interface{}:
		for kk, vv := range k {
			a.Set(kk, vv)
		}
	default:
		var (
			kk string
			vv interface{}
		)
		if k, y := key.(string); y {
			kk = k
		} else {
			kk = fmt.Sprint(key)
		}
		if len(value) > 0 {
			vv = value[0]
		}
		switch kk {
		case "RoleId":
			a.RoleId = param.AsUint(vv)
		case "Type":
			a.Type = param.AsString(vv)
		case "Permission":
			a.Permission = param.AsString(vv)
		}
	}
}

func (a *NgingUserRolePermission) AsRow(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["role_id"] = a.RoleId
		r["type"] = a.Type
		r["permission"] = a.Permission
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "role_id":
			r["role_id"] = a.RoleId
		case "type":
			r["type"] = a.Type
		case "permission":
			r["permission"] = a.Permission
		}
	}
	return r
}

func (a *NgingUserRolePermission) ListPage(cond *db.Compounds, sorts ...interface{}) error {
	_, err := pagination.NewLister(a, nil, func(r db.Result) db.Result {
		return r.OrderBy(sorts...)
	}, cond.And()).Paging(a.Context())
	return err
}

func (a *NgingUserRolePermission) ListPageAs(recv interface{}, cond *db.Compounds, sorts ...interface{}) error {
	_, err := pagination.NewLister(a, recv, func(r db.Result) db.Result {
		return r.OrderBy(sorts...)
	}, cond.And()).Paging(a.Context())
	return err
}

func (a *NgingUserRolePermission) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return DBI.Fields.BatchValidate(a.Short_(), kvset)
}

func (a *NgingUserRolePermission) Validate(field string, value interface{}) error {
	return DBI.Fields.Validate(a.Short_(), field, value)
}
