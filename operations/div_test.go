package operations

import "testing"

func TestDiv(t *testing.T) {
	got, err := Div(10, 2)
	want := 5
	if err != nil {
		t.Errorf("予期しないエラーが発生しました: %v", err)
	}
	if got != want {
		t.Errorf("Div(10,2) = %d; want %d", got, want)
	}

	_, err = Div(10, 0)
	if err == nil {
		t.Errorf("ゼロ除算エラーを期待しましたが、発生しませんでした")
	}
}
