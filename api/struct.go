package api

// TODO: https://github.com/go-validator/validator
// TODO: Functional config patterns
// TODO: Separate API types by section (eg. create, data, etc)
// TODO: Go through and ensure all of these are completed/match up correctly/etc

type FieldType string
type LogicType string
type ActionType string
type DetailToType string
type DetailTargetType string
type DetailTargetValue string
type ConditionOp string
type ConditionVarType string

const (
	FieldTypeDate           FieldType = "date"
	FieldTypeDropdown       FieldType = "dropdown"
	FieldTypeEmail          FieldType = "email"
	FieldTypeFileUpload     FieldType = "file_upload"
	FieldTypeGroup          FieldType = "group"
	FieldTypeLegal          FieldType = "legal"
	FieldTypeLongText       FieldType = "long_text"
	FieldTypeMultipleChoice FieldType = "multiple_choice"
	FieldTypeNumber         FieldType = "number"
	FieldTypeOpinionScale   FieldType = "opinion_scale"
	FieldTypePayment        FieldType = "payment"
	FieldTypePictureChoice  FieldType = "picture_choice"
	FieldTypeRating         FieldType = "rating"
	FieldTypeShortText      FieldType = "short_text"
	FieldTypeStatement      FieldType = "statement"
	FieldTypeWebsite        FieldType = "website"
	FieldTypeYesNo          FieldType = "yes_no"

	LogicTypeField  LogicType = "field"
	LogicTypeHidden LogicType = "hidden"

	ActionTypeJump     ActionType = "jump"
	ActionTypeAdd      ActionType = "add"
	ActionTypeSubtract ActionType = "subtract"
	ActionTypeMultiply ActionType = "multiply"
	ActionTypeDivide   ActionType = "divide"

	DetailToTypeField    DetailToType = "field"
	DetailToTypeHidden   DetailToType = "hidden"
	DetailToTypeThankyou DetailToType = "thankyou"

	DetailTargetTypeThankyou DetailTargetType = "variable"

	DetailTargetValueScore DetailTargetValue = "score"
	DetailTargetValuePrice DetailTargetValue = "price"

	ConditionOpBeginsWith       ConditionOp = "begins_with"
	ConditionOpEndsWith         ConditionOp = "ends_with"
	ConditionOpContains         ConditionOp = "contains"
	ConditionOpNotContains      ConditionOp = "not_contains"
	ConditionOpLowerThan        ConditionOp = "lower_than"
	ConditionOpLowerEqualThan   ConditionOp = "lower_equal_than"
	ConditionOpGreaterThan      ConditionOp = "greater_than"
	ConditionOpGreaterEqualThan ConditionOp = "greater_equal_than"
	ConditionOpIs               ConditionOp = "is"
	ConditionOpIsNot            ConditionOp = "is_not"
	ConditionOpEqual            ConditionOp = "equal"
	ConditionOpNotEqual         ConditionOp = "not_equal"
	ConditionOpAlways           ConditionOp = "always"
	ConditionOpOn               ConditionOp = "on"
	ConditionOpNotOn            ConditionOp = "not_on"
	ConditionOpEarlierThan      ConditionOp = "earlier_than"
	ConditionOpEarlierThanOrOn  ConditionOp = "earlier_than_or_on"
	ConditionOpLaterThan        ConditionOp = "later_than"
	ConditionOpLaterThanOrOn    ConditionOp = "later_than_or_on"

	ConditionVarTypeField    ConditionVarType = "field"
	ConditionVarTypeHidden   ConditionVarType = "hidden"
	ConditionVarTypeVariable ConditionVarType = "variable"
	ConditionVarTypeConstant ConditionVarType = "constant"
	ConditionVarTypeEnd      ConditionVarType = "end"
)

type Form struct {
	ID              string           `json:"id"`
	Title           string           `json:"title"`
	Hidden          []string         `json:"hidden"`
	WelcomeScreens  []WelcomeScreen  `json:"welcome_screens"`
	ThankyouScreens []ThankyouScreen `json:"thankyou_screens"`
	Fields          []Field          `json:"fields"`
	Logic           []Logic          `json:"logic"`

	Theme     Href     `json:"theme"`
	Workspace Href     `json:"workspace"`
	Links     Links    `json:"_links"`
	Language  string   `json:"language"`
	Settings  Settings `json:"settings"`
}

type Links struct {
	Display string `json:"display"`
}

type Href struct {
	Href string `json:"href"`
}

type Settings struct {
	IsPublic             bool   `json:"is_public"`
	IsTrial              bool   `json:"is_trial"`
	Language             string `json:"language"`
	ProgressBar          string `json:"progress_bar"`
	ShowProgressBar      bool   `json:"show_progress_bar"`
	ShowTypeformBranding bool   `json:"show_typeform_branding"`
	Meta                 struct {
		AllowIndexing bool `json:"allow_indexing"`
	} `json:"meta"`
	Notifications Notifications `json:"notifications"`
}

type Notifications struct {
	Self       NotificationSelf       `json:"self"`
	Respondent NotificationRespondent `json:"respondent"`
}

type notificationCommon struct {
	Subject string `json:"subject"`
	Message string `json:"message"`
	Enabled bool   `json:"enabled"`
}

type NotificationSelf struct {
	Recipients []string `json:"recipients"`
	notificationCommon
}

type NotificationRespondent struct {
	Recipient string `json:"recipient"`
	notificationCommon
}

type WelcomeScreen struct {
	Ref        string                  `json:"ref"`
	Title      string                  `json:"title"`
	Properties WelcomeScreenProperties `json:"properties"`
	Attachment Attachment              `json:"attachment"`
}

type WelcomeScreenProperties struct {
	ShowButton  bool   `json:"show_button"`
	Description string `json:"description"`
	ButtonText  string `json:"button_text"`
}

type ThankyouScreen struct {
	Ref        string                   `json:"ref"`
	Title      string                   `json:"title"`
	Properties ThankyouScreenProperties `json:"properties"`
	Attachment Attachment               `json:"attachment"`
}

type ThankyouScreenProperties struct {
	ShowButton  bool   `json:"show_button"`
	ShareIcons  bool   `json:"share_icons"`
	ButtonMode  string `json:"button_mode"`
	ButtonText  string `json:"button_text"`
	RedirectUrl string `json:"redirect_url"`
}

type Attachment struct {
	Type string `json:"type"`
	Href string `json:"href"`
}

type Field struct {
	ID          string           `json:"id"`
	Ref         string           `json:"ref"`
	Title       string           `json:"title"`
	Type        FieldType        `json:"type"`
	Properties  FieldProperties  `json:"properties"`
	Validations FieldValidations `json:"validations"`
	//Attachment FieldAttachment `json:"attachment"`
}

type FieldProperties struct {
	Randomize              bool           `json:"randomize"`
	AllowMultipleSelection bool           `json:"allow_multiple_selection"`
	AllowOtherChoice       bool           `json:"allow_other_choice"`
	VerticalAlignment      bool           `json:"vertical_alignment"`
	Choices                []FieldChoices `json:"choices"`

	// TODO: These might only be for groups?
	Description string  `json:"description"`
	ShowButton  bool    `json:"show_button"`
	ButtonText  string  `json:"button_text"`
	Fields      []Field `json:"fields"`
}

type FieldChoices struct {
	ID    string `json:"id"`
	Ref   string `json:"ref"`
	Label string `json:"label"`
}

type FieldValidations struct {
	Required bool `json:"required"`
	MinValue int  `json:"min_value"`
	MaxValue int  `json:"max_value"`
}

//type FieldAttachment struct{}

type Logic struct {
	Type    LogicType `json:"type"`
	Ref     string    `json:"ref"`
	Actions []Action  `json:"actions"`
}

type Action struct {
	Action    ActionType `json:"action"`
	Details   Detail     `json:"details"`
	Condition Condition  `json:"condition"`
}

type Detail struct {
	To     DetailTo     `json:"to"`
	Target DetailTarget `json:"target"`
	Value  DetailValue  `json:"value"`
}

type DetailTo struct {
	Type  DetailToType `json:"type"`
	Value string       `json:"value"`
}

type DetailTarget struct {
	Type  DetailTargetType  `json:"type"`
	Value DetailTargetValue `json:"value"`
}

type DetailValue struct {
	Type  string `json:"type"`
	Value int    `json:"value"`
}

type Condition struct {
	Op   ConditionOp    `json:"op"`
	Vars []ConditionVar `json:"vars"`
}

type ConditionVar struct {
	Type  ConditionVarType `json:"type"`
	Value interface{}      `json:"value"`
}
