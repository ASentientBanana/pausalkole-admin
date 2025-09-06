package util

import (
	"encoding/json"
	"fmt"
	"github.com/asentientbanana/pausalkole-admin/models"
	"os"
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

	itemLen := len(invoice.Items)

	activeArray := 0

	for i := 0; i < itemLen; i++ {
		if activeArray == 0 {
			if i < firstPageRowNumber {
				items[activeArray] = append(items[activeArray], i)
				continue
			} else {
				activeArray++
				items = append(items, []int{})
			}
		}

		items[activeArray] = append(items[activeArray], i)
		if len(items[activeArray]) >= limit {
			items = append(items, []int{})
			activeArray++
		}
	}
	jsonMap := make(map[string]interface{})
	for i := 0; i < len(items); i++ {
		fmt.Println(len(items[i]))
		jsonMap["page"+strconv.Itoa(i+1)] = items[i]
	}

	jsonFIle, err := os.Create("items.json")
	defer jsonFIle.Close()

	if err == nil {
		data, er := json.MarshalIndent(jsonMap, "", " ")
		if er == nil {
			jsonFIle.Write(data)
		}
	}

	return items
}
