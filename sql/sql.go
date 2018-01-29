package sql

import (
	"time"

	"github.com/OrderMyGear/order-service/order"
	"github.com/PRAgarawal/eet/eet"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	sq "github.com/lann/squirrel"
)

// Repository is a MySQL repository for orders
type Repository struct {
	db *sqlx.DB
}

// NewRepository returns a new instance of a mysql-backed order service repository.
func NewRepository(adapter string, dsn string) (*Repository, error) {
	db, err := sqlx.Connect(adapter, dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) Ping() error {
	return r.db.Ping()
}

func (r *Repository) FindMeetingMembers(filter *eet.MeetingMemberFilter) ([]*order.Item, error) {
	var items []*order.Item
	sb := sq.Select("meeting_member.*").
		From("meeting_member")
	q, args, _ := sb.ToSql()

	err := r.db.Select(&items, q, args...)
	if err != nil {
		return items, err
	}

	return items, nil
}

func (r *Repository) FindMeetingMembersByTeam(filter *eet.MeetingMemberFilter) ([]*order.Item, error) {
	var items []*order.Item
	sb := sq.Select("meeting_member.*").
		From("meeting_member")
	q, args, _ := sb.ToSql()

	err := r.db.Select(&items, q, args...)
	if err != nil {
		return items, err
	}

	return items, nil
}

func (r *Repository) StoreMeetingMember(i *order.Item) error {
	tx, err := r.db.Beginx()
	if nil != err {
		return err
	}

	// Insert
	i.CreatedAt = time.Now()
	_, err = tx.NamedExec("INSERT INTO `meeting_member` (name,type,quantity,unit_price,total,external_id,status,order_id,created_at) VALUES (:name,:type,:quantity,:unit_price,:total,:external_id,:status,:order_id,:created_at)", i)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *Repository) DeleteMeetingMembers(ids []int) error {
	return nil
}
