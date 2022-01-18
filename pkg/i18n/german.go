package i18n

func germanSet() TranslationSet {
	return TranslationSet{
		PruningStatus:              "zerstören",
		RemovingStatus:             "entfernen",
		RestartingStatus:           "neustarten",
		StoppingStatus:             "anhalten",
		RunningCustomCommandStatus: "führt benutzerdefinierten Befehl aus",

		RunningSubprocess:                          "Unterprozess ausführen",
		NoViewMachingNewLineFocusedSwitchStatement: "No view matching newLineFocused switch statement",

		ErrorOccurred:                     "Es ist ein Fehler aufgetreten! Bitte erstelle ein Issue hier: https://github.com/jesseduffield/lazydocker/issues",
		ConnectionFailed:                  "Verbindung zum Docker Client fehlgeschlagen. Du musst ggf. den Docker Client neustarten.",
		UnattachableContainerError:        "Der Container bietet keine Unterstützung für das Anbinden. Du musst den Dienst entweder mit der '-it' Flagge benutzen oder `stdin_open: true, tty: true` in der docker-compose.yml Datei setzen.",
		CannotAttachStoppedContainerError: "Du kannst keinen angehaltenen Container anbinden. Du musst ihn erst starten (was du tun kannst, indem du 'r' drückst), (ja, ich bin zu faul um das zu automatisieren) (aber ist schon cool, dass ich so eine Konversation durch eine Fehlermeldung mit dir führen kann)",
		CannotAccessDockerSocketError:     "Kann nicht auf den Socket zugreifen: unix:///var/run/docker.sock\nFühre lazydocker als root aus oder lese https://docs.docker.com/install/linux/linux-postinstall/",

		Donate:  "Spenden",
		Confirm: "Bestätigen",

		Return:             "zurück",
		FocusMain:          "fokussieren aufs Hauptpanel",
		Navigate:           "navigieren",
		Execute:            "ausführen",
		Close:              "schließen",
		Menu:               "Menü",
		Scroll:             "scrollen",
		OpenConfig:         "öffne lazydocker Konfiguration",
		EditConfig:         "bearbeite lazydocker Konfiguration",
		Cancel:             "abbrechen",
		Remove:             "entfernen",
		ForceRemove:        "Entfernen erzwingen",
		RemoveWithVolumes:  "entferne mit Volumes",
		RemoveService:      "entferne Container",
		Stop:               "anhalten",
		Restart:            "neustarten",
		Rebuild:            "neubauen",
		Recreate:           "neuerstellen",
		PreviousContext:    "vorheriges Tab",
		NextContext:        "nächstes Tab",
		Attach:             "anbinden",
		ViewLogs:           "zeige Protokolle",
		RemoveImage:        "entferne Image",
		RemoveVolume:       "entferne Volume",
		RemoveWithoutPrune: "entfernen, ohne die unmarkierten Eltern zu entfernen",
		PruneContainers:    "entferne verlassene Container",
		PruneVolumes:       "entferne unbenutzte Volumes",
		PruneImages:        "entferne unbenutzte Images",
		ViewRestartOptions: "zeige Neustartoptionen",
		RunCustomCommand:   "führe vordefinierten benutzerdefinierten Befehl aus",

		GlobalTitle:               "Global",
		MainTitle:                 "Haupt",
		ProjectTitle:              "Projekt",
		ServicesTitle:             "Dienste",
		ContainersTitle:           "Container",
		StandaloneContainersTitle: "Alleinstehende Container",
		ImagesTitle:               "Images",
		VolumesTitle:              "Volumes",
		CustomCommandTitle:        "Benutzerdefinierter Befehl",
		ErrorTitle:                "Fehler",
		LogsTitle:                 "Protokoll",
		ConfigTitle:               "Konfiguration",
		EnvTitle:                  "Env",
		DockerComposeConfigTitle:  "Docker-Compose Konfiguration",
		TopTitle:                  "Top",
		StatsTitle:                "Statistiken",
		CreditsTitle:              "Über Uns",
		ContainerConfigTitle:      "Container Konfiguration",
		ContainerEnvTitle:         "Container Env",
		NothingToDisplay:          "Nothing to display",
		CannotDisplayEnvVariables:  "Something went wrong while displaying environment variables",

		NoContainers: "Keine Container",
		NoContainer:  "Kein Container",
		NoImages:     "Keine Images",
		NoVolumes:    "Keine Volumes",

		ConfirmQuit:                "Bist du dir sicher, dass du verlassen möchtest?",
		MustForceToRemoveContainer: "Du kannst keinen Container entfernen, der noch ausgeführt wird außer du erzwingst es. Möchtest du es erzwingen?",
		NotEnoughSpace:             "Nicht genug Platz um die Panel darzustellen",
		ConfirmPruneImages:         "Bist du dir sicher, dass du alle unbenutzten Images entfernen möchtest?",
		ConfirmPruneContainers:     "Bist du dir sicher, dass du alle angehaltenen Container entfernen möchtes?",
		ConfirmPruneVolumes:        "Bist du dir sicher, dass du alle unbenutzen Volumes entfernen möchtest?",
		StopService:                "Bist du dir sicher, dass du den Dienst dieses Containers anhalten möchtest?",
		StopContainer:              "Bist du dir sicher, dass du den Container anhalten möchtest?",
		PressEnterToReturn:         "Drücke Eingabe um zu lazydocker zurückzukehren. (Diese Nachfrage kann in Deiner Konfiguration deaktiviert werden, indem du folgenden Wert setzt: `gui.returnImmediately: true`)",
	}
}
