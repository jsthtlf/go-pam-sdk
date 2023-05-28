package service

// API: Основные запросы
const (
	UserProfileURL       = "/api/v1/users/profile/"                   // Информация о пользователе
	TerminalRegisterURL  = "/api/v1/terminal/terminal-registrations/" // Регистрация терминала
	TerminalConfigURL    = "/api/v1/terminal/terminals/config/"       // Настройки терминала
	TerminalHeartBeatURL = "/api/v1/terminal/terminals/status/"       // Состояние терминала
)

// API: Аутентификация
const (
	TokenAssetURL      = "/api/v1/authentication/connection-token/%s/"         // Получение токена подключения к ресурсу, %s - токен
	UserTokenAuthURL   = "/api/v1/authentication/tokens/"                      // Получение токена пользователя для аутентификации
	UserConfirmAuthURL = "/api/v1/authentication/login-confirm-ticket/status/" // Использование токена пользователя
	AuthMFASelectURL   = "/api/v1/authentication/mfa/select/"                  // Выбор MFA

	TokenAuthInfoURL = "/api/v1/authentication/connection-token/secret-info/detail/" // Получение информации о токене подключения
	TokenRenewalURL  = "/api/v1/authentication/connection-token/renewal/"            // Сброс токена подключения
)

// API: Взаимодействие с информацией сессии
const (
	SessionListURL      = "/api/v1/terminal/sessions/"               // Создание сессии
	SessionDetailURL    = "/api/v1/terminal/sessions/%s/"            // Взаимодействие с сессией, %s - токен сессии
	SessionReplayURL    = "/api/v1/terminal/sessions/%s/replay/"     // Загрузка записи сессии, %s - токен сессии
	SessionCommandURL   = "/api/v1/terminal/commands/"               // Загрузка выполненной команды в сессии
	FinishTaskURL       = "/api/v1/terminal/tasks/%s/"               // Завершение задачи
	JoinRoomValidateURL = "/api/v1/terminal/sessions/join/validate/" // Проверка доступа подключения к сессии
	FTPLogListURL       = "/api/v1/audits/ftp-logs/"                 // Загрузка FTP журнала
)

// API: Взаимодействие с правами доступа
const (
	UserPermsAssetsURL                 = "/api/v1/perms/users/%s/assets/"                          // Получение ресурсов пользователя
	UserPermsNodesListURL              = "/api/v1/perms/users/%s/nodes/"                           // Получение узлов пользователя
	UserPermsNodeAssetsListURL         = "/api/v1/perms/users/%s/nodes/%s/assets/"                 // Получение ресурсов узла пользователя
	UserPermsNodeTreeWithAssetURL      = "/api/v1/perms/users/%s/nodes/children-with-assets/tree/" // Получение дерево узлов ресурсов пользователя
	UserPermsApplicationsURL           = "/api/v1/perms/users/%s/applications/?type=%s"            // Получение приложений пользователя по типу
	UserPermsAssetSystemUsersURL       = "/api/v1/perms/users/%s/assets/%s/system-users/"          // Получение системного пользователя
	UserPermsApplicationSystemUsersURL = "/api/v1/perms/users/%s/applications/%s/system-users/"    // Получение системных пользователей, которые имеют доступ к приложению
	ValidateUserAssetPermissionURL     = "/api/v1/perms/asset-permissions/user/validate/"          // Получение информации о доступе к ресурсу
	ValidateApplicationPermissionURL   = "/api/v1/perms/application-permissions/user/validate/"    // Получение информации о доступе к приложению

	UserPermsDatabaseURL = "/api/v1/perms/users/%s/applications/?category=db&type__in=%s" // Получение информации о доступе к базам данных
)

// API: Взаимодействие с аутентификацией системного пользователя
const (
	SystemUserAuthURL      = "/api/v1/assets/system-users/%s/auth-info/"                 // Получить данные аутентификации
	SystemUserAppAuthURL   = "/api/v1/assets/system-users/%s/applications/%s/auth-info/" // Получить данные аутентификации в приложении
	SystemUserAssetAuthURL = "/api/v1/assets/system-users/%s/assets/%s/auth-info/"       // Получить данные аутентификации в ресурсе
)

// API: Получение информации об отдельных объектах
const (
	UserListURL          = "/api/v1/users/users/"                  // Список пользователей
	UserDetailURL        = "/api/v1/users/users/%s/"               // Информация о пользователе
	AssetDetailURL       = "/api/v1/assets/assets/%s/"             // Информация о ресурсе
	AssetPlatFormURL     = "/api/v1/assets/assets/%s/platform/"    // Информация о платформе ресурса
	SystemUserDetailURL  = "/api/v1/assets/system-users/%s/"       // Информация о системном пользователе
	ApplicationDetailURL = "/api/v1/applications/applications/%s/" // Информация о приложении

	SystemUserCmdFilterRulesListURL = "/api/v1/assets/system-users/%s/cmd-filter-rules/" // Информация о правилах фильтрации команд

	CommandFilterRulesListURL = "/api/v1/assets/cmd-filter-rules/" // Список правил фильтрации команд

	DomainDetailWithGateways = "/api/v1/assets/domains/%s/?gateway=1" // Информация о домене
)

const (
	NotificationCommandURL = "/api/v1/terminal/commands/insecure-command/" //
)

const (
	PermissionURL = "/api/v1/perms/asset-permissions/user/actions/" // Получение информации о правах доступах к ресурсу

	RemoteAPPURL = "/api/v1/applications/remote-apps/%s/connection-info/" // Получение информации об удаленном приложении
)

const (
	AssetLoginConfirmURL = "/api/v1/acls/login-asset/check/"
)

// 命令复核

const (
	CommandConfirmURL = "/api/v1/assets/cmd-filters/command-confirm/"
)

const (
	ShareCreateURL        = "/api/v1/terminal/session-sharings/"
	ShareSessionJoinURL   = "/api/v1/terminal/session-join-records/"
	ShareSessionFinishURL = "/api/v1/terminal/session-join-records/%s/finished/"
)

const (
	PublicSettingURL = "/api/v1/settings/public/"
)

const (
	TicketSessionURL = "/api/v1/tickets/ticket-session-relation/"
)
