package flow052

import (
	"forkliftarchive/internal/domain"
	"forkliftarchive/internal/service"
	"forkliftarchive/internal/store"
	"path/filepath"
	"testing"
)

func openStore(t *testing.T) (*store.Store, *service.Service) {
	t.Helper()
	st, err := store.Open(filepath.Join(t.TempDir(), "t.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { st.Close() })
	return st, service.New(st, nil)
}

// TestDetailShowsValidTitleAfterSecondBatch reproduces the stale-title bug:
// after importing a second batch, each record's detail must show its own
// valid (stored) title rather than the first title of the latest batch.
func TestDetailShowsValidTitleAfterSecondBatch(t *testing.T) {
	_, svc := openStore(t)
	h := New(svc)

	// 第一批数据
	if err := h.ImportBatch([]domain.ImportRow{
		{Code: "FK-001", Title: "原标题", Location: "A", Capacity: 1000},
		{Code: "FK-002", Title: "第二批前标题", Location: "B", Capacity: 2000},
	}); err != nil {
		t.Fatalf("import batch1: %v", err)
	}

	// 第二批数据：导入后，旧 titles 切片会被重置（这是原 bug 的触发条件）
	if err := h.ImportBatch([]domain.ImportRow{
		{Code: "FK-003", Title: "第二批标题一", Location: "C", Capacity: 500},
		{Code: "FK-004", Title: "第二批标题二", Location: "D", Capacity: 3000},
	}); err != nil {
		t.Fatalf("import batch2: %v", err)
	}

	// FK-001 的详情应只显示它自己的有效标题"原标题"，
	// 而不是过期的"第二批标题一"（第二批首个标题）。
	got, err := h.Detail("batch-FK-001")
	if err != nil {
		t.Fatalf("Detail FK-001: %v", err)
	}
	if got.Title != "原标题" {
		t.Fatalf("FK-001 标题对应关系错误: got %q, want %q", got.Title, "原标题")
	}
	// 第二批各记录的标题对应关系也应正确。
	got3, _ := h.Detail("batch-FK-003")
	if got3.Title != "第二批标题一" {
		t.Fatalf("FK-003 标题对应关系错误: got %q, want %q", got3.Title, "第二批标题一")
	}
	got4, _ := h.Detail("batch-FK-004")
	if got4.Title != "第二批标题二" {
		t.Fatalf("FK-004 标题对应关系错误: got %q, want %q", got4.Title, "第二批标题二")
	}
}
