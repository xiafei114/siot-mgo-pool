// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package orm

import (
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// Fielder define field info
type Fielder interface {
	String() string
	FieldType() int
	SetRaw(interface{}) error
	RawValue() interface{}
}

// Ormer define the orm interface
type Ormer interface {
	Read(md interface{}, cols ...string) error
	ReadOrCreate(md interface{}, col1 string, cols ...string) (bool, int64, error)
	Insert(interface{}) error
	// InsertOrUpdate(md interface{}, colConflitAndArgs ...string) (int64, error)
	// InsertMulti(bulk int, mds interface{}) (int64, error)
	// Update(md interface{}, cols ...string) (int64, error)
	// Delete(md interface{}, cols ...string) (int64, error)

	QueryTable(ptrStructOrTableName interface{}) QuerySeter

	Begin() error
	Commit() error
	Rollback() error
	Using(name string) error
}
type QuerySeter interface {
	Filter(string, ...interface{}) QuerySeter
	Exclude(string, ...interface{}) QuerySeter
	SetCond(*Condition) QuerySeter
	GetCond() *Condition
	Limit(limit interface{}, args ...interface{}) QuerySeter
	Offset(offset interface{}) QuerySeter
	GroupBy(exprs ...string) QuerySeter
	OrderBy(exprs ...string) QuerySeter
	RelatedSel(params ...interface{}) QuerySeter
	Distinct() QuerySeter
	ForUpdate() QuerySeter
	Count() (int64, error)
	Exist() bool
	Update(Params) (int64, error)
	Delete() (int64, error)
	All(interface{}, ...string) (int64, error)
	One(interface{}, ...string) error
	Values(results *[]Params, exprs ...string) (int64, error)
	ValuesList(results *[]ParamsList, exprs ...string) (int64, error)
	ValuesFlat(result *ParamsList, expr string) (int64, error)
	RowsToMap(result *Params, keyCol, valueCol string) (int64, error)
	RowsToStruct(ptrStruct interface{}, keyCol, valueCol string) (int64, error)
}
type dbQuerier interface {
	GetDB() *mongo.Database
	Begin() error
	Commit() error
	Rollback() error
}

// base database struct
type dbBaser interface {
	Read(dbQuerier, *modelInfo, reflect.Value, interface{}, *time.Location, []string) error
	FindOne(*querySet, *modelInfo, *Condition, interface{}, *time.Location, []string) error
	Find(*querySet, *modelInfo, *Condition, interface{}, *time.Location, []string) (int64, error)
	TimeFromDB(*time.Time, *time.Location)
	TimeToDB(*time.Time, *time.Location)
}