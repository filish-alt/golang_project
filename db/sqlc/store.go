package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) (*Store){
	return &Store{
		db :db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error{
       tx,err := store.db.BeginTx(ctx,nil)

	   if err !=nil{
		return nil
	   }
q:=New(tx)
 err =fn(q)

 if err != nil{
	if rbErr := tx.Rollback(); rbErr !=nil{
		return fmt.Errorf("txt: %v rberr: %v",err,rbErr)
	}
	return err
}
return tx.Commit()
}

type TransferParam struct{
   Username pgtype
}

type TransferResult struct{
    User User
}
func(store *Store) TransferMoney(context context.Context,transferParam TransferParam) (TransferResult, error){
    var result TransferResult

	err:=store.execTx(context, func(q *Queries) error {
		result.User, err := q.CreateUsers(context, CreateUsersParams{
               Username: transferParam.Username,
		})
		if err!=nil{
			return err
		}
		return nil
	})

	return result, err
}