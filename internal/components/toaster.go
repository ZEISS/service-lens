package components

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/alpine"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/toasts"
)

const (
	INFO    = "info"
	SUCCESS = "success"
	ERROR   = "error"
)

// Toast is a message to display to the user.
type Toast struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}

// New returns a new Toast.
func New(level string, message string) Toast {
	return Toast{level, message}
}

// Error returns the error message.
func (t Toast) Error() string {
	return t.Message
}

// Info returns an info message.
func Info(message string) Toast {
	return New(INFO, message)
}

// Success returns a success message.
func Success(c *fiber.Ctx, message string) {
	New(SUCCESS, message).SetHXTriggerHeader(c)
}

// Error returns an error message.
func Warning(message string) Toast {
	return New(ERROR, message)
}

func (t Toast) jsonify() (string, error) {
	// Set the message to it's error representation to include the level
	t.Message = t.Error()

	// Create the map expected by HTMX
	eventMap := map[string]Toast{}
	eventMap["notify"] = t

	// Convert the structure to JSON
	jsonData, err := json.Marshal(eventMap)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

// SetHXTriggerHeader sets the HTMX trigger header.
func (t Toast) SetHXTriggerHeader(c *fiber.Ctx) error {
	jsonData, err := t.jsonify()
	if err != nil {
		return err
	}

	htmx.Trigger(c, jsonData)

	if t.Level != SUCCESS {
		htmx.ReSwap(c, "none")
	}

	return nil
}

// ToasterProps is the properties for the Toaster component.
type ToasterProps struct {
	// ClassNames are the class names for the toast.
	ClassNames htmx.ClassNames
}

// Toaster is the layout host for the toasts.
func Toaster(props ToasterProps, children ...htmx.Node) htmx.Node {
	return toasts.Toasts(
		toasts.ToastsProps{
			ClassNames: htmx.Merge(
				htmx.ClassNames{},
				props.ClassNames,
			),
		},
		alpine.XData(`{
			notifications: [],
			add(e) {
				this.notifications.push({
					id: e.timeStamp,
					level: e.detail.level,
					message: e.detail.message,
				})
			},
			remove(notification) {
				this.notifications = this.notifications.filter(i => i.id !== notification.id)
			},
		}`),
		htmx.Role("status"),
		htmx.Attribute("aria-live", "polite"),
		alpine.XOn("notify.window", "add($event)"),
		htmx.Template(
			alpine.XFor("notification in notifications"),
			htmx.Attribute(":key", "notification.id"),
			htmx.Div(
				htmx.ClassNames{},
				alpine.XData(`{
					show: false,
					init() {
						this.$nextTick(() => this.show = true)
						
						setTimeout(() => this.transitionOut(), 3000)
					},
					transitionOut() {
						this.show = false

						setTimeout(() => this.remove(this.notification), 500)
					},
				}`),
				alpine.XShow("show"),
				alpine.XTransition("duration.500ms"),
				htmx.Div(
					htmx.ClassNames{
						"alert":       true,
						"alert-error": true,
					},
					alpine.XShow("notification.level === 'error'"),
					htmx.Div(
						alpine.XText("notification.message"),
					),
					buttons.Outline(
						buttons.ButtonProps{
							ClassNames: htmx.ClassNames{
								"btn-sm": true,
							},
						},
						alpine.XOn("click", "transitionOut()"),
						htmx.Text("Close"),
					),
				),
				htmx.Div(
					htmx.ClassNames{
						"alert":         true,
						"alert-success": true,
					},
					alpine.XShow("notification.level === 'success'"),
					htmx.Div(
						alpine.XText("notification.message"),
					),
					buttons.Outline(
						buttons.ButtonProps{
							ClassNames: htmx.ClassNames{
								"btn-sm": true,
							},
						},
						alpine.XOn("click", "transitionOut()"),
						htmx.Text("Close"),
					),
				),
				htmx.Div(
					htmx.ClassNames{
						"alert":      true,
						"alert-info": true,
					},
					alpine.XShow("notification.level === 'info'"),
					htmx.Div(
						alpine.XText("notification.message"),
					),
					buttons.Outline(
						buttons.ButtonProps{
							ClassNames: htmx.ClassNames{
								"btn-sm": true,
							},
						},
						alpine.XOn("click", "transitionOut()"),
						htmx.Text("Close"),
					),
				),
			),
		),
		htmx.Group(children...),
	)
}
