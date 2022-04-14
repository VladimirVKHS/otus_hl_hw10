package counter

type Counter struct {
	UserId              int
	UnreadMessagesCount int
}

type CounterUpdateRequest struct {
	UnreadMessagesCountDelta int
}

func (c *Counter) ToResponse() map[string]interface{} {
	return map[string]interface{}{
		"UserId":              c.UserId,
		"UnreadMessagesCount": c.UnreadMessagesCount,
	}
}
