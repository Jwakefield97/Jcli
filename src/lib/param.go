package lib

import (
	"flag"
)

// JcliParams - struct that holds the command line parameters passed into the app
type JcliParams struct {
	Job string //name of the jenkins job
	Jobs bool //used to get a list of jobs on the server
	Build int //build number of a job
	Info bool //whether to display job info or not
	Trigger bool //whether you want to trigger the build or not
	Generate bool //bool indicating whether to generate a config or not 
	Config string //alternate config location
}

// HasMinParams - a function that returns a boolean indicating whether the minimum number of params where passed in
func (params *JcliParams) HasMinParams() bool {
	return !IsBlank(params.Job) || params.Jobs
}

// HasExtraParams - a function that returns a boolean indicating whether any parameters passed the minimum were passed in 
func (params *JcliParams) HasExtraParams() bool {
	//.Config is not added to this function because it is used before any other actions take place
	return (
		params.Build != BuildParamDefault || 
		params.Trigger != TriggerParamDefault || 
		params.Info != InfoParamDefault || 
		params.Generate != GenerateParamDefault)
}

// NewJcliParams - a constructor for JcliParams struct
func NewJcliParams() JcliParams {
	params := JcliParams{}
	params.Job = JobParamDefault
	params.Jobs = JobsParamDefault
	params.Build = BuildParamDefault
	params.Info = InfoParamDefault
	params.Trigger = TriggerParamDefault
	params.Generate = GenerateParamDefault
	params.Config = ConfigParamDefault
	return params
}

// GeParameters - a function to populate the JcliParams struct with command line argument values
func GetParameters() JcliParams {
	params := NewJcliParams()

	flag.StringVar(&params.Job, JobParam, JobParamDefault, JobParamHelp)
	flag.BoolVar(&params.Jobs, JobsParam, JobsParamDefault, JobsParamHelp)
	flag.IntVar(&params.Build, BuildParam, BuildParamDefault, BuildParamHelp)
	flag.BoolVar(&params.Info, InfoParam, InfoParamDefault, InfoParamHelp)
	flag.BoolVar(&params.Trigger, TriggerParam, TriggerParamDefault, TriggerParamHelp)	
	flag.BoolVar(&params.Generate, GenerateParam, GenerateParamDefault, GenerateParamHelp)	
	flag.StringVar(&params.Config, ConfigParam, ConfigParamDefault, ConfigParamHelp)
	flag.Parse()
	
	return params
}