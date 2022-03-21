package core

import (
    "bytes"
    "regexp"
)

var segmentRegex *regexp.Regexp

func init() {
    segmentRegex = regexp.MustCompile(`\{[a-z0-9_.]+\}`)
}

func (m *TemplateString) Parse(str string) error {
    if m != nil && len(str) > 0 {
        index := segmentRegex.FindAllStringIndex(str, -1)
        cur := 0
        for _, i := range index {
            left := i[0]
            right := i[1]

            // {{something}} will be skipped
            if left > 0 && right < len(str) && str[left-1] == '{' && str[right] == '}' {
                continue
            }

            if left > cur {
                m.Segments = append(m.Segments, &TemplateString_Segment{Content: str[cur:left]})
            }

            content := str[left+1 : right-1]
            m.Segments = append(m.Segments, &TemplateString_Segment{Content: content, Templated: true})
            cur = right
        }

        if cur < len(str) {
            m.Segments = append(m.Segments, &TemplateString_Segment{Content: str[cur:]})
        }
    }
    return nil
}

func (m *TemplateString) Format() string {
    if m != nil {
        buffer := bytes.Buffer{}
        for _, segment := range m.Segments {
            if segment.Templated {
                buffer.WriteByte('{')
                buffer.WriteString(segment.Content)
                buffer.WriteByte('}')
            } else {
                buffer.WriteString(segment.Content)
            }
        }
        return buffer.String()
    }
    return ""
}