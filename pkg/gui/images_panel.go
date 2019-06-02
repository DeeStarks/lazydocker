package gui

import (
	"fmt"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/fatih/color"
	"github.com/go-errors/errors"
	"github.com/jesseduffield/gocui"
	"github.com/jesseduffield/lazydocker/pkg/commands"
	"github.com/jesseduffield/lazydocker/pkg/utils"
)

// list panel functions

func (gui *Gui) getImageContexts() []string {
	return []string{"config"}
}

func (gui *Gui) getSelectedImage(g *gocui.Gui) (*commands.Image, error) {
	selectedLine := gui.State.Panels.Images.SelectedLine
	if selectedLine == -1 {
		return &commands.Image{}, gui.Errors.ErrNoImages
	}

	return gui.DockerCommand.Images[selectedLine], nil
}

func (gui *Gui) handleImagesFocus(g *gocui.Gui, v *gocui.View) error {
	if gui.popupPanelFocused() {
		return nil
	}

	cx, cy := v.Cursor()
	_, oy := v.Origin()

	prevSelectedLine := gui.State.Panels.Images.SelectedLine
	newSelectedLine := cy - oy

	if newSelectedLine > len(gui.DockerCommand.Images)-1 || len(utils.Decolorise(gui.DockerCommand.Images[newSelectedLine].Name)) < cx {
		return gui.handleImageSelect(gui.g, v)
	}

	gui.State.Panels.Images.SelectedLine = newSelectedLine

	if prevSelectedLine == newSelectedLine && gui.currentViewName() == v.Name() {
		return gui.handleImagePress(gui.g, v)
	} else {
		return gui.handleImageSelect(gui.g, v)
	}
}

func (gui *Gui) handleImageSelect(g *gocui.Gui, v *gocui.View) error {
	if _, err := gui.g.SetCurrentView(v.Name()); err != nil {
		return err
	}

	Image, err := gui.getSelectedImage(g)
	if err != nil {
		if err != gui.Errors.ErrNoImages {
			return err
		}
		return gui.renderString(g, "main", gui.Tr.SLocalize("NoImages"))
	}

	key := Image.ID + "-" + gui.getImageContexts()[gui.State.Panels.Images.ContextIndex]
	if gui.State.Panels.Main.ObjectKey == key {
		return nil
	} else {
		gui.State.Panels.Main.ObjectKey = key
	}

	if err := gui.focusPoint(0, gui.State.Panels.Images.SelectedLine, len(gui.DockerCommand.Images), v); err != nil {
		return err
	}

	mainView := gui.getMainView()

	mainView.Clear()
	mainView.SetOrigin(0, 0)
	mainView.SetCursor(0, 0)

	switch gui.getImageContexts()[gui.State.Panels.Images.ContextIndex] {
	case "config":
		if err := gui.renderImageConfig(mainView, Image); err != nil {
			return err
		}
	default:
		return errors.New("Unknown context for Images panel")
	}

	return nil
}

func (gui *Gui) renderImageConfig(mainView *gocui.View, image *commands.Image) error {
	mainView.Autoscroll = false
	mainView.Wrap = false
	mainView.Title = "Config"

	gui.T.NewTask(func(stop chan struct{}) {
		output := ""
		output += utils.WithPadding("ID: ", 10) + image.Image.ID + "\n"
		output += utils.WithPadding("Tags: ", 10) + utils.ColoredString(strings.Join(image.Image.RepoTags, ", "), color.FgGreen) + "\n"
		output += utils.WithPadding("Size: ", 10) + utils.FormatDecimalBytes(int(image.Image.Size)) + "\n"
		output += utils.WithPadding("Created: ", 10) + fmt.Sprintf("%v", time.Unix(image.Image.Created, 0).Format(time.RFC1123)) + "\n"

		history, err := image.RenderHistory()
		if err != nil {
			gui.Log.Error(err)
		}

		output += "\n\n" + history

		gui.renderString(gui.g, "main", output)
	})

	return nil
}

func (gui *Gui) refreshImages() error {
	ImagesView := gui.getImagesView()
	if ImagesView == nil {
		// if the ImagesView hasn't been instantiated yet we just return
		return nil
	}
	if err := gui.refreshStateImages(); err != nil {
		return err
	}

	if len(gui.DockerCommand.Images) > 0 && gui.State.Panels.Images.SelectedLine == -1 {
		gui.State.Panels.Images.SelectedLine = 0
	}
	if len(gui.DockerCommand.Images)-1 < gui.State.Panels.Images.SelectedLine {
		gui.State.Panels.Images.SelectedLine = len(gui.DockerCommand.Images) - 1
	}

	gui.g.Update(func(g *gocui.Gui) error {

		ImagesView.Clear()
		isFocused := gui.g.CurrentView().Name() == "Images"
		list, err := utils.RenderList(gui.DockerCommand.Images, utils.IsFocused(isFocused))
		if err != nil {
			return err
		}
		fmt.Fprint(ImagesView, list)

		if ImagesView == g.CurrentView() {
			return gui.handleImageSelect(g, ImagesView)
		}
		return nil
	})

	return nil
}

func (gui *Gui) refreshStateImages() error {
	Images, err := gui.DockerCommand.GetImages()
	if err != nil {
		return err
	}

	gui.DockerCommand.Images = Images

	return nil
}

func (gui *Gui) handleImagesNextLine(g *gocui.Gui, v *gocui.View) error {
	if gui.popupPanelFocused() {
		return nil
	}

	panelState := gui.State.Panels.Images
	gui.changeSelectedLine(&panelState.SelectedLine, len(gui.DockerCommand.Images), false)

	return gui.handleImageSelect(gui.g, v)
}

func (gui *Gui) handleImagesPrevLine(g *gocui.Gui, v *gocui.View) error {
	if gui.popupPanelFocused() {
		return nil
	}

	panelState := gui.State.Panels.Images
	gui.changeSelectedLine(&panelState.SelectedLine, len(gui.DockerCommand.Images), true)

	return gui.handleImageSelect(gui.g, v)
}

func (gui *Gui) handleImagePress(g *gocui.Gui, v *gocui.View) error {
	return nil
}

func (gui *Gui) handleImagesPrevContext(g *gocui.Gui, v *gocui.View) error {
	contexts := gui.getImageContexts()
	if gui.State.Panels.Images.ContextIndex >= len(contexts)-1 {
		gui.State.Panels.Images.ContextIndex = 0
	} else {
		gui.State.Panels.Images.ContextIndex++
	}

	gui.handleImageSelect(gui.g, v)

	return nil
}

func (gui *Gui) handleImagesNextContext(g *gocui.Gui, v *gocui.View) error {
	contexts := gui.getImageContexts()
	if gui.State.Panels.Images.ContextIndex <= 0 {
		gui.State.Panels.Images.ContextIndex = len(contexts) - 1
	} else {
		gui.State.Panels.Images.ContextIndex--
	}

	gui.handleImageSelect(gui.g, v)

	return nil
}

type removeImageOption struct {
	description   string
	command       string
	configOptions types.ImageRemoveOptions
	runCommand    bool
}

// GetDisplayStrings is a function.
func (r *removeImageOption) GetDisplayStrings(isFocused bool) []string {
	return []string{r.description, color.New(color.FgRed).Sprint(r.command)}
}

func (gui *Gui) handleImagesRemoveMenu(g *gocui.Gui, v *gocui.View) error {
	Image, err := gui.getSelectedImage(g)
	if err != nil {
		return nil
	}

	options := []*removeImageOption{
		{
			description:   gui.Tr.SLocalize("remove"),
			command:       "docker image rm " + Image.ID[1:20],
			configOptions: types.ImageRemoveOptions{PruneChildren: true},
			runCommand:    true,
		},
		{
			description:   gui.Tr.SLocalize("removeWithoutPrune"),
			command:       "docker image rm --no-prune " + Image.ID[1:20],
			configOptions: types.ImageRemoveOptions{PruneChildren: false},
			runCommand:    true,
		},
		{
			description: gui.Tr.SLocalize("cancel"),
			runCommand:  false,
		},
	}

	handleMenuPress := func(index int) error {
		if !options[index].runCommand {
			return nil
		}
		configOptions := options[index].configOptions
		if cerr := Image.Remove(configOptions); cerr != nil {
			return gui.createErrorPanel(gui.g, cerr.Error())
		}

		return gui.refreshImages()
	}

	return gui.createMenu("", options, len(options), handleMenuPress)
}

func (gui *Gui) handlePruneImages(g *gocui.Gui, v *gocui.View) error {
	return gui.createConfirmationPanel(gui.g, v, gui.Tr.SLocalize("Confirm"), gui.Tr.SLocalize("confirmPruneImages"), func(g *gocui.Gui, v *gocui.View) error {
		return gui.WithWaitingStatus(gui.Tr.SLocalize("PruningStatus"), func() error {
			err := gui.DockerCommand.PruneImages()
			if err != nil {
				return gui.createErrorPanel(gui.g, err.Error())
			}
			return gui.refreshImages()
		})
	}, nil)
}
