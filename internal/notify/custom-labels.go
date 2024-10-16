package notify

func markPriority(priorityID int) string {
	switch priorityID {
	case 5:
		return "\U00002B24 " // Эмодзи: 🔶
	case 4:
		return "\U0001F534 " // Эмодзи: 🔴
	case 3:
		return "\U0001F7E1 " // Эмодзи: 🟡
	case 2:
		return "\U0001F7E2 " // Эмодзи: 🟢
	default:
		return "?"
	}
}

func markTracker(trackerID int) string {
	switch trackerID {
	case 4:
		return "\U0001F4B0 " // Эмодзи: 💰
	default:
		return ""
	}
}

func oldPriorString(priorityID int) string {
	switch priorityID {
	case 5:
		return "Нулевого"
	case 4:
		return "Первого"
	case 3:
		return "Второго"
	case 2:
		return "Третьего"
	default:
		return "?"
	}
}

func newPriorString(priorityID int) string {
	switch priorityID {
	case 5:
		return "Нулевой"
	case 4:
		return "Первый"
	case 3:
		return "Второй"
	case 2:
		return "Третий"
	default:
		return "?"
	}
}
