// Code generated by https://github.com/src-d/go-kallax. DO NOT EDIT
// Please, do not touch the code below, and if you do, do it under your own
// risk. Take into account that all the code you write here will be completely
// erased from earth the next time you generate the kallax models.
package model

import (
	"database/sql"
	"fmt"
	"time"

	"gopkg.in/src-d/go-kallax.v1"
	"gopkg.in/src-d/go-kallax.v1/types"
)

var _ types.SQLType
var _ fmt.Formatter

type modelSaveFunc func(*kallax.Store) error

// NewCompany returns a new instance of Company.
func NewCompany() (record *Company) {
	return new(Company)
}

// GetID returns the primary key of the model.
func (r *Company) GetID() kallax.Identifier {
	return (*kallax.NumericID)(&r.ID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *Company) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return (*kallax.NumericID)(&r.ID), nil
	case "created_at":
		return &r.Timestamps.CreatedAt, nil
	case "updated_at":
		return &r.Timestamps.UpdatedAt, nil
	case "name":
		return &r.Name, nil
	case "address":
		return &r.Address, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Company: %s", col)
	}
}

// Value returns the value of the given column.
func (r *Company) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.ID, nil
	case "created_at":
		return r.Timestamps.CreatedAt, nil
	case "updated_at":
		return r.Timestamps.UpdatedAt, nil
	case "name":
		return r.Name, nil
	case "address":
		return r.Address, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Company: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *Company) NewRelationshipRecord(field string) (kallax.Record, error) {
	switch field {
	case "Users":
		return new(User), nil

	}
	return nil, fmt.Errorf("kallax: model Company has no relationship %s", field)
}

// SetRelationship sets the given relationship in the given field.
func (r *Company) SetRelationship(field string, rel interface{}) error {
	switch field {
	case "Users":
		records, ok := rel.([]kallax.Record)
		if !ok {
			return fmt.Errorf("kallax: relationship field %s needs a collection of records, not %T", field, rel)
		}

		r.Users = make([]*User, len(records))
		for i, record := range records {
			rel, ok := record.(*User)
			if !ok {
				return fmt.Errorf("kallax: element of type %T cannot be added to relationship %s", record, field)
			}
			r.Users[i] = rel
		}
		return nil

	}
	return fmt.Errorf("kallax: model Company has no relationship %s", field)
}

// CompanyStore is the entity to access the records of the type Company
// in the database.
type CompanyStore struct {
	*kallax.Store
}

// NewCompanyStore creates a new instance of CompanyStore
// using a SQL database.
func NewCompanyStore(db *sql.DB) *CompanyStore {
	return &CompanyStore{kallax.NewStore(db)}
}

// GenericStore returns the generic store of this store.
func (s *CompanyStore) GenericStore() *kallax.Store {
	return s.Store
}

// SetGenericStore changes the generic store of this store.
func (s *CompanyStore) SetGenericStore(store *kallax.Store) {
	s.Store = store
}

// Debug returns a new store that will print all SQL statements to stdout using
// the log.Printf function.
func (s *CompanyStore) Debug() *CompanyStore {
	return &CompanyStore{s.Store.Debug()}
}

// DebugWith returns a new store that will print all SQL statements using the
// given logger function.
func (s *CompanyStore) DebugWith(logger kallax.LoggerFunc) *CompanyStore {
	return &CompanyStore{s.Store.DebugWith(logger)}
}

func (s *CompanyStore) relationshipRecords(record *Company) []modelSaveFunc {
	var result []modelSaveFunc

	for i := range record.Users {
		r := record.Users[i]
		if !r.IsSaving() {
			r.AddVirtualColumn("company_id", record.GetID())
			result = append(result, func(store *kallax.Store) error {
				_, err := (&UserStore{store}).Save(r)
				return err
			})
		}
	}

	return result
}

// Insert inserts a Company in the database. A non-persisted object is
// required for this operation.
func (s *CompanyStore) Insert(record *Company) error {
	record.SetSaving(true)
	defer record.SetSaving(false)

	record.CreatedAt = record.CreatedAt.Truncate(time.Microsecond)
	record.UpdatedAt = record.UpdatedAt.Truncate(time.Microsecond)

	if err := record.BeforeSave(); err != nil {
		return err
	}

	records := s.relationshipRecords(record)

	if len(records) > 0 {
		return s.Store.Transaction(func(s *kallax.Store) error {
			if err := s.Insert(Schema.Company.BaseSchema, record); err != nil {
				return err
			}

			for _, r := range records {
				if err := r(s); err != nil {
					return err
				}
			}

			return nil
		})
	}

	return s.Store.Insert(Schema.Company.BaseSchema, record)
}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *CompanyStore) Update(record *Company, cols ...kallax.SchemaField) (updated int64, err error) {
	record.CreatedAt = record.CreatedAt.Truncate(time.Microsecond)
	record.UpdatedAt = record.UpdatedAt.Truncate(time.Microsecond)

	record.SetSaving(true)
	defer record.SetSaving(false)

	if err := record.BeforeSave(); err != nil {
		return 0, err
	}

	records := s.relationshipRecords(record)

	if len(records) > 0 {
		err = s.Store.Transaction(func(s *kallax.Store) error {
			updated, err = s.Update(Schema.Company.BaseSchema, record, cols...)
			if err != nil {
				return err
			}

			for _, r := range records {
				if err := r(s); err != nil {
					return err
				}
			}

			return nil
		})
		if err != nil {
			return 0, err
		}

		return updated, nil
	}

	return s.Store.Update(Schema.Company.BaseSchema, record, cols...)
}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *CompanyStore) Save(record *Company) (updated bool, err error) {
	if !record.IsPersisted() {
		return false, s.Insert(record)
	}

	rowsUpdated, err := s.Update(record)
	if err != nil {
		return false, err
	}

	return rowsUpdated > 0, nil
}

// Delete removes the given record from the database.
func (s *CompanyStore) Delete(record *Company) error {
	return s.Store.Delete(Schema.Company.BaseSchema, record)
}

// Find returns the set of results for the given query.
func (s *CompanyStore) Find(q *CompanyQuery) (*CompanyResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewCompanyResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *CompanyStore) MustFind(q *CompanyQuery) *CompanyResultSet {
	return NewCompanyResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *CompanyStore) Count(q *CompanyQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *CompanyStore) MustCount(q *CompanyQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `ErrNotFound` is returned if there are no results.
func (s *CompanyStore) FindOne(q *CompanyQuery) (*Company, error) {
	q.Limit(1)
	q.Offset(0)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// FindAll returns a list of all the rows returned by the given query.
func (s *CompanyStore) FindAll(q *CompanyQuery) ([]*Company, error) {
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	return rs.All()
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *CompanyStore) MustFindOne(q *CompanyQuery) *Company {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the Company with the data in the database and
// makes it writable.
func (s *CompanyStore) Reload(record *Company) error {
	return s.Store.Reload(Schema.Company.BaseSchema, record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *CompanyStore) Transaction(callback func(*CompanyStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&CompanyStore{store})
	})
}

// RemoveUsers removes the given items of the Users field of the
// model. If no items are given, it removes all of them.
// The items will also be removed from the passed record inside this method.
// Note that is required that `Users` is not empty. This method clears the
// the elements of Users in a model, it does not retrieve them to know
// what relationships the model has.
func (s *CompanyStore) RemoveUsers(record *Company, deleted ...*User) error {
	var updated []*User
	var clear bool
	if len(deleted) == 0 {
		clear = true
		deleted = record.Users
		if len(deleted) == 0 {
			return nil
		}
	}

	if len(deleted) > 1 {
		err := s.Store.Transaction(func(s *kallax.Store) error {
			for _, d := range deleted {
				var r kallax.Record = d

				if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
					if err := beforeDeleter.BeforeDelete(); err != nil {
						return err
					}
				}

				if err := s.Delete(Schema.User.BaseSchema, d); err != nil {
					return err
				}

				if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
					if err := afterDeleter.AfterDelete(); err != nil {
						return err
					}
				}
			}
			return nil
		})

		if err != nil {
			return err
		}

		if clear {
			record.Users = nil
			return nil
		}
	} else {
		var r kallax.Record = deleted[0]
		if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
			if err := beforeDeleter.BeforeDelete(); err != nil {
				return err
			}
		}

		var err error
		if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
			err = s.Store.Transaction(func(s *kallax.Store) error {
				err := s.Delete(Schema.User.BaseSchema, r)
				if err != nil {
					return err
				}

				return afterDeleter.AfterDelete()
			})
		} else {
			err = s.Store.Delete(Schema.User.BaseSchema, deleted[0])
		}

		if err != nil {
			return err
		}
	}

	for _, r := range record.Users {
		var found bool
		for _, d := range deleted {
			if d.GetID().Equals(r.GetID()) {
				found = true
				break
			}
		}
		if !found {
			updated = append(updated, r)
		}
	}
	record.Users = updated
	return nil
}

// CompanyQuery is the object used to create queries for the Company
// entity.
type CompanyQuery struct {
	*kallax.BaseQuery
}

// NewCompanyQuery returns a new instance of CompanyQuery.
func NewCompanyQuery() *CompanyQuery {
	return &CompanyQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.Company.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *CompanyQuery) Select(columns ...kallax.SchemaField) *CompanyQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *CompanyQuery) SelectNot(columns ...kallax.SchemaField) *CompanyQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *CompanyQuery) Copy() *CompanyQuery {
	return &CompanyQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *CompanyQuery) Order(cols ...kallax.ColumnOrder) *CompanyQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *CompanyQuery) BatchSize(size uint64) *CompanyQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *CompanyQuery) Limit(n uint64) *CompanyQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *CompanyQuery) Offset(n uint64) *CompanyQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *CompanyQuery) Where(cond kallax.Condition) *CompanyQuery {
	q.BaseQuery.Where(cond)
	return q
}

func (q *CompanyQuery) WithUsers(cond kallax.Condition) *CompanyQuery {
	q.AddRelation(Schema.User.BaseSchema, "Users", kallax.OneToMany, cond)
	return q
}

// FindByID adds a new filter to the query that will require that
// the ID property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *CompanyQuery) FindByID(v ...int64) *CompanyQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.In(Schema.Company.ID, values...))
}

// FindByCreatedAt adds a new filter to the query that will require that
// the CreatedAt property is equal to the passed value.
func (q *CompanyQuery) FindByCreatedAt(cond kallax.ScalarCond, v time.Time) *CompanyQuery {
	return q.Where(cond(Schema.Company.CreatedAt, v))
}

// FindByUpdatedAt adds a new filter to the query that will require that
// the UpdatedAt property is equal to the passed value.
func (q *CompanyQuery) FindByUpdatedAt(cond kallax.ScalarCond, v time.Time) *CompanyQuery {
	return q.Where(cond(Schema.Company.UpdatedAt, v))
}

// FindByName adds a new filter to the query that will require that
// the Name property is equal to the passed value.
func (q *CompanyQuery) FindByName(v string) *CompanyQuery {
	return q.Where(kallax.Eq(Schema.Company.Name, v))
}

// FindByAddress adds a new filter to the query that will require that
// the Address property is equal to the passed value.
func (q *CompanyQuery) FindByAddress(v string) *CompanyQuery {
	return q.Where(kallax.Eq(Schema.Company.Address, v))
}

// CompanyResultSet is the set of results returned by a query to the
// database.
type CompanyResultSet struct {
	ResultSet kallax.ResultSet
	last      *Company
	lastErr   error
}

// NewCompanyResultSet creates a new result set for rows of the type
// Company.
func NewCompanyResultSet(rs kallax.ResultSet) *CompanyResultSet {
	return &CompanyResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *CompanyResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.ResultSet.Close()
		rs.last = nil
		return false
	}

	var record kallax.Record
	record, rs.lastErr = rs.ResultSet.Get(Schema.Company.BaseSchema)
	if rs.lastErr != nil {
		rs.last = nil
	} else {
		var ok bool
		rs.last, ok = record.(*Company)
		if !ok {
			rs.lastErr = fmt.Errorf("kallax: unable to convert record to *Company")
			rs.last = nil
		}
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *CompanyResultSet) Get() (*Company, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *CompanyResultSet) ForEach(fn func(*Company) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *CompanyResultSet) All() ([]*Company, error) {
	var result []*Company
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *CompanyResultSet) One() (*Company, error) {
	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// Err returns the last error occurred.
func (rs *CompanyResultSet) Err() error {
	return rs.lastErr
}

// Close closes the result set.
func (rs *CompanyResultSet) Close() error {
	return rs.ResultSet.Close()
}

// NewUser returns a new instance of User.
func NewUser() (record *User) {
	return new(User)
}

// GetID returns the primary key of the model.
func (r *User) GetID() kallax.Identifier {
	return (*kallax.NumericID)(&r.ID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *User) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return (*kallax.NumericID)(&r.ID), nil
	case "created_at":
		return &r.Timestamps.CreatedAt, nil
	case "updated_at":
		return &r.Timestamps.UpdatedAt, nil
	case "email":
		return &r.Email, nil
	case "salt":
		return &r.Salt, nil
	case "passhash":
		return &r.Passhash, nil
	case "name":
		return &r.Name, nil
	case "phone":
		return &r.Phone, nil
	case "company_id":
		return types.Nullable(kallax.VirtualColumn("company_id", r, new(kallax.NumericID))), nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in User: %s", col)
	}
}

// Value returns the value of the given column.
func (r *User) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.ID, nil
	case "created_at":
		return r.Timestamps.CreatedAt, nil
	case "updated_at":
		return r.Timestamps.UpdatedAt, nil
	case "email":
		return r.Email, nil
	case "salt":
		return r.Salt, nil
	case "passhash":
		return r.Passhash, nil
	case "name":
		return r.Name, nil
	case "phone":
		return r.Phone, nil
	case "company_id":
		v := r.Model.VirtualColumn(col)
		if v == nil {
			return nil, kallax.ErrEmptyVirtualColumn
		}
		return v, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in User: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *User) NewRelationshipRecord(field string) (kallax.Record, error) {
	switch field {
	case "Company":
		return new(Company), nil

	}
	return nil, fmt.Errorf("kallax: model User has no relationship %s", field)
}

// SetRelationship sets the given relationship in the given field.
func (r *User) SetRelationship(field string, rel interface{}) error {
	switch field {
	case "Company":
		val, ok := rel.(*Company)
		if !ok {
			return fmt.Errorf("kallax: record of type %t can't be assigned to relationship Company", rel)
		}
		if !val.GetID().IsEmpty() {
			r.Company = val
		}

		return nil

	}
	return fmt.Errorf("kallax: model User has no relationship %s", field)
}

// UserStore is the entity to access the records of the type User
// in the database.
type UserStore struct {
	*kallax.Store
}

// NewUserStore creates a new instance of UserStore
// using a SQL database.
func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{kallax.NewStore(db)}
}

// GenericStore returns the generic store of this store.
func (s *UserStore) GenericStore() *kallax.Store {
	return s.Store
}

// SetGenericStore changes the generic store of this store.
func (s *UserStore) SetGenericStore(store *kallax.Store) {
	s.Store = store
}

// Debug returns a new store that will print all SQL statements to stdout using
// the log.Printf function.
func (s *UserStore) Debug() *UserStore {
	return &UserStore{s.Store.Debug()}
}

// DebugWith returns a new store that will print all SQL statements using the
// given logger function.
func (s *UserStore) DebugWith(logger kallax.LoggerFunc) *UserStore {
	return &UserStore{s.Store.DebugWith(logger)}
}

func (s *UserStore) inverseRecords(record *User) []modelSaveFunc {
	var result []modelSaveFunc

	if record.Company != nil && !record.Company.IsSaving() {
		record.AddVirtualColumn("company_id", record.Company.GetID())
		result = append(result, func(store *kallax.Store) error {
			_, err := (&CompanyStore{store}).Save(record.Company)
			return err
		})
	}

	return result
}

// Insert inserts a User in the database. A non-persisted object is
// required for this operation.
func (s *UserStore) Insert(record *User) error {
	record.SetSaving(true)
	defer record.SetSaving(false)

	record.CreatedAt = record.CreatedAt.Truncate(time.Microsecond)
	record.UpdatedAt = record.UpdatedAt.Truncate(time.Microsecond)

	if err := record.BeforeSave(); err != nil {
		return err
	}

	inverseRecords := s.inverseRecords(record)

	if len(inverseRecords) > 0 {
		return s.Store.Transaction(func(s *kallax.Store) error {
			for _, r := range inverseRecords {
				if err := r(s); err != nil {
					return err
				}
			}

			if err := s.Insert(Schema.User.BaseSchema, record); err != nil {
				return err
			}

			return nil
		})
	}

	return s.Store.Insert(Schema.User.BaseSchema, record)
}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *UserStore) Update(record *User, cols ...kallax.SchemaField) (updated int64, err error) {
	record.CreatedAt = record.CreatedAt.Truncate(time.Microsecond)
	record.UpdatedAt = record.UpdatedAt.Truncate(time.Microsecond)

	record.SetSaving(true)
	defer record.SetSaving(false)

	if err := record.BeforeSave(); err != nil {
		return 0, err
	}

	inverseRecords := s.inverseRecords(record)

	if len(inverseRecords) > 0 {
		err = s.Store.Transaction(func(s *kallax.Store) error {
			for _, r := range inverseRecords {
				if err := r(s); err != nil {
					return err
				}
			}

			updated, err = s.Update(Schema.User.BaseSchema, record, cols...)
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return 0, err
		}

		return updated, nil
	}

	return s.Store.Update(Schema.User.BaseSchema, record, cols...)
}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *UserStore) Save(record *User) (updated bool, err error) {
	if !record.IsPersisted() {
		return false, s.Insert(record)
	}

	rowsUpdated, err := s.Update(record)
	if err != nil {
		return false, err
	}

	return rowsUpdated > 0, nil
}

// Delete removes the given record from the database.
func (s *UserStore) Delete(record *User) error {
	return s.Store.Delete(Schema.User.BaseSchema, record)
}

// Find returns the set of results for the given query.
func (s *UserStore) Find(q *UserQuery) (*UserResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewUserResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *UserStore) MustFind(q *UserQuery) *UserResultSet {
	return NewUserResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *UserStore) Count(q *UserQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *UserStore) MustCount(q *UserQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `ErrNotFound` is returned if there are no results.
func (s *UserStore) FindOne(q *UserQuery) (*User, error) {
	q.Limit(1)
	q.Offset(0)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// FindAll returns a list of all the rows returned by the given query.
func (s *UserStore) FindAll(q *UserQuery) ([]*User, error) {
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	return rs.All()
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *UserStore) MustFindOne(q *UserQuery) *User {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the User with the data in the database and
// makes it writable.
func (s *UserStore) Reload(record *User) error {
	return s.Store.Reload(Schema.User.BaseSchema, record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *UserStore) Transaction(callback func(*UserStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&UserStore{store})
	})
}

// UserQuery is the object used to create queries for the User
// entity.
type UserQuery struct {
	*kallax.BaseQuery
}

// NewUserQuery returns a new instance of UserQuery.
func NewUserQuery() *UserQuery {
	return &UserQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.User.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *UserQuery) Select(columns ...kallax.SchemaField) *UserQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *UserQuery) SelectNot(columns ...kallax.SchemaField) *UserQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *UserQuery) Copy() *UserQuery {
	return &UserQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *UserQuery) Order(cols ...kallax.ColumnOrder) *UserQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *UserQuery) BatchSize(size uint64) *UserQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *UserQuery) Limit(n uint64) *UserQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *UserQuery) Offset(n uint64) *UserQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *UserQuery) Where(cond kallax.Condition) *UserQuery {
	q.BaseQuery.Where(cond)
	return q
}

func (q *UserQuery) WithCompany() *UserQuery {
	q.AddRelation(Schema.Company.BaseSchema, "Company", kallax.OneToOne, nil)
	return q
}

// FindByID adds a new filter to the query that will require that
// the ID property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *UserQuery) FindByID(v ...int64) *UserQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.In(Schema.User.ID, values...))
}

// FindByCreatedAt adds a new filter to the query that will require that
// the CreatedAt property is equal to the passed value.
func (q *UserQuery) FindByCreatedAt(cond kallax.ScalarCond, v time.Time) *UserQuery {
	return q.Where(cond(Schema.User.CreatedAt, v))
}

// FindByUpdatedAt adds a new filter to the query that will require that
// the UpdatedAt property is equal to the passed value.
func (q *UserQuery) FindByUpdatedAt(cond kallax.ScalarCond, v time.Time) *UserQuery {
	return q.Where(cond(Schema.User.UpdatedAt, v))
}

// FindByEmail adds a new filter to the query that will require that
// the Email property is equal to the passed value.
func (q *UserQuery) FindByEmail(v string) *UserQuery {
	return q.Where(kallax.Eq(Schema.User.Email, v))
}

// FindBySalt adds a new filter to the query that will require that
// the Salt property is equal to the passed value.
func (q *UserQuery) FindBySalt(v string) *UserQuery {
	return q.Where(kallax.Eq(Schema.User.Salt, v))
}

// FindByPasshash adds a new filter to the query that will require that
// the Passhash property is equal to the passed value.
func (q *UserQuery) FindByPasshash(v string) *UserQuery {
	return q.Where(kallax.Eq(Schema.User.Passhash, v))
}

// FindByName adds a new filter to the query that will require that
// the Name property is equal to the passed value.
func (q *UserQuery) FindByName(v string) *UserQuery {
	return q.Where(kallax.Eq(Schema.User.Name, v))
}

// FindByPhone adds a new filter to the query that will require that
// the Phone property is equal to the passed value.
func (q *UserQuery) FindByPhone(v string) *UserQuery {
	return q.Where(kallax.Eq(Schema.User.Phone, v))
}

// FindByCompany adds a new filter to the query that will require that
// the foreign key of Company is equal to the passed value.
func (q *UserQuery) FindByCompany(v int64) *UserQuery {
	return q.Where(kallax.Eq(Schema.User.CompanyFK, v))
}

// UserResultSet is the set of results returned by a query to the
// database.
type UserResultSet struct {
	ResultSet kallax.ResultSet
	last      *User
	lastErr   error
}

// NewUserResultSet creates a new result set for rows of the type
// User.
func NewUserResultSet(rs kallax.ResultSet) *UserResultSet {
	return &UserResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *UserResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.ResultSet.Close()
		rs.last = nil
		return false
	}

	var record kallax.Record
	record, rs.lastErr = rs.ResultSet.Get(Schema.User.BaseSchema)
	if rs.lastErr != nil {
		rs.last = nil
	} else {
		var ok bool
		rs.last, ok = record.(*User)
		if !ok {
			rs.lastErr = fmt.Errorf("kallax: unable to convert record to *User")
			rs.last = nil
		}
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *UserResultSet) Get() (*User, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *UserResultSet) ForEach(fn func(*User) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *UserResultSet) All() ([]*User, error) {
	var result []*User
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *UserResultSet) One() (*User, error) {
	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// Err returns the last error occurred.
func (rs *UserResultSet) Err() error {
	return rs.lastErr
}

// Close closes the result set.
func (rs *UserResultSet) Close() error {
	return rs.ResultSet.Close()
}

type schema struct {
	Company *schemaCompany
	User    *schemaUser
}

type schemaCompany struct {
	*kallax.BaseSchema
	ID        kallax.SchemaField
	CreatedAt kallax.SchemaField
	UpdatedAt kallax.SchemaField
	Name      kallax.SchemaField
	Address   kallax.SchemaField
}

type schemaUser struct {
	*kallax.BaseSchema
	ID        kallax.SchemaField
	CreatedAt kallax.SchemaField
	UpdatedAt kallax.SchemaField
	Email     kallax.SchemaField
	Salt      kallax.SchemaField
	Passhash  kallax.SchemaField
	Name      kallax.SchemaField
	Phone     kallax.SchemaField
	CompanyFK kallax.SchemaField
}

var Schema = &schema{
	Company: &schemaCompany{
		BaseSchema: kallax.NewBaseSchema(
			"companies",
			"__company",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{
				"Users": kallax.NewForeignKey("company_id", false),
			},
			func() kallax.Record {
				return new(Company)
			},
			true,
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("created_at"),
			kallax.NewSchemaField("updated_at"),
			kallax.NewSchemaField("name"),
			kallax.NewSchemaField("address"),
		),
		ID:        kallax.NewSchemaField("id"),
		CreatedAt: kallax.NewSchemaField("created_at"),
		UpdatedAt: kallax.NewSchemaField("updated_at"),
		Name:      kallax.NewSchemaField("name"),
		Address:   kallax.NewSchemaField("address"),
	},
	User: &schemaUser{
		BaseSchema: kallax.NewBaseSchema(
			"users",
			"__user",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{
				"Company": kallax.NewForeignKey("company_id", true),
			},
			func() kallax.Record {
				return new(User)
			},
			true,
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("created_at"),
			kallax.NewSchemaField("updated_at"),
			kallax.NewSchemaField("email"),
			kallax.NewSchemaField("salt"),
			kallax.NewSchemaField("passhash"),
			kallax.NewSchemaField("name"),
			kallax.NewSchemaField("phone"),
			kallax.NewSchemaField("company_id"),
		),
		ID:        kallax.NewSchemaField("id"),
		CreatedAt: kallax.NewSchemaField("created_at"),
		UpdatedAt: kallax.NewSchemaField("updated_at"),
		Email:     kallax.NewSchemaField("email"),
		Salt:      kallax.NewSchemaField("salt"),
		Passhash:  kallax.NewSchemaField("passhash"),
		Name:      kallax.NewSchemaField("name"),
		Phone:     kallax.NewSchemaField("phone"),
		CompanyFK: kallax.NewSchemaField("company_id"),
	},
}
