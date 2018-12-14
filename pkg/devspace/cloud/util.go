package cloud

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/covexo/devspace/pkg/util/log"
)

// PrintDevSpaces prints the users devspaces
func (p *Provider) PrintDevSpaces(name string) error {
	devspaces, err := p.GetDevSpaces()
	if err != nil {
		return fmt.Errorf("Error retrieving devspaces: %v", err)
	}

	headerColumnNames := []string{
		"DevSpaceID",
		"Name",
		"Created",
	}
	values := [][]string{}

	for _, devspace := range devspaces {
		if name == "" || name == devspace.Name {
			created, err := time.Parse(time.RFC3339, strings.Split(devspace.Created, ".")[0]+"Z")
			if err != nil {
				return err
			}

			values = append(values, []string{
				strconv.Itoa(devspace.DevSpaceID),
				devspace.Name,
				created.String(),
			})
		}
	}

	if len(values) > 0 {
		log.PrintTable(headerColumnNames, values)
	} else {
		log.Info("No devspaces found")
	}

	return nil
}