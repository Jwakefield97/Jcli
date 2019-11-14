package lib

const JobParam = "job"
const JobParamDefault = ""
const JobParamHelp = "The name of the job to retrieve (MUST INCLUDE). If lone param, downloads the last artifact built."

const JobsParam = "jobs"
const JobsParamDefault = false
const JobsParamHelp = "Get a list of all jenkins jobs on the server (MUST BE LONE PARAM)."

const BuildParam = "build"
const BuildParamDefault = -1
const BuildParamHelp = "Build number of the job to retrieve (OPTIONAL). If not included last successful is assumed."

const InfoParam = "info"
const InfoParamDefault = false
const InfoParamHelp = "Boolean param indicating whether you want information about the job displayed (OPTIONAL)."

const TriggerParam = "trigger"
const TriggerParamDefault = false
const TriggerParamHelp = "Boolean param indicating whether you want to trigger the next build or not (OPTIONAL)."

const GenerateParam = "generate"
const GenerateParamDefault = false
const GenerateParamHelp = "Boolean param used to generate config files (OPTIONAL). The program will ask for user input to generate the file at a location. If no location is provided the file will be created at ~/.jcli."

const ConfigParam = "config"
const ConfigParamDefault = ""
const ConfigParamHelp = "An alternate location for a .jcli config file (OPTIONAL)."

const ConfigFileTemplate = `Username = "{{.Username}}"
Apikey = "{{.Apikey}}"
JenkinsUrl = "{{.JenkinsUrl}}"
`