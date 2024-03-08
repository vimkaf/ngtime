package ngtime

import "time"

func getTimeInNigeria() (time.Time, error) {
	// Set the time zone to WAT (West Africa Time - Nigeria Standard Time)
	nigeriaLocation, err := time.LoadLocation("Africa/Lagos")
	if err != nil {
		return time.Time{}, err
	}
	// Get the current time in Nigeria
	nigeriaTime := time.Now().In(nigeriaLocation)
	return nigeriaTime, nil
}
