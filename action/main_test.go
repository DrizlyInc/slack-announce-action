package main

import (
	"testing"
)

func Test_GetColorForStatus(t *testing.T) {
	color := GetColorForStatus("success")
	assertEquals(t, "#4caf50", color)

	color = GetColorForStatus("failure")
	assertEquals(t, "#f44336", color)

	color = GetColorForStatus("cancelled")
	assertEquals(t, "#808080", color)

	color = GetColorForStatus("skipped")
	assertEquals(t, "#808080", color)
}

func Test_GetTitleSuffixForStatus(t *testing.T) {
	suffix := GetTitleSuffixForStatus("success")
	assertEquals(t, "completed successfully!", suffix)

	suffix = GetTitleSuffixForStatus("failure")
	assertEquals(t, "failed!", suffix)

	suffix = GetTitleSuffixForStatus("cancelled")
	assertEquals(t, "was cancelled.", suffix)

	suffix = GetTitleSuffixForStatus("skipped")
	assertEquals(t, "was skipped.", suffix)
}

func Test_GetEmojiForStatus(t *testing.T) {
	emoji := GetEmojiForStatus("success")
	assertEquals(t, ":white_check_mark:", emoji)

	emoji = GetEmojiForStatus("failure")
	assertEquals(t, ":x:", emoji)

	emoji = GetEmojiForStatus("cancelled")
	assertEquals(t, ":grey_exclamation:", emoji)

	emoji = GetEmojiForStatus("skipped")
	assertEquals(t, ":heavy_minus_sign:", emoji)
}
