package sql

import (
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	sq "github.com/lann/squirrel"
	"github.com/OrderMyGear/order-service/order"
	"github.com/PRAgarawal/eet/eet"
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

func (r *Repository) StoreMeetingMember(i *order.Item) error {
	tx, err := r.db.Beginx()
	if nil != err {
		return err
	}

	// Insert
	i.CreatedAt = time.Now()
	result, err := tx.NamedExec("INSERT INTO `meeting_member` (name,type,quantity,unit_price,total,external_id,status,order_id,created_at) VALUES (:name,:type,:quantity,:unit_price,:total,:external_id,:status,:order_id,:created_at)", i)
	if err != nil {
		tx.Rollback()
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}
	i.Id = strconv.FormatInt(id, 10)

	if err = storeAttributes(tx, i); err != nil {
		tx.Rollback()
		return err
	}

	itemId, _ := strconv.Atoi(i.Id)
	if i.Status != "" {
		if _, err := tx.NamedExec("INSERT INTO `status` (item_id,status,created_at) VALUES (:item_id,:status,:created_at)", order.Status{
			ItemId:    itemId,
			Status:    i.Status,
			CreatedAt: time.Now(),
		}); err != nil {
			return err
		}
	}

	if err = updateOrderTotalWithItem(tx, i); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}