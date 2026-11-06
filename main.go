package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/glofoxinc/glofox-go-takehome-task/config"
)

func main() {
	logSeededPlans()

	mux := http.NewServeMux()
	// Candidates wire their HTTP handlers here.

	addr := ":8080"
	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}

func logSeededPlans() {
	log.Println("Seeded membership plans (use these ids in your API):")
	for _, plan := range config.Plans {
		log.Printf(
			"  id=%s name=%q type=%s active=%t price_cents=%d credits=%s remaining_slots=%s",
			plan.ID,
			plan.Name,
			plan.Type,
			plan.Active,
			plan.PriceCents,
			formatOptionalInt(plan.Credits),
			formatOptionalInt(plan.RemainingSlots),
		)
	}
}

func formatOptionalInt(v *int) string {
	if v == nil {
		return "null"
	}
	return fmt.Sprintf("%d", *v)
}
