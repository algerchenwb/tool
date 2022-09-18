package random

import "fmt"

type TemplateString string

const (
	TemplateString_0_9         TemplateString = `1234567890`
	TemplateString_a_z         TemplateString = `abcdefghijklmnopqrstuvwxyz`
	TemplateString_A_Z         TemplateString = `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
	TemplateString_0_9_a_f     TemplateString = `1234567890abcdef`
	TemplateString_0_9_a_z     TemplateString = `1234567890abcdefghijklmnopqrstuvwxyz`
	TemplateString_0_9_A_F     TemplateString = `1234567890ABCDEF`
	TemplateString_0_9_A_Z     TemplateString = `1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ`
	TemplateString_a_z_A_Z     TemplateString = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`
	TemplateString_0_9_a_z_A_Z TemplateString = `1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`
)

// String 生成随机的字符串
func String(t TemplateString, length int) string {
	if length <= 0 {
		return ""
	}
	var rt string
	var tLen = len(t)
	for i := 0; i < length; i++ {
		rt += fmt.Sprintf("%c", t[Rand.Intn(tLen)])
	}
	return rt
}
