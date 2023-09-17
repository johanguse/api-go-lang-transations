package enums

type Status string //@name Status

const (
	Processing Status = "processing"
	Processed  Status = "processed"
	Created    Status = "created"
)
