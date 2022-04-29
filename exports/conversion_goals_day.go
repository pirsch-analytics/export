package exports

import (
	"fmt"
	"github.com/pirsch-analytics/pirsch-go-sdk"
	"log"
	"os"
	"strings"
	"time"
)

// ExportConversionGoalsDays exports all conversion goal unique visitor statistics grouped by day.
func ExportConversionGoalsDays(client *pirsch.Client, from, to time.Time) error {
	log.Println("Exporting conversion goals unique visitors grouped by day")
	domain, err := client.Domain()

	if err != nil {
		return err
	}

	goals, err := client.ConversionGoals(&pirsch.Filter{
		DomainID: domain.ID,
	})

	if err != nil {
		return err
	}

	var out strings.Builder
	out.WriteString("Conversion Goal,")
	period := to.Sub(from)
	days := int(period.Hours()/24) + 1

	for i := 0; i < days; i++ {
		out.WriteString(from.Add(time.Duration(i) * time.Hour * 24).Format("2006-01-02"))

		if i != days-1 {
			out.WriteString(",")
		}
	}

	out.WriteString("\n")

	for _, goal := range goals {
		if err := exportConversionGoal(&out, client, domain.ID, goal.PageGoal.Name, goal.PageGoal.PathPattern, from, to); err != nil {
			return err
		}
	}

	return os.WriteFile("export/conversion_goals_day.csv", []byte(out.String()), 0644)
}

func exportConversionGoal(out *strings.Builder, client *pirsch.Client, domainID, name, pattern string, from, to time.Time) error {
	visitors, err := client.Visitors(&pirsch.Filter{
		DomainID: domainID,
		From:     from,
		To:       to,
		Pattern:  pattern,
		Scale:    pirsch.ScaleDay,
	})

	if err != nil {
		return err
	}

	out.WriteString(fmt.Sprintf(`"%s",`, name))
	n := len(visitors)

	for i, stats := range visitors {
		out.WriteString(fmt.Sprintf("%d", stats.Visitors))

		if i < n-1 {
			out.WriteString(",")
		}
	}

	out.WriteString("\n")
	return nil
}
