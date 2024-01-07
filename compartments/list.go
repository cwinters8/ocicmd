package compartments

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/identity"
	"github.com/spf13/cobra"
)

func ListCmd() *cobra.Command {
	recursive := false
	var filters []string
	cmd := cobra.Command{
		Use:     "compartments",
		Aliases: []string{"cmp"},
		Short:   "Get a list of all top level compartments",
		RunE: func(cmd *cobra.Command, args []string) error {
			compartments, err := List(&recursive, &filters)
			if err != nil {
				return err
			}
			marshalled, err := json.MarshalIndent(compartments, "", "    ")
			if err != nil {
				return fmt.Errorf("failed to marshal json: %w", err)
			}
			fmt.Printf("%v\n", string(marshalled))
			return nil
		},
	}
	cmd.Flags().BoolVarP(&recursive, "recursive", "r", false, "recursively retrieve all child compartments")
	cmd.Flags().StringSliceVarP(&filters, "filter", "f", nil, "case insensitive. A list of names to filter results. Partial matches will be included")
	return &cmd
}

func List(recursive *bool, filters *[]string) ([]identity.Compartment, error) {
	cfg := common.DefaultConfigProvider()
	client, err := identity.NewIdentityClientWithConfigurationProvider(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to initiate identity client: %w", err)
	}
	tenancy, err := cfg.TenancyOCID()
	if err != nil {
		return nil, fmt.Errorf("failed to get tenancy ocid from config: %w", err)
	}
	req := identity.ListCompartmentsRequest{
		CompartmentId:          &tenancy,
		CompartmentIdInSubtree: recursive,
	}
	resp, err := client.ListCompartments(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("failed to list compartments: %w", err)
	}
	items := resp.Items
	if filters != nil {
		f := *filters
		if len(f) > 0 {
			filteredItems := []identity.Compartment{}
			for idx := 0; idx < len(items); idx++ {
				name := strings.ToLower(*items[idx].Name)
				for fIdx := 0; fIdx < len(f); fIdx++ {
					if strings.Contains(name, strings.ToLower(f[fIdx])) {
						filteredItems = append(filteredItems, items[idx])
						// prevent the item from matching more than one filter and being added to the slice multiple times
						break
					}
				}
			}
			return filteredItems, nil
		}
	}
	return items, nil
}
