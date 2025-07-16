package ui

import (
	"math"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"stagfoo.pennyblack/app/files"
)

func List(dbBooks []files.Book, selectedIndex int) *widget.List {
	var data []string
	for _, book := range dbBooks {
		data = append(data, book.Title)
	}
	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		},
	)
	// Set initial selection
	list.Select(selectedIndex)
	return list
}

type PaginatedScroll struct {
	*container.Scroll
	pageHeight  float32
	isAnimating bool
}

func NewPaginatedScroll(content fyne.CanvasObject, pageHeight float32) *PaginatedScroll {
	scroll := container.NewVScroll(content)

	ps := &PaginatedScroll{
		Scroll:      scroll,
		pageHeight:  pageHeight,
		isAnimating: false,
	}

	return ps
}

func (ps *PaginatedScroll) Scrolled(ev *fyne.ScrollEvent) {
	if ps.isAnimating {
		return // Ignore scroll events during animation
	}

	// Get current scroll position
	currentY := ps.Offset.Y

	// Calculate which page we should snap to
	var targetPage int
	if ev.Scrolled.DY > 0 {
		// Scrolling down
		targetPage = int(math.Ceil(float64(currentY / ps.pageHeight)))
	} else {
		// Scrolling up
		targetPage = int(math.Floor(float64(currentY / ps.pageHeight)))
	}

	// Calculate target position
	targetY := float32(targetPage) * ps.pageHeight

	// Ensure we don't scroll past content bounds
	maxScroll := ps.Content.Size().Height - ps.Size().Height
	if targetY > maxScroll {
		targetY = maxScroll
	}
	if targetY < 0 {
		targetY = 0
	}

	// Animate to target position
	ps.animateToPosition(targetY)
}

func (ps *PaginatedScroll) animateToPosition(targetY float32) {
	if ps.isAnimating {
		return
	}

	ps.isAnimating = true
	startY := ps.Offset.Y
	duration := 300 * time.Millisecond
	startTime := time.Now()

	// Simple easing function
	easeOutQuart := func(t float32) float32 {
		t = t - 1
		return 1 - t*t*t*t
	}

	go func() {
		ticker := time.NewTicker(16 * time.Millisecond) // ~60fps
		defer ticker.Stop()

		for range ticker.C {
			elapsed := time.Since(startTime)
			if elapsed >= duration {
				// Animation complete
				ps.Offset = fyne.NewPos(0, targetY)
				ps.Refresh()
				ps.isAnimating = false
				return
			}

			// Calculate interpolated position
			progress := float32(elapsed) / float32(duration)
			easedProgress := easeOutQuart(progress)
			currentY := startY + (targetY-startY)*easedProgress

			// Update scroll position
			ps.Offset = fyne.NewPos(0, currentY)
			ps.Refresh()
		}
	}()
}
