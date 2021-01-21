package types

// GuacConnection base guacamole connection info
type GuacConnection struct {
	Name              string                   `json:"name"`
	Identifier        string                   `json:"identifier,omitempty"`
	ParentIdentifier  string                   `json:"parentIdentifier"`
	Protocol          string                   `json:"protocol"`
	Attributes        GuacConnectionAttributes `json:"attributes"`
	Parameters        GuacConnectionParameters `json:"parameters"`
	ActiveConnections int                      `json:"activeConnections,omitempty"`
}

// GuacConnectionAttributes guacd attributes
type GuacConnectionAttributes struct {
	GuacdEncryption       string `json:"guacd-encryption"`
	FailoverOnly          string `json:"failover-only"`
	Weight                string `json:"weight"`
	MaxConnections        string `json:"max-connections"`
	GuacdHostname         string `json:"guacd-hostname,omitempty"`
	GuacdPort             string `json:"guacd-port"`
	MaxConnectionsPerUser string `json:"max-connections-per-user"`
}

// GuacConnectionParameters defines guacamole connection parameters
type GuacConnectionParameters struct {
	/*** Network ***/
	// All
	Hostname string `json:"hostname,omitempty"`
	Port     string `json:"port,omitempty"`
	// SSH
	PublicHostKey string `json:"host-key,omitempty"`
	// Kubernetes
	UseSSL string `json:"use-ssl,omitempty"`
	CACert string `json:"ca-cert,omitempty"`

	/*** Authentication ***/
	// SSH
	Username   string `json:"username,omitempty"`
	Password   string `json:"password,omitempty"`
	PrivateKey string `json:"private-key,omitempty"`
	Passphrase string `json:"passphrase,omitempty"`
	// RDP
	Domain                string `json:"domain,omitempty"`
	Security              string `json:"security,omitempty"`
	DisableAuthentication string `json:"disable-auth,omitempty"`
	IgnoreCert            string `json:"ignore-cert,omitempty"`
	// Telnet
	UsernameRegex     string `json:"username-regex,omitempty"`
	PasswordRegex     string `json:"password-regex,omitempty"`
	LoginSuccessRegex string `json:"login-success-regex,omitempty"`
	LoginFailureRegex string `json:"login-failure-regex,omitempty"`
	// Kubernetes
	ClientCert string `json:"client-cert,omitempty"`
	ClientKey  string `json:"client-key,omitempty"`

	/*** Display ***/
	// SSH
	ColorScheme string `json:"color-scheme,omitempty"`
	FontName    string `json:"font-name,omitempty"`
	FontSize    string `json:"font-size,omitempty"`
	Scrollback  string `json:"scrollback,omitempty"`
	ReadOnly    string `json:"read-only,omitempty"`
	// RDP
	Width        string `json:"width,omitempty"`
	Height       string `json:"height,omitempty"`
	DPI          string `json:"dpi,omitempty"`
	ColorDepth   string `json:"color-depth,omitempty"`
	ResizeMethod string `json:"resize-method,omitempty"`
	// VNC
	SwapRedBlue string `json:"swap-red-blue,omitempty"`
	Cursor      string `json:"cursor,omitempty"`

	/*** Clipboard ***/
	DisableCopy  string `json:"disable-copy,omitempty"`
	DisablePaste string `json:"disable-paste,omitempty"`
	// VNC
	ClipboardEncoding string `json:"clipboard-encoding,omitempty"`

	/*** Session Environment, Basic Settings ***/
	ExecuteCommand          string `json:"command,omitempty"`
	Locale                  string `json:"locale,omitempty"`
	Timezone                string `json:"timezone,omitempty"`
	ServerKeepaliveInterval string `json:"server-alive-interval,omitempty"`
	InitialProgram          string `json:"initial-program,omitempty"`
	ClientName              string `json:"client-name,omitempty"`
	KeyboardLayout          string `json:"server-layout,omitempty"`
	AdministratorConsole    string `json:"console,omitempty"`

	/*** Terminal Behavior ***/
	Backspace    string `json:"backspace,omitempty"`
	TerminalType string `json:"terminal-type,omitempty"`

	/*** Typescript (Text Session Recording) ***/
	TypescriptPath       string `json:"typescript-path,omitempty"`
	TypescriptName       string `json:"typescript-name,omitempty"`
	CreateTypescriptPath string `json:"create-typescript-path,omitempty"`

	/*** Screen Recording ***/
	RecordingPath          string `json:"recording-path,omitempty"`
	RecordingName          string `json:"recording-name,omitempty"`
	RecordingExcludeOutput string `json:"recording-exclude-output,omitempty"`
	RecordingExcludeMouse  string `json:"recording-exclude-mouse,omitempty"`
	RecordingIncludeKeys   string `json:"recording-include-keys,omitempty"`
	CreateRecordingPath    string `json:"create-recording-path,omitempty"`

	/*** SFTP ***/
	EnableSFTP              string `json:"enable-sftp,omitempty"`
	SFTPRootDirectory       string `json:"sftp-root-directory,omitempty"`
	SFTPDisableFileDownload string `json:"sftp-disable-download,omitempty"`
	SFTPDisableFileUpload   string `json:"sftp-disable-upload,omitempty"`
	SFTPHostname            string `json:"sftp-hostname,omitempty"`
	SFTPPort                string `json:"sftp-port,omitempty"`
	SFTPHostKey             string `json:"sftp-host-key,omitempty"`
	SFTPUsername            string `json:"sftp-username,omitempty"`
	SFTPPassword            string `json:"sftp-password,omitempty"`
	SFTPPrivateKey          string `json:"sftp-private-key,omitempty"`
	SFTPPassphrase          string `json:"sftp-passphrase,omitempty"`
	SFTPUploadDirectory     string `json:"sftp-directory,omitempty"`
	SFTPKeepAliveInterval   string `json:"sftp-server-alive-interval,omitempty"`

	/*** Wake-on-LAN ***/
	WOLSendPacket       string `json:"wol-send-packet,omitempty"`
	WOLMacAddress       string `json:"wol-mac-addr,omitempty"`
	WOLBroadcastAddress string `json:"wol-broadcast-addr,omitempty"`
	WOLBootWaitTime     string `json:"wol-wait-time,omitempty"`

	/*** RDP - Remote Desktop Gateway ***/
	GatewayHostname string `json:"gateway-hostname,omitempty"`
	GatewayPort     string `json:"gateway-port,omitempty"`
	GatewayUsername string `json:"gateway-username,omitempty"`
	GatewayPassword string `json:"gateway-password,omitempty"`
	GatewayDomain   string `json:"gateway-domain,omitempty"`

	/*** RDP - Device Redirection ***/
	ConsoleAudio        string `json:"console-audio,omitempty"`
	DisableAudio        string `json:"disable-audio,omitempty"`
	EnableAudioInput    string `json:"enable-audio-input,omitempty"`
	EnablePrinting      string `json:"enable-printing,omitempty"`
	PrinterName         string `json:"printer-name,omitempty"`
	EnableDrive         string `json:"enable-drive,omitempty"`
	DriveName           string `json:"drive-name,omitempty"`
	DisableFileDownload string `json:"disable-download,omitempty"`
	DisableFileUpload   string `json:"disable-upload,omitempty"`
	DrivePath           string `json:"drive-path,omitempty"`
	CreateDrivePath     string `json:"create-drive-path,omitempty"`
	StaticChannels      string `json:"static-channels,omitempty"`

	/*** RDP - Performance ***/
	EnableWallpaper          string `json:"enable-wallpaper,omitempty"`
	EnableTheming            string `json:"enable-theming,omitempty"`
	EnableFontSmoothing      string `json:"enable-font-smoothing,omitempty"`
	EnableFullWindowDrag     string `json:"enable-full-window-drag,omitempty"`
	EnableDesktopComposition string `json:"enable-desktop-composition,omitempty"`
	EnableMenuAnimations     string `json:"enable-menu-animations,omitempty"`
	DisableBitmapCaching     string `json:"disable-bitmap-caching,omitempty"`
	DisableOffscreenCaching  string `json:"disable-offscreen-caching,omitempty"`
	DisableGlyphCaching      string `json:"disable-glyph-caching,omitempty"`

	/*** RDP - RemoteApp ***/
	RemoteApp                 string `json:"remote-app,omitempty"`
	RemoteAppWorkingDirectory string `json:"remote-app-dir,omitempty"`
	RemoteAppParameters       string `json:"remote-app-args,omitempty"`

	/*** RDP - Preconnection PDU/Hyper-V ***/
	PreconnectionID   string `json:"preconnection-id,omitempty"`
	PreconnectionBLOB string `json:"preconnection-blob,omitempty"`

	/*** RDP - Load Balancing ***/
	LoadBalanceInfo string `json:"load-balance-info,omitempty"`

	/*** VNC Repeater ***/
	DestinationHost string `json:"dest-host,omitempty"`
	DestinationPort string `json:"dest-port,omitempty"`

	/*** Audio ***/
	// VNC
	AudioServerName string `json:"audio-servername,omitempty"`
	EnableAudio     string `json:"enable-audio,omitempty"`

	/*** Container ***/
	Container string `json:"container,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Pod       string `json:"pod,omitempty"`
}

// ValidEncryptionTypes returns a list of valid encryption types
func (GuacConnectionAttributes) ValidEncryptionTypes() []string {
	return []string{
		"none",
		"ssl",
	}
}

// ValidColorSchemes returns a list of valid color schemes
func (GuacConnectionParameters) ValidColorSchemes() []string {
	return []string{
		"black-white",
		"gray-black",
		"green-black",
		"white-black",
	}
}

// ValidFontSizes returns a list of valid font sizes
func (GuacConnectionParameters) ValidFontSizes() []string {
	return []string{
		"8",
		"9",
		"10",
		"11",
		"12",
		"14",
		"18",
		"24",
		"30",
		"36",
		"48",
		"60",
		"72",
		"96",
	}
}

// ValidBackspaceCodes returns a list of valid backspace codes
func (GuacConnectionParameters) ValidBackspaceCodes() []string {
	return []string{
		"127",
		"8",
	}
}

// ValidTerminalTypes returns a list of valid terminaly types
func (GuacConnectionParameters) ValidTerminalTypes() []string {
	return []string{
		"ansi",
		"linux",
		"vt100",
		"vt220",
		"xterm",
		"xterm-25color",
	}
}

// ValidCursors returns a list of valid cursors
func (GuacConnectionParameters) ValidCursors() []string {
	return []string{
		"local",
		"remote",
	}
}

// ValidColorDepths returns a list of valid color depths
func (GuacConnectionParameters) ValidColorDepths() []string {
	return []string{
		"8",
		"16",
		"24",
		"32",
	}
}

// ValidClipboardEncodings returns a list of valid clipboard encodings
func (GuacConnectionParameters) ValidClipboardEncodings() []string {
	return []string{
		"CP1252",
		"ISO8859-1",
		"UTF-16",
		"UTF-8",
	}
}

// ValidSecurityModes returns a list of valid security modes
func (GuacConnectionParameters) ValidSecurityModes() []string {
	return []string{
		"any",
		"nla",
		"rdp",
		"tls",
		"vmconnect",
	}
}

// ValidKeyboardLayouts returns a list of valid keyboard layouts
func (GuacConnectionParameters) ValidKeyboardLayouts() []string {
	return []string{
		"da-dk-qwerty",
		"de-ch-qwertz",
		"de-de-qwertz",
		"en-gb-qwerty",
		"en-us-qwerty",
		"es-es-qwerty",
		"es-latam-qwerty",
		"failsafe",
		"fr-be-azerty",
		"fr-ch-qwertz",
		"fr-fr-azerty",
		"hu-hu-qwertz",
		"it-it-qwerty",
		"ja-jp-qwerty",
		"pt-br-qwerty",
		"sv-se-qwerty",
		"tr-tr-qwerty",
	}
}

// ValidResizeMethods returns a list of valid resize methods
func (GuacConnectionParameters) ValidResizeMethods() []string {
	return []string{
		"display-update",
		"reconnect",
	}
}
