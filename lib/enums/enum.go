package enums

type Status string

const (
	Processing Status = "processing"
	Processed  Status = "processed"
	Created    Status = "created"
)
