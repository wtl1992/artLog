package artlog

import (
	"artlog/artlog/colors"
	"artlog/artlog/common"
	"artlog/artlog/constants"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

type logger struct {
	Writer io.Writer `json:"writer"`
}

func NewLogger(writer io.Writer) *logger {
	if writer == nil {
		writer = os.Stderr
	}

	_logger := logger{Writer: writer}
	return &_logger
}

var (
	once          = sync.Once{}
	DefaultLogger *logger
)

const (
	ShowContentWithFormatAndColorsWithAttrStr = "__ATTR____BACKGROUND_COLOR____FONT_COLOR__CONTENT\033[0m"
)

func init() {
	once.Do(func() {
		DefaultLogger = NewLogger(nil)
	})
}

func SetOutput(writer io.Writer) {
	if writer == nil {
		writer = os.Stderr
	}
	DefaultLogger.Writer = writer
}

func infoLogFormat(format string, msg interface{}) string {
	formatStr := fmt.Sprintf(format, msg)
	return formatStr
}

func defaultSettings(_attrs *colors.Attrs) {
	var options = []common.Option{colors.Color_Font_White, colors.Attr_HideCursor}
	for _, option := range options {
		option.SetOptionValue(_attrs)
	}
}

func InfoLog(format string, msg interface{}, options ...common.Option) {
	content := infoLogFormat(format, msg)
	var showContentFormat = ShowContentWithFormatAndColorsWithAttrStr
	var attrs = new(colors.Attrs)

	if options != nil && len(options) > 0 {
		for _, option := range options {
			option.SetOptionValue(attrs)
		}
	}

	if attrs.Attr != nil {
		showContentFormat = strings.ReplaceAll(showContentFormat, "__ATTR__", "\033["+attrs.Attr.Attr)
	} else {
		showContentFormat = strings.ReplaceAll(showContentFormat, "__ATTR__", constants.EMPTY_STRING)
	}

	if attrs.BackgroundColor != nil {
		showContentFormat = strings.ReplaceAll(showContentFormat, "__BACKGROUND_COLOR__", "\033["+attrs.BackgroundColor.Color)
	} else {
		showContentFormat = strings.ReplaceAll(showContentFormat, "__BACKGROUND_COLOR__", constants.EMPTY_STRING)
	}

	if attrs.FontColor != nil {
		showContentFormat = strings.ReplaceAll(showContentFormat, "__FONT_COLOR__", "\033["+attrs.FontColor.Color)
	} else {
		showContentFormat = strings.ReplaceAll(showContentFormat, "__FONT_COLOR__", constants.EMPTY_STRING)
	}

	showContentFormat = strings.ReplaceAll(showContentFormat, "CONTENT", content)

	_, err := DefaultLogger.Writer.Write([]byte(showContentFormat))
	if err != nil {
		panic(err)
	}
}
