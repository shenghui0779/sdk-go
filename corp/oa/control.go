package oa

type ControlType string

const (
	ControlText            ControlType = "Text"
	ControlTextarea        ControlType = "Textarea"
	ControlNumber          ControlType = "Number"
	ControlMoney           ControlType = "Money"
	ControlDate            ControlType = "Date"
	ControlSelector        ControlType = "Selector"
	ControlContact         ControlType = "Contact"
	ControlTips            ControlType = "Tips"
	ControlFile            ControlType = "File"
	ControlTable           ControlType = "Table"
	ControlAttendance      ControlType = "Attendance"
	ControlVacation        ControlType = "Vacation"
	ControlDateRange       ControlType = "DateRange"
	ControlLocation        ControlType = "Location"
	ControlFormula         ControlType = "Formula"
	ControlSchoolContact   ControlType = "SchoolContact"
	ControlPunchCorrection ControlType = "PunchCorrection"
	ControlRelatedApproval ControlType = "RelatedApproval"
)

type ControlProperty struct {
	Control     ControlType    `json:"control"`
	ID          string         `json:"id"`
	Title       []*DisplayText `json:"title"`
	Placeholder []*DisplayText `json:"placeholder"`
	Require     int            `json:"require"`
	UnPrint     int            `json:"un_print"`
}

type ControlConfig struct {
	Date         *DateConfig       `json:"date,omitempty"`
	Selector     *SelectorConfig   `json:"selector,omitempty"`
	Table        *TableConfig      `json:"table,omitempty"`
	Attendance   *AttendanceConfig `json:"attendance,omitempty"`
	VacationList *VacationConfig   `json:"vacation_list,omitempty"`
}

type ControlValue struct {
	Text            string                  `json:"text"`
	NewNumber       string                  `json:"new_number"`
	NewMoney        string                  `json:"new_money"`
	Tips            interface{}             `json:"tips"`
	Date            *DateValue              `json:"date"`
	Seletor         *SelectorValue          `json:"seletor"`
	Members         []*ContactMember        `json:"members"`
	Departments     []*ContactDepartment    `json:"departments"`
	Files           []*FileValue            `json:"files"`
	Children        []*TableValue           `json:"children"`
	StatField       interface{}             `json:"stat_field"`
	SumField        interface{}             `json:"sum_field"`
	Students        []*SchoolContactStudent `json:"students"`
	Classes         []*SchoolContactClass   `json:"classes"`
	DateRange       *DateRangeValue         `json:"date_range"`
	Location        *LocationValue          `json:"location"`
	Formula         *FormulaValue           `json:"formula"`
	Vacation        *VacationValue          `json:"vacation"`
	Attendance      *AttendanceValue        `json:"attendance"`
	PunchCorrection *PunchCorrectionValue   `json:"punch_correction"`
	RelatedApproval []*RelatedApprovalValue `json:"related_approval"`
}

type DateConfig struct {
	Type string `json:"type"` // 时间展示类型：day-日期；hour-日期+时间
}

type DateValue struct {
	Type       string `json:"type"`
	STimestamp string `json:"s_timestamp"`
}

type SelectorConfig struct {
	Type    string // 选择类型：single-单选；multi-多选
	Options []*SelectorOption
}

type SelectorValue struct {
	Type    string
	Options []*SelectorOption
}

type SelectorOption struct {
	Key   string         `json:"key"`
	Value []*DisplayText `json:"value"`
}

type ContactConfig struct {
	Type string `json:"type"` // 选择方式：single-单选；multi-多选
	Mode string `json:"mode"` // 选择对象：user-成员；department-部门
}

type ContactMember struct {
	UserID string `json:"userid"`
	Name   string `json:"name"`
}

type ContactDepartment struct {
	OpenApiID string `json:"openapi_id"`
	Name      string `json:"name"`
}

type FileValue struct {
	FileID string `json:"file_id"`
}

type TableConfig struct {
	Children  []*TableChildConfig `json:"children"`
	StatField interface{}         `json:"stat_field"`
}

type TableChildConfig struct {
	Property *ControlProperty `json:"property"`
}

type TableValue struct {
	List []*TableChildValue `json:"list"`
}

type TableChildValue struct {
	Control ControlType    `json:"control"`
	ID      string         `json:"id"`
	Title   []*DisplayText `json:"title"`
	Value   *ControlValue  `json:"value"`
}

type DateRangeValue struct {
	Type        string `json:"type"`
	NewBegin    int64  `json:"new_begin"`
	NewEnd      int64  `json:"new_end"`
	NewDuration int    `json:"new_duration"`
}

type LocationValue struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Title     string `json:"title"`
	Address   string `json:"address"`
	Time      int64  `json:"time"`
}

type FormulaValue struct {
	Value string `json:"value"`
}

type SchoolContactStudent struct {
	Name string `json:"name"`
}

type SchoolContactClass struct {
	Name string `json:"name"`
}

type VacationConfig struct {
	Item []*VacationConfigItem `json:"item"`
}

type VacationConfigItem struct {
	ID   int            `json:"id"`
	Name []*DisplayText `json:"name"`
}

type VacationValue struct {
	Selector   *SelectorValue   `json:"selector"`
	Attendance *AttendanceValue `json:"attendance"`
}

type AttendanceConfig struct {
	DateRange *DateRangeValue `json:"date_range"`
	Type      int             `json:"type"`
}

type AttendanceValue struct {
	DateRange *DateRangeValue      `json:"date_range"`
	Type      int                  `json:"type"`
	SliceInfo *AttendanceSliceInfo `json:"slice_info"`
}

type AttendanceSliceInfo struct {
	DayItems []*AttendanceDay `json:"day_items"`
	Duration int              `json:"duration"`
	State    int              `json:"state"`
}

type AttendanceDay struct {
	DayTime  int64 `json:"daytime"`
	Duration int   `json:"duration"`
}

type PunchCorrectionValue struct {
	State string `json:"state"`
	Time  int64  `json:"time"`
}

type RelatedApprovalValue struct {
	TemplateNames []*DisplayText `json:"template_names"`
	SPStatus      int            `json:"sp_status"`
	Name          string         `json:"name"`
	CreateTime    int64          `json:"create_time"`
	SPNO          string         `json:"sp_no"`
}

type OAUser struct {
	UserID string `json:"userid"`
}

type DisplayText struct {
	Text string `json:"text"`
	Lang string `json:"lang"`
}

type KeyValue struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type ApplyData struct {
	Contents []*ApplyContent `json:"contents"`
}

type ApplyContent struct {
	Control ControlType    `json:"control"`
	ID      string         `json:"id"`
	Title   []*DisplayText `json:"title"`
	Value   *ControlValue  `json:"value"`
}
