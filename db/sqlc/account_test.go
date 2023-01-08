package db

import (
	"Documents/project/gotest/go3/util"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)      // untuk cek error seharusnya tidak ada error
	require.NotEmpty(t, account) // untuk cek return nilai, seharusnya ada return nilai

	require.Equal(t, arg.Owner, account.Owner)       // untuk compare nilai input dan nilai yang sudah masuk
	require.Equal(t, arg.Balance, account.Balance)   // untuk compare nilai input dan	 nilai	 yang	 sudah	 masuk
	require.Equal(t, arg.Currency, account.Currency) // untuk compare nilai input dan nilai yang sudah masuk

	require.NotZero(t, account.ID)        // untuk cek tidak ada yang zero
	require.NotZero(t, account.CreatedAt) // untuk cek tidak ada yang zero

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomMoney(),
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)

	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)                             // untuk cek jika hasilnya ada error yaitu error data kosong
	require.EqualError(t, err, sql.ErrNoRows.Error()) // untuk cek jika error itu merupakan error tanpa result data
	require.Empty(t, account2)                        // untuk cek bahwa account2 itu kosong
}
