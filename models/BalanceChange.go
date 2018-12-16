package models

type BalanceChange struct {
	Good              int
	Balance           int
	User              *User
	Match             *Match // be provided if it is trade
	Order             *Order // be provided if it is margin
	IsDepositWithdrew bool
}

type BalanceChanges []*BalanceChange

func (balanceChanges *BalanceChanges) Add(balanceChange *BalanceChange) {
	*balanceChanges = append(*balanceChanges, balanceChange)

	// TODO: use for look to calculate all changes for the user's balance cache
	balanceChange.User.GoodAmount = uint(int(balanceChange.User.GoodAmount) + balanceChange.Good)
	balanceChange.User.Balance = uint(int(balanceChange.User.Balance) + balanceChange.Balance)
}
