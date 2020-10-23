package main

import (
  "github.com/mitchellh/go-ps"
  //"fmt"
  "log"
  "strings"
  "os/exec"
)


func findPlex()bool{
  process, err := ps.Processes()
  if err != nil {
    log.Fatalf("uh oh, stinkyyyyy: %s", err)
}
  //fmt.Printf("%#v", process[0].Executable())
  var plexIsUp = false
  for i := 0; i < len(process); i++{
    //fmt.Println(process[i].Executable()) Prints names of programs running.
  if strings.Contains(process[i].Executable(), "Plex_Media_Server"){
    plexIsUp = true
    return plexIsUp
  }
  }
return false}

func plexRestarter(){
  cmd := exec.Command("service", "plexmediaserver", "start")
  log.Println("Restarting 'plexmediaserver', please wait.")
  err := cmd.Run()
  log.Println("Command finished with error: %v", err)
}


func main() {
	//function to detect plex isn't runnning

  if findPlex(){
  } else{
    plexRestarter()
  }
	//plex restarter function
}
