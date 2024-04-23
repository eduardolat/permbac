package generate

import (
	_ "embed"
	"fmt"
	"go/format"
	"strings"

	"github.com/eduardolat/permbac/internal/schema"
)

//go:embed template/permbac.go
var permbacTpl string

func GeneratePerms(
	version string,
	config schema.Config,
	perms schema.Perms,
) ([]byte, error) {
	permTpl := `// Perm%s %s.
	//
	// Name: %s
	var Perm%s = Perm{
		Name: "%s",
		Desc: "%s",
	}`
	permNameTpl := `Perm%s,`

	permsSl := []string{}
	namesSl := []string{}
	for _, perm := range perms {
		permsSl = append(permsSl, fmt.Sprintf(
			permTpl,
			perm.Name, perm.Desc, perm.Name, perm.Name, perm.Name, perm.Desc,
		))
		namesSl = append(namesSl, fmt.Sprintf(permNameTpl, perm.Name))
	}

	comment := fmt.Sprintf("// Code generated by PermBAC %s\n", version)
	comment += "// https://github.com/eduardolat/permbac\n"
	comment += "// DO NOT EDIT\n"

	code := []string{
		comment,
		replacePkg(permbacTpl, "template", config.Package),
	}

	codeStr := strings.Join(code, "\n\n")
	codeStr = strings.ReplaceAll(codeStr, "//*NAMES_HERE*//", strings.Join(namesSl, "\n"))
	codeStr = strings.ReplaceAll(codeStr, "//*PERMS_HERE*//", strings.Join(permsSl, "\n\n"))

	b, err := format.Source([]byte(codeStr))
	if err != nil {
		return nil, fmt.Errorf("error formatting the perms source: %w", err)
	}

	return b, nil
}
