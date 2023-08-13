package interpreter

type ResultStatus string

const (
	ResultStatusNone     ResultStatus = "none"
	ResultStatusBreak    ResultStatus = "break"
	ResultStatusContinue ResultStatus = "continue"
	ResultStatusReturn   ResultStatus = "return"
)

type BoxedResult struct {
	Value  any
	Status ResultStatus
}

func NewBoxedResult(value any, status ResultStatus) BoxedResult {
	return BoxedResult{
		Value:  value,
		Status: status,
	}
}

func NewBoxedResults(values ...any) []BoxedResult {
	results := make([]BoxedResult, 0)

	for _, v := range values {
		results = append(results, BoxedResult{
			Value:  v,
			Status: ResultStatusNone,
		})
	}

	return results
}

type Result struct {
	Values []BoxedResult
}

func NewResult(values ...any) *Result {
	r := &Result{
		Values: make([]BoxedResult, 0),
	}

	if values != nil {
		r.Values = NewBoxedResults(values...)
	}

	return r
}
func NewResultWithStatus(value any, status ResultStatus) *Result {
	r := &Result{
		Values: make([]BoxedResult, 0),
	}

	r.Values = append(r.Values, NewBoxedResult(value, status))

	return r
}

func (self *Result) Add(value any, resultStatus ...ResultStatus) *Result {
	status := ResultStatusNone
	if resultStatus != nil {
		status = resultStatus[0]
	}
	self.Values = append(self.Values, NewBoxedResult(value, status))

	return self
}

func (self *Result) AddExisting(result BoxedResult) *Result {
	self.Values = append(self.Values, result)
	return self
}

func (self *Result) Merge(results ...*Result) *Result {
	if results == nil {
		return nil
	}

	for _, r := range results {
		if r == nil {
			continue
		}

		self.Values = append(self.Values, r.Values...)
	}

	return self
}

func (self *Result) FirstBoxedResult() (BoxedResult, bool) {
	if len(self.Values) == 0 {
		return BoxedResult{}, false
	}

	return self.Values[0], true
}

func (self *Result) First() any {
	if len(self.Values) == 0 {
		return nil
	}

	return self.Values[0].Value
}

func (self *Result) HasValue() bool {
	if self.Values == nil {
		return false
	}
	if len(self.Values) <= 0 {
		return false
	}

	return true
}

func (self *Result) Clear() {
	self.Values = nil
}

func (self *Result) Last() any {
	if len(self.Values) == 0 {
		return nil
	}

	return self.Values[len(self.Values)-1]
}

func (self *Result) HasStatus(statusBreak ResultStatus) (BoxedResult, bool) {
	if self.Values == nil || len(self.Values) <= 0 {
		return BoxedResult{}, false
	}

	for _, value := range self.Values {
		if value.Status == statusBreak {
			return value, true
		}
	}

	return BoxedResult{}, false
}

func (self *Result) HasAnyStatus(status ...ResultStatus) bool {
	if self.Values == nil || len(self.Values) <= 0 {
		return false
	}

	for _, value := range self.Values {
		for _, s := range status {
			if value.Status == s {
				return true
			}
		}
	}

	return false
}
