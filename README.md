# Fluxy

Fluxy is designed to be a simple CLI tool to use with [FluxCD v1](https://docs.fluxcd.io/en/1.21.1/) configured repositories.
It will enable you to:
- Enable automatic deploys.
- Change container versions in deployments.

## :bomb: Installation

* Clone the project
* Being at the root level project do:
````shell
go install
````
* Enjoy

## :computer: Commands
The majority of commands are pretty self-explanatory but here is a more detailed explanation
of what each one of them do exactly.

### Create
````shell
fluxy create -p projectName -a annotationsFilePath -fp fluxPatchFilePath
````

**Description**: creates the settings configuration for a specific project.

**Details**: the command needs as an input the name of the project (that you're going
to use in other commands) and two important paths to files. The path to where you specified
your fluxcd annotations, and the file for the flux-patch. The former is for disabling the 
automatic deployment configuration, and the second if to update the image tag.

This configuration is saved at your home in a file `.fluxy`, if you want to configure manually:
- Format expected: yaml.
- Strucutre: 
```
projects:
  example:
    annotations: someAbsolutePath
    fluxPatch: otherAbsolutePath
```

### Deploy
````shell
fluxy deploy -p projectName -c containerName -t tag
````

**Description**: changes the tag version in configuration files.

**Details**: with the information given previously (using `create`) It is going to
update the files with the new tag for the corresponding container. Bear in mind that
the tag is updated in the image considering that the tag image doesn't contain a `:`.
In that case the replacement won't be as expected.

### Automated
````shell
fluxy automated -p projectName
````

**Description**: enables automated deploys

**Details**: updates the annotations in your configuration files to enable again FluxCD
normal flow.


## :notebook: TL;DR
This tool is to change a few files with a single command.


:wave: