package cmdconfig

import (
	"github.com/turbot/go-kit/files"
	"github.com/turbot/pipe-fittings/app_specific"
	"github.com/turbot/pipe-fittings/cmdconfig"
	"github.com/turbot/pipe-fittings/constants"
	steampipeversion "github.com/turbot/steampipe/pkg/version"
)

// SetAppSpecificConstants sets app specific constants defined in pipe-fittings
func SetAppSpecificConstants() {
	// set the default install dir
	installDir, err := files.Tildefy("~/.steampipe")
	if err != nil {
		panic(err)
	}
	app_specific.AppName = "steampipe"
	app_specific.AppVersion = steampipeversion.SteampipeVersion
	app_specific.AutoVariablesExtension = ".auto.spvars"
	app_specific.ClientConnectionAppNamePrefix = "steampipe_client"
	app_specific.ClientSystemConnectionAppNamePrefix = "steampipe_client_system"
	app_specific.DefaultInstallDir = installDir
	app_specific.DefaultVarsFileName = "steampipe.spvars"
	app_specific.DefaultWorkspaceDatabase = "local"
	app_specific.ModDataExtension = ".sp"
	app_specific.ModFileName = "mod.sp"
	app_specific.ServiceConnectionAppNamePrefix = "steampipe_service"
	app_specific.VariablesExtension = ".spvars"
	app_specific.WorkspaceIgnoreFile = ".steampipeignore"
	app_specific.WorkspaceDataDir = ".steampipe"
	app_specific.EnvAppPrefix = "STEAMPIPE_"
	// EnvInputVarPrefix is the prefix for environment variables that represent values for input variables.
	app_specific.EnvInputVarPrefix = "SP_VAR_"
	// set the command pre and post hooks
	cmdconfig.CustomPreRunHook = PreRunHook
	cmdconfig.CustomPostRunHook = PostRunHook

	// set path arguments
	app_specific.FilePathKeys = []string{
		constants.ArgSnapshotLocation,
		constants.ArgModLocation,
		constants.ArgInstallDir,
		constants.ArgOutputDir,
		constants.ArgLogDir,
	}
}
