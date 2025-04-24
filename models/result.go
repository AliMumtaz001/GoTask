package model

type Result struct {
    ID     int    `json:"id"`
    UserID string `json:"user_id"`
    Data   string `json:"data"`
}