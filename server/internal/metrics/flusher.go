package metrics

import (
	"time"

	"SeeAll/internal/database"
)

func StartFlusher() {

	go func() {

		ticker := time.NewTicker(10 * time.Second)

		for range ticker.C {

			total, endpoints, visitors := Snapshot()
			if total == 0 && len(endpoints) == 0 && len(visitors) == 0 {
				continue
			}
			tx, err := database.DB.Begin()
			if err != nil {
				continue
			}

			// endpoint visits
			for ep, count := range endpoints {

				if count == 0 {
					continue
				}

				tx.Exec(`
				INSERT INTO endpoint_visits(endpoint, visits)
				VALUES(?, ?)
				ON CONFLICT(endpoint)
				DO UPDATE SET visits = visits + excluded.visits
				`, ep, count)
			}

			// total visits
			if total > 0 {

				tx.Exec(`
				UPDATE totals
				SET total_visits = total_visits + ?
				WHERE id = 1
				`, total)
			}

			// visitors
			for _, v := range visitors {

				tx.Exec(`
				INSERT OR IGNORE INTO visitors_daily(day, ip_hash)
				VALUES(date('now'), ?)
				`, v)
			}

			tx.Commit()
		}
	}()
}
