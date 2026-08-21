package app

import "testing"

// imageFullResURL 逻辑自检：剥离已有选项、补 =s0、保持 PNG。无网络依赖。
func TestImageFullResURL(t *testing.T) {
	cases := []struct{ in, want string }{
		// 上游 issue #14: 带 =s1024-rj 选项 → 去掉后补 =s0
		{"https://lh3.example.com/gg-dl/abc123=w640-h640-s1024-rj", "https://lh3.example.com/gg-dl/abc123=s0"},
		// 无选项 → 直接补 =s0
		{"https://lh3.example.com/gg-dl/xyz", "https://lh3.example.com/gg-dl/xyz=s0"},
		// token 本身 base64url 不含 '='，只有选项段才可能有 =
		{"https://lh3.example.com/gg-dl/token-with-dash-sDmk", "https://lh3.example.com/gg-dl/token-with-dash-sDmk=s0"},
	}
	for _, c := range cases {
		if got := imageFullResURL(c.in); got != c.want {
			t.Errorf("imageFullResURL(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}