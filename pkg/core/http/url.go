package http

// UserProvider urls
const (
	UrlUserProfile = "/api/v1/users/profile/"
	UrlUserList    = "/api/v1/users/users/"
	UrlUserDetail  = "/api/v1/users/users/%s/"
)

// TerminalProvider urls
const (
	UrlTerminalRegister  = "/api/v1/terminal/terminal-registrations/"
	UrlTerminalConfig    = "/api/v1/terminal/terminals/config/"
	UrlTerminalHeartBeat = "/api/v1/terminal/terminals/status/"
	UrlTerminalTask      = "/api/v1/terminal/tasks/%s/"
	UrlPublicSetting     = "/api/v1/settings/public/"
)

// AssetProvider urls
const (
	UrlAssetDetail         = "/api/v1/assets/assets/%s/"
	UrlAssetPlatFormDetail = "/api/v1/assets/assets/%s/platform/"
	UrlAssetDomainDetail   = "/api/v1/assets/domains/%s/?gateway=1"
)

// ApplicationProvider urls
const (
	UrlAppDetail       = "/api/v1/applications/applications/%s/"
	UrlRemoteAppDetail = "/api/v1/applications/remote-apps/%s/connection-info/"
)

// SystemUserProvider urls
const (
	UrlSystemUserDetail        = "/api/v1/assets/system-users/%s/"
	UrlSystemUserAppAuth       = "/api/v1/assets/system-users/%s/applications/%s/auth-info/"
	UrlUserPermsAppSystemUsers = "/api/v1/perms/users/%s/applications/%s/system-users/"
	UrlSystemUserAuth          = "/api/v1/assets/system-users/%s/auth-info/"
	UrlSystemUserAssetAuth     = "/api/v1/assets/system-users/%s/assets/%s/auth-info/"
)

// AuditProvider urls
const (
	UrlFtpLogList                   = "/api/v1/audits/ftp-logs/"
	UrlSessionCommand               = "/api/v1/terminal/commands/"
	UrlSessionNotifyCommand         = "/api/v1/terminal/commands/insecure-command/"
	UrlSystemUserCmdFilterRulesList = "/api/v1/assets/system-users/%s/cmd-filter-rules/"
	UrlCmdFilterRulesList           = "/api/v1/assets/cmd-filter-rules/"
)

// SessionProvider urls
const (
	UrlSessionList           = "/api/v1/terminal/sessions/"
	UrlSessionReplay         = "/api/v1/terminal/sessions/%s/replay/"
	UrlSessionDetail         = "/api/v1/terminal/sessions/%s/"
	UrlSessionTicketRelation = "/api/v1/tickets/ticket-session-relation/"
)

// PermissionProvider urls
const (
	UrlAssetPermsDetail           = "/api/v1/perms/asset-permissions/user/actions/"
	UrlValidateAssetPerms         = "/api/v1/perms/asset-permissions/user/validate/"
	UrlValidateAppPerms           = "/api/v1/perms/application-permissions/user/validate/"
	UrlValidateJoinRoom           = "/api/v1/terminal/sessions/join/validate/"
	UrlUserPermsAssets            = "/api/v1/perms/users/%s/assets/"
	UrlUserPermsAssetSystemUsers  = "/api/v1/perms/users/%s/assets/%s/system-users/"
	UrlUserPermsApplications      = "/api/v1/perms/users/%s/applications/?type=%s"
	UrlUserPermsDatabase          = "/api/v1/perms/users/%s/applications/?category=db&type__in=%s"
	UrlUserPermsNodesList         = "/api/v1/perms/users/%s/nodes/"
	UrlUserPermsNodeAssetsList    = "/api/v1/perms/users/%s/nodes/%s/assets/"
	UrlUserPermsNodeTreeWithAsset = "/api/v1/perms/users/%s/nodes/children-with-assets/tree/"
)

// ShareRoomProvider urls
const (
	UrlShareCreate        = "/api/v1/terminal/session-sharings/"
	UrlShareSessionJoin   = "/api/v1/terminal/session-join-records/"
	UrlShareSessionFinish = "/api/v1/terminal/session-join-records/%s/finished/"
)

// TicketProvider url
const (
	UrlCommandConfirm    = "/api/v1/assets/cmd-filters/command-confirm/"
	UrlAssetLoginConfirm = "/api/v1/acls/login-asset/check/"
)

// TokenProvider url
const (
	UrlTokenAsset    = "/api/v1/authentication/connection-token/%s/"
	UrlTokenAuthInfo = "/api/v1/authentication/connection-token/secret-info/detail/"
	UrlTokenRenewal  = "/api/v1/authentication/connection-token/renewal/"
)
