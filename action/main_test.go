package main

import (
	"testing"
)

func Test_GetColorAndTitleSuffix(t *testing.T) {
	color, suffix := GetColorAndTitleSuffix("success")
	assertEquals(t, "#4caf50", color)
	assertEquals(t, "completed successfully!", suffix)

	color, suffix = GetColorAndTitleSuffix("failure")
	assertEquals(t, "#f44336", color)
	assertEquals(t, "failed!", suffix)

	color, suffix = GetColorAndTitleSuffix("cancelled")
	assertEquals(t, "#808080", color)
	assertEquals(t, "was cancelled.", suffix)

	color, suffix = GetColorAndTitleSuffix("skipped")
	assertEquals(t, "#808080", color)
	assertEquals(t, "was skipped.", suffix)
}

func Test_GetStatusEmoji(t *testing.T) {
	emoji := GetStatusEmoji("success")
	assertEquals(t, ":white_check_mark:", emoji)

	emoji = GetStatusEmoji("failure")
	assertEquals(t, ":x:", emoji)

	emoji = GetStatusEmoji("cancelled")
	assertEquals(t, ":grey_exclamation:", emoji)

	emoji = GetStatusEmoji("skipped")
	assertEquals(t, ":heavy_minus_sign:", emoji)
}
