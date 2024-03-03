package view

type Key struct {
	Label  string
	Sound  string
	Events map[string]func()
}
type Keyboard struct {
	Keys []Key
}

func CreateKeyboard() *Keyboard {
	// Define the keys
	keys := []Key{
		{"C", "sound_C", nil},
		{"C#", "sound_C#", nil},
		{"D", "sound_D", nil},
		{"D#", "sound_D#", nil},
		{"E", "sound_E", nil},
		{"E#", "sound_E#", nil},
		{"F", "sound_F", nil},
		{"F#", "sound_F#", nil},
		{"G", "sound_G", nil},
		{"G#", "sound_G#", nil},
		{"A", "sound_A", nil},
		{"A#", "sound_A#", nil},
		{"B", "sound_B", nil},
		{"B#", "sound_B#", nil},
	}

	return &Keyboard{Keys: keys}
}

func (k *Keyboard) PlayKey(label string) {
	for _, key := range k.Keys {
		if key.Label == label && key.Events["play"] != nil {
			key.Events["play"]()
			break
		}
	}
}
