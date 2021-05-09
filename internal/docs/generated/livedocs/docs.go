// Code generated by "mdtogo"; DO NOT EDIT.
package livedocs

var LiveShort = `Deploy local packages to a cluster.`
var LiveLong = `
The ` + "`" + `live` + "`" + ` command group contains subcommands for deploying local
` + "`" + `kpt` + "`" + ` packages to a cluster.
`

var ApplyShort = `Apply a package to the cluster (create, update, prune).`
var ApplyLong = `
  kpt live apply [PKG_PATH|-] [flags]

Args:
  PKG_PATH|-:
    Path to the local package which should be applied to the cluster. It must 
    contain a Kptfile with inventory information. Defaults to the current working
    directory.
    Using '-' as the package path will cause kpt to read resources from stdin.

Flags:
  --field-manager:
    Identifier for the **owner** of the fields being applied. Only usable
    when --server-side flag is specified. Default value is kubectl.
  
  --force-conflicts:
    Force overwrite of field conflicts during apply due to different field
    managers. Only usable when --server-side flag is specified.
    Default value is false (error and failure when field managers conflict).
    
  --install-resource-group:
    Install the ResourceGroup CRD into the cluster if it isn't already
    available. Default is false.
  
  --inventory-policy:
    Determines how to handle overlaps between the package being currently applied
    and existing resources in the cluster. The available options are:
    
      * strict: If any of the resources already exist in the cluster, but doesn't
        belong to the current package, it is considered an error.
      * adopt: If a resource already exist in the cluster, but belongs to a 
        different package, it is considered an error. Resources that doesn't belong
        to other packages are adopted into the current package.
        
    The default value is ` + "`" + `strict` + "`" + `.
    
  --output:
    Determines the output format for the status information. Must be one of the following:
    
      * events: The output will be a list of the status events as they become available.
      * json: The output will be a list of the status events as they become available,
        each formatted as a json object.
      * table: The output will be presented as a table that will be updated inline
        as the status of resources become available.
  
    The default value is ‘events’.
    
  --poll-period:
    The frequency with which the cluster will be polled to determine
    the status of the applied resources. The default value is 2 seconds.
  
  --prune-propagation-policy:
    The propagation policy that should be used when pruning resources. The
    default value here is 'Background'. The other options are 'Foreground' and 'Orphan'.
  
  --prune-timeout:
    The threshold for how long to wait for all pruned resources to be
    deleted before giving up. If this flag is not set, kpt live apply will not
    wait. In most cases, it would also make sense to set the
    --prune-propagation-policy to Foreground when this flag is set.
  
  --reconcile-timeout:
    The threshold for how long to wait for all resources to reconcile before
    giving up. If this flag is not set, kpt live apply will not wait for
    resources to reconcile.
  
  --server-side:
    Perform the apply operation server-side rather than client-side.
    Default value is false (client-side).
`

var DestroyShort = `Remove all previously applied resources in a package from the cluster`
var DestroyLong = `
  kpt live destroy [PKG_PATH|-]

Args:

  PKG_PATH|-:
    Path to the local package which should be deleted from the cluster. It must
    contain a Kptfile with inventory information. Defaults to the current working
    directory.
    Using '-' as the package path will cause kpt to read resources from stdin.

Flags:
  --inventory-policy:
    Determines how to handle overlaps between the package being currently applied
    and existing resources in the cluster. The available options are:
    
      * strict: If any of the resources already exist in the cluster, but doesn't
        belong to the current package, it is considered an error.
      * adopt: If a resource already exist in the cluster, but belongs to a 
        different package, it is considered an error. Resources that doesn't belong
        to other packages are adopted into the current package.
        
    The default value is ` + "`" + `strict` + "`" + `.
  
  --output:
    Determines the output format for the status information. Must be one of the following:
    
      * events: The output will be a list of the status events as they become available.
      * json: The output will be a list of the status events as they become available,
        each formatted as a json object.
      * table: The output will be presented as a table that will be updated inline
        as the status of resources become available.
  
    The default value is ‘events’.
`
var DestroyExamples = `
  # remove all resources in the current package from the cluster.
  kpt live destroy
`

var DiffShort = `Display the diff between the local package and the live cluster resources.`
var DiffLong = `
  kpt live diff [PKG_PATH|-]

Args:
  PKG_PATH|-:
    Path to the local package which should be diffed against the cluster. It must
    contain a Kptfile with inventory information. Defaults to the current working
    directory.
    Using '-' as the package path will cause kpt to read resources from stdin.

Environment Variables:
  KUBECTL_EXTERNAL_DIFF:
    Commandline diffing tool ('diff; by default) that will be used to show
    changes.
    
    # Use meld to show changes
    KPT_EXTERNAL_DIFF=meld kpt live diff

Exit statuses:
  The following exit values are returned:
  
    * 0: No differences were found.
    * 1: Differences were found.
    * >1 kpt live or diff failed with an error.
`
var DiffExamples = `
  # diff the config in the current directory against the live cluster resources.
  kpt live diff
  
  # specify the local diff program to use.
  export KUBECTL_EXTERNAL_DIFF=meld; kpt live diff my-dir
`

var InitShort = `Initialize a package with the information needed for inventory tracking.`
var InitLong = `
  kpt live init [PKG_PATH] [flags]

Args:
  PKG_PATH:
    Path to the local package which should be updated with inventory information.
    It must contain a Kptfile. Defaults to the current working directory.

Flags:
  --force:
    Forces the inventory values to be updated, even if they are already set.
    Defaults to false.
  
  --inventory-id:
    Inventory identifier for the package. This is used to detect overlap between 
    packages that might use the same name and namespace for the inventory object.
    Defaults to an auto-generated value.
  
  --name:
    The name for the ResourceGroup resource that contains the inventory
    for the package. Defaults to the name of the package.
  
  --namespace:
    The namespace for the ResourceGroup resource that contains the inventory
    for the package. If not provided, kpt will check if all the resources
    in the package belong in the same namespace. If they do, that namespace will 
    be used. If they do not, the namespace in the user's context will be chosen.
`
var InitExamples = `
  # initialize a package in the current directory.
  kpt live init

  # initialize a package with a specific name for the group of resources.
  kpt live init --namespace=test my-dir
`

var PreviewShort = `Preview the changes apply would make to the cluster`
var PreviewLong = `
  kpt live preview [PKG_PATH|-] [flags]

Args:

  PKG_PATH|-:
    Path to the local package for which a preview of the operations of apply
    or destroy should be displayed. It must contain a Kptfile with inventory
    information. Defaults to the current working directory.
    Using '-' as the package path will cause kpt to read resources from stdin.

Flags:

  --destroy:
    If true, preview deletion of all resources.
  
  --field-manager:
    Identifier for the **owner** of the fields being applied. Only usable
    when --server-side flag is specified. Default value is kubectl.
  
  --force-conflicts:
    Force overwrite of field conflicts during apply due to different field
    managers. Only usable when --server-side flag is specified.
    Default value is false (error and failure when field managers conflict).
  
  --install-resource-group:
    Install the ResourceGroup CRD into the cluster if it isn't already
    available. Default is false.
  
  --inventory-policy:
    Determines how to handle overlaps between the package being currently applied
    and existing resources in the cluster. The available options are:
    
      * strict: If any of the resources already exist in the cluster, but doesn't
        belong to the current package, it is considered an error.
      * adopt: If a resource already exist in the cluster, but belongs to a 
        different package, it is considered an error. Resources that doesn't belong
        to other packages are adopted into the current package.
        
    The default value is ` + "`" + `strict` + "`" + `.
  
  --output:
    Determines the output format for the status information. Must be one of the following:
    
      * events: The output will be a list of the status events as they become available.
      * json: The output will be a list of the status events as they become available,
        each formatted as a json object.
      * table: The output will be presented as a table that will be updated inline
        as the status of resources become available.
  
    The default value is ‘events’.
  
  --server-side:
    Perform the apply operation server-side rather than client-side.
    Default value is false (client-side).
`
var PreviewExamples = `
  # preview apply for the package in the current directory. 
  kpt live preview

  # preview destroy for a package in the my-dir directory.
  kpt live preview --destroy my-dir
`

var StatusShort = `Display shows the status for the resources in the cluster`
var StatusLong = `
  kpt live status [PKG_PATH|-] [flags]

Args:

  PKG_PATH|-:
    Path to the local package for which the status of the package in the cluster
    should be displayed. It must contain a Kptfile with inventory information.
    Defaults to the current working directory.
    Using '-' as the package path will cause kpt to read resources from stdin.

Flags:

  --output:
    Determines the output format for the status information. Must be one of the following:
    
      * events: The output will be a list of the status events as they become available.
      * json: The output will be a list of the status events as they become available,
        each formatted as a json object.
      * table: The output will be presented as a table that will be updated inline
        as the status of resources become available.
  
    The default value is ‘events’.
    
  --poll-period:
    The frequency with which the cluster will be polled to determine the status
    of the applied resources. The default value is 2 seconds.
  
  --poll-until:
    When to stop polling for status and exist. Must be one of the following:
    
      * known: Exit when the status for all resources have been found.
      * current: Exit when the status for all resources have reached the Current status.
      * deleted: Exit when the status for all resources have reached the NotFound
        status, i.e. all the resources have been deleted from the live state.
      * forever: Keep polling for status until interrupted.
    
    The default value is ‘known’.
  
  --timeout:
    Determines how long the command should run before exiting. This deadline will
    be enforced regardless of the value of the --poll-until flag. The default is
    to wait forever.
`
var StatusExamples = `
  # Monitor status for the resources belonging to the package in the current
  # directory. Wait until all resources have reconciled.
  kpt live status

  # Monitor status for the resources belonging to the package in the my-app
  # directory. Output in table format:
  kpt live status my-app --poll-until=forever --output=table
`
