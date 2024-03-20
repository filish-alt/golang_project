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

func NewStore(db *sql.DB) *Store{
	return &Store{
		db :db,
		Queries: New(db),
	}
}
func NewSQLDBWrapper(db *sql.DB) *Store {
    return &Store{db: db}
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

type TransferTxParams struct{
	FromAccountID int64 `json:"from_account_id`
	ToAccountID   int64 `json:"To_account_id`
	Amount        int64  `json:"amount`
}
type TranferTxResult struct{
	FromAccount Account `json:"from_account`
	ToAccount   Account `json:"To_account`
	Transfer    Transfer`json:"transfer`
	FromEntry   Entry `json:from_entry`
	ToEntry     Entry `json:To_entry`

}

func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TranferTxResult, error){
	var result TranferTxResult
	
	err := store.execTx(ctx,func(q *Queries) error{
		var err error
		result.Transfer,err =q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID :arg.FromAccountID,
			ToAccountID: arg.ToAccountID,
			Amount: arg.Amount,
		})
		if err != nil{
			return err
		}
		result.FromEntry,err =q.CreateEntry(ctx,CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount: -arg.Amount,
		})
		if err != nil{
			return err
		}
		result.ToEntry,err =q.CreateEntry(ctx,CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount: arg.Amount,
		})
		if err != nil{
			return err
		}
		return nil
	})

	return result ,err
}
// type TransferParam struct{
//    Username pgtype
// }

// type TransferResult struct{
//     User User
// }
// func(store *Store) TransferMoney(context context.Context,transferParam TransferParam) (TransferResult, error){
//     var result TransferResult

// 	err:=store.execTx(context, func(q *Queries) error {
// 		result.User, err := q.CreateUsers(context, CreateUsersParams{
//                Username: transferParam.Username,
// 		})
// 		if err!=nil{
// 			return err
// 		}
// 		return nil
// 	})

// 	return result, err
// }