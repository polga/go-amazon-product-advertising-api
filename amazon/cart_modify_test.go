package amazon

import "testing"

func TestCartModify(t *testing.T) {
	client, _ := New("AK", "SK", "JP")
	req := client.CartModify()
	if req == nil {
		t.Error("Expected not nil but got nil")
	}
	Test{client, req.Client}.Compare(t)
}

func TestCartModifyBuildQuery(t *testing.T) {
	client, _ := New("AK", "SK", "JP")
	q := client.CartModify().buildQuery()
	Test{0, len(q)}.Compare(t)
}