package lib

import (
	"fmt"
	"os"

	"github.com/medisafe/jenkins-api/jenkins"
)

// StartJenkinsCommands - a function that starts the process of making requests to jenkins
func StartJenkinsCommands(conf *JcliConfig, params *JcliParams) {
	jenkinsApi := jenkins.Init(&jenkins.Connection{
		Username:    conf.Username,
		AccessToken: conf.Apikey,
		BaseUrl:     conf.JenkinsUrl,
	})

	if params.Jobs {
		jobs, _ := jenkinsApi.GetJobs()
		prettyPrintJobsList(jobs)
	} else {
		job, _ := jenkinsApi.GetJob(params.Job)

		if params.HasExtraParams() && params.Info { //if job was the only option or if info flag is present
			prettyPrintJob(job)
		} else if !params.HasExtraParams() {
			downloadArtifactsForBuild(jenkinsApi, job, job.LastBuild.Number)
		}

		if params.Build != BuildParamDefault && params.Info { //if a build number is supplied and user wants info about the build
			build, err := jenkinsApi.GetBuild(params.Job, params.Build)
			if err != nil {
				fmt.Println("Build could not be found :(.")
				os.Exit(1)
			}
			prettyPrintBuild(build)
		} else if params.Build != BuildParamDefault { //if just a build number is supplied then download the artifacts from the specific build
			downloadArtifactsForBuild(jenkinsApi, job, params.Build)
		}

		if params.Trigger { //if the user is triggering the next build
			jenkinsApi.StartBuild(params.Job, map[string]interface{}{})
		}
	}
}

// prettyPrintJobsList - a function to print out the names of the jobs on the jenkins server
func prettyPrintJobsList(jobs []jenkins.Job) {
	fmt.Println("----------------------Jobs List-----------------------")
	for _, job := range jobs {
		fmt.Printf("\t%s\n", job.Name)
	}
	fmt.Println("---------------------End Jobs List---------------------")
}

// prettyPrintJob - a function to print out a jenkins.Job struct (only useful info)
func prettyPrintJob(job *jenkins.Job) {
	fmt.Println("---------------------Job Info----------------------")
	fmt.Printf("\tDisplayName: %s\n", job.DisplayName)
	fmt.Printf("\tName: %s\n", job.Name)
	fmt.Printf("\tDescription: %s\n", job.Description)
	fmt.Printf("\tUrl: %s\n", job.Url)
	fmt.Printf("\tBuildable: %t\n", job.Buildable)
	fmt.Printf("\tColor: %s\n", job.Color)
	fmt.Printf("\tInQueue: %t\n", job.InQueue)
	fmt.Printf("\tNextBuildNumber: %d\n", job.NextBuildNumber)
	fmt.Println("--------------------End Job Info--------------------")
}

// prettPrintBuild - a function to print out a jenkins.Build struct (only useful info)
func prettyPrintBuild(build *jenkins.Build) {
	fmt.Println("---------------------Build Info----------------------")
	fmt.Printf("\t FullDisplayName: %s\n", build.FullDisplayName)
	fmt.Printf("\t DisplayName: %s\n", build.DisplayName)
	fmt.Printf("\t Number: %d\n", build.Number)
	fmt.Printf("\t Id: %s\n", build.Id)
	fmt.Printf("\t Description: %s\n", build.Description)
	fmt.Printf("\t Result: %s\n", build.Result)
	fmt.Printf("\t Duration: %f sec\n", float64(build.Duration)/1000)
	fmt.Printf("\t Estimated Duration: %f sec\n", float64(build.EstimatedDuration)/1000)
	fmt.Printf("\t Timestamp: %d\n", build.Timestamp) //TODO: replace with actual time stamp string
	fmt.Printf("\t Url: %s\n", build.Url)
	fmt.Printf("\t Building: %t\n", build.Building)
	fmt.Printf("\t QueueId: %d\n", build.QueueId)
	fmt.Printf("\t Artifacts: \n")
	for i, artifact := range build.Artifacts {
		fmt.Printf("\t\t%d :%sartifact/%s\n", i, build.Url, artifact.RelativePath)
	}
	fmt.Println("--------------------End Build Info--------------------")
}

func downloadArtifactsForBuild(jenkinsApi *jenkins.JenkinsApi, job *jenkins.Job, buildNum int) {
	lastBuild, err := jenkinsApi.GetBuild(job.Name, buildNum)
	fmt.Printf("Downloading artifacts for job: %s from build: %d\n", job.DisplayName, lastBuild.Number) //get the last build to obtain the artifacts

	if err != nil {
		fmt.Printf("Error: could not retreive build for job: %s from build: %d\n", job.DisplayName, buildNum) //get the last build to obtain the artifacts
		os.Exit(1)
	}

	//download all artifacts from build
	for i, artifact := range lastBuild.Artifacts {
		fmt.Printf("%d: %s\n", i, artifact.FileName)
		url := getArtifactDownloadUrl(lastBuild.Url, &artifact)
		fmt.Println(url)
		if err := DownloadFile(url, artifact.FileName); err != nil {
			fmt.Printf("Error: An error occured while downloading artifact %s\n", artifact.FileName)
		}
	}
}

// getArtifactDownloadUrl - get all of the web urls for each artifact
func getArtifactDownloadUrl(buildUrl string, artifact *jenkins.Artifact) string {
	return fmt.Sprintf("%sartifact/%s", buildUrl, artifact.RelativePath)
}
