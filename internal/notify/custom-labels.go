package notify

func oldPriorString(pId int) string {
	switch pId {
	case 5:
		return "нулевого"
	case 4:
		return "первого"
	case 3:
		return "второго"
	case 2:
		return "третьего"
	default:
		return "?"
	}
}

func newPriorString(pId int) string {
	switch pId {
	case 5:
		return "нулевой"
	case 4:
		return "первый"
	case 3:
		return "второй"
	case 2:
		return "третий"
	default:
		return "?"
	}
}

func markPriority(pId int) string {
	switch pId {
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

func markTracker(tId int) string {
	switch tId {
	case 4:
		return "\U0001F4B0 " // Эмодзи: 💰
	default:
		return ""
	}
}
