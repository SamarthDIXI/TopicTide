package communication_protocol

type Message struct {
    Topic   string `json:"Topic"`
    Content string `json:"Content"`
}
