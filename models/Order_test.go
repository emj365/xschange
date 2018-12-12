package models

import (
	"testing"
)

func TestChoosePeers(t *testing.T) {
	peers := []*Order{
		&Order{Selling: true, Quantity: 1, Remain: 1},
		&Order{Selling: true, Quantity: 1, Remain: 1},
		&Order{Selling: false, Quantity: 1, Remain: 1},
		&Order{Selling: false, Quantity: 1, Remain: 1},
		&Order{Selling: false, Quantity: 1, Remain: 1},
		&Order{Selling: false, Quantity: 1, Remain: 0},
	}

	result := choosePeers(&peers, false)
	if want := 3; len(result) != want {
		t.Errorf("len(result) == %v, want %v", len(result), want)
	}
}

func TestFilterByPrice(t *testing.T) {
	peers := []*Order{
		&Order{Price: 1},
		&Order{Price: 2},
		&Order{Price: 3},
		&Order{Price: 4},
		&Order{Price: 5},
		&Order{Price: 6},
	}

	result := filterByPrice(&peers, true, 4) // for buyer
	if want := 4; len(result) != want {
		t.Errorf("len(result ) == %v, want %v", len(result), want)
	}
	if want := 1; result[0].Price != want {
		t.Errorf("result [0].Price == %v, want %v", result[0].Price, want)
	}
	if want := 2; result[1].Price != want {
		t.Errorf("result [1].Price == %v, want %v", result[1].Price, want)
	}
	if want := 3; result[2].Price != want {
		t.Errorf("result [2].Price == %v, want %v", result[2].Price, want)
	}
	if want := 4; result[3].Price != want {
		t.Errorf("result [3].Price == %v, want %v", result[3].Price, want)
	}

	result = filterByPrice(&peers, false, 4) // for seller
	if want := 3; len(result) != want {
		t.Errorf("len(result ) == %v, want %v", len(result), want)
	}
	if want := 4; result[0].Price != want {
		t.Errorf("result [0].Price == %v, want %v", result[0].Price, want)
	}
	if want := 5; result[1].Price != want {
		t.Errorf("result [1].Price == %v, want %v", result[1].Price, want)
	}
	if want := 6; result[2].Price != want {
		t.Errorf("result [2].Price == %v, want %v", result[2].Price, want)
	}
}

func TestSortPeers(t *testing.T) {
	peers := []*Order{
		&Order{Price: 10},
		&Order{Price: 11},
		&Order{Price: 12},
	}

	sortPeers(&peers, false)
	if want := 12; peers[0].Price != want {
		t.Errorf("peers[0].Price == %v, want %v", peers[0].Price, want)
	}
	if want := 11; peers[1].Price != want {
		t.Errorf("peers[0].Price == %v, want %v", peers[0].Price, want)
	}
	if want := 10; peers[2].Price != want {
		t.Errorf("peers[0].Price == %v, want %v", peers[0].Price, want)
	}

	sortPeers(&peers, true)

	if want := 10; peers[0].Price != want {
		t.Errorf("peers[0].Price == %v, want %v", peers[0].Price, want)
	}
	if want := 11; peers[1].Price != want {
		t.Errorf("peers[0].Price == %v, want %v", peers[0].Price, want)
	}
	if want := 12; peers[2].Price != want {
		t.Errorf("peers[0].Price == %v, want %v", peers[0].Price, want)
	}
}
