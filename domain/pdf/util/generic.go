package util

import (
	"fmt"
	"github.com/asentientbanana/pausalkole-admin/models"
	"strconv"
)

//func splitRowsPerPage(items []models.InvoiceItem) [][]models.InvoiceItem {
//	itemsPerPage := 15
//
//	return [][]models.InvoiceItem{}
//}

func CalculatePdfEntityInfoHeight(r *models.Entity, a *models.Entity) (float64, int) {
	lineHeight := 8.0
	rLen := len(r.Fields)
	aLen := len(a.Fields)
	if rLen > aLen {
		return lineHeight * float64(rLen), rLen
	} else {
		return lineHeight * float64(aLen), aLen
	}
}

func GetPdfRowsPerPage(invoice models.Invoice, limit int, firstPageRowNumber int) [][]int {

	items := [][]int{[]int{}}

	fmt.Println("LIMIT::" + strconv.Itoa(limit))

	itemLen := len(invoice.Items)

	activeArray := 0

	for i := 0; i < itemLen; i++ {
		if activeArray == 0 {
			if i < firstPageRowNumber {
				items[activeArray] = append(items[activeArray], i)
				continue
			}
			activeArray++
			items = append(items, []int{})
		}

		items[activeArray] = append(items[activeArray], i)
		if len(items[activeArray]) == limit {
			items = append(items, []int{})
			activeArray++
		}
	}

	return items
}
