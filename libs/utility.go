package libs

import (
	"fmt"
	"log"

	"github.com/emj365/xschange/models"
)

func LogOrders(orders *[]*models.Order) {
	fmt.Println("\033[2J")
	log.Printf("orders: %#v\n\n", orders)
	for i, o := range *orders {
		log.Printf("orders[%v]: %#v\n", i, *o)
		for j, p := range (*o).Matchs {
			log.Printf("orders[%v].Matchs[%v]: %#v\n", i, j, *p)
		}
	}

}
