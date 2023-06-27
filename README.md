# Exads Go

Example configuration (config.yaml):

    exads-core:
        name: exads-core
        path: /tmp/repos/foo
    www-network-api-exads:
        name: www-network-api-exads
        path: /tmp/repos/foo1
    www-api-exads:
        name: www-api-exads
        path: /tmp/repos/foo2
    jobs:
        name: jobs
        path: /tmp/repos/foo3
        
## Available Commands

    checkout    Checkout and pull the branch given as argument
    completion  Generate the autocompletion script for the specified shell
    help        Help about any command
    status      Display current branch and last update

## TODO

- [ ] Installation script (curl | bash ) style
- [ ] Consider various setups (symlink vs no symlink)
- [ ] Update dependencies as well (vendor)
- [ ] Add symlink functionality
- [ ] Add logs functionality
- [ ] Add self-update functionality
- [ ] Create changelog and version number (use semver)
- [ ] Generate config command
- [ ] Correct help message and 'took ...' message
