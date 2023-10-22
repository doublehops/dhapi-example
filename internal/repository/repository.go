package repository

import (
	"fmt"
	"github.com/doublehops/dhapi-example/internal/handlers/pagination"
	"log"
	"strings"
)

func SubstitutePaginationVars(q string, p *pagination.RequestPagination) string {
	wc := fmt.Sprintf(" LIMIT %d, %d", p.Offset, p.PerPage)

	log.Println(">>>>>>>" + wc + "\n")
	return strings.Replace(q, "__PAGINATION__", wc, -1)
}
