package config

import "github.com/google/uuid"

// Pre-seeded membership plans for the exercise.
// Load these into your store at startup — candidates do not need a "create plan" endpoint.

var (
	MonthlyUnlimitedID = uuid.MustParse("11111111-1111-1111-1111-111111111111")
	TenClassPackID     = uuid.MustParse("22222222-2222-2222-2222-222222222222")
	LaunchPromoPackID  = uuid.MustParse("33333333-3333-3333-3333-333333333333")
	LegacyPlanID       = uuid.MustParse("44444444-4444-4444-4444-444444444444")
)

// PlanSeed describes a seeded membership plan.
// RemainingSlots is only set for the Launch Promo Pack; nil means unlimited inventory.
// Credits is class credits for pack plans; nil for unlimited plans.
type PlanSeed struct {
	ID             uuid.UUID
	Name           string
	Type           string
	Active         bool
	PriceCents     int
	Credits        *int
	RemainingSlots *int
}

func intPtr(v int) *int { return &v }

// Plans are the pre-seeded membership plans for the exercise.
var Plans = []PlanSeed{
	{
		ID:             MonthlyUnlimitedID,
		Name:           "Monthly Unlimited",
		Type:           "unlimited",
		Active:         true,
		PriceCents:     8000,
		Credits:        nil,
		RemainingSlots: nil,
	},
	{
		ID:             TenClassPackID,
		Name:           "10-Class Pack",
		Type:           "pack",
		Active:         true,
		PriceCents:     5000,
		Credits:        intPtr(10),
		RemainingSlots: nil,
	},
	{
		ID:             LaunchPromoPackID,
		Name:           "Launch Promo Pack",
		Type:           "pack",
		Active:         true,
		PriceCents:     2500,
		Credits:        intPtr(5),
		RemainingSlots: intPtr(5),
	},
	{
		ID:             LegacyPlanID,
		Name:           "Legacy Plan",
		Type:           "unlimited",
		Active:         false,
		PriceCents:     6000,
		Credits:        nil,
		RemainingSlots: nil,
	},
}
