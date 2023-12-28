package updates

const (
	LongpollResponseFailed  = "failed"
	LongpollResponseUpdates = "updates"
	LongpollResponseTS      = "ts"
)

const (
	LongpollUnknownStatus       = 0
	LongpollHistoryMissedStatus = 1
	LongpollKeyExpiredStatus    = 2
	LongpollInfoMissedStatus    = 3
)
