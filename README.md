Here is the file for my booking app. 
A Go (golang) port of the Ruby dotenv project (which loads env vars from a .env file).

From the original Library:

Storing configuration in the environment is one of the tenets of a twelve-factor app. Anything that is likely to change between deployment environments–such as resource handles for databases or credentials for external services–should be extracted from the code into environment variables.

But it is not always practical to set environment variables on development machines or continuous integration servers where multiple projects are run. Dotenv load variables from a .env file into ENV when the environment is bootstrapped.

It can be used as a library (for loading in env for your own daemons etc.) or as a bin command.

There is test coverage and CI for both linuxish and Windows environments, but I make no guarantees about the bin version working on Windows.

Installation
As a library

go get github.com/joho/godotenv
or if you want to use it as a bin command

go >= 1.17

go install github.com/joho/godotenv/cmd/godotenv@latest
go < 1.17

go get github.com/joho/godotenv/cmd/godotenv
Usage
Add your application configuration to your .env file in the root of your project:

S3_BUCKET=YOURS3BUCKET
SECRET_KEY=YOURSECRETKEYGOESHERE
Then in your Go app you can do something like

package main

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  s3Bucket := os.Getenv("S3_BUCKET")
  secretKey := os.Getenv("SECRET_KEY")

  // now do something with s3 or whatever
}
If you're even lazier than that, you can just take advantage of the autoload package which will read in .env on import

import _ "github.com/joho/godotenv/autoload"
While .env in the project root is the default, you don't have to be constrained, both examples below are 100% legit

godotenv.Load("somerandomfile")
godotenv.Load("filenumberone.env", "filenumbertwo.env")
If you want to be really fancy with your env file you can do comments and exports (below is a valid env file)

# I am a comment and that is OK
SOME_VAR=someval
FOO=BAR # comments at line end are OK too
export BAR=BAZ
Or finally you can do YAML(ish) style

FOO: bar
BAR: baz
as a final aside, if you don't want godotenv munging your env you can just get a map back instead

var myEnv map[string]string
myEnv, err := godotenv.Read()

s3Bucket := myEnv["S3_BUCKET"]
... or from an io.Reader instead of a local file

reader := getRemoteFile()
myEnv, err := godotenv.Parse(reader)
... or from a string if you so desire

content := getRemoteFileContent()
myEnv, err := godotenv.Unmarshal(content)
Precedence & Conventions
Existing envs take precedence of envs that are loaded later.

The convention for managing multiple environments (i.e. development, test, production) is to create an env named {YOURAPP}_ENV and load envs in this order:

env := os.Getenv("FOO_ENV")
if "" == env {
  env = "development"
}

godotenv.Load(".env." + env + ".local")
if "test" != env {
  godotenv.Load(".env.local")
}
godotenv.Load(".env." + env)
godotenv.Load() // The Original .env
If you need to, you can also use godotenv.Overload() to defy this convention and overwrite existing envs instead of only supplanting them. Use with caution.

Command Mode
Assuming you've installed the command as above and you've got $GOPATH/bin in your $PATH

godotenv -f /some/path/to/.env some_command with some args
If you don't specify -f it will fall back on the default of loading .env in PWD

By default, it won't override existing environment variables; you can do that with the -o flag.

Writing Env Files
Godotenv can also write a map representing the environment to a correctly-formatted and escaped file

env, err := godotenv.Unmarshal("KEY=value")
err := godotenv.Write(env, "./.env")
... or to a string

env, err := godotenv.Unmarshal("KEY=value")
content, err := godotenv.Marshal(env)
