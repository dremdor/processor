package assembler

import (
	"fmt"
	"strings"
)

type LineKind int

const (
	LineEmpty LineKind = iota
	LineOrg
	LineDb
	LineInstr
)

type Line struct {
	LineNo int

	Label string
	Kind  LineKind

	Mnemonic string
	Arg0     string

	Args []string
}

func LexLines(src string) ([]Line, error) {
	raw := strings.Split(src, "\n")
	out := make([]Line, 0, len(raw))

	for i, s := range raw {
		ln := Line{LineNo: i + 1}
		line := stripComment(strings.TrimSpace(s))
		if line == "" {
			ln.Kind = LineEmpty
			out = append(out, ln)
			continue
		}

		if idx := strings.Index(line, ":"); idx != -1 {
			before := strings.TrimSpace(line[:idx])
			after := strings.TrimSpace(line[idx+1:])
			if before == "" {
				return nil, fmt.Errorf("%d: empty label", ln.LineNo)
			}
			ln.Label = before
			line = after
			if line == "" {
				ln.Kind = LineEmpty
				out = append(out, ln)
				continue
			}
		}

		if strings.HasPrefix(line, ".") {
			fields := splitFields(line)
			switch strings.ToLower(fields[0]) {
			case ".org":
				if len(fields) != 2 {
					return nil, fmt.Errorf("%d: .org expects one argument", ln.LineNo)
				}
				ln.Kind = LineOrg
				ln.Arg0 = fields[1]
			case ".db":
				ln.Kind = LineDb
				args := splitCSV(strings.TrimSpace(line[len(fields[0]):]))
				if len(args) == 0 {
					return nil, fmt.Errorf("%d: .db expects at least one value", ln.LineNo)
				}
				ln.Args = args
			default:
				return nil, fmt.Errorf("%d: unknown directive %q", ln.LineNo, fields[0])
			}

			out = append(out, ln)
			continue
		}

		fields := splitFields(line)
		if len(fields) == 0 {
			ln.Kind = LineEmpty
			out = append(out, ln)
			continue
		}
		ln.Kind = LineInstr
		ln.Mnemonic = strings.ToUpper(fields[0])

		if len(fields) >= 2 {
			ln.Arg0 = fields[1]
		} else {
			ln.Arg0 = ""
		}
		if len(fields) > 2 {
			return nil, fmt.Errorf("%d: too many tokens in instruction", ln.LineNo)
		}

		out = append(out, ln)
	}

	return out, nil
}

func stripComment(s string) string {
	if idx := strings.Index(s, ";"); idx != -1 {
		return strings.TrimSpace(s[:idx])
	}
	return s
}

func splitFields(s string) []string {
	return strings.Fields(s)
}

func splitCSV(s string) []string {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		v := strings.TrimSpace(p)
		if v != "" {
			out = append(out, v)
		}
	}
	return out
}
