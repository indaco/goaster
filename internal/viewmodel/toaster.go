package viewmodel

type ToastViewModel struct {
	Message string
	Level   string
	Icon    string
}

type ToasterViewModel struct {
	Variant     string
	Border      bool
	Rounded     bool
	ShowIcon    bool
	Button      bool
	AutoDismiss bool
	Animation   bool
	ProgressBar bool
	Position    string
	Toasts      []ToastViewModel
}
