package model

import "time"

type Plan struct {
	ID           	string
	Name         	string
	DailyLimit   	int
	WeeklyLimit  	int
	MonthlyLimit 	int
	PriceUSD     	float64
}

type User struct {
	ID     				string
	Email  				string
	APIKey 				string
	PlanID 				string
}

type UsageLog struct {
	ID        		string
	UserID    		string
	Timestamp 		time.Time
}
