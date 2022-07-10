package colors

type ColorEnum struct {
	Color string `json:"color"`
	Desc  string `json:"desc"`
}

type BackgroundColor ColorEnum
type FontColor ColorEnum

type Attrs struct {
	BackgroundColor *BackgroundColor `json:"backgroundColor"`
	FontColor       *FontColor       `json:"fontColor"`
	Attr            *AttrEnum        `json:"attr"`
}

func (color *BackgroundColor) SetOptionValue(attrs *Attrs) {
	attrs.BackgroundColor = color
}

func (color *FontColor) SetOptionValue(attrs *Attrs) {
	attrs.FontColor = color
}

var (
	// 前景色
	Color_Font_Black   = &FontColor{Color: "30m", Desc: "字体黑色"}
	Color_Font_Red     = &FontColor{Color: "31m", Desc: "字体红色"}
	Color_Font_Green   = &FontColor{Color: "32m", Desc: "字体绿色"}
	Color_Font_Yellow  = &FontColor{Color: "33m", Desc: "字体黄色"}
	Color_Font_Blue    = &FontColor{Color: "34m", Desc: "字体蓝色"}
	Color_Font_Purple  = &FontColor{Color: "35m", Desc: "字体紫色"}
	Color_Font_Skyblue = &FontColor{Color: "36m", Desc: "字体天蓝色"}
	Color_Font_White   = &FontColor{Color: "37m", Desc: "字体白色"}

	//背景色
	Color_Background_Black   = &BackgroundColor{Color: "40m", Desc: "背景黑色"}
	Color_Background_Red     = &BackgroundColor{Color: "41m", Desc: "背景红色"}
	Color_Background_Green   = &BackgroundColor{Color: "42m", Desc: "背景绿色"}
	Color_Background_Yellow  = &BackgroundColor{Color: "43m", Desc: "背景黄色"}
	Color_Background_Blue    = &BackgroundColor{Color: "44m", Desc: "背景蓝色"}
	Color_Background_Purple  = &BackgroundColor{Color: "45m", Desc: "背景紫色"}
	Color_Background_Skyblue = &BackgroundColor{Color: "46m", Desc: "背景天蓝色"}
	Color_Background_White   = &BackgroundColor{Color: "47m", Desc: "背景白色"}
)

type AttrEnum struct {
	Attr string `json:"attr"`
	Num  uint8  `json:"num"`
	Desc string `json:"desc"`
}

func (attr *AttrEnum) SetOptionValue(attrs *Attrs) {
	attrs.Attr = attr
}

var (
	Attr_CloseAllAttrs                    = &AttrEnum{Attr: "0m", Desc: "关闭所有属性"}
	Attr_SetHighBrightness                = &AttrEnum{Attr: "1m", Desc: "设置高亮度"}
	Attr_Underline                        = &AttrEnum{Attr: "4m", Desc: "下划线"}
	Attr_Twinkle                          = &AttrEnum{Attr: "5m", Desc: "闪烁"}
	Attr_ReverseDisplay                   = &AttrEnum{Attr: "7m", Desc: "反显"}
	Attr_Blanking                         = &AttrEnum{Attr: "8m", Desc: "消隐"}
	Attr_MoveTheCursorUpNLines            = &AttrEnum{Attr: "NA", Desc: "光标上移n行"}
	Attr_MoveTheCursorDownNLines          = &AttrEnum{Attr: "NB", Desc: "光标下移n行"}
	Attr_MoveTheCursorRightNLines         = &AttrEnum{Attr: "NC", Desc: "光标右移n行"}
	Attr_MoveTheCursorLeftNLines          = &AttrEnum{Attr: "ND", Desc: "光标左移n行"}
	Attr_ClearCcontentsFromCursor2Endline = &AttrEnum{Attr: "K", Desc: "清除从光标到行尾的内容"}
	Attr_SaveCursorPosition               = &AttrEnum{Attr: "s", Desc: "保存光标位置"}
	Attr_RestoreCursorPosition            = &AttrEnum{Attr: "u", Desc: "恢复光标位置"}
	Attr_HideCursor                       = &AttrEnum{Attr: "?25l", Desc: "隐藏光标"}
	Attr_ShowCursor                       = &AttrEnum{Attr: "?25h", Desc: "显示光标"}
)
