package enum

type Task int64

const (
	TaskSingleFileUpload Task = 1
)

func (task Task) ToInt64() int64 {
	switch task {
	case TaskSingleFileUpload:
		return 1
	default:
		return 0
	}
}
