// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"time"

	"github.com/webx-top/com"
	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/db/lib/factory/pagination"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/param"
)

type Slice_NgingFileMoved []*NgingFileMoved

func (s Slice_NgingFileMoved) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingFileMoved) RangeRaw(fn func(m *NgingFileMoved) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingFileMoved) GroupBy(keyField string) map[string][]*NgingFileMoved {
	r := map[string][]*NgingFileMoved{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*NgingFileMoved{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (s Slice_NgingFileMoved) KeyBy(keyField string) map[string]*NgingFileMoved {
	r := map[string]*NgingFileMoved{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (s Slice_NgingFileMoved) AsKV(keyField string, valueField string) param.Store {
	r := param.Store{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (s Slice_NgingFileMoved) Transform(transfers map[string]param.Transfer) []param.Store {
	r := make([]param.Store, len(s))
	for idx, row := range s {
		r[idx] = row.AsMap().Transform(transfers)
	}
	return r
}

func (s Slice_NgingFileMoved) FromList(data interface{}) Slice_NgingFileMoved {
	values, ok := data.([]*NgingFileMoved)
	if !ok {
		for _, value := range data.([]interface{}) {
			row := &NgingFileMoved{}
			row.FromRow(value.(map[string]interface{}))
			s = append(s, row)
		}
		return s
	}
	s = append(s, values...)

	return s
}

func NewNgingFileMoved(ctx echo.Context) *NgingFileMoved {
	m := &NgingFileMoved{}
	m.SetContext(ctx)
	return m
}

// NgingFileMoved 文件移动记录
type NgingFileMoved struct {
	base    factory.Base
	objects []*NgingFileMoved

	Id      uint64 `db:"id,omitempty,pk" bson:"id,omitempty" comment:"主键" json:"id" xml:"id"`
	FileId  uint64 `db:"file_id" bson:"file_id" comment:"文件ID" json:"file_id" xml:"file_id"`
	From    string `db:"from" bson:"from" comment:"文件原路径" json:"from" xml:"from"`
	To      string `db:"to" bson:"to" comment:"文件新路径" json:"to" xml:"to"`
	ThumbId uint64 `db:"thumb_id" bson:"thumb_id" comment:"缩略图ID(缩略图时有效)" json:"thumb_id" xml:"thumb_id"`
	Created uint   `db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
}

// - base function

func (a *NgingFileMoved) Trans() *factory.Transaction {
	return a.base.Trans()
}

func (a *NgingFileMoved) Use(trans *factory.Transaction) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *NgingFileMoved) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *NgingFileMoved) EventON(on ...bool) factory.Model {
	a.base.EventON(on...)
	return a
}

func (a *NgingFileMoved) EventOFF(off ...bool) factory.Model {
	a.base.EventOFF(off...)
	return a
}

func (a *NgingFileMoved) Context() echo.Context {
	return a.base.Context()
}

func (a *NgingFileMoved) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *NgingFileMoved) SetNamer(namer func(factory.Model) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *NgingFileMoved) Namer() func(factory.Model) string {
	return a.base.Namer()
}

func (a *NgingFileMoved) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *NgingFileMoved) Param(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam().SetMiddleware(mw).SetArgs(args...)
	}
	return a.base.Param().SetMiddleware(mw).SetArgs(args...)
}

// - current function

func (a *NgingFileMoved) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName, connID[0]).Use(a.base.Trans())
	}
	return factory.NewModel(structName, a.base.ConnID()).Use(a.base.Trans())
}

func (a *NgingFileMoved) Objects() []*NgingFileMoved {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *NgingFileMoved) XObjects() Slice_NgingFileMoved {
	return Slice_NgingFileMoved(a.Objects())
}

func (a *NgingFileMoved) NewObjects() factory.Ranger {
	return &Slice_NgingFileMoved{}
}

func (a *NgingFileMoved) InitObjects() *[]*NgingFileMoved {
	a.objects = []*NgingFileMoved{}
	return &a.objects
}

func (a *NgingFileMoved) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *NgingFileMoved) Short_() string {
	return "nging_file_moved"
}

func (a *NgingFileMoved) Struct_() string {
	return "NgingFileMoved"
}

func (a *NgingFileMoved) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *NgingFileMoved) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.Use(source.Trans())
	a.SetNamer(source.Namer())
	return a
}

func (a *NgingFileMoved) Get(mw func(db.Result) db.Result, args ...interface{}) (err error) {
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

func (a *NgingFileMoved) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
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
		case *[]*NgingFileMoved:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFileMoved(*v))
		case []*NgingFileMoved:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFileMoved(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingFileMoved) GroupBy(keyField string, inputRows ...[]*NgingFileMoved) map[string][]*NgingFileMoved {
	var rows Slice_NgingFileMoved
	if len(inputRows) > 0 {
		rows = Slice_NgingFileMoved(inputRows[0])
	} else {
		rows = Slice_NgingFileMoved(a.Objects())
	}
	return rows.GroupBy(keyField)
}

func (a *NgingFileMoved) KeyBy(keyField string, inputRows ...[]*NgingFileMoved) map[string]*NgingFileMoved {
	var rows Slice_NgingFileMoved
	if len(inputRows) > 0 {
		rows = Slice_NgingFileMoved(inputRows[0])
	} else {
		rows = Slice_NgingFileMoved(a.Objects())
	}
	return rows.KeyBy(keyField)
}

func (a *NgingFileMoved) AsKV(keyField string, valueField string, inputRows ...[]*NgingFileMoved) param.Store {
	var rows Slice_NgingFileMoved
	if len(inputRows) > 0 {
		rows = Slice_NgingFileMoved(inputRows[0])
	} else {
		rows = Slice_NgingFileMoved(a.Objects())
	}
	return rows.AsKV(keyField, valueField)
}

func (a *NgingFileMoved) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
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
		case *[]*NgingFileMoved:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFileMoved(*v))
		case []*NgingFileMoved:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFileMoved(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingFileMoved) Insert() (pk interface{}, err error) {
	a.Created = uint(time.Now().Unix())
	a.Id = 0
	if a.base.Eventable() {
		err = DBI.Fire("creating", a, nil)
		if err != nil {
			return
		}
	}
	pk, err = a.Param(nil).SetSend(a).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint64(v)
		}
	}
	if err == nil && a.base.Eventable() {
		err = DBI.Fire("created", a, nil)
	}
	return
}

func (a *NgingFileMoved) Update(mw func(db.Result) db.Result, args ...interface{}) (err error) {

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

func (a *NgingFileMoved) Updatex(mw func(db.Result) db.Result, args ...interface{}) (affected int64, err error) {

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

func (a *NgingFileMoved) UpdateByFields(mw func(db.Result) db.Result, fields []string, args ...interface{}) (err error) {

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

func (a *NgingFileMoved) UpdatexByFields(mw func(db.Result) db.Result, fields []string, args ...interface{}) (affected int64, err error) {

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

func (a *NgingFileMoved) UpdateField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.UpdateFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *NgingFileMoved) UpdateFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {

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

func (a *NgingFileMoved) UpdateValues(mw func(db.Result) db.Result, keysValues *db.KeysValues, args ...interface{}) (err error) {
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

func (a *NgingFileMoved) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param(mw, args...).SetSend(a).Upsert(func() error {
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("updating", a, mw, args...)
	}, func() error {
		a.Created = uint(time.Now().Unix())
		a.Id = 0
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("creating", a, nil)
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint64(v)
		}
	}
	if err == nil && a.base.Eventable() {
		if pk == nil {
			err = DBI.Fire("updated", a, mw, args...)
		} else {
			err = DBI.Fire("created", a, nil)
		}
	}
	return
}

func (a *NgingFileMoved) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

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

func (a *NgingFileMoved) Deletex(mw func(db.Result) db.Result, args ...interface{}) (affected int64, err error) {

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

func (a *NgingFileMoved) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param(mw, args...).Count()
}

func (a *NgingFileMoved) Exists(mw func(db.Result) db.Result, args ...interface{}) (bool, error) {
	return a.Param(mw, args...).Exists()
}

func (a *NgingFileMoved) Reset() *NgingFileMoved {
	a.Id = 0
	a.FileId = 0
	a.From = ``
	a.To = ``
	a.ThumbId = 0
	a.Created = 0
	return a
}

func (a *NgingFileMoved) AsMap(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["Id"] = a.Id
		r["FileId"] = a.FileId
		r["From"] = a.From
		r["To"] = a.To
		r["ThumbId"] = a.ThumbId
		r["Created"] = a.Created
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "Id":
			r["Id"] = a.Id
		case "FileId":
			r["FileId"] = a.FileId
		case "From":
			r["From"] = a.From
		case "To":
			r["To"] = a.To
		case "ThumbId":
			r["ThumbId"] = a.ThumbId
		case "Created":
			r["Created"] = a.Created
		}
	}
	return r
}

func (a *NgingFileMoved) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "id":
			a.Id = param.AsUint64(value)
		case "file_id":
			a.FileId = param.AsUint64(value)
		case "from":
			a.From = param.AsString(value)
		case "to":
			a.To = param.AsString(value)
		case "thumb_id":
			a.ThumbId = param.AsUint64(value)
		case "created":
			a.Created = param.AsUint(value)
		}
	}
}

func (a *NgingFileMoved) Set(key interface{}, value ...interface{}) {
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
		case "Id":
			a.Id = param.AsUint64(vv)
		case "FileId":
			a.FileId = param.AsUint64(vv)
		case "From":
			a.From = param.AsString(vv)
		case "To":
			a.To = param.AsString(vv)
		case "ThumbId":
			a.ThumbId = param.AsUint64(vv)
		case "Created":
			a.Created = param.AsUint(vv)
		}
	}
}

func (a *NgingFileMoved) AsRow(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["id"] = a.Id
		r["file_id"] = a.FileId
		r["from"] = a.From
		r["to"] = a.To
		r["thumb_id"] = a.ThumbId
		r["created"] = a.Created
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "id":
			r["id"] = a.Id
		case "file_id":
			r["file_id"] = a.FileId
		case "from":
			r["from"] = a.From
		case "to":
			r["to"] = a.To
		case "thumb_id":
			r["thumb_id"] = a.ThumbId
		case "created":
			r["created"] = a.Created
		}
	}
	return r
}

func (a *NgingFileMoved) ListPage(cond *db.Compounds, sorts ...interface{}) error {
	_, err := pagination.NewLister(a, nil, func(r db.Result) db.Result {
		return r.OrderBy(sorts...)
	}, cond.And()).Paging(a.Context())
	return err
}

func (a *NgingFileMoved) ListPageAs(recv interface{}, cond *db.Compounds, sorts ...interface{}) error {
	_, err := pagination.NewLister(a, recv, func(r db.Result) db.Result {
		return r.OrderBy(sorts...)
	}, cond.And()).Paging(a.Context())
	return err
}

func (a *NgingFileMoved) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return DBI.Fields.BatchValidate(a.Short_(), kvset)
}

func (a *NgingFileMoved) Validate(field string, value interface{}) error {
	return DBI.Fields.Validate(a.Short_(), field, value)
}
