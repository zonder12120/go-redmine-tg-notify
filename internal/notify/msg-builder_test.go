package notify

import "testing"

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

		t.Logf("Calling AddStatusTxt(%s, %s), result: %s", tstCase.oldStatusName, tstCase.newStatusName, result)

		if err != nil {
			t.Errorf("Should not produce an error %v", err)
		}

		if result != tstCase.expected {
			t.Errorf("Incorrect result, want: %s, got: %s.", tstCase.expected, result)
		}
	}

}

func TestAddPriorityTxt(t *testing.T) {

}
