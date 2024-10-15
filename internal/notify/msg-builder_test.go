package notify

import (
	"testing"
)

func TestAddStatusTxt(t *testing.T) {
	testTable := []struct {
		oldStatusName string
		newStatusName string
		expected      string
	}{
		{
			oldStatusName: "В работе",
			newStatusName: "Решена",
			expected:      "\\\n\\-изменился статус c *В работе* на *Решена*",
		},
	}

	for _, tstCase := range testTable {
		result, err := AddStatusTxt(tstCase.oldStatusName, tstCase.newStatusName)

		//t.Logf("Calling AddStatusTxt(%s, %s), result: %s", tstCase.oldStatusName, tstCase.newStatusName, result)

		if err != nil {
			t.Errorf("Should not produce an error %s", err)
		}

		if result != tstCase.expected {
			t.Errorf("Incorrect result, want: %s, got: %s", tstCase.expected, result)
		}
	}
}

func TestAddPriorityTxt(t *testing.T) {
	testTable := []struct {
		oldPriorityID int
		newPriorityID int
		expected      string
	}{
		{
			oldPriorityID: 4,
			newPriorityID: 5,
			expected:      "\\\n\\-изменился приоритет c *Первого* на *Нулевой*",
		},
		{
			oldPriorityID: 6,
			newPriorityID: 5,
			expected:      "\\\n\\-изменился приоритет c *?* на *Нулевой*",
		},
		{
			oldPriorityID: 4,
			newPriorityID: 1,
			expected:      "\\\n\\-изменился приоритет c *Первого* на *?*",
		},
	}

	for _, tstCase := range testTable {
		result, err := AddPriorityTxt(tstCase.oldPriorityID, tstCase.newPriorityID)

		//t.Logf("Calling AddStatusTxt(%d, %d), result: %s", tstCase.oldPriorityID, tstCase.newPriorityID, result)

		if err != nil {
			t.Errorf("Should not produce an error %s", err)
		}

		if result != tstCase.expected {
			t.Errorf("Incorrect result, want: %s, got: %s", tstCase.expected, result)
		}
	}
}

func TestAddAssignedTxt(t *testing.T) {
	testTable := []struct {
		oldAssignedToName string
		newAssignedToName string
		expected          string
	}{
		{
			oldAssignedToName: "Старый исполнитель",
			newAssignedToName: "Новый исполнитель",
			expected:          "\\\n\\-изменился исполнитель c *Старый исполнитель* на *Новый исполнитель*",
		},
		{
			oldAssignedToName: "",
			newAssignedToName: "Новый исполнитель",
			expected:          "\\\n\\-изменился исполнитель c ** на *Новый исполнитель*",
		},
		{
			oldAssignedToName: "Старый исполнитель",
			newAssignedToName: "",
			expected:          "\\\n\\-изменился исполнитель c *Старый исполнитель* на **",
		},
		{
			oldAssignedToName: "",
			newAssignedToName: "",
			expected:          "\\\n\\-изменился исполнитель c ** на **",
		},
	}

	for _, tstCase := range testTable {
		result, err := AddAssignedTxt(tstCase.oldAssignedToName, tstCase.newAssignedToName)

		//t.Logf("Calling AddStatusTxt(%s, %s), result: %s", tstCase.oldAssignedToName, tstCase.newAssignedToName, result)

		if err != nil {
			t.Errorf("Should not produce an error %s", err)
		}

		if result != tstCase.expected {
			t.Errorf("Incorrect result, want: %s, got: %s", tstCase.expected, result)
		}
	}
}

func TestAddNewCommentTxt(t *testing.T) {
	testTable := []struct {
		str      string
		expected string
	}{
		{
			str:      "Тест",
			expected: "\\\n\\-был добавлен комментарий: *\\\"Тест\\\"*",
		},
		{
			str:      "",
			expected: "\\\n\\-был добавлен комментарий: *\\\"\\\"*",
		},
	}

	for _, tstCase := range testTable {
		result, err := AddNewCommentTxt(tstCase.str)

		//t.Logf("Calling AddStatusTxt(%s), result: %s", tstCase.str, result)

		if err != nil {
			t.Errorf("Should not produce an error %s", err)
		}

		if result != tstCase.expected {
			t.Errorf("Incorrect result, want: %s, got: %s", tstCase.expected, result)
		}
	}
}

func TestCreateMsg(t *testing.T) {
	testTable := []struct {
		issueID      int
		priorityID   int
		trackerID    int
		title        string
		text         string
		assignToName string
		expected     string
	}{
		{
			issueID:      0,
			priorityID:   3,
			trackerID:    4,
			title:        "Тестовая задача",
			text:         "\\\n\\-изменился приоритет c *Первого* на *Нулевой*\\\n\\-был добавлен комментарий: *\"Тест\"*",
			assignToName: "Тестов Тест",
			expected:     "\U0001F4B0 \U0001F7E1 В задаче [0](/issues/0) \\- Тестовая задача\\\n\\-изменился приоритет c *Первого* на *Нулевой*\\\n\\-был добавлен комментарий: *\"Тест\"*\\\nИсполнитель *Тестов Тест*",
		},
		{
			issueID:      0,
			priorityID:   3,
			trackerID:    4,
			title:        "*Тестовая_задача[]()~><#+-=|.!",
			text:         "\\\n\\-изменился приоритет c *Первого* на *Нулевой*\\\n\\-был добавлен комментарий: *\"Тест\"*",
			assignToName: "Тестов Тест",
			expected:     "\U0001F4B0 \U0001F7E1 В задаче [0](/issues/0) \\- \\*Тестовая\\_задача\\[\\]\\(\\)\\~\\>\\<\\#\\+\\-\\=\\|\\.\\!\\\n\\-изменился приоритет c *Первого* на *Нулевой*\\\n\\-был добавлен комментарий: *\"Тест\"*\\\nИсполнитель *Тестов Тест*",
		},
	}

	for _, tstCase := range testTable {
		result, err := CreateMsg(tstCase.issueID, tstCase.priorityID, tstCase.trackerID, tstCase.title, tstCase.text, tstCase.assignToName)

		//t.Logf("Calling AddStatusTxt(%d, %d, %d, %s, %s, %s), result: %s", tstCase.issueID, tstCase.priorityID, tstCase.trackerID, tstCase.title, tstCase.text, tstCase.assignToName, result)

		if err != nil {
			t.Errorf("Should not produce an error %s", err)
		}

		if result != tstCase.expected {
			t.Errorf("Incorrect result, want: %s, got: %s", tstCase.expected, result)
		}
	}
}
