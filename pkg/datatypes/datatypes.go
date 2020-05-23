package datatypes

// Schedule consists whole week schedule
type Schedule map[string]Day

// Day consists time blocks of day of the week
type Day []*Block

// Block consists fields of time block
type Block struct {
	StartHour, StartMinute,
	FinishHour, FinishMinute,
	Description string
	NumStartTime, NumFinishTime float32
}
