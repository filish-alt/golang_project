package db

import (
	"context"
	

)
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