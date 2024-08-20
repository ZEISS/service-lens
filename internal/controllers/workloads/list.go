package workloads

import (
	"context"
	"fmt"

	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/components"
	"github.com/zeiss/service-lens/internal/models"
	"github.com/zeiss/service-lens/internal/ports"
	"github.com/zeiss/service-lens/internal/utils"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/icons"
	"github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/fiber-htmx/components/tables"
)

// WorkloadListControllerImpl ...
type WorkloadListControllerImpl struct {
	workloads tables.Results[models.Workload]
	store     seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewWorkloadListController ...
func NewWorkloadListController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *WorkloadListControllerImpl {
	return &WorkloadListControllerImpl{store: store}
}

// Prepare ...
func (w *WorkloadListControllerImpl) Prepare() error {
	if err := w.BindQuery(&w.workloads); err != nil {
		return err
	}

	return w.store.ReadTx(w.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListWorkloads(ctx, &w.workloads)
	})
}

// Get ...
func (w *WorkloadListControllerImpl) Get() error {
	return w.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Path: w.Path(),
				User: w.Session().User,
			},
			func() htmx.Node {
				return cards.CardBordered(
					cards.CardProps{
						ClassNames: htmx.ClassNames{
							"m-2": true,
						},
					},
					cards.Body(
						cards.BodyProps{},
						components.Table(
							components.TableProps[*models.Workload]{
								Rows:   w.workloads.GetRows(),
								Offset: w.workloads.GetOffset(),
								Limit:  w.workloads.GetLimit(),
								Total:  w.workloads.GetTotalRows(),
							},
							[]tables.ColumnDef[*models.Workload]{
								{
									ID:          "id",
									AccessorKey: "id",
									Header: func(p tables.TableProps) htmx.Node {
										return htmx.Th(htmx.Text("ID"))
									},
									Cell: func(p tables.TableProps, row *models.Workload) htmx.Node {
										return htmx.Td(
											htmx.Text(row.ID.String()),
										)
									},
								},
								{
									ID:          "name",
									AccessorKey: "name",
									Header: func(p tables.TableProps) htmx.Node {
										return htmx.Th(htmx.Text("Name"))
									},
									Cell: func(p tables.TableProps, row *models.Workload) htmx.Node {
										return htmx.Td(
											links.Link(
												links.LinkProps{
													Href: "/workloads/" + row.ID.String(),
												},
												htmx.Text(row.Name),
											),
										)
									},
								},
								{
									ID:          "profile",
									AccessorKey: "profile",
									Header: func(p tables.TableProps) htmx.Node {
										return htmx.Th(htmx.Text("Profile"))
									},
									Cell: func(p tables.TableProps, row *models.Workload) htmx.Node {
										return htmx.Td(
											links.Link(
												links.LinkProps{
													Href: fmt.Sprintf(utils.ShowProfileUrlFormat, row.Profile.ID),
												},
												htmx.Text(row.Profile.Name),
											),
										)
									},
								},
								{
									ID:          "environment",
									AccessorKey: "environment",
									Header: func(p tables.TableProps) htmx.Node {
										return htmx.Th(htmx.Text("Environment"))
									},
									Cell: func(p tables.TableProps, row *models.Workload) htmx.Node {
										return htmx.Td(
											links.Link(
												links.LinkProps{
													Href: fmt.Sprintf(utils.ShowEnvironmentUrlFormat, row.Environment.ID),
												},
												htmx.Text(row.Environment.Name),
											),
										)
									},
								},
								{
									Header: func(p tables.TableProps) htmx.Node {
										return nil
									},
									Cell: func(p tables.TableProps, row *models.Workload) htmx.Node {
										return htmx.Td(
											buttons.Button(
												buttons.ButtonProps{
													ClassNames: htmx.ClassNames{
														"btn-sm": true,
													},
												},
												htmx.HxDelete(fmt.Sprintf(utils.DeleteWorkloadUrlFormat, row.ID)),
												htmx.HxConfirm("Are you sure you want to delete workload?"),
												htmx.HxTarget("closest tr"),
												htmx.HxSwap("outerHTML swap:1s"),
												icons.TrashOutline(
													icons.IconProps{
														ClassNames: htmx.ClassNames{
															"w-6 h-6": false,
															"w-4":     true,
															"h-4":     true,
														},
													},
												),
											),
										)
									},
								},
							}...,
						),
					),
				)
			},
		),
	)
}
