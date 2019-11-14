# Jcli
A slim command line app for triggering jenkins builds and downloading artifacts.

## Quick Reference
***
### Download a copy of Jcli 
To download a copy of Jcli for your OS, you can visit http://jacobstevenwakefield.com/.

### Get Jenkins Api Key
Go to you profile in jenkins > Configure > Api Token > Add New Token .

### Generate Config File:
To generate a config file you can run the command ```./jcli -generate```. You will then be prompted to enter the fields necessary to connect to jenkins. You'll then be prompted for a file location to store the config. ~/.jcli is the default location. If you store it anywhere else you will have to add the parameters -config to specify its location when running jcli. 

### Get All Jobs on the Jenkins Server:
To display a list of jenkins jobs you can run the command ```./jcli -jobs```. 

### Get Information on a Job: 
To get a job's information run the command ```./jcli -job <job name> -info```. Where \<job name\> is the exact name of the jenkins job. Note: most commands require a -job parameter. 

### Get Artifacts From Last Build:
To get the artifacts from the last build for a given job you can just run ```./jcli -job <job name>```. This will download all the artifacts from that build into the current directory.

### Get Information on a Specific Build: 
To get information on a specific build run the command ```./jcli -job <job name> -build <build number> -info```. Where \<job name\> is the job you want to target and \<build number\> is the build being targeted. 

### Get Artifacts from a Specific Build: 
To get the artifacts from a specific build run the command ```./jcli -job <job name> -build <build number>```. This command will download all the artifacts from a specific build into the current directory. 

### Trigger a Build for a Given Job
To trigger a build run the command ```./jcli -job <job name> -trigger```. This will trigger the next build for the job.

### The Help Command
To get a help page on the options available to jcli run ```./jcli -help```. Warning: There are options that have not been implemented yet. However, all of the above commands should function.

## TODO:
***
* add -o/output option to view the console output for a given build or live output of a build that's building.
* add an output option for downloading artifacts so they aren't just crammed into the current directory.
* figure out date format so printing the date looks better on output.
* add last stable, last successful etc for downloading artifacts.