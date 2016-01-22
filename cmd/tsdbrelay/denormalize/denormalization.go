package denormalize

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"bosun.org/opentsdb"
)

type DenormalizationRule struct {
	Metric   string
	TagNames []string
}

func (d *DenormalizationRule) String() string {
	var inputTags, outputTags, resBuf bytes.Buffer
	val := 'a'
	for i, tagk := range d.TagNames {
		if i != 0 {
			inputTags.WriteRune(',')
			outputTags.WriteRune('.')
		}
		inputTags.WriteString(tagk)
		inputTags.WriteRune('=')
		inputTags.WriteRune(val)
		outputTags.WriteRune(val)
		val++
	}
	resBuf.WriteString(d.Metric)
	resBuf.WriteRune('{')
	resBuf.Write(inputTags.Bytes())
	resBuf.WriteString("} -> __")
	resBuf.Write(outputTags.Bytes())
	resBuf.WriteRune('.')
	resBuf.WriteString(d.Metric)
	return resBuf.String()
}

func ParseDenormalizationRules(config string) (map[string]*DenormalizationRule, error) {
	m := make(map[string]*DenormalizationRule)
	rules := strings.Split(config, ",")
	for _, r := range rules {
		parts := strings.Split(r, "__")
		if len(parts) < 2 {
			return nil, fmt.Errorf("Denormalization rules must have at least one tag name specified.")
		}
		rule := &DenormalizationRule{Metric: parts[0]}
		for _, part := range parts[1:] {
			rule.TagNames = append(rule.TagNames, part)
		}
		log.Println("Denormalizing", rule)
		m[rule.Metric] = rule
	}
	return m, nil
}

func (d *DenormalizationRule) Translate(dp *opentsdb.DataPoint) error {
	tagString := bytes.NewBufferString("__")
	for i, tagName := range d.TagNames {
		val, ok := dp.Tags[tagName]
		if !ok {
			return fmt.Errorf("tag %s not present in data point for %s.", tagName, dp.Metric)
		}
		if i > 0 {
			tagString.WriteRune('.')
		}
		tagString.WriteString(val)
	}
	tagString.WriteRune('.')
	tagString.WriteString(dp.Metric)
	dp.Metric = tagString.String()
	return nil
}
