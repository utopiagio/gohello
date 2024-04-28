// SPDX-License-Identifier: Unlicense OR MIT

/* github.com/utopiagio/gohello/hello.go */

package main

import (
    "strconv"
    
    ui "github.com/utopiagio/utopia"
    "github.com/utopiagio/utopia-x/uireference"
    "github.com/utopiagio/utopia/desktop"
    "github.com/utopiagio/utopia/metrics"
)

var mainwin *ui.GoWindowObj
var viewer *ui.GoWindowObj
var popup *ui.GoPopupWindowObj
var lblWindowProperties *ui.GoLabelObj
var richText *ui.GoRichTextObj
var navigator *ui.GoListViewObj
var hdrSection *ui.GoLabelObj
var layoutWindowProperties *ui.GoLayoutObj

func main() {
    // create application instance before any other objects
    app := ui.GoApplication("GoHelloDemo")
    // create application window
    mainwin = ui.GoMainWindow("GoHello Demo - UtopiaGio Package")
    // set the window layout style to stack widgets vertically
    mainwin.SetLayoutStyle(ui.VFlexBoxLayout)
    mainwin.SetMargin(10,10,10,10)
    mainwin.SetBorder(ui.BorderSingleLine, 2, 10, ui.Color_Blue)
    mainwin.SetPadding(10,10,10,10)
    mainwin.SetSize(900, 640)
    mainwin.SetPos(100,100)

    // setup the main window MenuBar
    menuBar := mainwin.MenuBar()
    menuBar.Show()
    mnuCode := menuBar.AddMenu("Documentation")
    mnuCode.AddAction("Overview", LoadOverview)
        
    // add the header layout to align widgets horizontally
    layoutHeader := ui.GoHFlexBoxLayout(mainwin.Layout())
    layoutHeader.SetSizePolicy(ui.ExpandingWidth, ui.PreferredHeight)
    layoutHeader.SetPadding(0,0,0,5)

        lblVBox := ui.GoLabel(layoutHeader, "GoVBoxLayout")
        lblVBox.SetSizePolicy(ui.FixedWidth, ui.PreferredHeight)
        lblVBox.SetWidth(270)
        lblVBox.SetBackgroundColor(ui.Color_LightBlue)
        lblVBox.SetBorder(ui.BorderSingleLine, 1, 6, ui.Color_LightGray)

        ui.GoSpacer(layoutHeader, 10)

        lblVFlexBox := ui.GoLabel(layoutHeader, "GoVFlexBoxLayout")
        lblVFlexBox.SetSizePolicy(ui.ExpandingWidth, ui.PreferredHeight)
        lblVFlexBox.SetBackgroundColor(ui.Color_LightBlue)
        lblVFlexBox.SetBorder(ui.BorderSingleLine, 1, 6, ui.Color_LightGray)

    // add the Content layout to align the content horizontally
    layoutContent := ui.GoHFlexBoxLayout(mainwin.Layout())
    layoutContent.SetMargin(0,0,0,0)
    layoutContent.SetPadding(0,0,0,0)
    
        // add the VBox layout to contain the Label text providing scrollbars
        layoutWindowProperties = ui.GoVBoxLayout(layoutContent)
        layoutWindowProperties.SetSizePolicy(ui.FixedWidth, ui.ExpandingHeight)
        layoutWindowProperties.SetWidth(270)
        layoutWindowProperties.SetBorder(ui.BorderSingleLine, 2, 6, ui.Color_LightGray)

            lblWindowProperties = ui.GoLabel(layoutWindowProperties, "")
            lblWindowProperties.SetWrap(false)
            lblWindowProperties.SetPadding(8,8,8,8)

        // add fixed spacer between content
        ui.GoSpacer(layoutContent, 10)

        // add the VFlexBox layout to contain the Label layout options
        layoutLblSizing := ui.GoVFlexBoxLayout(layoutContent)
        layoutLblSizing.SetBorder(ui.BorderSingleLine, 2, 6, ui.Color_LightGray)
        layoutLblSizing.SetPadding(10,10,10,10)

            ui.GoLabel(layoutLblSizing, "GoLabel...GoSizePolicy{Horiz: FixedWidth, Vert: FixedHeight}")
            lblLabel0 := ui.GoLabel(layoutLblSizing, "Hello from UtopiaGio")
            lblLabel0.SetSizePolicy(ui.FixedWidth, ui.FixedHeight)
            lblLabel0.SetBorder(ui.BorderSingleLine, 2, 6, ui.Color_LightBlue)
            lblLabel0.SetHeight(50)
            lblLabel0.SetWidth(500)
            lblLabel0.SetMaxLines(1)
            lblLabel0.SetFontSize(24)
            lblLabel0.SetFontBold(true)
            lblLabel0.SetTextColor(ui.Color_Blue)
           
            ui.GoSpacer(layoutLblSizing, 10)

            ui.GoLabel(layoutLblSizing, "GoLabel...GoSizePolicy{Horiz: PreferredWidth, Vert: PreferredHeight}")
            lblLabel1 := ui.GoLabel(layoutLblSizing, "Hello from UtopiaGio")
            lblLabel1.SetSizePolicy(ui.PreferredWidth, ui.PreferredHeight)
            lblLabel1.SetBorder(ui.BorderSingleLine, 2, 6, ui.Color_LightBlue)
            lblLabel1.SetMaxLines(0)
            lblLabel1.SetFontSize(36)
            lblLabel1.SetFontBold(true)
            lblLabel1.SetTextColor(ui.Color_Blue)
         
            ui.GoSpacer(layoutLblSizing, 10)
         
            ui.GoLabel(layoutLblSizing, "GoLabel...GoSizePolicy{Horiz: ExpandingWidth, Vert: ExpandingHeight}")
            lblLabel2 := ui.GoLabel(layoutLblSizing, "Hello from UtopiaGio")
            lblLabel2.SetSizePolicy(ui.ExpandingWidth, ui.ExpandingHeight)
            lblLabel2.SetBorder(ui.BorderSingleLine, 2, 6, ui.Color_LightBlue)
            lblLabel2.SetMaxLines(0)
            lblLabel2.SetFontSize(48)
            lblLabel2.SetFontBold(true)
            lblLabel2.SetTextColor(ui.Color_Blue)
    
    // add the Action Bar layout to contain button controls
    layoutBottom := ui.GoHFlexBoxLayout(mainwin.Layout())
    layoutBottom.SetSizePolicy(ui.ExpandingWidth, ui.PreferredHeight)   // Note: ui.FixedHeight
    layoutBottom.SetMargin(0,10,0,0)
    layoutBottom.SetPadding(0,0,0,0)
    layoutBottom.SetBorder(ui.BorderSingleLine, 2, 10, ui.Color_Blue)

        lblHint := ui.GoLabel(layoutBottom, "Try Resizing the Window...")
        lblHint.SetMargin(10, 6, 0, 0)
        lblHint.SetMaxLines(1)
        lblHint.SetFontSize(24)
        lblHint.SetFontBold(true)
        // add expanding spacer
        padding := ui.GoSpacer(layoutBottom, 0)
        padding.SetSizePolicy(ui.ExpandingWidth, ui.FixedHeight)

        btnClose := ui.GoButton(layoutBottom, "Close")
        btnClose.SetWidth(260)
        btnClose.SetHeight(160)
        btnClose.SetMargin(4,4,4,4)
        btnClose.SetPadding(4,4,4,4)
        btnClose.SetOnClick(ActionExit_Clicked)
    
    lblWindowProperties.SetText("Click the Refresh Button........\n\n   to see the window properties.")
    
    mainwin.SetOnConfig(UpdateWindowProperties)
    // show the application window
    mainwin.Show()
    // run the application
    app.Run()
}

func ActionExit_Clicked() {
    mainwin.Close()
}

func GetWindowProperties() (text string) {
    text = "WINDOW PROPERTIES>\n\n"
    text += "Screen Geometry :" + "\n"
    text += "    ScreenWidth:       " + strconv.Itoa(metrics.DpToPx(ui.GoDpr, desktop.Width())) + " px\n"    // * ui.GoDpr)) + "\n"
    text += "    ScreenHeight:      " + strconv.Itoa(metrics.DpToPx(ui.GoDpr, desktop.Height())) + " px\n"    // * ui.GoDpr)) + "\n\n"
    text += "    HorizontalRes:       " + strconv.Itoa(desktop.HorizontalRes()) + " dpi\n"
    text += "    VerticalRes:           " + strconv.Itoa(desktop.VerticalRes()) + " dpi\n\n"

    text += "Screen Available :" + "\n"
    text += "    ClientWidth:        " + strconv.Itoa(metrics.DpToPx(ui.GoDpr, desktop.ClientWidth())) + " px\n"  // * ui.GoDpr)) + "\n"
    text += "    ClientHeight:         " + strconv.Itoa(metrics.DpToPx(ui.GoDpr, desktop.ClientHeight())) + " px\n\n"    // * ui.GoDpr)) + "\n"
    
    wX, wY := mainwin.Pos()
    wWidth, wHeight := mainwin.Size()
    text += "Window Geometry :" + "\n"
    text += "    WindowPos:     " + " (" + strconv.Itoa(wX) + ", " + strconv.Itoa(wY) + ")" + " dp\n"
    text += "    WindowSize:    " + " (" + strconv.Itoa(wWidth) + ", " + strconv.Itoa(wHeight) + ")" + " dp\n\n"
    
    cX, cY := mainwin.ClientPos()
    cWidth, cHeight := mainwin.ClientSize()
    text += "Window Client Geometry :" + "\n"
    text += "    ClientPos:     " + " (" + strconv.Itoa(cX) + ", " + strconv.Itoa(cY) + ")" + " dp\n"
    text += "    ClientSize:    " + " (" + strconv.Itoa(cWidth) + ", " + strconv.Itoa(cHeight) + ")" + " dp\n\n"

    text += "Window Geometry Screen Pixels:" + "\n"
    text += "    WindowPos:     " + " (" + strconv.Itoa(metrics.DpToPx(ui.GoDpr, wX)) + ", " + strconv.Itoa(metrics.DpToPx(ui.GoDpr, wY)) + ")" + " px\n"
    text += "    WindowSize:    " + " (" + strconv.Itoa(metrics.DpToPx(ui.GoDpr, wWidth)) + ", " + strconv.Itoa(metrics.DpToPx(ui.GoDpr, wHeight)) + ")" + " px\n"

    return text
}

func LaunchViewer(section string, content string) {
    // create viewer window
    viewer = ui.GoWindow("UtopiaGio: Reference Documentation")
    viewer.SetPos(100, 100)
    viewer.SetSize(1000, 600)
    viewer.SetLayoutStyle(ui.VFlexBoxLayout)
    viewer.Layout().SetPadding(10,10,10,10)
    /*page := */uireference.Page(viewer.Layout(), "UtopiaGio", section, content)
    // show the viewer window
    viewer.Show()
}

func LoadOverview() {
    LaunchViewer("GoWindowObj", "")
}

func UpdateWindowProperties() {
    lblWindowProperties.SetText(GetWindowProperties())
    mainwin.Refresh()
}

