// Copyright 2020 NGR Softlab
//
package errorLib

import "errors"

// Global Errors store (with default error texts)
var GlobalErrors = ErrMap{
	//500 (for updater)
	errorFailedMakeReport: errors.New("code: 500, message: failed make report"),

	//rights
	errorForbidden: errors.New("code: 403, message: forbidden"),

	//auth
	errorUnauthorized:  errors.New("code: 401, message: unauthorized"),
	errorTechnicalUser: errors.New("code: 400, message: user is technical"),

	//actions
	errorInvalidSysAction: errors.New("code: 400, message: invalid action with system data"),

	//database
	errorBadDatabaseConnection:  errors.New("code: 400, message: bad database connection"),
	errorBadDatabaseScheme:      errors.New("code: 400, message: bad database scheme"),
	errorBadDatabaseOperation:   errors.New("code: 400, message: bad query syntax or connection"),
	errorBadDatabaseSelect:      errors.New("code: 400, message: bad select from database"),
	errorBadDatabaseInsert:      errors.New("code: 400, message: bad insertion to database"),
	errorBadDatabaseUpdate:      errors.New("code: 400, message: bad updation in database"),
	errorBadDatabaseDelete:      errors.New("code: 400, message: data's not deleted from database"),
	errorBadDatabaseScan:        errors.New("code: 400, message: bad scan from database"),
	errorBadDatabaseTransaction: errors.New("code: 400, message: bad database transaction"),

	//elements
	errorDuplicateName:    errors.New("code: 400, message: duplicate element name"),
	errorElementNotFound:  errors.New("code: 400, message: element not found"),
	errorNotElementAuthor: errors.New("code: 400, message: only author can change element flags"),
	errorCanNotUnpublish:  errors.New("code: 400, message: can not unpublish element"),

	//queries
	errorNonSelectQuery:  errors.New("code: 400, message: can not use modifying query"),
	errorForbiddenQuery:  errors.New("code: 400, message: forbidden or table doesn't exist"),
	errorWhileCreateView: errors.New("code: 400, message: error while create view"),

	//taskplan
	errorNoSchedule:     errors.New("code: 400, message: schedule not found"),
	errorFailedSendTask: errors.New("code: 400, message: failed send task to executor"),

	//settings
	errorBadDatabaseSave:     errors.New("code: 400, message: can not save db settings"),
	errorBadADSave:           errors.New("code: 400, message: can not save domain controller settings"),
	errorBadAuthData:         errors.New("code: 400, message: invalid login or password"),
	errorBadIpOrPort:         errors.New("code: 400, message: invalid ip or port for connection"),
	errorBadBaseDn:           errors.New("code: 400, message: invalid base_dn search string"),
	errorBadNodeAddition:     errors.New("code: 400, message: bad nodes settings"),
	errorBadNodeNumberOrType: errors.New("code: 400, message: no such nodes type or node number"),
	errorBadSettingsRead:     errors.New("code: 400, message: can not read current settings"),
	errorBadSettingsWrite:    errors.New("code: 400, message: can not write new settings"),
	errorNoServices:          errors.New("code: 400, message: no special services"),

	//data
	errorBadUnmarshal: errors.New("code: 400, message: can not unmarshal request data"),
	errorBadParams:    errors.New("code: 400, message: bad params"),
	errorBadUrlParam:  errors.New("code: 400, message: can not parse url param"),
	errorSameData:     errors.New("code: 400, message: same data"),
	errorNoData:       errors.New("code: 400, message: no data"),
	errorBadData:      errors.New("code: 400, message: bad data"),
	errorInvalidData:  errors.New("code: 400, message: invalid data"),

	//files
	errorFileTooLarge:    errors.New("code: 400, message: file too large"),
	errorInvalidFileData: errors.New("code: 400, message: invalid file data"),
	errorFileRead:        errors.New("code: 400, message: bad file reading"),
	errorFileWrite:       errors.New("code: 400, message: bad file writing"),
	errorFileRemove:      errors.New("code: 400, message: bad file removing"),
	errorFileCreate:      errors.New("code: 400, message: bad file create"),
	errorFileNotFound:    errors.New("code: 400, message: file not found"),

	//custom
	errorConnectionTimeout: errors.New("code: 400, message: connection timeout"),
	errorBeingDeployed:     errors.New("code: 400, message: configs are being deployed now"),
	errorConnection:        errors.New("code: 400, message: bad connection"),
	errorBadSshCommands:    errors.New("code: 400, message: bad ssh commands"),

	//scripts
	errorJobAlreadyStarted: errors.New("code: 400, message: job already started"),
	errorJobWasNotStarted:  errors.New("code: 400, message: job was not started"),
	errorScriptPipInstall:  errors.New("code: 400, message: python pip install error"),
}

// json tags for the future
type ErrMap struct {
	//500 (for updater)
	errorFailedMakeReport error `json:"err_failed_make_report"`

	//rights
	errorForbidden error `json:"err_forbidden"`

	//auth
	errorUnauthorized  error `json:"err_unauthorized"`
	errorTechnicalUser error `json:"err_tech_user"`

	//actions
	errorInvalidSysAction error `json:"err_invalid_sys_action"`

	//database
	errorBadDatabaseConnection  error `json:"err_bad_db_conn"`
	errorBadDatabaseScheme      error `json:"err_bad_db_scheme"`
	errorBadDatabaseOperation   error `json:"err_bad_db_operation"`
	errorBadDatabaseSelect      error `json:"err_bad_db_select"`
	errorBadDatabaseInsert      error `json:"err_bad_db_insert"`
	errorBadDatabaseUpdate      error `json:"err_bad_db_update"`
	errorBadDatabaseDelete      error `json:"err_bad_db_delete"`
	errorBadDatabaseScan        error `json:"err_bad_db_scan"`
	errorBadDatabaseTransaction error `json:"err_bad_db_tx"`

	//elements
	errorDuplicateName    error `json:"err_duplicate_name"`
	errorElementNotFound  error `json:"err_element_not_found"`
	errorNotElementAuthor error `json:"err_not_element_author"`
	errorCanNotUnpublish  error `json:"err_can_not_unpublish"`

	//queries
	errorNonSelectQuery  error `json:"err_non_select_query"`
	errorForbiddenQuery  error `json:"err_forbidden_query"`
	errorWhileCreateView error `json:"err_while_create_view"`

	// taskplan
	errorNoSchedule     error `json:"err_no_schedule"`
	errorFailedSendTask error `json:"err_failed_send_task"`

	// settings (in core)
	errorBadDatabaseSave     error `json:"err_bad_db_save"`
	errorBadADSave           error `json:"err_bad_ad_save"`
	errorBadAuthData         error `json:"err_bad_auth_data"`
	errorBadIpOrPort         error `json:"err_bad_ip_port"`
	errorBadBaseDn           error `json:"err_bad_base_dn"`
	errorBadNodeAddition     error `json:"err_bad_node_addition"`
	errorBadNodeNumberOrType error `json:"err_bad_node_num"`
	errorBadSettingsRead     error `json:"err_bad_settings_read"`
	errorBadSettingsWrite    error `json:"err_bad_settings_write"`
	errorNoServices          error `json:"err_no_services`

	//data
	errorBadUnmarshal error `json:"err_bad_unmarshal"`
	errorBadParams    error `json:"err_bad_params"`
	errorBadUrlParam  error `json:"err_bad_url_param"`
	errorSameData     error `json:"err_same_data"`
	errorNoData       error `json:"err_no_data"`
	errorBadData      error `json:"err_bad_data"`
	errorInvalidData  error `json:"err_invalid_data"`

	//files
	errorFileTooLarge    error `json:"err_file_too_large"`
	errorInvalidFileData error `json:"err_invalid_format"`
	errorFileRead        error `json:"err_bad_file_read"`
	errorFileWrite       error `json:"err_bad_file_write"`
	errorFileRemove      error `json:"err_bad_file_remove"`
	errorFileCreate      error `json:"err_bad_file_create"`
	errorFileNotFound    error `json:"err_file_not_found"`

	//custom
	errorConnectionTimeout error `json:"err_connection_timeout"`
	errorBeingDeployed     error `json:"err_being_deployed"`
	errorConnection        error `json:"err_connection"`
	errorBadSshCommands    error `json:"err_ssh_commands"`

	//scripts
	errorJobAlreadyStarted error `json:"err_job_already_started"`
	errorScriptPipInstall  error `json:"err_script_pip_install"`
	errorJobWasNotStarted  error `json:"err_job_was_not_started"`
}

func LoadErrorMap() error {
	// TODO: load to GlobalErrors from database (custom load)
	return nil
}

//////////////////////////////////////////////
//////////////////////////////////////////////
//////////////////////////////////////////////

//500 (for updater)
func (globErrs *ErrMap) ErrFailedMakeReportFile() error {
	return globErrs.errorFailedMakeReport
}

//rights (403)
func (globErrs *ErrMap) ErrForbidden() error {
	return globErrs.errorForbidden
}

//auth (401)
func (globErrs *ErrMap) ErrUnauthorized() error {
	return globErrs.errorUnauthorized
}

func (globErrs *ErrMap) ErrTechnicalUser() error {
	return globErrs.errorTechnicalUser
}

//actions
func (globErrs *ErrMap) ErrInvalidSysAction() error {
	return globErrs.errorInvalidSysAction
}

//database
func (globErrs *ErrMap) ErrBadDbConn() error {
	return globErrs.errorBadDatabaseConnection
}

func (globErrs *ErrMap) ErrBadDbScheme() error {
	return globErrs.errorBadDatabaseScheme
}

func (globErrs *ErrMap) ErrDbOperation() error {
	return globErrs.errorBadDatabaseOperation
}

func (globErrs *ErrMap) ErrDbSelect() error {
	return globErrs.errorBadDatabaseSelect
}

func (globErrs *ErrMap) ErrDbInsert() error {
	return globErrs.errorBadDatabaseInsert
}

func (globErrs *ErrMap) ErrDbUpdate() error {
	return globErrs.errorBadDatabaseUpdate
}

func (globErrs *ErrMap) ErrDbDelete() error {
	return globErrs.errorBadDatabaseDelete
}

func (globErrs *ErrMap) ErrDbScan() error {
	return globErrs.errorBadDatabaseScan
}

func (globErrs *ErrMap) ErrDbTx() error {
	return globErrs.errorBadDatabaseTransaction
}

//elements
func (globErrs *ErrMap) ErrDuplicateName() error {
	return globErrs.errorDuplicateName
}

func (globErrs *ErrMap) ErrElementNotFound() error {
	return globErrs.errorElementNotFound
}

func (globErrs *ErrMap) ErrNotElementAuthor() error {
	return globErrs.errorNotElementAuthor
}

func (globErrs *ErrMap) ErrCanNotUnpublish() error {
	return globErrs.errorCanNotUnpublish
}

//queries
func (globErrs *ErrMap) ErrNonSelectQuery() error {
	return globErrs.errorNonSelectQuery
}

func (globErrs *ErrMap) ErrForbiddenQuery() error {
	return globErrs.errorForbiddenQuery
}

func (globErrs *ErrMap) ErrWhileCreateView() error {
	return globErrs.errorWhileCreateView
}

//taskplan
func (globErrs *ErrMap) ErrNoSchedule() error {
	return globErrs.errorNoSchedule
}

func (globErrs *ErrMap) ErrFailedSendTask() error {
	return globErrs.errorFailedSendTask
}

//settings (in core)
func (globErrs *ErrMap) ErrBadDatabaseSave() error {
	return globErrs.errorBadDatabaseSave
}

func (globErrs *ErrMap) ErrBadADSave() error {
	return globErrs.errorBadADSave
}

func (globErrs *ErrMap) ErrBadAuthData() error {
	return globErrs.errorBadAuthData
}

func (globErrs *ErrMap) ErrBadIpOrPort() error {
	return globErrs.errorBadIpOrPort
}

func (globErrs *ErrMap) ErrBadBaseDn() error {
	return globErrs.errorBadBaseDn
}

func (globErrs *ErrMap) ErrBadNodeAdd() error {
	return globErrs.errorBadNodeAddition
}

func (globErrs *ErrMap) ErrBadNodes() error {
	return globErrs.errorBadNodeNumberOrType
}

func (globErrs *ErrMap) ErrBadSettsRead() error {
	return globErrs.errorBadSettingsRead
}

func (globErrs *ErrMap) ErrBadSettsWrite() error {
	return globErrs.errorBadSettingsWrite
}

func (globErrs *ErrMap) ErrNoServices() error {
	return globErrs.errorNoServices
}

//data
func (globErrs *ErrMap) ErrBadUnmarshal() error {
	return globErrs.errorBadUnmarshal
}

func (globErrs *ErrMap) ErrBadParams() error {
	return globErrs.errorBadParams
}

func (globErrs *ErrMap) ErrBadUrlParam() error {
	return globErrs.errorBadUrlParam
}

func (globErrs *ErrMap) ErrSameData() error {
	return globErrs.errorSameData
}

func (globErrs *ErrMap) ErrNoData() error {
	return globErrs.errorNoData
}

func (globErrs *ErrMap) ErrBadData() error {
	return globErrs.errorBadData
}

func (globErrs *ErrMap) ErrInvalidData() error {
	return globErrs.errorInvalidData
}

//files
func (globErrs *ErrMap) ErrFileTooLarge() error {
	return globErrs.errorFileTooLarge
}

func (globErrs *ErrMap) ErrInvalidFileData() error {
	return globErrs.errorInvalidFileData
}

func (globErrs *ErrMap) ErrFileRead() error {
	return globErrs.errorFileRead
}

func (globErrs *ErrMap) ErrFileWrite() error {
	return globErrs.errorFileWrite
}

func (globErrs *ErrMap) ErrFileRemove() error {
	return globErrs.errorFileRemove
}

func (globErrs *ErrMap) ErrFileCreate() error {
	return globErrs.errorFileCreate
}

func (globErrs *ErrMap) ErrFileNotFound() error {
	return globErrs.errorFileNotFound
}

//custom
func (globErrs *ErrMap) ErrConnectionTimeout() error {
	return globErrs.errorConnectionTimeout
}

func (globErrs *ErrMap) ErrBeingDeployed() error {
	return globErrs.errorBeingDeployed
}

func (globErrs *ErrMap) ErrConnection() error {
	return globErrs.errorConnection
}

func (globErrs *ErrMap) ErrSshCommands() error {
	return globErrs.errorBadSshCommands
}

//scripts
func (globErrs *ErrMap) ErrJobAlreadyStarted() error {
	return globErrs.errorJobAlreadyStarted
}

func (globErrs *ErrMap) ErrScriptPipInstall() error {
	return globErrs.errorScriptPipInstall
}

func (globErrs *ErrMap) ErrJobWasNotStarted() error {
	return globErrs.errorJobWasNotStarted
}
